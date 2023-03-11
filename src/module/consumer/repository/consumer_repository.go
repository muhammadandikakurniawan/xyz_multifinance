package repository

import (
	"context"

	"github.com/muhammadandikakurniawan/xyz_multifinance/src/module/consumer/aggregate/consumer"
	"github.com/muhammadandikakurniawan/xyz_multifinance/src/module/consumer/entity"
	"github.com/muhammadandikakurniawan/xyz_multifinance/src/shared/model"
)

type ConsumerRepository interface {
	Save(ctx context.Context, ag *consumer.ConsumerAggregate) (err error)
	FindTenorLimitByConsumerId(ctx context.Context, consumerId string) (ag *consumer.ConsumerAggregate, err error)
	FindRequestLoanByConsumerId(ctx context.Context, consumerId string) (ag *consumer.ConsumerAggregate, err error)
	FindRequestLoanById(ctx context.Context, requestLoanId int64) (ag *consumer.ConsumerAggregate, err error)
	SearchListRequestLoan(ctx context.Context, paginationreq model.PaginationRequestModel, filter entity.SearchRequestLoanFilterModel) (res []entity.ConsumerEntity, paginationReq model.PaginationResponseModel, err error)
}
