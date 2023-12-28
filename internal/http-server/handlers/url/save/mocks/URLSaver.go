// Code generated by mockery v2.28.2. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// URLSaver is an autogenerated mock type for the URLSaver type
type URLSaver struct {
	mock.Mock
}

// SaveURL provides a mock function with given fields: urlToSave
func (_m *URLSaver) SaveURL(urlToSave string) (int64, error) {
	ret := _m.Called(urlToSave)

	var r0 int64
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (int64, error)); ok {
		return rf(urlToSave)
	}
	if rf, ok := ret.Get(0).(func(string) int64); ok {
		r0 = rf(urlToSave)
	} else {
		r0 = ret.Get(0).(int64)
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(urlToSave)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewURLSaver interface {
	mock.TestingT
	Cleanup(func())
}

// NewURLSaver creates a new instance of URLSaver. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewURLSaver(t mockConstructorTestingTNewURLSaver) *URLSaver {
	mock := &URLSaver{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
