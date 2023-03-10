package abstraction

type Aggregate interface {
	GetAggregate() interface{}
	GetEvents() []AggregateEvent
}

type BaseAggregate struct {
	aggreagateRoot interface{}
	events         []AggregateEvent
}

func (o BaseAggregate) GetAggregate() interface{} {
	return o.aggreagateRoot
}

func (o *BaseAggregate) AddEvent(event AggregateEvent) {
	o.events = append(o.events, event)
}

func (o BaseAggregate) GetEvents() []AggregateEvent {
	return o.events
}
