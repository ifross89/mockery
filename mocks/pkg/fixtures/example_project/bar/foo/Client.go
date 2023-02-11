// Code generated by mockery. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// Client is an autogenerated mock type for the Client type
type Client struct {
	mock.Mock
}

type Client_Expecter struct {
	mock *mock.Mock
}

func (_m *Client) EXPECT() *Client_Expecter {
	return &Client_Expecter{mock: &_m.Mock}
}

// Search provides a mock function with given fields: query
func (_m *Client) Search(query string) ([]string, error) {
	ret := _m.Called(query)

	var r0 []string
	var r1 error
	if rf, ok := ret.Get(0).(func(string) ([]string, error)); ok {
		return rf(query)
	}
	if rf, ok := ret.Get(0).(func(string) []string); ok {
		r0 = rf(query)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(query)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Client_Search_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Search'
type Client_Search_Call struct {
	*mock.Call
}

// Search is a helper method to define mock.On call
//   - query string
func (_e *Client_Expecter) Search(query interface{}) *Client_Search_Call {
	return &Client_Search_Call{Call: _e.mock.On("Search", query)}
}

func (_c *Client_Search_Call) Run(run func(query string)) *Client_Search_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *Client_Search_Call) Return(_a0 []string, _a1 error) *Client_Search_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Client_Search_Call) RunAndReturn(run func(string) ([]string, error)) *Client_Search_Call {
	_c.Call.Return(run)
	return _c
}

type mockConstructorTestingTNewClient interface {
	mock.TestingT
	Cleanup(func())
}

// NewClient creates a new instance of Client. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewClient(t mockConstructorTestingTNewClient) *Client {
	mock := &Client{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
