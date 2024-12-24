package settlements

import (
	"net/http"
	"time"

	"github.com/irwinarruda/pro-cris-server/libs/proinject"
	"github.com/irwinarruda/pro-cris-server/modules/appointments"
	"github.com/irwinarruda/pro-cris-server/modules/students"
	"github.com/irwinarruda/pro-cris-server/shared/configs"
	"github.com/irwinarruda/pro-cris-server/shared/models"
	"github.com/irwinarruda/pro-cris-server/shared/utils"
)

type SettlementService struct {
	SettlementRepository ISettlementRepository            `inject:"settlement_repository"`
	StudentService       students.IStudentService         `inject:"student_service"`
	AppointmentService   appointments.IAppointmentService `inject:"appointment_service"`
	Validate             configs.Validate                 `inject:"validate"`
}

type ISettlementService = *SettlementService

func NewSettlementService() *SettlementService {
	return proinject.Resolve(&SettlementService{})
}

func (s *SettlementService) GetLastSettlementByStudent(data GetLastSettlementByStudentDTO) (Settlement, error) {
	if err := s.Validate.Struct(data); err != nil {
		return Settlement{}, err
	}
	settlements, err := s.SettlementRepository.GetSettlementsByStudent(GetSettlementsByStudentDTO(data))
	if err != nil {
		return Settlement{}, err
	}
	if len(settlements) == 0 {
		return Settlement{}, utils.NewAppError("No settlements found.", true, http.StatusNotFound)
	}
	last := Settlement{}
	for _, settlement := range settlements {
		if settlement.StartDate.After(last.StartDate) {
			last = settlement
		}
	}
	return last, nil
}

func (s *SettlementService) CreateSettlement(settlement CreateSettlementDTO) (int, error) {
	if err := s.Validate.Struct(settlement); err != nil {
		return 0, err
	}
	student, err := s.StudentService.GetStudentByID(students.GetStudentDTO{
		IDAccount: settlement.IDAccount,
		ID:        settlement.IDStudent,
	})
	if err != nil {
		return 0, err
	}
	settlement.CreateSettlementOptionsDTO = CreateSettlementOptionsDTO{
		PaymentStyle:         student.PaymentStyle,
		PaymentType:          student.PaymentType,
		PaymentTypeValue:     student.PaymentTypeValue,
		SettlementStyle:      student.SettlementStyle,
		SettlementStyleValue: student.SettlementStyleValue,
		SettlementStyleDay:   student.SettlementStyleDay,
	}

	lastSettlement, err := s.GetLastSettlementByStudent(GetLastSettlementByStudentDTO{
		IDAccount: settlement.IDAccount,
		IDStudent: settlement.IDStudent,
	})
	if err == nil {
		if settlement.StartDate.Before(lastSettlement.EndDate) {
			settlement.StartDate = lastSettlement.EndDate.AddDate(0, 0, 1)
		}
	}

	notSettledAppointments, err := s.AppointmentService.GetNotSettledAppointmentsByStudent(appointments.GetNotSettledAppointmentsByStudentDTO{
		IDAccount: settlement.IDAccount,
		IDStudent: settlement.IDStudent,
	})
	if err != nil {
		return 0, err
	}
	if len(notSettledAppointments) == 0 {
		return 0, utils.NewAppError("Settlement cannot be created without appointments.", true, http.StatusBadRequest)
	}
	includedAppointments := []appointments.Appointment{}
	if settlement.SettlementStyle == models.SettlementStyleMonthly {
		timeNow := time.Now()
		settlement.StartDate = time.Date(timeNow.Year(), timeNow.Month(), *settlement.SettlementStyleDay, 0, 0, 0, 0, time.Local)
		settlement.EndDate = settlement.StartDate.AddDate(0, *settlement.SettlementStyleValue, -1)
	} else if settlement.SettlementStyle == models.SettlementStyleWeekly {
		settlement.EndDate = settlement.StartDate.AddDate(0, 0, *settlement.SettlementStyleValue*7)
	}
	if settlement.SettlementStyle == models.SettlementStyleAppointments {
		for i := 0; i < *settlement.SettlementStyleValue; i++ {
			if i >= len(notSettledAppointments) {
				break
			}
			includedAppointments = append(includedAppointments, notSettledAppointments[i])
		}
	} else {
		for _, appointment := range notSettledAppointments {
			if !utils.IsOverlappingDate(settlement.StartDate, settlement.EndDate, appointment.CalendarDay, appointment.CalendarDay) {
				continue
			}
			includedAppointments = append(includedAppointments, appointment)
		}
	}
	if settlement.PaymentType == models.PaymentTypeFixed {
		settlement.TotalAmount = *settlement.PaymentTypeValue
	} else {
		settlement.TotalAmount = 0
		for _, appointment := range includedAppointments {
			settlement.TotalAmount += appointment.Price
		}
	}

	return 0, nil
}

func (s *SettlementService) UpdateSettlementAppointments(appointments []int) {
}

func (s *SettlementService) SettleSettlement() {
}

func (s *SettlementService) DeleteSettlement() {
}
