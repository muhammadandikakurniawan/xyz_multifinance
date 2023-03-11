package event

import (
	"github.com/muhammadandikakurniawan/xyz_multifinance/src/module/consumer/entity"
	"github.com/muhammadandikakurniawan/xyz_multifinance/src/shared/abstraction"
)

func NewApproveRequestLoanEvent(data *entity.RequestLoanEntity) ApproveRequestLoanEvent {
	event := ApproveRequestLoanEvent{abstraction.NewBaseAggregateEvent(data)}
	return event
}

type ApproveRequestLoanEvent struct {
	abstraction.BaseAggregateEvent[*entity.RequestLoanEntity]
}
