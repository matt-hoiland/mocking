// Code generated by mockery 2.9.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// Doer is an autogenerated mock type for the Doer type
type Doer struct {
	mock.Mock
}

// Do provides a mock function with given fields: s, i
func (_m *Doer) Do(s string, i int) error {
	ret := _m.Called(s, i)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, int) error); ok {
		r0 = rf(s, i)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
