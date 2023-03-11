package consumer

import (
	"context"

	"github.com/muhammadandikakurniawan/xyz_multifinance/src/module/consumer/usecase/consumer/dto"
	"github.com/muhammadandikakurniawan/xyz_multifinance/src/shared/model"
)

type ConsumerUsecase interface {
	Register(ctx context.Context, requestData dto.RequestCreateNewConsumerDto) (result model.BaseResponseModel[dto.ConsumerId], err error)

	RequestLoan(ctx context.Context, requestData dto.RequestLoanDto) (result model.BaseResponseModel[dto.RequestLoanDto], err error)

	ApproveRequestLoan(ctx context.Context, requestData dto.ApprovalResponseDataDto) (result model.BaseResponseModel[dto.ApprovalResponseDataDto], err error)

	AddTenorLimit(ctx context.Context, requestData dto.AddTenorLmitRequestDto) (result model.BaseResponseModel[dto.AddTenorLmitRequestDto], err error)

	GetListRequestLoan(ctx context.Context, requestData dto.GetListRequestLoanRequestDto) (result model.BaseResponseModel[[]dto.RequestLoanDto], err error)

	GetConsumer(ctx context.Context, consumerId string) (result model.BaseResponseModel[*dto.ConsumerDto], err error)
}
