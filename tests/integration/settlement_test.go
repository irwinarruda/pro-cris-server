package integration

import (
	"testing"

	"github.com/irwinarruda/pro-cris-server/libs/proinject"
	"github.com/irwinarruda/pro-cris-server/modules/settlements"
	"github.com/stretchr/testify/assert"
)

func TestSettlementService(t *testing.T) {
	Init()
	var assert = assert.New(t)

	t.Run("Happy Path", func(t *testing.T) {
		beforeEachSettlement()
		assert.Equal(1+1, 2, "1 + 1 should be equal to 2")
		afterEachSettlement()
	})

	t.Run("Error Path", func(t *testing.T) {
		beforeEachSettlement()
		afterEachSettlement()
	})

}

func beforeEachSettlement() {
	var settlementRepository = proinject.Get[settlements.ISettlementRepository]("settlement_repository")
	settlementRepository.ResetSettlement()
}

func afterEachSettlement() {
	var settlementRepository = proinject.Get[settlements.ISettlementRepository]("settlement_repository")
	settlementRepository.ResetSettlement()
}
