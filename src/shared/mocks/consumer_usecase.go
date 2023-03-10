// Code generated by mockery v2.15.0. DO NOT EDIT.

package mocks

import (
	context "context"

	dto "github.com/muhammadandikakurniawan/xyz_multifinance/src/module/consumer/usecase/consumer/dto"
	mock "github.com/stretchr/testify/mock"

	model "github.com/muhammadandikakurniawan/xyz_multifinance/src/shared/model"
)

// ConsumerUsecase is an autogenerated mock type for the ConsumerUsecase type
type ConsumerUsecase struct {
	mock.Mock
}

// AddTenorLimit provides a mock function with given fields: ctx, requestData
func (_m *ConsumerUsecase) AddTenorLimit(ctx context.Context, requestData dto.AddTenorLmitRequestDto) (model.BaseResponseModel[dto.AddTenorLmitRequestDto], error) {
	ret := _m.Called(ctx, requestData)

	var r0 model.BaseResponseModel[dto.AddTenorLmitRequestDto]
	if rf, ok := ret.Get(0).(func(context.Context, dto.AddTenorLmitRequestDto) model.BaseResponseModel[dto.AddTenorLmitRequestDto]); ok {
		r0 = rf(ctx, requestData)
	} else {
		r0 = ret.Get(0).(model.BaseResponseModel[dto.AddTenorLmitRequestDto])
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, dto.AddTenorLmitRequestDto) error); ok {
		r1 = rf(ctx, requestData)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ApproveRequestLoan provides a mock function with given fields: ctx, requestData
func (_m *ConsumerUsecase) ApproveRequestLoan(ctx context.Context, requestData dto.ConsumerDto) (model.BaseResponseModel[dto.ApprovalResponseDataDto], error) {
	ret := _m.Called(ctx, requestData)

	var r0 model.BaseResponseModel[dto.ApprovalResponseDataDto]
	if rf, ok := ret.Get(0).(func(context.Context, dto.ConsumerDto) model.BaseResponseModel[dto.ApprovalResponseDataDto]); ok {
		r0 = rf(ctx, requestData)
	} else {
		r0 = ret.Get(0).(model.BaseResponseModel[dto.ApprovalResponseDataDto])
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, dto.ConsumerDto) error); ok {
		r1 = rf(ctx, requestData)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Register provides a mock function with given fields: ctx, requestData
func (_m *ConsumerUsecase) Register(ctx context.Context, requestData dto.RequestCreateNewConsumerDto) (model.BaseResponseModel[dto.ConsumerId], error) {
	ret := _m.Called(ctx, requestData)

	var r0 model.BaseResponseModel[dto.ConsumerId]
	if rf, ok := ret.Get(0).(func(context.Context, dto.RequestCreateNewConsumerDto) model.BaseResponseModel[dto.ConsumerId]); ok {
		r0 = rf(ctx, requestData)
	} else {
		r0 = ret.Get(0).(model.BaseResponseModel[dto.ConsumerId])
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, dto.RequestCreateNewConsumerDto) error); ok {
		r1 = rf(ctx, requestData)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RequestLoan provides a mock function with given fields: ctx, requestData
func (_m *ConsumerUsecase) RequestLoan(ctx context.Context, requestData dto.RequestLoanDto) (model.BaseResponseModel[dto.RequestLoanDto], error) {
	ret := _m.Called(ctx, requestData)

	var r0 model.BaseResponseModel[dto.RequestLoanDto]
	if rf, ok := ret.Get(0).(func(context.Context, dto.RequestLoanDto) model.BaseResponseModel[dto.RequestLoanDto]); ok {
		r0 = rf(ctx, requestData)
	} else {
		r0 = ret.Get(0).(model.BaseResponseModel[dto.RequestLoanDto])
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, dto.RequestLoanDto) error); ok {
		r1 = rf(ctx, requestData)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewConsumerUsecase interface {
	mock.TestingT
	Cleanup(func())
}

// NewConsumerUsecase creates a new instance of ConsumerUsecase. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewConsumerUsecase(t mockConstructorTestingTNewConsumerUsecase) *ConsumerUsecase {
	mock := &ConsumerUsecase{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
