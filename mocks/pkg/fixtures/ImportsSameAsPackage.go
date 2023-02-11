// Code generated by mockery. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"
	fixtures "github.com/vektra/mockery/v2/pkg/fixtures"

	test "github.com/vektra/mockery/v2/pkg/fixtures/test"
)

// ImportsSameAsPackage is an autogenerated mock type for the ImportsSameAsPackage type
type ImportsSameAsPackage struct {
	mock.Mock
}

type ImportsSameAsPackage_Expecter struct {
	mock *mock.Mock
}

func (_m *ImportsSameAsPackage) EXPECT() *ImportsSameAsPackage_Expecter {
	return &ImportsSameAsPackage_Expecter{mock: &_m.Mock}
}

// A provides a mock function with given fields:
func (_m *ImportsSameAsPackage) A() test.B {
	ret := _m.Called()

	var r0 test.B
	if rf, ok := ret.Get(0).(func() test.B); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(test.B)
	}

	return r0
}

// ImportsSameAsPackage_A_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'A'
type ImportsSameAsPackage_A_Call struct {
	*mock.Call
}

// A is a helper method to define mock.On call
func (_e *ImportsSameAsPackage_Expecter) A() *ImportsSameAsPackage_A_Call {
	return &ImportsSameAsPackage_A_Call{Call: _e.mock.On("A")}
}

func (_c *ImportsSameAsPackage_A_Call) Run(run func()) *ImportsSameAsPackage_A_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *ImportsSameAsPackage_A_Call) Return(_a0 test.B) *ImportsSameAsPackage_A_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *ImportsSameAsPackage_A_Call) RunAndReturn(run func() test.B) *ImportsSameAsPackage_A_Call {
	_c.Call.Return(run)
	return _c
}

// B provides a mock function with given fields:
func (_m *ImportsSameAsPackage) B() fixtures.KeyManager {
	ret := _m.Called()

	var r0 fixtures.KeyManager
	if rf, ok := ret.Get(0).(func() fixtures.KeyManager); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(fixtures.KeyManager)
		}
	}

	return r0
}

// ImportsSameAsPackage_B_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'B'
type ImportsSameAsPackage_B_Call struct {
	*mock.Call
}

// B is a helper method to define mock.On call
func (_e *ImportsSameAsPackage_Expecter) B() *ImportsSameAsPackage_B_Call {
	return &ImportsSameAsPackage_B_Call{Call: _e.mock.On("B")}
}

func (_c *ImportsSameAsPackage_B_Call) Run(run func()) *ImportsSameAsPackage_B_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *ImportsSameAsPackage_B_Call) Return(_a0 fixtures.KeyManager) *ImportsSameAsPackage_B_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *ImportsSameAsPackage_B_Call) RunAndReturn(run func() fixtures.KeyManager) *ImportsSameAsPackage_B_Call {
	_c.Call.Return(run)
	return _c
}

// C provides a mock function with given fields: _a0
func (_m *ImportsSameAsPackage) C(_a0 fixtures.C) {
	_m.Called(_a0)
}

// ImportsSameAsPackage_C_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'C'
type ImportsSameAsPackage_C_Call struct {
	*mock.Call
}

// C is a helper method to define mock.On call
//   - _a0 fixtures.C
func (_e *ImportsSameAsPackage_Expecter) C(_a0 interface{}) *ImportsSameAsPackage_C_Call {
	return &ImportsSameAsPackage_C_Call{Call: _e.mock.On("C", _a0)}
}

func (_c *ImportsSameAsPackage_C_Call) Run(run func(_a0 fixtures.C)) *ImportsSameAsPackage_C_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(fixtures.C))
	})
	return _c
}

func (_c *ImportsSameAsPackage_C_Call) Return() *ImportsSameAsPackage_C_Call {
	_c.Call.Return()
	return _c
}

func (_c *ImportsSameAsPackage_C_Call) RunAndReturn(run func(fixtures.C)) *ImportsSameAsPackage_C_Call {
	_c.Call.Return(run)
	return _c
}

type mockConstructorTestingTNewImportsSameAsPackage interface {
	mock.TestingT
	Cleanup(func())
}

// NewImportsSameAsPackage creates a new instance of ImportsSameAsPackage. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewImportsSameAsPackage(t mockConstructorTestingTNewImportsSameAsPackage) *ImportsSameAsPackage {
	mock := &ImportsSameAsPackage{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
