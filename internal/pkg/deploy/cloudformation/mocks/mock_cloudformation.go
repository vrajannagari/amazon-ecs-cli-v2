// Code generated by MockGen. DO NOT EDIT.
// Source: ./internal/pkg/deploy/cloudformation/cloudformation.go

// Package mocks is a generated GoMock package.
package mocks

import (
	cloudformation "github.com/aws/amazon-ecs-cli-v2/internal/pkg/aws/cloudformation"
	cloudformation0 "github.com/aws/aws-sdk-go/service/cloudformation"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockstackConfiguration is a mock of stackConfiguration interface
type MockstackConfiguration struct {
	ctrl     *gomock.Controller
	recorder *MockstackConfigurationMockRecorder
}

// MockstackConfigurationMockRecorder is the mock recorder for MockstackConfiguration
type MockstackConfigurationMockRecorder struct {
	mock *MockstackConfiguration
}

// NewMockstackConfiguration creates a new mock instance
func NewMockstackConfiguration(ctrl *gomock.Controller) *MockstackConfiguration {
	mock := &MockstackConfiguration{ctrl: ctrl}
	mock.recorder = &MockstackConfigurationMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockstackConfiguration) EXPECT() *MockstackConfigurationMockRecorder {
	return m.recorder
}

// StackName mocks base method
func (m *MockstackConfiguration) StackName() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StackName")
	ret0, _ := ret[0].(string)
	return ret0
}

// StackName indicates an expected call of StackName
func (mr *MockstackConfigurationMockRecorder) StackName() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StackName", reflect.TypeOf((*MockstackConfiguration)(nil).StackName))
}

// Template mocks base method
func (m *MockstackConfiguration) Template() (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Template")
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Template indicates an expected call of Template
func (mr *MockstackConfigurationMockRecorder) Template() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Template", reflect.TypeOf((*MockstackConfiguration)(nil).Template))
}

// Parameters mocks base method
func (m *MockstackConfiguration) Parameters() []*cloudformation0.Parameter {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Parameters")
	ret0, _ := ret[0].([]*cloudformation0.Parameter)
	return ret0
}

// Parameters indicates an expected call of Parameters
func (mr *MockstackConfigurationMockRecorder) Parameters() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Parameters", reflect.TypeOf((*MockstackConfiguration)(nil).Parameters))
}

// Tags mocks base method
func (m *MockstackConfiguration) Tags() []*cloudformation0.Tag {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Tags")
	ret0, _ := ret[0].([]*cloudformation0.Tag)
	return ret0
}

// Tags indicates an expected call of Tags
func (mr *MockstackConfigurationMockRecorder) Tags() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Tags", reflect.TypeOf((*MockstackConfiguration)(nil).Tags))
}

// MockcfnClient is a mock of cfnClient interface
type MockcfnClient struct {
	ctrl     *gomock.Controller
	recorder *MockcfnClientMockRecorder
}

// MockcfnClientMockRecorder is the mock recorder for MockcfnClient
type MockcfnClientMockRecorder struct {
	mock *MockcfnClient
}

// NewMockcfnClient creates a new mock instance
func NewMockcfnClient(ctrl *gomock.Controller) *MockcfnClient {
	mock := &MockcfnClient{ctrl: ctrl}
	mock.recorder = &MockcfnClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockcfnClient) EXPECT() *MockcfnClientMockRecorder {
	return m.recorder
}

// Create mocks base method
func (m *MockcfnClient) Create(arg0 *cloudformation.Stack) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create
func (mr *MockcfnClientMockRecorder) Create(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockcfnClient)(nil).Create), arg0)
}

// CreateAndWait mocks base method
func (m *MockcfnClient) CreateAndWait(arg0 *cloudformation.Stack) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateAndWait", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateAndWait indicates an expected call of CreateAndWait
func (mr *MockcfnClientMockRecorder) CreateAndWait(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateAndWait", reflect.TypeOf((*MockcfnClient)(nil).CreateAndWait), arg0)
}

// WaitForCreate mocks base method
func (m *MockcfnClient) WaitForCreate(stackName string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WaitForCreate", stackName)
	ret0, _ := ret[0].(error)
	return ret0
}

// WaitForCreate indicates an expected call of WaitForCreate
func (mr *MockcfnClientMockRecorder) WaitForCreate(stackName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WaitForCreate", reflect.TypeOf((*MockcfnClient)(nil).WaitForCreate), stackName)
}

// Update mocks base method
func (m *MockcfnClient) Update(arg0 *cloudformation.Stack) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update
func (mr *MockcfnClientMockRecorder) Update(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockcfnClient)(nil).Update), arg0)
}

// UpdateAndWait mocks base method
func (m *MockcfnClient) UpdateAndWait(arg0 *cloudformation.Stack) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateAndWait", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateAndWait indicates an expected call of UpdateAndWait
func (mr *MockcfnClientMockRecorder) UpdateAndWait(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateAndWait", reflect.TypeOf((*MockcfnClient)(nil).UpdateAndWait), arg0)
}

// Delete mocks base method
func (m *MockcfnClient) Delete(stackName string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", stackName)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete
func (mr *MockcfnClientMockRecorder) Delete(stackName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockcfnClient)(nil).Delete), stackName)
}

// DeleteAndWait mocks base method
func (m *MockcfnClient) DeleteAndWait(stackName string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteAndWait", stackName)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteAndWait indicates an expected call of DeleteAndWait
func (mr *MockcfnClientMockRecorder) DeleteAndWait(stackName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteAndWait", reflect.TypeOf((*MockcfnClient)(nil).DeleteAndWait), stackName)
}

// Describe mocks base method
func (m *MockcfnClient) Describe(stackName string) (*cloudformation.StackDescription, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Describe", stackName)
	ret0, _ := ret[0].(*cloudformation.StackDescription)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Describe indicates an expected call of Describe
func (mr *MockcfnClientMockRecorder) Describe(stackName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Describe", reflect.TypeOf((*MockcfnClient)(nil).Describe), stackName)
}

// Events mocks base method
func (m *MockcfnClient) Events(stackName string) ([]cloudformation.StackEvent, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Events", stackName)
	ret0, _ := ret[0].([]cloudformation.StackEvent)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Events indicates an expected call of Events
func (mr *MockcfnClientMockRecorder) Events(stackName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Events", reflect.TypeOf((*MockcfnClient)(nil).Events), stackName)
}
