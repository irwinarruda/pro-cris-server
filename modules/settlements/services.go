package settlements

import "github.com/irwinarruda/pro-cris-server/libs/proinject"

type SettlementService struct {
	SettlementRepository ISettlementRepository `inject:"settlement_repository"`
}

type ISettlementService = *SettlementService

func NewSettlementService() *SettlementService {
	return proinject.Resolve(&SettlementService{})
}

func (s *SettlementService) GetStudentNextSettlement(data GetStudentNextSettlementDTO) {
}

func (s *SettlementService) CreateSettlement(settlement CreateSettlementDTO) {
}

func (s *SettlementService) UpdateSettlementAppointments(appointments []int) {
}

func (s *SettlementService) SettleSettlement() {
}

func (s *SettlementService) DeleteSettlement() {
}
