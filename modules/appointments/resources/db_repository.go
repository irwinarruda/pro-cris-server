package appointmentsresources

import (
	"fmt"
	"time"

	"github.com/irwinarruda/pro-cris-server/libs/proinject"
	"github.com/irwinarruda/pro-cris-server/modules/appointments"
	"github.com/irwinarruda/pro-cris-server/shared/configs"
	"github.com/irwinarruda/pro-cris-server/shared/utils"
)

type DbAppointment struct {
	ID            int       `json:"id"`
	IDCalendarDay int       `json:"idCalendarDay"`
	IDStudent     int       `json:"idStudent"`
	StartHour     string    `json:"startHour"`
	Duration      int       `json:"duration"`
	Price         float64   `json:"price"`
	IsExtra       bool      `json:"isExtra"`
	IsDeleted     bool      `json:"isDeleted"`
	CreatedAt     time.Time `json:"createdAt"`
	UpdatedAt     time.Time `json:"updatedAt"`
}

func (a *DbAppointment) FromCreateAppointmentDTO(appointment appointments.CreateAppointmentDTO) {
	a.IDStudent = appointment.IDStudent
	a.StartHour = appointment.StartHour
	a.Duration = appointment.Duration
	a.Price = appointment.Price
	a.IsExtra = appointment.IsExtra
}

func (a *DbAppointment) ToAppointment(day appointments.CalendarDay, student appointments.AppointmentStudent) appointments.Appointment {
	return appointments.Appointment{
		ID:          a.ID,
		CalendarDay: day,
		StartHour:   a.StartHour,
		Duration:    a.Duration,
		Price:       a.Price,
		IsExtra:     a.IsExtra,
		Student:     student,
		IsDeleted:   a.IsDeleted,
		CreatedAt:   a.CreatedAt,
		UpdatedAt:   a.UpdatedAt,
	}
}

type DbAppointmentRepository struct {
	Db configs.Db `inject:"db"`
}

func NewDbAppointmentRepository() *DbAppointmentRepository {
	return proinject.Resolve(&DbAppointmentRepository{})
}

func (a *DbAppointmentRepository) CreateAppointment(appointment appointments.CreateAppointmentDTO) int {
	appointmentE := DbAppointment{}
	appointmentE.FromCreateAppointmentDTO(appointment)
	sql := fmt.Sprintf(`
    INSERT INTO "calendar_day"(
      day,
      month,
      year
    ) %s
    RETURNING id;
  `, utils.SqlValues(1, 3))
	a.Db.Raw(
		sql,
		appointment.CalendarDay.Day,
		appointment.CalendarDay.Month,
		appointment.CalendarDay.Year,
	).Scan(&appointmentE.IDCalendarDay)
	sql = fmt.Sprintf(`
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
	a.Db.Raw(
		sql,
		appointmentE.IDCalendarDay,
		appointmentE.IDStudent,
		appointmentE.StartHour,
		appointmentE.Duration,
		appointmentE.Price,
		appointmentE.IsExtra,
	).Scan(&appointmentE.ID)
	return appointmentE.ID
}

func (a *DbAppointmentRepository) GetAppointmentByID(id int) (appointments.Appointment, error) {
	sql := "SELECT * FROM \"appointment\" WHERE id = ?;"
	appointmentE := DbAppointment{}
	a.Db.Raw(sql, id).Scan(&appointmentE)

	sql = "SELECT * FROM \"calendar_day\" WHERE id = ?;"
	day := appointments.CalendarDay{}
	a.Db.Raw(sql, appointmentE.IDCalendarDay).Scan(&day)

	sql = "SELECT id, name, display_color, picture FROM \"student\" WHERE id = ?;"
	appointmentStudent := appointments.AppointmentStudent{}
	a.Db.Raw(sql, appointmentE.IDStudent).Scan(&appointmentStudent)

	return appointmentE.ToAppointment(day, appointmentStudent), nil
}

func (a *DbAppointmentRepository) CreateAppointmentsByRoutine() int {
	return 0
}

func (a *DbAppointmentRepository) ResetAppointments() {

}
