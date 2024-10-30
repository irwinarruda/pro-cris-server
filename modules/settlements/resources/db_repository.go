package settlementsresources

import (
	"time"

	"github.com/irwinarruda/pro-cris-server/libs/proinject"
	"github.com/irwinarruda/pro-cris-server/shared/configs"
	"github.com/irwinarruda/pro-cris-server/shared/models"
)

type DbSettlement struct {
	ID                   int
	IDAccount            int
	PaymentStyle         models.PaymentStyle
	PaymentType          models.PaymentType
	PaymentTypeValue     *float64
	SettlementStyle      models.SettlementStyle
	SettlementStyleValue *int
	SettlementStyleDay   *int
	StartDate            time.Time
	EndDate              time.Time
	IsSettled            bool
	IsDeleted            bool
	CreatedAt            time.Time
	UpdatedAt            time.Time
	IDStudent            int
	Name                 string
	DisplayColor         string
	Picture              *string
}

type DbSettlementRepository struct {
	Db configs.Db `inject:"db"`
}

func NewDbSettlementRepository() *DbSettlementRepository {
	return proinject.Resolve(&DbSettlementRepository{})
}

func (a *DbSettlementRepository) ResetSettlement() {
	a.Db.Exec(`DELETE FROM "settlement";`)
	a.Db.Exec(`ALTER SEQUENCE settlement_id_seq RESTART WITH 1;`)
}
