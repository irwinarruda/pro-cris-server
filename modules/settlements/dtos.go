package settlements

import (
	"time"

	"github.com/irwinarruda/pro-cris-server/shared/models"
)

type GetLastSettlementByStudentDTO struct {
	IDAccount int `json:"idAccount" validate:"required"`
	IDStudent int `json:"idStudent" validate:"required"`
}

type GetSettlementsByStudentDTO struct {
	IDAccount int `json:"idAccount" validate:"required"`
	IDStudent int `json:"idStudent" validate:"required"`
}

type CreateSettlementDTO struct {
	CreateSettlementOptionsDTO
	IDAccount int       `json:"idAccount" validate:"required"`
	IDStudent int       `json:"idStudent" validate:"required"`
	StartDate time.Time `json:"startDate" validate:"required"`
	EndDate   time.Time `json:"endDate" validate:"required"`
}

type CreateSettlementOptionsDTO struct {
	PaymentStyle         models.PaymentStyle    `json:"paymentStyle" validate:"payment_style"`
	PaymentType          models.PaymentType     `json:"paymentType" validate:"payment_type"`
	PaymentTypeValue     *float64               `json:"paymentTypeValue"`
	SettlementStyle      models.SettlementStyle `json:"settlementStyle" validate:"settlement_style"`
	SettlementStyleValue *int                   `json:"settlementStyleValue"`
	SettlementStyleDay   *int                   `json:"settlementStyleDay"`
}
