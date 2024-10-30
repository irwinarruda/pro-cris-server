package settlements

import "github.com/irwinarruda/pro-cris-server/libs/proinject"

type SettlementService struct{}

type ISettlementService = *SettlementService

func NewSettlementService() *SettlementService {
	return proinject.Resolve(&SettlementService{})
}

func (s *SettlementService) GetStudentNextSettlement() {
}

func (s *SettlementService) CreateSettlementByStudent() {
}

func (s *SettlementService) SettleSettlement() {
}

func (s *SettlementService) DeleteSettlement() {
}
