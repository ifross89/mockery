// Code generated by mockery. DO NOT EDIT.

package mocks

import (
	io "io"

	constraints "github.com/vektra/mockery/v2/pkg/fixtures/constraints"

	mock "github.com/stretchr/testify/mock"

	test "github.com/vektra/mockery/v2/pkg/fixtures"
)

// RequesterGenerics is an autogenerated mock type for the RequesterGenerics type
type RequesterGenerics[TAny interface{}, TComparable comparable, TSigned constraints.Signed, TIntf test.GetInt, TExternalIntf io.Writer, TGenIntf test.GetGeneric[TSigned], TInlineType interface{ ~int | ~uint }, TInlineTypeGeneric interface {
	~int | test.GenericType[int, test.GetInt]
	comparable
}] struct {
	mock.Mock
}

// GenericAnonymousStructs provides a mock function with given fields: _a0
func (_m *RequesterGenerics[TAny, TComparable, TSigned, TIntf, TExternalIntf, TGenIntf, TInlineType, TInlineTypeGeneric]) GenericAnonymousStructs(_a0 struct{ Type1 TExternalIntf }) struct {
	Type2 test.GenericType[string, test.EmbeddedGet[int]]
} {
	ret := _m.Called(_a0)

	var r0 struct {
		Type2 test.GenericType[string, test.EmbeddedGet[int]]
	}
	if rf, ok := ret.Get(0).(func(struct{ Type1 TExternalIntf }) struct {
		Type2 test.GenericType[string, test.EmbeddedGet[int]]
	}); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(struct {
			Type2 test.GenericType[string, test.EmbeddedGet[int]]
		})
	}

	return r0
}

// GenericArguments provides a mock function with given fields: _a0, _a1
func (_m *RequesterGenerics[TAny, TComparable, TSigned, TIntf, TExternalIntf, TGenIntf, TInlineType, TInlineTypeGeneric]) GenericArguments(_a0 TAny, _a1 TComparable) (TSigned, TIntf) {
	ret := _m.Called(_a0, _a1)

	var r0 TSigned
	if rf, ok := ret.Get(0).(func(TAny, TComparable) TSigned); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Get(0).(TSigned)
	}

	var r1 TIntf
	if rf, ok := ret.Get(1).(func(TAny, TComparable) TIntf); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Get(1).(TIntf)
	}

	return r0, r1
}

// GenericStructs provides a mock function with given fields: _a0
func (_m *RequesterGenerics[TAny, TComparable, TSigned, TIntf, TExternalIntf, TGenIntf, TInlineType, TInlineTypeGeneric]) GenericStructs(_a0 test.GenericType[TAny, TIntf]) test.GenericType[TSigned, TIntf] {
	ret := _m.Called(_a0)

	var r0 test.GenericType[TSigned, TIntf]
	if rf, ok := ret.Get(0).(func(test.GenericType[TAny, TIntf]) test.GenericType[TSigned, TIntf]); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(test.GenericType[TSigned, TIntf])
	}

	return r0
}

type NewRequesterGenericsT interface {
	mock.TestingT
	Cleanup(func())
}

// NewRequesterGenerics creates a new instance of RequesterGenerics. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewRequesterGenerics[TAny interface{}, TComparable comparable, TSigned constraints.Signed, TIntf test.GetInt, TExternalIntf io.Writer, TGenIntf test.GetGeneric[TSigned], TInlineType interface{ ~int | ~uint }, TInlineTypeGeneric interface {
	~int | test.GenericType[int, test.GetInt]
	comparable
}](t NewRequesterGenericsT) *RequesterGenerics[TAny, TComparable, TSigned, TIntf, TExternalIntf, TGenIntf, TInlineType, TInlineTypeGeneric] {
	mock := &RequesterGenerics[TAny, TComparable, TSigned, TIntf, TExternalIntf, TGenIntf, TInlineType, TInlineTypeGeneric]{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
