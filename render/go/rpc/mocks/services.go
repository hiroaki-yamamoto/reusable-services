// Code generated by MockGen. DO NOT EDIT.
// Source: ../go/rpc/services.pb.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	rpc "github.com/hiroaki-yamamoto/reusable-services/render/go/rpc"
	grpc "google.golang.org/grpc"
	reflect "reflect"
)

// MockTemplateServiceClient is a mock of TemplateServiceClient interface
type MockTemplateServiceClient struct {
	ctrl     *gomock.Controller
	recorder *MockTemplateServiceClientMockRecorder
}

// MockTemplateServiceClientMockRecorder is the mock recorder for MockTemplateServiceClient
type MockTemplateServiceClientMockRecorder struct {
	mock *MockTemplateServiceClient
}

// NewMockTemplateServiceClient creates a new mock instance
func NewMockTemplateServiceClient(ctrl *gomock.Controller) *MockTemplateServiceClient {
	mock := &MockTemplateServiceClient{ctrl: ctrl}
	mock.recorder = &MockTemplateServiceClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockTemplateServiceClient) EXPECT() *MockTemplateServiceClientMockRecorder {
	return m.recorder
}

// Render mocks base method
func (m *MockTemplateServiceClient) Render(ctx context.Context, in *rpc.RenderingRequest, opts ...grpc.CallOption) (*rpc.RenderingResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Render", varargs...)
	ret0, _ := ret[0].(*rpc.RenderingResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Render indicates an expected call of Render
func (mr *MockTemplateServiceClientMockRecorder) Render(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Render", reflect.TypeOf((*MockTemplateServiceClient)(nil).Render), varargs...)
}

// MockTemplateServiceServer is a mock of TemplateServiceServer interface
type MockTemplateServiceServer struct {
	ctrl     *gomock.Controller
	recorder *MockTemplateServiceServerMockRecorder
}

// MockTemplateServiceServerMockRecorder is the mock recorder for MockTemplateServiceServer
type MockTemplateServiceServerMockRecorder struct {
	mock *MockTemplateServiceServer
}

// NewMockTemplateServiceServer creates a new mock instance
func NewMockTemplateServiceServer(ctrl *gomock.Controller) *MockTemplateServiceServer {
	mock := &MockTemplateServiceServer{ctrl: ctrl}
	mock.recorder = &MockTemplateServiceServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockTemplateServiceServer) EXPECT() *MockTemplateServiceServerMockRecorder {
	return m.recorder
}

// Render mocks base method
func (m *MockTemplateServiceServer) Render(arg0 context.Context, arg1 *rpc.RenderingRequest) (*rpc.RenderingResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Render", arg0, arg1)
	ret0, _ := ret[0].(*rpc.RenderingResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Render indicates an expected call of Render
func (mr *MockTemplateServiceServerMockRecorder) Render(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Render", reflect.TypeOf((*MockTemplateServiceServer)(nil).Render), arg0, arg1)
}
