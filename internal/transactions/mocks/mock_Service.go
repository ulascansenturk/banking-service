// Code generated by mockery. DO NOT EDIT.

package mocks

import (
	context "context"
	constants "ulascansenturk/service/internal/constants"

	gorm "gorm.io/gorm"

	mock "github.com/stretchr/testify/mock"

	time "time"

	transactions "ulascansenturk/service/internal/transactions"

	uuid "github.com/google/uuid"
)

// MockService is an autogenerated mock type for the Service type
type MockService struct {
	mock.Mock
}

// BeginTransaction provides a mock function with given fields: ctx
func (_m *MockService) BeginTransaction(ctx context.Context) (*gorm.DB, error) {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for BeginTransaction")
	}

	var r0 *gorm.DB
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) (*gorm.DB, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) *gorm.DB); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*gorm.DB)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteTransaction provides a mock function with given fields: ctx, id
func (_m *MockService) DeleteTransaction(ctx context.Context, id uuid.UUID) error {
	ret := _m.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for DeleteTransaction")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID) error); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetTransactionByID provides a mock function with given fields: ctx, id
func (_m *MockService) GetTransactionByID(ctx context.Context, id uuid.UUID) (*transactions.Transaction, error) {
	ret := _m.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for GetTransactionByID")
	}

	var r0 *transactions.Transaction
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID) (*transactions.Transaction, error)); ok {
		return rf(ctx, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID) *transactions.Transaction); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*transactions.Transaction)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, uuid.UUID) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetTransactionByReferenceID provides a mock function with given fields: ctx, referenceID
func (_m *MockService) GetTransactionByReferenceID(ctx context.Context, referenceID uuid.UUID) (*transactions.Transaction, error) {
	ret := _m.Called(ctx, referenceID)

	if len(ret) == 0 {
		panic("no return value specified for GetTransactionByReferenceID")
	}

	var r0 *transactions.Transaction
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID) (*transactions.Transaction, error)); ok {
		return rf(ctx, referenceID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID) *transactions.Transaction); ok {
		r0 = rf(ctx, referenceID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*transactions.Transaction)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, uuid.UUID) error); ok {
		r1 = rf(ctx, referenceID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetTransactionsByCreatedAt provides a mock function with given fields: ctx, createdAt
func (_m *MockService) GetTransactionsByCreatedAt(ctx context.Context, createdAt time.Time) ([]*transactions.Transaction, error) {
	ret := _m.Called(ctx, createdAt)

	if len(ret) == 0 {
		panic("no return value specified for GetTransactionsByCreatedAt")
	}

	var r0 []*transactions.Transaction
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, time.Time) ([]*transactions.Transaction, error)); ok {
		return rf(ctx, createdAt)
	}
	if rf, ok := ret.Get(0).(func(context.Context, time.Time) []*transactions.Transaction); ok {
		r0 = rf(ctx, createdAt)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*transactions.Transaction)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, time.Time) error); ok {
		r1 = rf(ctx, createdAt)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetTransactionsByFromAccountID provides a mock function with given fields: ctx, fromAccountID
func (_m *MockService) GetTransactionsByFromAccountID(ctx context.Context, fromAccountID uuid.UUID) ([]*transactions.Transaction, error) {
	ret := _m.Called(ctx, fromAccountID)

	if len(ret) == 0 {
		panic("no return value specified for GetTransactionsByFromAccountID")
	}

	var r0 []*transactions.Transaction
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID) ([]*transactions.Transaction, error)); ok {
		return rf(ctx, fromAccountID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID) []*transactions.Transaction); ok {
		r0 = rf(ctx, fromAccountID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*transactions.Transaction)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, uuid.UUID) error); ok {
		r1 = rf(ctx, fromAccountID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetTransactionsByToAccountID provides a mock function with given fields: ctx, toAccountID
func (_m *MockService) GetTransactionsByToAccountID(ctx context.Context, toAccountID uuid.UUID) ([]*transactions.Transaction, error) {
	ret := _m.Called(ctx, toAccountID)

	if len(ret) == 0 {
		panic("no return value specified for GetTransactionsByToAccountID")
	}

	var r0 []*transactions.Transaction
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID) ([]*transactions.Transaction, error)); ok {
		return rf(ctx, toAccountID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID) []*transactions.Transaction); ok {
		r0 = rf(ctx, toAccountID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*transactions.Transaction)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, uuid.UUID) error); ok {
		r1 = rf(ctx, toAccountID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateTransaction provides a mock function with given fields: ctx, transaction, tx
func (_m *MockService) UpdateTransaction(ctx context.Context, transaction *transactions.Transaction, tx *gorm.DB) error {
	ret := _m.Called(ctx, transaction, tx)

	if len(ret) == 0 {
		panic("no return value specified for UpdateTransaction")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *transactions.Transaction, *gorm.DB) error); ok {
		r0 = rf(ctx, transaction, tx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateTransactionStatus provides a mock function with given fields: ctx, id, status
func (_m *MockService) UpdateTransactionStatus(ctx context.Context, id uuid.UUID, status constants.TransactionStatus) (*transactions.Transaction, error) {
	ret := _m.Called(ctx, id, status)

	if len(ret) == 0 {
		panic("no return value specified for UpdateTransactionStatus")
	}

	var r0 *transactions.Transaction
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID, constants.TransactionStatus) (*transactions.Transaction, error)); ok {
		return rf(ctx, id, status)
	}
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID, constants.TransactionStatus) *transactions.Transaction); ok {
		r0 = rf(ctx, id, status)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*transactions.Transaction)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, uuid.UUID, constants.TransactionStatus) error); ok {
		r1 = rf(ctx, id, status)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewMockService creates a new instance of MockService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockService(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockService {
	mock := &MockService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
