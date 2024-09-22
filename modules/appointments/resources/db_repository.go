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
	IDAccount    int
	ID           int
	StartHour    string
	Duration     int
	Price        float64
	IsExtra      bool
	IsPaid       bool
	IsDeleted    bool
	CalendarDay  time.Time
	Name         string
	IDStudent    int
	DisplayColor string
	Picture      *string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

func (a *DbAppointment) FromCreateAppointmentDTO(appointment appointments.CreateAppointmentDTO) {
	a.IDAccount = appointment.IDAccount
	a.IDStudent = appointment.IDStudent
	a.CalendarDay = appointment.CalendarDay
	a.StartHour = appointment.StartHour
	a.Duration = appointment.Duration
	a.Price = appointment.Price
	a.IsExtra = appointment.IsExtra
	a.IsPaid = appointment.IsPaid
}

func (a *DbAppointment) FromUpdateAppointmentDTO(appointment appointments.UpdateAppointmentDTO) {
	a.IDAccount = appointment.IDAccount
	a.ID = appointment.ID
	a.Price = appointment.Price
	a.IsExtra = appointment.IsExtra
	a.IsPaid = appointment.IsPaid
}

func (a *DbAppointment) ToAppointment() appointments.Appointment {
	return appointments.Appointment{
		ID:          a.ID,
		StartHour:   a.StartHour,
		Duration:    a.Duration,
		Price:       a.Price,
		IsExtra:     a.IsExtra,
		IsPaid:      a.IsPaid,
		IsDeleted:   a.IsDeleted,
		CreatedAt:   a.CreatedAt,
		UpdatedAt:   a.UpdatedAt,
		CalendarDay: a.CalendarDay,
		Student:     appointments.AppointmentStudent{ID: a.IDStudent, Name: a.Name, DisplayColor: a.DisplayColor, Picture: a.Picture},
	}
}

type DbAppointmentRepository struct {
	Db configs.Db `inject:"db"`
}

func NewDbAppointmentRepository() *DbAppointmentRepository {
	return proinject.Resolve(&DbAppointmentRepository{})
}

func (a *DbAppointmentRepository) GetAppointmentByID(data appointments.GetAppointmentDTO) (appointments.Appointment, error) {
	sql := `
    SELECT
      ap.*,
      st.id_account,
      st.name,
      st.display_color,
      st.picture
    FROM "appointment" ap
    LEFT JOIN "student" st ON ap.id_student = st.id
    WHERE ap.id = ?
    AND ap.id_account = ?
    AND ap.is_deleted = false;
  `
	appointmentE := []DbAppointment{}
	result := a.Db.Raw(sql, data.ID, data.IDAccount).Scan(&appointmentE)
	if result.Error != nil {
		return appointments.Appointment{}, utils.NewAppError("Database query error", false, result.Error)
	}
	if len(appointmentE) == 0 {
		return appointments.Appointment{}, utils.NewAppError("Appointment not found.", true, nil)
	}

	return appointmentE[0].ToAppointment(), nil
}

func (a *DbAppointmentRepository) CreateAppointment(appointment appointments.CreateAppointmentDTO) (int, error) {
	appointmentE := DbAppointment{}
	appointmentE.FromCreateAppointmentDTO(appointment)
	sql := fmt.Sprintf(`
    INSERT INTO "appointment"(
      id_account,
      id_student,
      calendar_day,
      start_hour,
      duration,
      price,
      is_extra,
      is_paid
    ) %s
    RETURNING id;
  `, utils.SqlValues(1, 8))
	result := a.Db.Raw(
		sql,
		appointmentE.IDAccount,
		appointmentE.IDStudent,
		appointmentE.CalendarDay,
		appointmentE.StartHour,
		appointmentE.Duration,
		appointmentE.Price,
		appointmentE.IsExtra,
		appointmentE.IsPaid,
	).Scan(&appointmentE.ID)
	if result.Error != nil {
		return 0, utils.NewAppError("Database query error", false, result.Error)
	}
	return appointmentE.ID, nil
}

func (a *DbAppointmentRepository) UpdateAppointment(appointment appointments.UpdateAppointmentDTO) (int, error) {
	appointmentE := DbAppointment{}
	appointmentE.FromUpdateAppointmentDTO(appointment)
	sql := `
    UPDATE "appointment" ap
    SET
      price = ?,
      is_extra = ?,
      is_paid = ?,
      updated_at = CURRENT_TIMESTAMP
    WHERE ap.id = ?
    AND ap.id_account = ?;
  `
	result := a.Db.Exec(sql, appointmentE.Price, appointmentE.IsExtra, appointmentE.IsPaid, appointmentE.ID, appointmentE.IDAccount)
	if result.Error != nil {
		return 0, utils.NewAppError("Database query error", false, result.Error)
	}
	if result.RowsAffected == 0 {
		return 0, utils.NewAppError("Appointment not found.", true, nil)
	}
	return appointmentE.ID, nil
}

func (a *DbAppointmentRepository) DeleteAppointment(data appointments.DeleteAppointmentDTO) (int, error) {
	sql := `
    UPDATE "appointment" ap
    SET
      is_deleted = true,
      updated_at = CURRENT_TIMESTAMP
    WHERE ap.id = ?
    AND ap.id_account = ?;
  `
	result := a.Db.Exec(sql, data.ID, data.IDAccount)
	if result.Error != nil {
		return 0, utils.NewAppError("Database query error", false, result.Error)
	}
	if result.RowsAffected == 0 {
		return 0, utils.NewAppError("Appointment not found.", true, nil)
	}
	return data.ID, nil
}

func (a *DbAppointmentRepository) ResetAppointments() {
	a.Db.Exec(`DELETE FROM "appointment";`)
	a.Db.Exec(`ALTER SEQUENCE appointment_id_seq RESTART WITH 1;`)
}
