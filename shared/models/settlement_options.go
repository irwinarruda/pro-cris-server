package models

import "github.com/irwinarruda/pro-cris-server/shared/utils"

type PaymentStyle string

const (
	PaymentStyleUpfront PaymentStyle = "Upfront"
	PaymentStyleLater   PaymentStyle = "Later"
)

func (s *PaymentStyle) UnmarshalJSON(b []byte) (err error) {
	return utils.UnmarshalEnum(s, GetPaymentStylesString(), b)
}

func (p PaymentStyle) String() string { return string(p) }

func GetPaymentStylesString() []string {
	return []string{PaymentStyleUpfront.String(), PaymentStyleLater.String()}
}

type PaymentType string

const (
	PaymentTypeFixed    PaymentType = "Fixed"
	PaymentTypeVariable PaymentType = "Variable"
)

func (p *PaymentType) UnmarshalJSON(b []byte) (err error) {
	return utils.UnmarshalEnum(p, GetPaymentTypesString(), b)
}

func (p PaymentType) String() string { return string(p) }

func GetPaymentTypesString() []string {
	return []string{PaymentTypeFixed.String(), PaymentTypeVariable.String()}
}

type SettlementStyle string

const (
	SettlementStyleAppointments SettlementStyle = "Appointments"
	SettlementStyleWeekly       SettlementStyle = "Weekly"
	SettlementStyleMonthly      SettlementStyle = "Monthly"
)

func (s *SettlementStyle) UnmarshalJSON(b []byte) (err error) {
	return utils.UnmarshalEnum(s, GetSettlementStylesString(), b)
}

func (p SettlementStyle) String() string { return string(p) }

func GetSettlementStylesString() []string {
	return []string{SettlementStyleAppointments.String(), SettlementStyleWeekly.String(), SettlementStyleMonthly.String()}
}

type SettlementOptions struct {
	PaymentStyle         PaymentStyle    `json:"paymentStyle"`
	PaymentType          PaymentType     `json:"paymentType"`
	PaymentTypeValue     *float64        `json:"paymentTypeValue"`
	SettlementStyle      SettlementStyle `json:"settlementStyle"`
	SettlementStyleValue *int            `json:"settlementStyleValue"`
	SettlementStyleDay   *int            `json:"settlementStyleDay"`
}
