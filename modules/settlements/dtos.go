package settlements

type GetStudentNextSettlementDTO struct {
	IDAccount int `json:"idAccount" validate:"required"`
	IDStudent int `json:"idStudent" validate:"required"`
}

type CreateSettlementDTO struct {
	IDAccount int `json:"idAccount" validate:"required"`
	IDStudent int `json:"idStudent" validate:"required"`
}
