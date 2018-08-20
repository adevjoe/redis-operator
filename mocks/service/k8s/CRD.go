// Code generated by mockery v1.0.0. DO NOT EDIT.
package mocks

import crd "github.com/spotahome/kooper/client/crd"

import mock "github.com/stretchr/testify/mock"

// CRD is an autogenerated mock type for the CRD type
type CRD struct {
	mock.Mock
}

// EnsureCRD provides a mock function with given fields: conf
func (_m *CRD) EnsureCRD(conf crd.Conf) error {
	ret := _m.Called(conf)

	var r0 error
	if rf, ok := ret.Get(0).(func(crd.Conf) error); ok {
		r0 = rf(conf)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
