// Code generated by mockery v2.15.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// PkgCrypto is an autogenerated mock type for the PkgCrypto type
type PkgCrypto struct {
	mock.Mock
}

// Decrypt provides a mock function with given fields: encryptedTxt
func (_m *PkgCrypto) Decrypt(encryptedTxt []byte) ([]byte, error) {
	ret := _m.Called(encryptedTxt)

	var r0 []byte
	if rf, ok := ret.Get(0).(func([]byte) []byte); ok {
		r0 = rf(encryptedTxt)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func([]byte) error); ok {
		r1 = rf(encryptedTxt)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Encrypt provides a mock function with given fields: plainTxt
func (_m *PkgCrypto) Encrypt(plainTxt []byte) ([]byte, error) {
	ret := _m.Called(plainTxt)

	var r0 []byte
	if rf, ok := ret.Get(0).(func([]byte) []byte); ok {
		r0 = rf(plainTxt)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func([]byte) error); ok {
		r1 = rf(plainTxt)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewPkgCrypto interface {
	mock.TestingT
	Cleanup(func())
}

// NewPkgCrypto creates a new instance of PkgCrypto. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewPkgCrypto(t mockConstructorTestingTNewPkgCrypto) *PkgCrypto {
	mock := &PkgCrypto{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}