// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/libopenstorage/openstorage/schedpolicy (interfaces: SchedulePolicy)

// Package mock is a generated GoMock package.
package mock

import (
	gomock "github.com/golang/mock/gomock"
	schedpolicy "github.com/libopenstorage/openstorage/schedpolicy"
	reflect "reflect"
)

// MockSchedulePolicy is a mock of SchedulePolicy interface
type MockSchedulePolicy struct {
	ctrl     *gomock.Controller
	recorder *MockSchedulePolicyMockRecorder
}

// MockSchedulePolicyMockRecorder is the mock recorder for MockSchedulePolicy
type MockSchedulePolicyMockRecorder struct {
	mock *MockSchedulePolicy
}

// NewMockSchedulePolicy creates a new mock instance
func NewMockSchedulePolicy(ctrl *gomock.Controller) *MockSchedulePolicy {
	mock := &MockSchedulePolicy{ctrl: ctrl}
	mock.recorder = &MockSchedulePolicyMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockSchedulePolicy) EXPECT() *MockSchedulePolicyMockRecorder {
	return m.recorder
}

// SchedPolicyCreate mocks base method
func (m *MockSchedulePolicy) SchedPolicyCreate(arg0, arg1 string) error {
	ret := m.ctrl.Call(m, "SchedPolicyCreate", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// SchedPolicyCreate indicates an expected call of SchedPolicyCreate
func (mr *MockSchedulePolicyMockRecorder) SchedPolicyCreate(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SchedPolicyCreate", reflect.TypeOf((*MockSchedulePolicy)(nil).SchedPolicyCreate), arg0, arg1)
}

// SchedPolicyDelete mocks base method
func (m *MockSchedulePolicy) SchedPolicyDelete(arg0 string) error {
	ret := m.ctrl.Call(m, "SchedPolicyDelete", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SchedPolicyDelete indicates an expected call of SchedPolicyDelete
func (mr *MockSchedulePolicyMockRecorder) SchedPolicyDelete(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SchedPolicyDelete", reflect.TypeOf((*MockSchedulePolicy)(nil).SchedPolicyDelete), arg0)
}

// SchedPolicyEnumerate mocks base method
func (m *MockSchedulePolicy) SchedPolicyEnumerate() ([]*schedpolicy.SchedPolicy, error) {
	ret := m.ctrl.Call(m, "SchedPolicyEnumerate")
	ret0, _ := ret[0].([]*schedpolicy.SchedPolicy)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SchedPolicyEnumerate indicates an expected call of SchedPolicyEnumerate
func (mr *MockSchedulePolicyMockRecorder) SchedPolicyEnumerate() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SchedPolicyEnumerate", reflect.TypeOf((*MockSchedulePolicy)(nil).SchedPolicyEnumerate))
}

// SchedPolicyGet mocks base method
func (m *MockSchedulePolicy) SchedPolicyGet(arg0 string) (*schedpolicy.SchedPolicy, error) {
	ret := m.ctrl.Call(m, "SchedPolicyGet", arg0)
	ret0, _ := ret[0].(*schedpolicy.SchedPolicy)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SchedPolicyGet indicates an expected call of SchedPolicyGet
func (mr *MockSchedulePolicyMockRecorder) SchedPolicyGet(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SchedPolicyGet", reflect.TypeOf((*MockSchedulePolicy)(nil).SchedPolicyGet), arg0)
}

// SchedPolicyUpdate mocks base method
func (m *MockSchedulePolicy) SchedPolicyUpdate(arg0, arg1 string) error {
	ret := m.ctrl.Call(m, "SchedPolicyUpdate", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// SchedPolicyUpdate indicates an expected call of SchedPolicyUpdate
func (mr *MockSchedulePolicyMockRecorder) SchedPolicyUpdate(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SchedPolicyUpdate", reflect.TypeOf((*MockSchedulePolicy)(nil).SchedPolicyUpdate), arg0, arg1)
}
