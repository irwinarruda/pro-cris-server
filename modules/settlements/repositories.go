package settlements

type ISettlementRepository interface {
	GetSettlementsByStudent(data GetSettlementsByStudentDTO) ([]Settlement, error)
	CreateSettlement(settlement CreateSettlementDTO) (int, error)
	ResetSettlement()
}
