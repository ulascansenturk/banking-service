// Code generated by mockery. DO NOT EDIT.

package mocks

import (
	time "time"

	mock "github.com/stretchr/testify/mock"
)

// MockTimeProvider is an autogenerated mock type for the TimeProvider type
type MockTimeProvider struct {
	mock.Mock
}

// Now provides a mock function with given fields:
func (_m *MockTimeProvider) Now() time.Time {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Now")
	}

	var r0 time.Time
	if rf, ok := ret.Get(0).(func() time.Time); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(time.Time)
	}

	return r0
}

// NewMockTimeProvider creates a new instance of MockTimeProvider. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockTimeProvider(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockTimeProvider {
	mock := &MockTimeProvider{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
