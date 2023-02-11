// Code generated by mockery. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"
	foo "github.com/vektra/mockery/v2/pkg/fixtures/example_project/foo"
)

// Foo is an autogenerated mock type for the Foo type
type Foo struct {
	mock.Mock
}

type Foo_Expecter struct {
	mock *mock.Mock
}

func (_m *Foo) EXPECT() *Foo_Expecter {
	return &Foo_Expecter{mock: &_m.Mock}
}

// DoFoo provides a mock function with given fields:
func (_m *Foo) DoFoo() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// Foo_DoFoo_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DoFoo'
type Foo_DoFoo_Call struct {
	*mock.Call
}

// DoFoo is a helper method to define mock.On call
func (_e *Foo_Expecter) DoFoo() *Foo_DoFoo_Call {
	return &Foo_DoFoo_Call{Call: _e.mock.On("DoFoo")}
}

func (_c *Foo_DoFoo_Call) Run(run func()) *Foo_DoFoo_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *Foo_DoFoo_Call) Return(_a0 string) *Foo_DoFoo_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Foo_DoFoo_Call) RunAndReturn(run func() string) *Foo_DoFoo_Call {
	_c.Call.Return(run)
	return _c
}

// GetBaz provides a mock function with given fields:
func (_m *Foo) GetBaz() (*foo.Baz, error) {
	ret := _m.Called()

	var r0 *foo.Baz
	var r1 error
	if rf, ok := ret.Get(0).(func() (*foo.Baz, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() *foo.Baz); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*foo.Baz)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Foo_GetBaz_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetBaz'
type Foo_GetBaz_Call struct {
	*mock.Call
}

// GetBaz is a helper method to define mock.On call
func (_e *Foo_Expecter) GetBaz() *Foo_GetBaz_Call {
	return &Foo_GetBaz_Call{Call: _e.mock.On("GetBaz")}
}

func (_c *Foo_GetBaz_Call) Run(run func()) *Foo_GetBaz_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *Foo_GetBaz_Call) Return(_a0 *foo.Baz, _a1 error) *Foo_GetBaz_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Foo_GetBaz_Call) RunAndReturn(run func() (*foo.Baz, error)) *Foo_GetBaz_Call {
	_c.Call.Return(run)
	return _c
}

type mockConstructorTestingTNewFoo interface {
	mock.TestingT
	Cleanup(func())
}

// NewFoo creates a new instance of Foo. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewFoo(t mockConstructorTestingTNewFoo) *Foo {
	mock := &Foo{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
