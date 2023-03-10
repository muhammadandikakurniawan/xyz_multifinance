package event

import (
	"github.com/muhammadandikakurniawan/xyz_multifinance/src/module/consumer/entity"
	"github.com/muhammadandikakurniawan/xyz_multifinance/src/shared/abstraction"
)

func NewInsertNewConsumerEvent(data *entity.ConsumerEntity) InsertNewConsumerEvent {
	event := InsertNewConsumerEvent{abstraction.NewBaseAggregateEvent(data)}
	return event
}

type InsertNewConsumerEvent struct {
	abstraction.BaseAggregateEvent[*entity.ConsumerEntity]
}
