package consumer_test

import (
	"testing"

	"github.com/muhammadandikakurniawan/xyz_multifinance/src/module/consumer/aggregate/consumer"
	"github.com/muhammadandikakurniawan/xyz_multifinance/src/module/consumer/aggregate/consumer/event"
	"github.com/muhammadandikakurniawan/xyz_multifinance/src/module/consumer/entity"
	sharedErr "github.com/muhammadandikakurniawan/xyz_multifinance/src/shared/error"
	"github.com/stretchr/testify/assert"
)

func Test_RequestLoan(t *testing.T) {
	t.Run("successs", func(t *testing.T) {
		aggregateRoot := entity.ConsumerEntity{
			Id: "d066a2a3-a8cf-47f5-b5ae-d8cffa7071e2",
		}

		aggregateRoot.AddTenorLimit(entity.TenorLimitEntity{
			Month:      1,
			LimitValue: 1500000,
		})

		aggregate := consumer.BuildConsumerAggregate(aggregateRoot)

		requestLoan := entity.RequestLoanEntity{
			ConsumerId:  aggregateRoot.Id,
			OTR:         25000000,
			AdminFee:    500000,
			Installment: 700000,
			Interest:    500000,
			AssetName:   "asset test",
		}

		err := aggregate.AddRequestLoan(&requestLoan)
		assert.Nil(t, err)
		assert.Equal(t, false, requestLoan.ContractNumber == "")

		events := aggregate.GetEvents()
		assert.Equal(t, 1, len(events))

		requestLoanEvent, ok := events[0].(event.RequestLoanEvent)
		assert.Equal(t, true, ok)
		assert.NotNil(t, requestLoanEvent)
		assert.Equal(t, requestLoan.ContractNumber, requestLoanEvent.GetData().ContractNumber)
	})

	t.Run("failed, tenor limit not found", func(t *testing.T) {
		aggregateRoot := entity.ConsumerEntity{
			Id: "d066a2a3-a8cf-47f5-b5ae-d8cffa7071e2",
		}

		aggregate := consumer.BuildConsumerAggregate(aggregateRoot)

		requestLoan := entity.RequestLoanEntity{
			ConsumerId:  aggregateRoot.Id,
			OTR:         25000000,
			AdminFee:    500000,
			Installment: 700000,
			Interest:    500000,
			AssetName:   "asset test",
		}

		err := aggregate.AddRequestLoan(&requestLoan)
		assert.NotNil(t, err)

		appErr, ok := err.(sharedErr.AppError)
		assert.Equal(t, true, ok)
		assert.Equal(t, sharedErr.ERROR_BAD_REQUEST, appErr.ErrorCode)
		assert.Equal(t, "tenor limit not found", appErr.SystemErrorMessage)
	})
}

func Test_AddTenorLimit(t *testing.T) {
	t.Run("successs", func(t *testing.T) {
		aggregateRoot := entity.ConsumerEntity{
			Id: "d066a2a3-a8cf-47f5-b5ae-d8cffa7071e2",
		}

		aggregateRoot.AddTenorLimit(entity.TenorLimitEntity{
			Month:      1,
			LimitValue: 1500000,
		})
		aggregateRoot.AddTenorLimit(entity.TenorLimitEntity{
			Month:      3,
			LimitValue: 3000000,
		})

		aggregate := consumer.BuildConsumerAggregate(aggregateRoot)

		requestLoan := []*entity.TenorLimitEntity{
			&entity.TenorLimitEntity{
				ConsumerId: aggregateRoot.Id,
				Month:      1,
				LimitValue: 3000000,
			},
			&entity.TenorLimitEntity{
				ConsumerId: aggregateRoot.Id,
				Month:      2,
				LimitValue: 5000000,
			},
			&entity.TenorLimitEntity{
				ConsumerId: aggregateRoot.Id,
				Month:      3,
				LimitValue: 0,
			},
		}

		err := aggregate.AddTenorLimit(requestLoan...)
		assert.Nil(t, err)

		events := aggregate.GetEvents()
		assert.Equal(t, 3, len(events))

		// === add event ===
		addTenorLimitEvent, validAddEvent := events[0].(event.AddTenorLimitEvent)
		assert.Equal(t, true, validAddEvent)
		assert.NotNil(t, addTenorLimitEvent)

		addTenorData := addTenorLimitEvent.GetData()
		assert.Equal(t, 1, len(addTenorData))
		assert.Equal(t, 2, addTenorData[0].Month)
		assert.Equal(t, float64(5000000), addTenorData[0].LimitValue)
		// ======

		// === update event ===
		updateTenorLimitEvent, validUpdateEvent := events[1].(event.UpdateTenorLimitEvent)
		assert.Equal(t, true, validUpdateEvent)
		assert.NotNil(t, updateTenorLimitEvent)

		updateTenorData := updateTenorLimitEvent.GetData()
		assert.Equal(t, 1, len(updateTenorData))
		assert.Equal(t, 1, updateTenorData[0].Month)
		assert.Equal(t, float64(3000000), updateTenorData[0].LimitValue)
		// ======

		// === update event ===
		deleteTenorLimitEvent, validDeleteEvent := events[2].(event.DeleteTenorLimitEvent)
		assert.Equal(t, true, validDeleteEvent)
		assert.NotNil(t, deleteTenorLimitEvent)

		deleteTenorData := deleteTenorLimitEvent.GetData()
		assert.Equal(t, 1, len(updateTenorData))
		assert.Equal(t, 3, deleteTenorData[0].Month)
		assert.Equal(t, float64(0), deleteTenorData[0].LimitValue)
		// ======
	})
}
