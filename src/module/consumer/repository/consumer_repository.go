package repository

import (
	"context"

	"github.com/muhammadandikakurniawan/xyz_multifinance/src/module/consumer/aggregate/consumer"
)

type ConsumerRepository interface {
	Save(ctx context.Context, ag *consumer.ConsumerAggregate) (err error)
	FindTenorLimitByConsumerId(ctx context.Context, consumerId string) (ag *consumer.ConsumerAggregate, err error)
	FindRequestLoanByConsumerId(ctx context.Context, consumerId string) (ag *consumer.ConsumerAggregate, err error)
}
