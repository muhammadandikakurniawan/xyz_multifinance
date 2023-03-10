package event

import (
	"github.com/muhammadandikakurniawan/xyz_multifinance/src/module/consumer/entity"
	"github.com/muhammadandikakurniawan/xyz_multifinance/src/shared/abstraction"
)

func NewRequestLoanEvent(data *entity.RequestLoanEntity) RequestLoanEvent {
	event := RequestLoanEvent{abstraction.NewBaseAggregateEvent(data)}
	return event
}

type RequestLoanEvent struct {
	abstraction.BaseAggregateEvent[*entity.RequestLoanEntity]
}
