// Code generated by MockGen. DO NOT EDIT.
// Source: interface.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	adapter "github.com/hiroaki-yamamoto/reusable-services/adapter"
	reflect "reflect"
)

// MockIAdapter is a mock of IAdapter interface
type MockIAdapter struct {
	ctrl     *gomock.Controller
	recorder *MockIAdapterMockRecorder
}

// MockIAdapterMockRecorder is the mock recorder for MockIAdapter
type MockIAdapterMockRecorder struct {
	mock *MockIAdapter
}

// NewMockIAdapter creates a new mock instance
func NewMockIAdapter(ctrl *gomock.Controller) *MockIAdapter {
	mock := &MockIAdapter{ctrl: ctrl}
	mock.recorder = &MockIAdapterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockIAdapter) EXPECT() *MockIAdapterMockRecorder {
	return m.recorder
}

// Count mocks base method
func (m *MockIAdapter) Count(ctx context.Context, query interface{}) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Count", ctx, query)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Count indicates an expected call of Count
func (mr *MockIAdapterMockRecorder) Count(ctx, query interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Count", reflect.TypeOf((*MockIAdapter)(nil).Count), ctx, query)
}

// Find mocks base method
func (m *MockIAdapter) Find(ctx context.Context, query, docs interface{}, opts ...interface{}) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, query, docs}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Find", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// Find indicates an expected call of Find
func (mr *MockIAdapterMockRecorder) Find(ctx, query, docs interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, query, docs}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Find", reflect.TypeOf((*MockIAdapter)(nil).Find), varargs...)
}

// FindOne mocks base method
func (m *MockIAdapter) FindOne(ctx context.Context, query, doc interface{}, opts ...interface{}) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, query, doc}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "FindOne", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// FindOne indicates an expected call of FindOne
func (mr *MockIAdapterMockRecorder) FindOne(ctx, query, doc interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, query, doc}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindOne", reflect.TypeOf((*MockIAdapter)(nil).FindOne), varargs...)
}

// Insert mocks base method
func (m *MockIAdapter) Insert(ctx context.Context, doc interface{}) (interface{}, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Insert", ctx, doc)
	ret0, _ := ret[0].(interface{})
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Insert indicates an expected call of Insert
func (mr *MockIAdapterMockRecorder) Insert(ctx, doc interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Insert", reflect.TypeOf((*MockIAdapter)(nil).Insert), ctx, doc)
}

// InsertMany mocks base method
func (m *MockIAdapter) InsertMany(ctx context.Context, docs []interface{}) ([]interface{}, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InsertMany", ctx, docs)
	ret0, _ := ret[0].([]interface{})
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// InsertMany indicates an expected call of InsertMany
func (mr *MockIAdapterMockRecorder) InsertMany(ctx, docs interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InsertMany", reflect.TypeOf((*MockIAdapter)(nil).InsertMany), ctx, docs)
}

// Update mocks base method
func (m *MockIAdapter) Update(ctx context.Context, query, update interface{}) (*adapter.UpdateSummary, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", ctx, query, update)
	ret0, _ := ret[0].(*adapter.UpdateSummary)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update
func (mr *MockIAdapterMockRecorder) Update(ctx, query, update interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockIAdapter)(nil).Update), ctx, query, update)
}

// UpdateMany mocks base method
func (m *MockIAdapter) UpdateMany(ctx context.Context, query, update interface{}) (*adapter.UpdateSummary, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateMany", ctx, query, update)
	ret0, _ := ret[0].(*adapter.UpdateSummary)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateMany indicates an expected call of UpdateMany
func (mr *MockIAdapterMockRecorder) UpdateMany(ctx, query, update interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateMany", reflect.TypeOf((*MockIAdapter)(nil).UpdateMany), ctx, query, update)
}

// Delete mocks base method
func (m *MockIAdapter) Delete(ctx context.Context, filter interface{}) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", ctx, filter)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Delete indicates an expected call of Delete
func (mr *MockIAdapterMockRecorder) Delete(ctx, filter interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockIAdapter)(nil).Delete), ctx, filter)
}

// DeleteMany mocks base method
func (m *MockIAdapter) DeleteMany(ctx context.Context, filter interface{}) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteMany", ctx, filter)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteMany indicates an expected call of DeleteMany
func (mr *MockIAdapterMockRecorder) DeleteMany(ctx, filter interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteMany", reflect.TypeOf((*MockIAdapter)(nil).DeleteMany), ctx, filter)
}
