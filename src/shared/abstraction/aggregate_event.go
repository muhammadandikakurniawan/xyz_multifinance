package abstraction

type AggregateEvent interface {
	GetEventData() interface{}
}

func NewBaseAggregateEvent[T any](data T) BaseAggregateEvent[T] {
	return BaseAggregateEvent[T]{
		eventData: data,
	}
}

type BaseAggregateEvent[T any] struct {
	eventData T
}

func (o BaseAggregateEvent[any]) GetEventData() interface{} {
	return o.eventData
}

func (c *BaseAggregateEvent[T]) GetData() T {
	return c.eventData
}

func (c *BaseAggregateEvent[T]) SetData(data T) {
	c.eventData = data
}
