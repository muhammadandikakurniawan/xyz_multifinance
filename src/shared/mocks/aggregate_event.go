// Code generated by mockery v2.15.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// AggregateEvent is an autogenerated mock type for the AggregateEvent type
type AggregateEvent struct {
	mock.Mock
}

// GetEventData provides a mock function with given fields:
func (_m *AggregateEvent) GetEventData() interface{} {
	ret := _m.Called()

	var r0 interface{}
	if rf, ok := ret.Get(0).(func() interface{}); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(interface{})
		}
	}

	return r0
}

type mockConstructorTestingTNewAggregateEvent interface {
	mock.TestingT
	Cleanup(func())
}

// NewAggregateEvent creates a new instance of AggregateEvent. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewAggregateEvent(t mockConstructorTestingTNewAggregateEvent) *AggregateEvent {
	mock := &AggregateEvent{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
