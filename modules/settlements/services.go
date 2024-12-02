package settlements

import (
	"net/http"
	"time"

	"github.com/irwinarruda/pro-cris-server/libs/proinject"
	"github.com/irwinarruda/pro-cris-server/modules/students"
	"github.com/irwinarruda/pro-cris-server/shared/configs"
	"github.com/irwinarruda/pro-cris-server/shared/models"
	"github.com/irwinarruda/pro-cris-server/shared/utils"
)

type SettlementService struct {
	SettlementRepository ISettlementRepository    `inject:"settlement_repository"`
	StudentService       students.IStudentService `inject:"student_service"`
	Validate             configs.Validate         `inject:"validate"`
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
		if last.ID == 0 {
			last = settlement
			continue
		}
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
	if err != nil {
		settlement.StartDate = time.Now()
		return s.SettlementRepository.CreateSettlement(settlement)
	}

	if student.SettlementStyle == models.SettlementStyleAppointments {
		settlement.StartDate = lastSettlement.EndDate
		// Can use the last appointments date
		settlement.EndDate = settlement.StartDate.AddDate(0, 0, 30)
	} else if student.SettlementStyle == models.SettlementStyleMonthly {
		settlement.StartDate = lastSettlement.EndDate
		settlement.EndDate = settlement.StartDate.AddDate(0, 1, 0)
	} else if student.SettlementStyle == models.SettlementStyleWeekly {
		settlement.StartDate = lastSettlement.EndDate
		settlement.EndDate = settlement.StartDate.AddDate(0, 0, 7)
	}

	return 0, nil
}

func (s *SettlementService) UpdateSettlementAppointments(appointments []int) {
}

func (s *SettlementService) SettleSettlement() {
}

func (s *SettlementService) DeleteSettlement() {
}
