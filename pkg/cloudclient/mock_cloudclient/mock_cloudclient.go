// Code generated by MockGen. DO NOT EDIT.
// Source: pkg/cloudclient/cloudclient.go

// Package mock_cloudclient is a generated GoMock package.
package mock_cloudclient

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockCloudClient is a mock of CloudClient interface.
type MockCloudClient struct {
	ctrl     *gomock.Controller
	recorder *MockCloudClientMockRecorder
}

// MockCloudClientMockRecorder is the mock recorder for MockCloudClient.
type MockCloudClientMockRecorder struct {
	mock *MockCloudClient
}

// NewMockCloudClient creates a new mock instance.
func NewMockCloudClient(ctrl *gomock.Controller) *MockCloudClient {
	mock := &MockCloudClient{ctrl: ctrl}
	mock.recorder = &MockCloudClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCloudClient) EXPECT() *MockCloudClientMockRecorder {
	return m.recorder
}

// ByoVPCValidator mocks base method.
func (m *MockCloudClient) ByoVPCValidator(ctx context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ByoVPCValidator", ctx)
	ret0, _ := ret[0].(error)
	return ret0
}

// ByoVPCValidator indicates an expected call of ByoVPCValidator.
func (mr *MockCloudClientMockRecorder) ByoVPCValidator(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ByoVPCValidator", reflect.TypeOf((*MockCloudClient)(nil).ByoVPCValidator), ctx)
}

// ValidateEgress mocks base method.
func (m *MockCloudClient) ValidateEgress(ctx context.Context, vpcSubnetID, cloudImageID string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ValidateEgress", ctx, vpcSubnetID, cloudImageID)
	ret0, _ := ret[0].(error)
	return ret0
}

// ValidateEgress indicates an expected call of ValidateEgress.
func (mr *MockCloudClientMockRecorder) ValidateEgress(ctx, vpcSubnetID, cloudImageID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ValidateEgress", reflect.TypeOf((*MockCloudClient)(nil).ValidateEgress), ctx, vpcSubnetID, cloudImageID)
}
