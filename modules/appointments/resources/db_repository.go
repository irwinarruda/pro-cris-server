package appointmentsresources

import (
	"fmt"
	"time"

	"github.com/irwinarruda/pro-cris-server/libs/proinject"
	"github.com/irwinarruda/pro-cris-server/modules/appointments"
	"github.com/irwinarruda/pro-cris-server/modules/calendar"
	"github.com/irwinarruda/pro-cris-server/shared/configs"
	"github.com/irwinarruda/pro-cris-server/shared/utils"
)

type DbAppointment struct {
	ID            int
	StartHour     string
	Duration      int
	Price         float64
	IsExtra       bool
	IsDeleted     bool
	IDCalendarDay int
	Day           int
	Month         int
	Year          int
	Name          string
	IDStudent     int
	DisplayColor  string
	Picture       *string
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

func (a *DbAppointment) FromCreateAppointmentDTO(appointment appointments.CreateAppointmentDTO) {
	a.IDStudent = appointment.IDStudent
	a.IDCalendarDay = appointment.CalendarDay.ID
	a.StartHour = appointment.StartHour
	a.Duration = appointment.Duration
	a.Price = appointment.Price
	a.IsExtra = appointment.IsExtra
	a.Day = appointment.CalendarDay.Day
	a.Month = appointment.CalendarDay.Month
	a.Year = appointment.CalendarDay.Year
}

func (a *DbAppointment) ToAppointment() appointments.Appointment {
	return appointments.Appointment{
		ID:          a.ID,
		StartHour:   a.StartHour,
		Duration:    a.Duration,
		Price:       a.Price,
		IsExtra:     a.IsExtra,
		IsDeleted:   a.IsDeleted,
		CreatedAt:   a.CreatedAt,
		UpdatedAt:   a.UpdatedAt,
		CalendarDay: calendar.CalendarDay{ID: a.IDCalendarDay, Day: a.Day, Month: a.Month, Year: a.Year},
		Student:     appointments.AppointmentStudent{ID: a.IDStudent, Name: a.Name, DisplayColor: a.DisplayColor, Picture: a.Picture},
	}
}

type DbAppointmentRepository struct {
	Db configs.Db `inject:"db"`
}

func NewDbAppointmentRepository() *DbAppointmentRepository {
	return proinject.Resolve(&DbAppointmentRepository{})
}

func (a *DbAppointmentRepository) CreateAppointment(appointment appointments.CreateAppointmentDTO) (int, error) {
	appointmentE := DbAppointment{}
	appointmentE.FromCreateAppointmentDTO(appointment)
	sql := fmt.Sprintf(`
    INSERT INTO "appointment"(
      id_calendar_day,
      id_student,
      start_hour,
      duration,
      price,
      is_extra
    ) %s
    RETURNING id;
  `, utils.SqlValues(1, 6))
	err := a.Db.Raw(
		sql,
		appointmentE.IDCalendarDay,
		appointmentE.IDStudent,
		appointmentE.StartHour,
		appointmentE.Duration,
		appointmentE.Price,
		appointmentE.IsExtra,
	).Scan(&appointmentE.ID).Error
	if err != nil {
		return 0, err
	}
	return appointmentE.ID, nil
}

func (a *DbAppointmentRepository) GetAppointmentByID(id int) (appointments.Appointment, error) {
	sql := `
    SELECT
      "appointment".*,
      "calendar_day".day,
      "calendar_day".month,
      "calendar_day".year,
      "student".name,
      "student".display_color,
      "student".picture
    FROM "appointment"
    LEFT JOIN "calendar_day" ON "appointment".id_calendar_day = "calendar_day".id
    LEFT JOIN "student" ON "appointment".id_student = "student".id
    WHERE "appointment".id = ?;
  `
	appointmentE := DbAppointment{}
	result := a.Db.Raw(sql, id).Scan(&appointmentE)
	if result.Error != nil {
		return appointments.Appointment{}, result.Error
	}

	return appointmentE.ToAppointment(), nil
}

func (a *DbAppointmentRepository) CreateAppointmentsByRoutine() (int, error) {
	return 0, nil
}

func (a *DbAppointmentRepository) ResetAppointments() {
	a.Db.Exec(`DELETE FROM "appointment";`)
	a.Db.Exec(`ALTER SEQUENCE appointment_id_seq RESTART WITH 1;`)
}
