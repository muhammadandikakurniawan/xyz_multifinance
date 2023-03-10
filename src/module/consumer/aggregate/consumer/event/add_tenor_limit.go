package event

import (
	"github.com/muhammadandikakurniawan/xyz_multifinance/src/module/consumer/entity"
	"github.com/muhammadandikakurniawan/xyz_multifinance/src/shared/abstraction"
)

func NewAddTenorLimitEvent(data []*entity.TenorLimitEntity) AddTenorLimitEvent {
	event := AddTenorLimitEvent{abstraction.NewBaseAggregateEvent(data)}
	return event
}

type AddTenorLimitEvent struct {
	abstraction.BaseAggregateEvent[[]*entity.TenorLimitEntity]
}

func NewUpdateTenorLimitEvent(data []*entity.TenorLimitEntity) UpdateTenorLimitEvent {
	event := UpdateTenorLimitEvent{abstraction.NewBaseAggregateEvent(data)}
	return event
}

type UpdateTenorLimitEvent struct {
	abstraction.BaseAggregateEvent[[]*entity.TenorLimitEntity]
}

func NewDeleteTenorLimitEvent(data []*entity.TenorLimitEntity) DeleteTenorLimitEvent {
	event := DeleteTenorLimitEvent{abstraction.NewBaseAggregateEvent(data)}
	return event
}

type DeleteTenorLimitEvent struct {
	abstraction.BaseAggregateEvent[[]*entity.TenorLimitEntity]
}
