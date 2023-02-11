// Code generated by mockery. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// Stringer is an autogenerated mock type for the Stringer type
type Stringer struct {
	mock.Mock
}

type Stringer_Expecter struct {
	mock *mock.Mock
}

func (_m *Stringer) EXPECT() *Stringer_Expecter {
	return &Stringer_Expecter{mock: &_m.Mock}
}

// String provides a mock function with given fields:
func (_m *Stringer) String() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// Stringer_String_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'String'
type Stringer_String_Call struct {
	*mock.Call
}

// String is a helper method to define mock.On call
func (_e *Stringer_Expecter) String() *Stringer_String_Call {
	return &Stringer_String_Call{Call: _e.mock.On("String")}
}

func (_c *Stringer_String_Call) Run(run func()) *Stringer_String_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *Stringer_String_Call) Return(_a0 string) *Stringer_String_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Stringer_String_Call) RunAndReturn(run func() string) *Stringer_String_Call {
	_c.Call.Return(run)
	return _c
}

type mockConstructorTestingTNewStringer interface {
	mock.TestingT
	Cleanup(func())
}

// NewStringer creates a new instance of Stringer. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewStringer(t mockConstructorTestingTNewStringer) *Stringer {
	mock := &Stringer{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
