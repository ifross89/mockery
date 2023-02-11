// Code generated by mockery. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// Sibling is an autogenerated mock type for the Sibling type
type Sibling struct {
	mock.Mock
}

type Sibling_Expecter struct {
	mock *mock.Mock
}

func (_m *Sibling) EXPECT() *Sibling_Expecter {
	return &Sibling_Expecter{mock: &_m.Mock}
}

// DoSomething provides a mock function with given fields:
func (_m *Sibling) DoSomething() {
	_m.Called()
}

// Sibling_DoSomething_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DoSomething'
type Sibling_DoSomething_Call struct {
	*mock.Call
}

// DoSomething is a helper method to define mock.On call
func (_e *Sibling_Expecter) DoSomething() *Sibling_DoSomething_Call {
	return &Sibling_DoSomething_Call{Call: _e.mock.On("DoSomething")}
}

func (_c *Sibling_DoSomething_Call) Run(run func()) *Sibling_DoSomething_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *Sibling_DoSomething_Call) Return() *Sibling_DoSomething_Call {
	_c.Call.Return()
	return _c
}

func (_c *Sibling_DoSomething_Call) RunAndReturn(run func()) *Sibling_DoSomething_Call {
	_c.Call.Return(run)
	return _c
}

type mockConstructorTestingTNewSibling interface {
	mock.TestingT
	Cleanup(func())
}

// NewSibling creates a new instance of Sibling. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewSibling(t mockConstructorTestingTNewSibling) *Sibling {
	mock := &Sibling{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
