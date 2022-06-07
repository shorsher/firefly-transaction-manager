// Code generated by mockery v1.0.0. DO NOT EDIT.

package policyenginemocks

import (
	context "context"

	ffcapi "github.com/hyperledger/firefly-transaction-manager/pkg/ffcapi"
	mock "github.com/stretchr/testify/mock"

	policyengine "github.com/hyperledger/firefly-transaction-manager/pkg/policyengine"
)

// PolicyEngine is an autogenerated mock type for the PolicyEngine type
type PolicyEngine struct {
	mock.Mock
}

// Execute provides a mock function with given fields: ctx, cAPI, mtx
func (_m *PolicyEngine) Execute(ctx context.Context, cAPI ffcapi.API, mtx *policyengine.ManagedTXOutput) (bool, ffcapi.ErrorReason, error) {
	ret := _m.Called(ctx, cAPI, mtx)

	var r0 bool
	if rf, ok := ret.Get(0).(func(context.Context, ffcapi.API, *policyengine.ManagedTXOutput) bool); ok {
		r0 = rf(ctx, cAPI, mtx)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 ffcapi.ErrorReason
	if rf, ok := ret.Get(1).(func(context.Context, ffcapi.API, *policyengine.ManagedTXOutput) ffcapi.ErrorReason); ok {
		r1 = rf(ctx, cAPI, mtx)
	} else {
		r1 = ret.Get(1).(ffcapi.ErrorReason)
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(context.Context, ffcapi.API, *policyengine.ManagedTXOutput) error); ok {
		r2 = rf(ctx, cAPI, mtx)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}
