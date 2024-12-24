package settlements

import (
	"time"

	"github.com/irwinarruda/pro-cris-server/modules/appointments"
	"github.com/irwinarruda/pro-cris-server/shared/models"
)

type Settlement struct {
	models.SettlementOptions
	ID           int                        `json:"id"`
	Student      SettlementStudent          `json:"student"`
	StartDate    time.Time                  `json:"startDate"`
	EndDate      time.Time                  `json:"endDate"`
	IsSettled    bool                       `json:"isSettled"`
	IsDeleted    bool                       `json:"isDeleted"`
	TotalAmount  float64                    `json:"totalAmount"`
	CreatedAt    time.Time                  `json:"createdAt"`
	UpdatedAt    time.Time                  `json:"updatedAt"`
	Appointments []appointments.Appointment `json:"appointments"`
}

type SettlementStudent struct {
	ID           int     `json:"id"`
	Name         string  `json:"name"`
	DisplayColor string  `json:"displayColor"`
	Picture      *string `json:"picture"`
}
