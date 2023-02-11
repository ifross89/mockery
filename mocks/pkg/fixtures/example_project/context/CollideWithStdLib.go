// Code generated by mockery. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
)

// CollideWithStdLib is an autogenerated mock type for the CollideWithStdLib type
type CollideWithStdLib struct {
	mock.Mock
}

type CollideWithStdLib_Expecter struct {
	mock *mock.Mock
}

func (_m *CollideWithStdLib) EXPECT() *CollideWithStdLib_Expecter {
	return &CollideWithStdLib_Expecter{mock: &_m.Mock}
}

// NewClient provides a mock function with given fields: ctx
func (_m *CollideWithStdLib) NewClient(ctx context.Context) {
	_m.Called(ctx)
}

// CollideWithStdLib_NewClient_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'NewClient'
type CollideWithStdLib_NewClient_Call struct {
	*mock.Call
}

// NewClient is a helper method to define mock.On call
//   - ctx context.Context
func (_e *CollideWithStdLib_Expecter) NewClient(ctx interface{}) *CollideWithStdLib_NewClient_Call {
	return &CollideWithStdLib_NewClient_Call{Call: _e.mock.On("NewClient", ctx)}
}

func (_c *CollideWithStdLib_NewClient_Call) Run(run func(ctx context.Context)) *CollideWithStdLib_NewClient_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *CollideWithStdLib_NewClient_Call) Return() *CollideWithStdLib_NewClient_Call {
	_c.Call.Return()
	return _c
}

func (_c *CollideWithStdLib_NewClient_Call) RunAndReturn(run func(context.Context)) *CollideWithStdLib_NewClient_Call {
	_c.Call.Return(run)
	return _c
}

type mockConstructorTestingTNewCollideWithStdLib interface {
	mock.TestingT
	Cleanup(func())
}

// NewCollideWithStdLib creates a new instance of CollideWithStdLib. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewCollideWithStdLib(t mockConstructorTestingTNewCollideWithStdLib) *CollideWithStdLib {
	mock := &CollideWithStdLib{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
