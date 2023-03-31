// Code generated by MockGen. DO NOT EDIT.
// Source: client.go

// Package mock_lbs is a generated GoMock package.
package mock_lbs

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	loadbalancer "github.com/oracle/oci-go-sdk/v65/loadbalancer"
)

// MockLoadBalancerServiceClient is a mock of LoadBalancerServiceClient interface.
type MockLoadBalancerServiceClient struct {
	ctrl     *gomock.Controller
	recorder *MockLoadBalancerServiceClientMockRecorder
}

// MockLoadBalancerServiceClientMockRecorder is the mock recorder for MockLoadBalancerServiceClient.
type MockLoadBalancerServiceClientMockRecorder struct {
	mock *MockLoadBalancerServiceClient
}

// NewMockLoadBalancerServiceClient creates a new mock instance.
func NewMockLoadBalancerServiceClient(ctrl *gomock.Controller) *MockLoadBalancerServiceClient {
	mock := &MockLoadBalancerServiceClient{ctrl: ctrl}
	mock.recorder = &MockLoadBalancerServiceClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockLoadBalancerServiceClient) EXPECT() *MockLoadBalancerServiceClientMockRecorder {
	return m.recorder
}

// CreateBackend mocks base method.
func (m *MockLoadBalancerServiceClient) CreateBackend(ctx context.Context, request loadbalancer.CreateBackendRequest) (loadbalancer.CreateBackendResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateBackend", ctx, request)
	ret0, _ := ret[0].(loadbalancer.CreateBackendResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateBackend indicates an expected call of CreateBackend.
func (mr *MockLoadBalancerServiceClientMockRecorder) CreateBackend(ctx, request interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateBackend", reflect.TypeOf((*MockLoadBalancerServiceClient)(nil).CreateBackend), ctx, request)
}

// CreateLoadBalancer mocks base method.
func (m *MockLoadBalancerServiceClient) CreateLoadBalancer(ctx context.Context, request loadbalancer.CreateLoadBalancerRequest) (loadbalancer.CreateLoadBalancerResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateLoadBalancer", ctx, request)
	ret0, _ := ret[0].(loadbalancer.CreateLoadBalancerResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateLoadBalancer indicates an expected call of CreateLoadBalancer.
func (mr *MockLoadBalancerServiceClientMockRecorder) CreateLoadBalancer(ctx, request interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateLoadBalancer", reflect.TypeOf((*MockLoadBalancerServiceClient)(nil).CreateLoadBalancer), ctx, request)
}

// DeleteBackend mocks base method.
func (m *MockLoadBalancerServiceClient) DeleteBackend(ctx context.Context, request loadbalancer.DeleteBackendRequest) (loadbalancer.DeleteBackendResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteBackend", ctx, request)
	ret0, _ := ret[0].(loadbalancer.DeleteBackendResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteBackend indicates an expected call of DeleteBackend.
func (mr *MockLoadBalancerServiceClientMockRecorder) DeleteBackend(ctx, request interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteBackend", reflect.TypeOf((*MockLoadBalancerServiceClient)(nil).DeleteBackend), ctx, request)
}

// DeleteLoadBalancer mocks base method.
func (m *MockLoadBalancerServiceClient) DeleteLoadBalancer(ctx context.Context, request loadbalancer.DeleteLoadBalancerRequest) (loadbalancer.DeleteLoadBalancerResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteLoadBalancer", ctx, request)
	ret0, _ := ret[0].(loadbalancer.DeleteLoadBalancerResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteLoadBalancer indicates an expected call of DeleteLoadBalancer.
func (mr *MockLoadBalancerServiceClientMockRecorder) DeleteLoadBalancer(ctx, request interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteLoadBalancer", reflect.TypeOf((*MockLoadBalancerServiceClient)(nil).DeleteLoadBalancer), ctx, request)
}

// GetLoadBalancer mocks base method.
func (m *MockLoadBalancerServiceClient) GetLoadBalancer(ctx context.Context, request loadbalancer.GetLoadBalancerRequest) (loadbalancer.GetLoadBalancerResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetLoadBalancer", ctx, request)
	ret0, _ := ret[0].(loadbalancer.GetLoadBalancerResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetLoadBalancer indicates an expected call of GetLoadBalancer.
func (mr *MockLoadBalancerServiceClientMockRecorder) GetLoadBalancer(ctx, request interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLoadBalancer", reflect.TypeOf((*MockLoadBalancerServiceClient)(nil).GetLoadBalancer), ctx, request)
}

// GetWorkRequest mocks base method.
func (m *MockLoadBalancerServiceClient) GetWorkRequest(ctx context.Context, request loadbalancer.GetWorkRequestRequest) (loadbalancer.GetWorkRequestResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetWorkRequest", ctx, request)
	ret0, _ := ret[0].(loadbalancer.GetWorkRequestResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetWorkRequest indicates an expected call of GetWorkRequest.
func (mr *MockLoadBalancerServiceClientMockRecorder) GetWorkRequest(ctx, request interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetWorkRequest", reflect.TypeOf((*MockLoadBalancerServiceClient)(nil).GetWorkRequest), ctx, request)
}

// ListLoadBalancers mocks base method.
func (m *MockLoadBalancerServiceClient) ListLoadBalancers(ctx context.Context, request loadbalancer.ListLoadBalancersRequest) (loadbalancer.ListLoadBalancersResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListLoadBalancers", ctx, request)
	ret0, _ := ret[0].(loadbalancer.ListLoadBalancersResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListLoadBalancers indicates an expected call of ListLoadBalancers.
func (mr *MockLoadBalancerServiceClientMockRecorder) ListLoadBalancers(ctx, request interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListLoadBalancers", reflect.TypeOf((*MockLoadBalancerServiceClient)(nil).ListLoadBalancers), ctx, request)
}

// UpdateLoadBalancer mocks base method.
func (m *MockLoadBalancerServiceClient) UpdateLoadBalancer(ctx context.Context, request loadbalancer.UpdateLoadBalancerRequest) (loadbalancer.UpdateLoadBalancerResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateLoadBalancer", ctx, request)
	ret0, _ := ret[0].(loadbalancer.UpdateLoadBalancerResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateLoadBalancer indicates an expected call of UpdateLoadBalancer.
func (mr *MockLoadBalancerServiceClientMockRecorder) UpdateLoadBalancer(ctx, request interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateLoadBalancer", reflect.TypeOf((*MockLoadBalancerServiceClient)(nil).UpdateLoadBalancer), ctx, request)
}
