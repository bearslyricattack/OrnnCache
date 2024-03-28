package mock

import (
	"context"
	"github.com/golang/mock/gomock"
	"reflect"
	"time"
)

// MockClient is a mock of Client cacheinterface.
type MockClient struct {
	ctrl     *gomock.Controller
	recorder *MockClientMockRecorder
}

// MockClientMockRecorder is the mock recorder for MockClient.
type MockClientMockRecorder struct {
	mock *MockClient
}

// NewMockClient creates a new mock instance.
func NewMockClient(ctrl *gomock.Controller) *MockClient {
	mock := &MockClient{ctrl: ctrl}
	mock.recorder = &MockClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockClient) EXPECT() *MockClientMockRecorder {
	return m.recorder
}

// Delete mocks base method.
func (m *MockClient) Delete(ctx context.Context, k string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Delete", ctx, k)
}

// Delete indicates an expected call of Delete.
func (mr *MockClientMockRecorder) Delete(ctx, k interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockClient)(nil).Delete), ctx, k)
}

// DeleteExpired mocks base method.
func (m *MockClient) DeleteExpired(ctx context.Context) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "DeleteExpired", ctx)
}

// DeleteExpired indicates an expected call of DeleteExpired.
func (mr *MockClientMockRecorder) DeleteExpired(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteExpired", reflect.TypeOf((*MockClient)(nil).DeleteExpired), ctx)
}

// Flush mocks base method.
func (m *MockClient) Flush(ctx context.Context) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Flush", ctx)
}

// Flush indicates an expected call of Flush.
func (mr *MockClientMockRecorder) Flush(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Flush", reflect.TypeOf((*MockClient)(nil).Flush), ctx)
}

// Get mocks base method.
func (m *MockClient) Get(ctx context.Context, k string) (interface{}, bool) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", ctx, k)
	ret0, _ := ret[0].(interface{})
	ret1, _ := ret[1].(bool)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockClientMockRecorder) Get(ctx, k interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockClient)(nil).Get), ctx, k)
}

// ItemCount mocks base method.
func (m *MockClient) ItemCount(ctx context.Context) int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ItemCount", ctx)
	ret0, _ := ret[0].(int)
	return ret0
}

// ItemCount indicates an expected call of ItemCount.
func (mr *MockClientMockRecorder) ItemCount(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ItemCount", reflect.TypeOf((*MockClient)(nil).ItemCount), ctx)
}

// Replace mocks base method.
func (m *MockClient) Replace(ctx context.Context, k string, x interface{}, d time.Duration) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Replace", ctx, k, x, d)
	ret0, _ := ret[0].(error)
	return ret0
}

// Replace indicates an expected call of Replace.
func (mr *MockClientMockRecorder) Replace(ctx, k, x, d interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Replace", reflect.TypeOf((*MockClient)(nil).Replace), ctx, k, x, d)
}

// Set mocks base method.
func (m *MockClient) Set(ctx context.Context, k string, x interface{}, d time.Duration) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Set", ctx, k, x, d)
}

// Set indicates an expected call of Set.
func (mr *MockClientMockRecorder) Set(ctx, k, x, d interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Set", reflect.TypeOf((*MockClient)(nil).Set), ctx, k, x, d)
}
