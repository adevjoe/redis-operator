// Code generated by mockery v1.0.0. DO NOT EDIT.
package mocks

import mock "github.com/stretchr/testify/mock"

import v1alpha2 "github.com/spotahome/redis-operator/api/redisfailover/v1alpha2"

// RedisFailoverHeal is an autogenerated mock type for the RedisFailoverHeal type
type RedisFailoverHeal struct {
	mock.Mock
}

// NewSentinelMonitor provides a mock function with given fields: ip, monitor, rFailover
func (_m *RedisFailoverHeal) NewSentinelMonitor(ip string, monitor string, rFailover *v1alpha2.RedisFailover) error {
	ret := _m.Called(ip, monitor, rFailover)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string, *v1alpha2.RedisFailover) error); ok {
		r0 = rf(ip, monitor, rFailover)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// RestoreSentinel provides a mock function with given fields: ip
func (_m *RedisFailoverHeal) RestoreSentinel(ip string) error {
	ret := _m.Called(ip)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(ip)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SetMasterOnAll provides a mock function with given fields: masterIP, rFailover
func (_m *RedisFailoverHeal) SetMasterOnAll(masterIP string, rFailover *v1alpha2.RedisFailover) error {
	ret := _m.Called(masterIP, rFailover)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, *v1alpha2.RedisFailover) error); ok {
		r0 = rf(masterIP, rFailover)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SetRandomMaster provides a mock function with given fields: rFailover
func (_m *RedisFailoverHeal) SetRandomMaster(rFailover *v1alpha2.RedisFailover) error {
	ret := _m.Called(rFailover)

	var r0 error
	if rf, ok := ret.Get(0).(func(*v1alpha2.RedisFailover) error); ok {
		r0 = rf(rFailover)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SetSentinelCustomConfig provides a mock function with given fields: ip, rFailover
func (_m *RedisFailoverHeal) SetSentinelCustomConfig(ip string, rFailover *v1alpha2.RedisFailover) error {
	ret := _m.Called(ip, rFailover)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, *v1alpha2.RedisFailover) error); ok {
		r0 = rf(ip, rFailover)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
