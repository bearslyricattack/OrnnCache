package mock

import (
	"context"
	"github.com/go-redis/redis/v8"
	"github.com/golang/mock/gomock"
	"reflect"
	"time"
)

// MockUniversalClient is a mock of UniversalClient cacheinterface.
type MockUniversalClient struct {
	ctrl     *gomock.Controller
	recorder *MockUniversalClientMockRecorder
}

// MockUniversalClientMockRecorder is the mock recorder for MockUniversalClient.
type MockUniversalClientMockRecorder struct {
	mock *MockUniversalClient
}

// NewMockUniversalClient creates a new mock instance.
func NewMockUniversalClient(ctrl *gomock.Controller) *MockUniversalClient {
	mock := &MockUniversalClient{ctrl: ctrl}
	mock.recorder = &MockUniversalClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUniversalClient) EXPECT() *MockUniversalClientMockRecorder {
	return m.recorder
}

// AddHook mocks base method.
func (m *MockUniversalClient) AddHook(arg0 redis.Hook) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "AddHook", arg0)
}

// AddHook indicates an expected call of AddHook.
func (mr *MockUniversalClientMockRecorder) AddHook(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddHook", reflect.TypeOf((*MockUniversalClient)(nil).AddHook), arg0)
}

// Append mocks base method.
func (m *MockUniversalClient) Append(ctx context.Context, key, value string) *redis.IntCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Append", ctx, key, value)
	ret0, _ := ret[0].(*redis.IntCmd)
	return ret0
}

// Append indicates an expected call of Append.
func (mr *MockUniversalClientMockRecorder) Append(ctx, key, value interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Append", reflect.TypeOf((*MockUniversalClient)(nil).Append), ctx, key, value)
}

// BLMove mocks base method.
func (m *MockUniversalClient) BLMove(ctx context.Context, source, destination, srcpos, destpos string, timeout time.Duration) *redis.StringCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BLMove", ctx, source, destination, srcpos, destpos, timeout)
	ret0, _ := ret[0].(*redis.StringCmd)
	return ret0
}

// BLMove indicates an expected call of BLMove.
func (mr *MockUniversalClientMockRecorder) BLMove(ctx, source, destination, srcpos, destpos, timeout interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BLMove", reflect.TypeOf((*MockUniversalClient)(nil).BLMove), ctx, source, destination, srcpos, destpos, timeout)
}

// BLPop mocks base method.
func (m *MockUniversalClient) BLPop(ctx context.Context, timeout time.Duration, keys ...string) *redis.StringSliceCmd {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, timeout}
	for _, a := range keys {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "BLPop", varargs...)
	ret0, _ := ret[0].(*redis.StringSliceCmd)
	return ret0
}

// BLPop indicates an expected call of BLPop.
func (mr *MockUniversalClientMockRecorder) BLPop(ctx, timeout interface{}, keys ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, timeout}, keys...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BLPop", reflect.TypeOf((*MockUniversalClient)(nil).BLPop), varargs...)
}

// BRPop mocks base method.
func (m *MockUniversalClient) BRPop(ctx context.Context, timeout time.Duration, keys ...string) *redis.StringSliceCmd {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, timeout}
	for _, a := range keys {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "BRPop", varargs...)
	ret0, _ := ret[0].(*redis.StringSliceCmd)
	return ret0
}

// BRPop indicates an expected call of BRPop.
func (mr *MockUniversalClientMockRecorder) BRPop(ctx, timeout interface{}, keys ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, timeout}, keys...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BRPop", reflect.TypeOf((*MockUniversalClient)(nil).BRPop), varargs...)
}

// BRPopLPush mocks base method.
func (m *MockUniversalClient) BRPopLPush(ctx context.Context, source, destination string, timeout time.Duration) *redis.StringCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BRPopLPush", ctx, source, destination, timeout)
	ret0, _ := ret[0].(*redis.StringCmd)
	return ret0
}

// BRPopLPush indicates an expected call of BRPopLPush.
func (mr *MockUniversalClientMockRecorder) BRPopLPush(ctx, source, destination, timeout interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BRPopLPush", reflect.TypeOf((*MockUniversalClient)(nil).BRPopLPush), ctx, source, destination, timeout)
}

// BZPopMax mocks base method.
func (m *MockUniversalClient) BZPopMax(ctx context.Context, timeout time.Duration, keys ...string) *redis.ZWithKeyCmd {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, timeout}
	for _, a := range keys {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "BZPopMax", varargs...)
	ret0, _ := ret[0].(*redis.ZWithKeyCmd)
	return ret0
}

// BZPopMax indicates an expected call of BZPopMax.
func (mr *MockUniversalClientMockRecorder) BZPopMax(ctx, timeout interface{}, keys ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, timeout}, keys...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BZPopMax", reflect.TypeOf((*MockUniversalClient)(nil).BZPopMax), varargs...)
}

// BZPopMin mocks base method.
func (m *MockUniversalClient) BZPopMin(ctx context.Context, timeout time.Duration, keys ...string) *redis.ZWithKeyCmd {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, timeout}
	for _, a := range keys {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "BZPopMin", varargs...)
	ret0, _ := ret[0].(*redis.ZWithKeyCmd)
	return ret0
}

// BZPopMin indicates an expected call of BZPopMin.
func (mr *MockUniversalClientMockRecorder) BZPopMin(ctx, timeout interface{}, keys ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, timeout}, keys...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BZPopMin", reflect.TypeOf((*MockUniversalClient)(nil).BZPopMin), varargs...)
}

// BgRewriteAOF mocks base method.
func (m *MockUniversalClient) BgRewriteAOF(ctx context.Context) *redis.StatusCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BgRewriteAOF", ctx)
	ret0, _ := ret[0].(*redis.StatusCmd)
	return ret0
}

// BgRewriteAOF indicates an expected call of BgRewriteAOF.
func (mr *MockUniversalClientMockRecorder) BgRewriteAOF(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BgRewriteAOF", reflect.TypeOf((*MockUniversalClient)(nil).BgRewriteAOF), ctx)
}

// BgSave mocks base method.
func (m *MockUniversalClient) BgSave(ctx context.Context) *redis.StatusCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BgSave", ctx)
	ret0, _ := ret[0].(*redis.StatusCmd)
	return ret0
}

// BgSave indicates an expected call of BgSave.
func (mr *MockUniversalClientMockRecorder) BgSave(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BgSave", reflect.TypeOf((*MockUniversalClient)(nil).BgSave), ctx)
}

// BitCount mocks base method.
func (m *MockUniversalClient) BitCount(ctx context.Context, key string, bitCount *redis.BitCount) *redis.IntCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BitCount", ctx, key, bitCount)
	ret0, _ := ret[0].(*redis.IntCmd)
	return ret0
}

// BitCount indicates an expected call of BitCount.
func (mr *MockUniversalClientMockRecorder) BitCount(ctx, key, bitCount interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BitCount", reflect.TypeOf((*MockUniversalClient)(nil).BitCount), ctx, key, bitCount)
}

// BitField mocks base method.
func (m *MockUniversalClient) BitField(ctx context.Context, key string, args ...interface{}) *redis.IntSliceCmd {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, key}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "BitField", varargs...)
	ret0, _ := ret[0].(*redis.IntSliceCmd)
	return ret0
}

// BitField indicates an expected call of BitField.
func (mr *MockUniversalClientMockRecorder) BitField(ctx, key interface{}, args ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, key}, args...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BitField", reflect.TypeOf((*MockUniversalClient)(nil).BitField), varargs...)
}

// BitOpAnd mocks base method.
func (m *MockUniversalClient) BitOpAnd(ctx context.Context, destKey string, keys ...string) *redis.IntCmd {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, destKey}
	for _, a := range keys {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "BitOpAnd", varargs...)
	ret0, _ := ret[0].(*redis.IntCmd)
	return ret0
}

// BitOpAnd indicates an expected call of BitOpAnd.
func (mr *MockUniversalClientMockRecorder) BitOpAnd(ctx, destKey interface{}, keys ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, destKey}, keys...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BitOpAnd", reflect.TypeOf((*MockUniversalClient)(nil).BitOpAnd), varargs...)
}

// BitOpNot mocks base method.
func (m *MockUniversalClient) BitOpNot(ctx context.Context, destKey, key string) *redis.IntCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BitOpNot", ctx, destKey, key)
	ret0, _ := ret[0].(*redis.IntCmd)
	return ret0
}

// BitOpNot indicates an expected call of BitOpNot.
func (mr *MockUniversalClientMockRecorder) BitOpNot(ctx, destKey, key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BitOpNot", reflect.TypeOf((*MockUniversalClient)(nil).BitOpNot), ctx, destKey, key)
}

// BitOpOr mocks base method.
func (m *MockUniversalClient) BitOpOr(ctx context.Context, destKey string, keys ...string) *redis.IntCmd {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, destKey}
	for _, a := range keys {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "BitOpOr", varargs...)
	ret0, _ := ret[0].(*redis.IntCmd)
	return ret0
}

// BitOpOr indicates an expected call of BitOpOr.
func (mr *MockUniversalClientMockRecorder) BitOpOr(ctx, destKey interface{}, keys ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, destKey}, keys...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BitOpOr", reflect.TypeOf((*MockUniversalClient)(nil).BitOpOr), varargs...)
}

// BitOpXor mocks base method.
func (m *MockUniversalClient) BitOpXor(ctx context.Context, destKey string, keys ...string) *redis.IntCmd {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, destKey}
	for _, a := range keys {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "BitOpXor", varargs...)
	ret0, _ := ret[0].(*redis.IntCmd)
	return ret0
}

// BitOpXor indicates an expected call of BitOpXor.
func (mr *MockUniversalClientMockRecorder) BitOpXor(ctx, destKey interface{}, keys ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, destKey}, keys...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BitOpXor", reflect.TypeOf((*MockUniversalClient)(nil).BitOpXor), varargs...)
}

// BitPos mocks base method.
func (m *MockUniversalClient) BitPos(ctx context.Context, key string, bit int64, pos ...int64) *redis.IntCmd {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, key, bit}
	for _, a := range pos {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "BitPos", varargs...)
	ret0, _ := ret[0].(*redis.IntCmd)
	return ret0
}

// BitPos indicates an expected call of BitPos.
func (mr *MockUniversalClientMockRecorder) BitPos(ctx, key, bit interface{}, pos ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, key, bit}, pos...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BitPos", reflect.TypeOf((*MockUniversalClient)(nil).BitPos), varargs...)
}

// ClientGetName mocks base method.
func (m *MockUniversalClient) ClientGetName(ctx context.Context) *redis.StringCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ClientGetName", ctx)
	ret0, _ := ret[0].(*redis.StringCmd)
	return ret0
}

// ClientGetName indicates an expected call of ClientGetName.
func (mr *MockUniversalClientMockRecorder) ClientGetName(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ClientGetName", reflect.TypeOf((*MockUniversalClient)(nil).ClientGetName), ctx)
}

// ClientID mocks base method.
func (m *MockUniversalClient) ClientID(ctx context.Context) *redis.IntCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ClientID", ctx)
	ret0, _ := ret[0].(*redis.IntCmd)
	return ret0
}

// ClientID indicates an expected call of ClientID.
func (mr *MockUniversalClientMockRecorder) ClientID(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ClientID", reflect.TypeOf((*MockUniversalClient)(nil).ClientID), ctx)
}

// ClientKill mocks base method.
func (m *MockUniversalClient) ClientKill(ctx context.Context, ipPort string) *redis.StatusCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ClientKill", ctx, ipPort)
	ret0, _ := ret[0].(*redis.StatusCmd)
	return ret0
}

// ClientKill indicates an expected call of ClientKill.
func (mr *MockUniversalClientMockRecorder) ClientKill(ctx, ipPort interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ClientKill", reflect.TypeOf((*MockUniversalClient)(nil).ClientKill), ctx, ipPort)
}

// ClientKillByFilter mocks base method.
func (m *MockUniversalClient) ClientKillByFilter(ctx context.Context, keys ...string) *redis.IntCmd {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx}
	for _, a := range keys {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ClientKillByFilter", varargs...)
	ret0, _ := ret[0].(*redis.IntCmd)
	return ret0
}

// ClientKillByFilter indicates an expected call of ClientKillByFilter.
func (mr *MockUniversalClientMockRecorder) ClientKillByFilter(ctx interface{}, keys ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx}, keys...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ClientKillByFilter", reflect.TypeOf((*MockUniversalClient)(nil).ClientKillByFilter), varargs...)
}

// ClientList mocks base method.
func (m *MockUniversalClient) ClientList(ctx context.Context) *redis.StringCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ClientList", ctx)
	ret0, _ := ret[0].(*redis.StringCmd)
	return ret0
}

// ClientList indicates an expected call of ClientList.
func (mr *MockUniversalClientMockRecorder) ClientList(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ClientList", reflect.TypeOf((*MockUniversalClient)(nil).ClientList), ctx)
}

// ClientPause mocks base method.
func (m *MockUniversalClient) ClientPause(ctx context.Context, dur time.Duration) *redis.BoolCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ClientPause", ctx, dur)
	ret0, _ := ret[0].(*redis.BoolCmd)
	return ret0
}

// ClientPause indicates an expected call of ClientPause.
func (mr *MockUniversalClientMockRecorder) ClientPause(ctx, dur interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ClientPause", reflect.TypeOf((*MockUniversalClient)(nil).ClientPause), ctx, dur)
}

// Close mocks base method.
func (m *MockUniversalClient) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close.
func (mr *MockUniversalClientMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockUniversalClient)(nil).Close))
}

// ClusterAddSlots mocks base method.
func (m *MockUniversalClient) ClusterAddSlots(ctx context.Context, slots ...int) *redis.StatusCmd {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx}
	for _, a := range slots {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ClusterAddSlots", varargs...)
	ret0, _ := ret[0].(*redis.StatusCmd)
	return ret0
}

// ClusterAddSlots indicates an expected call of ClusterAddSlots.
func (mr *MockUniversalClientMockRecorder) ClusterAddSlots(ctx interface{}, slots ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx}, slots...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ClusterAddSlots", reflect.TypeOf((*MockUniversalClient)(nil).ClusterAddSlots), varargs...)
}

// ClusterAddSlotsRange mocks base method.
func (m *MockUniversalClient) ClusterAddSlotsRange(ctx context.Context, min, max int) *redis.StatusCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ClusterAddSlotsRange", ctx, min, max)
	ret0, _ := ret[0].(*redis.StatusCmd)
	return ret0
}

// ClusterAddSlotsRange indicates an expected call of ClusterAddSlotsRange.
func (mr *MockUniversalClientMockRecorder) ClusterAddSlotsRange(ctx, min, max interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ClusterAddSlotsRange", reflect.TypeOf((*MockUniversalClient)(nil).ClusterAddSlotsRange), ctx, min, max)
}

// ClusterCountFailureReports mocks base method.
func (m *MockUniversalClient) ClusterCountFailureReports(ctx context.Context, nodeID string) *redis.IntCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ClusterCountFailureReports", ctx, nodeID)
	ret0, _ := ret[0].(*redis.IntCmd)
	return ret0
}

// ClusterCountFailureReports indicates an expected call of ClusterCountFailureReports.
func (mr *MockUniversalClientMockRecorder) ClusterCountFailureReports(ctx, nodeID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ClusterCountFailureReports", reflect.TypeOf((*MockUniversalClient)(nil).ClusterCountFailureReports), ctx, nodeID)
}

// ClusterCountKeysInSlot mocks base method.
func (m *MockUniversalClient) ClusterCountKeysInSlot(ctx context.Context, slot int) *redis.IntCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ClusterCountKeysInSlot", ctx, slot)
	ret0, _ := ret[0].(*redis.IntCmd)
	return ret0
}

// ClusterCountKeysInSlot indicates an expected call of ClusterCountKeysInSlot.
func (mr *MockUniversalClientMockRecorder) ClusterCountKeysInSlot(ctx, slot interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ClusterCountKeysInSlot", reflect.TypeOf((*MockUniversalClient)(nil).ClusterCountKeysInSlot), ctx, slot)
}

// ClusterDelSlots mocks base method.
func (m *MockUniversalClient) ClusterDelSlots(ctx context.Context, slots ...int) *redis.StatusCmd {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx}
	for _, a := range slots {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ClusterDelSlots", varargs...)
	ret0, _ := ret[0].(*redis.StatusCmd)
	return ret0
}

// ClusterDelSlots indicates an expected call of ClusterDelSlots.
func (mr *MockUniversalClientMockRecorder) ClusterDelSlots(ctx interface{}, slots ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx}, slots...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ClusterDelSlots", reflect.TypeOf((*MockUniversalClient)(nil).ClusterDelSlots), varargs...)
}

// ClusterDelSlotsRange mocks base method.
func (m *MockUniversalClient) ClusterDelSlotsRange(ctx context.Context, min, max int) *redis.StatusCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ClusterDelSlotsRange", ctx, min, max)
	ret0, _ := ret[0].(*redis.StatusCmd)
	return ret0
}

// ClusterDelSlotsRange indicates an expected call of ClusterDelSlotsRange.
func (mr *MockUniversalClientMockRecorder) ClusterDelSlotsRange(ctx, min, max interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ClusterDelSlotsRange", reflect.TypeOf((*MockUniversalClient)(nil).ClusterDelSlotsRange), ctx, min, max)
}

// ClusterFailover mocks base method.
func (m *MockUniversalClient) ClusterFailover(ctx context.Context) *redis.StatusCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ClusterFailover", ctx)
	ret0, _ := ret[0].(*redis.StatusCmd)
	return ret0
}

// ClusterFailover indicates an expected call of ClusterFailover.
func (mr *MockUniversalClientMockRecorder) ClusterFailover(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ClusterFailover", reflect.TypeOf((*MockUniversalClient)(nil).ClusterFailover), ctx)
}

// ClusterForget mocks base method.
func (m *MockUniversalClient) ClusterForget(ctx context.Context, nodeID string) *redis.StatusCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ClusterForget", ctx, nodeID)
	ret0, _ := ret[0].(*redis.StatusCmd)
	return ret0
}

// ClusterForget indicates an expected call of ClusterForget.
func (mr *MockUniversalClientMockRecorder) ClusterForget(ctx, nodeID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ClusterForget", reflect.TypeOf((*MockUniversalClient)(nil).ClusterForget), ctx, nodeID)
}

// ClusterGetKeysInSlot mocks base method.
func (m *MockUniversalClient) ClusterGetKeysInSlot(ctx context.Context, slot, count int) *redis.StringSliceCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ClusterGetKeysInSlot", ctx, slot, count)
	ret0, _ := ret[0].(*redis.StringSliceCmd)
	return ret0
}

// ClusterGetKeysInSlot indicates an expected call of ClusterGetKeysInSlot.
func (mr *MockUniversalClientMockRecorder) ClusterGetKeysInSlot(ctx, slot, count interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ClusterGetKeysInSlot", reflect.TypeOf((*MockUniversalClient)(nil).ClusterGetKeysInSlot), ctx, slot, count)
}

// ClusterInfo mocks base method.
func (m *MockUniversalClient) ClusterInfo(ctx context.Context) *redis.StringCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ClusterInfo", ctx)
	ret0, _ := ret[0].(*redis.StringCmd)
	return ret0
}

// ClusterInfo indicates an expected call of ClusterInfo.
func (mr *MockUniversalClientMockRecorder) ClusterInfo(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ClusterInfo", reflect.TypeOf((*MockUniversalClient)(nil).ClusterInfo), ctx)
}

// ClusterKeySlot mocks base method.
func (m *MockUniversalClient) ClusterKeySlot(ctx context.Context, key string) *redis.IntCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ClusterKeySlot", ctx, key)
	ret0, _ := ret[0].(*redis.IntCmd)
	return ret0
}

// ClusterKeySlot indicates an expected call of ClusterKeySlot.
func (mr *MockUniversalClientMockRecorder) ClusterKeySlot(ctx, key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ClusterKeySlot", reflect.TypeOf((*MockUniversalClient)(nil).ClusterKeySlot), ctx, key)
}

// ClusterMeet mocks base method.
func (m *MockUniversalClient) ClusterMeet(ctx context.Context, host, port string) *redis.StatusCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ClusterMeet", ctx, host, port)
	ret0, _ := ret[0].(*redis.StatusCmd)
	return ret0
}

// ClusterMeet indicates an expected call of ClusterMeet.
func (mr *MockUniversalClientMockRecorder) ClusterMeet(ctx, host, port interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ClusterMeet", reflect.TypeOf((*MockUniversalClient)(nil).ClusterMeet), ctx, host, port)
}

// ClusterNodes mocks base method.
func (m *MockUniversalClient) ClusterNodes(ctx context.Context) *redis.StringCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ClusterNodes", ctx)
	ret0, _ := ret[0].(*redis.StringCmd)
	return ret0
}

// ClusterNodes indicates an expected call of ClusterNodes.
func (mr *MockUniversalClientMockRecorder) ClusterNodes(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ClusterNodes", reflect.TypeOf((*MockUniversalClient)(nil).ClusterNodes), ctx)
}

// ClusterReplicate mocks base method.
func (m *MockUniversalClient) ClusterReplicate(ctx context.Context, nodeID string) *redis.StatusCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ClusterReplicate", ctx, nodeID)
	ret0, _ := ret[0].(*redis.StatusCmd)
	return ret0
}

// ClusterReplicate indicates an expected call of ClusterReplicate.
func (mr *MockUniversalClientMockRecorder) ClusterReplicate(ctx, nodeID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ClusterReplicate", reflect.TypeOf((*MockUniversalClient)(nil).ClusterReplicate), ctx, nodeID)
}

// ClusterResetHard mocks base method.
func (m *MockUniversalClient) ClusterResetHard(ctx context.Context) *redis.StatusCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ClusterResetHard", ctx)
	ret0, _ := ret[0].(*redis.StatusCmd)
	return ret0
}

// ClusterResetHard indicates an expected call of ClusterResetHard.
func (mr *MockUniversalClientMockRecorder) ClusterResetHard(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ClusterResetHard", reflect.TypeOf((*MockUniversalClient)(nil).ClusterResetHard), ctx)
}

// ClusterResetSoft mocks base method.
func (m *MockUniversalClient) ClusterResetSoft(ctx context.Context) *redis.StatusCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ClusterResetSoft", ctx)
	ret0, _ := ret[0].(*redis.StatusCmd)
	return ret0
}

// ClusterResetSoft indicates an expected call of ClusterResetSoft.
func (mr *MockUniversalClientMockRecorder) ClusterResetSoft(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ClusterResetSoft", reflect.TypeOf((*MockUniversalClient)(nil).ClusterResetSoft), ctx)
}

// ClusterSaveConfig mocks base method.
func (m *MockUniversalClient) ClusterSaveConfig(ctx context.Context) *redis.StatusCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ClusterSaveConfig", ctx)
	ret0, _ := ret[0].(*redis.StatusCmd)
	return ret0
}

// ClusterSaveConfig indicates an expected call of ClusterSaveConfig.
func (mr *MockUniversalClientMockRecorder) ClusterSaveConfig(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ClusterSaveConfig", reflect.TypeOf((*MockUniversalClient)(nil).ClusterSaveConfig), ctx)
}

// ClusterSlaves mocks base method.
func (m *MockUniversalClient) ClusterSlaves(ctx context.Context, nodeID string) *redis.StringSliceCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ClusterSlaves", ctx, nodeID)
	ret0, _ := ret[0].(*redis.StringSliceCmd)
	return ret0
}

// ClusterSlaves indicates an expected call of ClusterSlaves.
func (mr *MockUniversalClientMockRecorder) ClusterSlaves(ctx, nodeID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ClusterSlaves", reflect.TypeOf((*MockUniversalClient)(nil).ClusterSlaves), ctx, nodeID)
}

// ClusterSlots mocks base method.
func (m *MockUniversalClient) ClusterSlots(ctx context.Context) *redis.ClusterSlotsCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ClusterSlots", ctx)
	ret0, _ := ret[0].(*redis.ClusterSlotsCmd)
	return ret0
}

// ClusterSlots indicates an expected call of ClusterSlots.
func (mr *MockUniversalClientMockRecorder) ClusterSlots(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ClusterSlots", reflect.TypeOf((*MockUniversalClient)(nil).ClusterSlots), ctx)
}

// Command mocks base method.
func (m *MockUniversalClient) Command(ctx context.Context) *redis.CommandsInfoCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Command", ctx)
	ret0, _ := ret[0].(*redis.CommandsInfoCmd)
	return ret0
}

// Command indicates an expected call of Command.
func (mr *MockUniversalClientMockRecorder) Command(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Command", reflect.TypeOf((*MockUniversalClient)(nil).Command), ctx)
}

// ConfigGet mocks base method.
func (m *MockUniversalClient) ConfigGet(ctx context.Context, parameter string) *redis.SliceCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ConfigGet", ctx, parameter)
	ret0, _ := ret[0].(*redis.SliceCmd)
	return ret0
}

// ConfigGet indicates an expected call of ConfigGet.
func (mr *MockUniversalClientMockRecorder) ConfigGet(ctx, parameter interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ConfigGet", reflect.TypeOf((*MockUniversalClient)(nil).ConfigGet), ctx, parameter)
}

// ConfigResetStat mocks base method.
func (m *MockUniversalClient) ConfigResetStat(ctx context.Context) *redis.StatusCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ConfigResetStat", ctx)
	ret0, _ := ret[0].(*redis.StatusCmd)
	return ret0
}

// ConfigResetStat indicates an expected call of ConfigResetStat.
func (mr *MockUniversalClientMockRecorder) ConfigResetStat(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ConfigResetStat", reflect.TypeOf((*MockUniversalClient)(nil).ConfigResetStat), ctx)
}

// ConfigRewrite mocks base method.
func (m *MockUniversalClient) ConfigRewrite(ctx context.Context) *redis.StatusCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ConfigRewrite", ctx)
	ret0, _ := ret[0].(*redis.StatusCmd)
	return ret0
}

// ConfigRewrite indicates an expected call of ConfigRewrite.
func (mr *MockUniversalClientMockRecorder) ConfigRewrite(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ConfigRewrite", reflect.TypeOf((*MockUniversalClient)(nil).ConfigRewrite), ctx)
}

// ConfigSet mocks base method.
func (m *MockUniversalClient) ConfigSet(ctx context.Context, parameter, value string) *redis.StatusCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ConfigSet", ctx, parameter, value)
	ret0, _ := ret[0].(*redis.StatusCmd)
	return ret0
}

// ConfigSet indicates an expected call of ConfigSet.
func (mr *MockUniversalClientMockRecorder) ConfigSet(ctx, parameter, value interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ConfigSet", reflect.TypeOf((*MockUniversalClient)(nil).ConfigSet), ctx, parameter, value)
}

// Context mocks base method.
func (m *MockUniversalClient) Context() context.Context {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Context")
	ret0, _ := ret[0].(context.Context)
	return ret0
}

// Context indicates an expected call of Context.
func (mr *MockUniversalClientMockRecorder) Context() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Context", reflect.TypeOf((*MockUniversalClient)(nil).Context))
}

// Copy mocks base method.
func (m *MockUniversalClient) Copy(ctx context.Context, sourceKey, destKey string, db int, replace bool) *redis.IntCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Copy", ctx, sourceKey, destKey, db, replace)
	ret0, _ := ret[0].(*redis.IntCmd)
	return ret0
}

// Copy indicates an expected call of Copy.
func (mr *MockUniversalClientMockRecorder) Copy(ctx, sourceKey, destKey, db, replace interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Copy", reflect.TypeOf((*MockUniversalClient)(nil).Copy), ctx, sourceKey, destKey, db, replace)
}

// DBSize mocks base method.
func (m *MockUniversalClient) DBSize(ctx context.Context) *redis.IntCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DBSize", ctx)
	ret0, _ := ret[0].(*redis.IntCmd)
	return ret0
}

// DBSize indicates an expected call of DBSize.
func (mr *MockUniversalClientMockRecorder) DBSize(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DBSize", reflect.TypeOf((*MockUniversalClient)(nil).DBSize), ctx)
}

// DebugObject mocks base method.
func (m *MockUniversalClient) DebugObject(ctx context.Context, key string) *redis.StringCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DebugObject", ctx, key)
	ret0, _ := ret[0].(*redis.StringCmd)
	return ret0
}

// DebugObject indicates an expected call of DebugObject.
func (mr *MockUniversalClientMockRecorder) DebugObject(ctx, key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DebugObject", reflect.TypeOf((*MockUniversalClient)(nil).DebugObject), ctx, key)
}

// Decr mocks base method.
func (m *MockUniversalClient) Decr(ctx context.Context, key string) *redis.IntCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Decr", ctx, key)
	ret0, _ := ret[0].(*redis.IntCmd)
	return ret0
}

// Decr indicates an expected call of Decr.
func (mr *MockUniversalClientMockRecorder) Decr(ctx, key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Decr", reflect.TypeOf((*MockUniversalClient)(nil).Decr), ctx, key)
}

// DecrBy mocks base method.
func (m *MockUniversalClient) DecrBy(ctx context.Context, key string, decrement int64) *redis.IntCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DecrBy", ctx, key, decrement)
	ret0, _ := ret[0].(*redis.IntCmd)
	return ret0
}

// DecrBy indicates an expected call of DecrBy.
func (mr *MockUniversalClientMockRecorder) DecrBy(ctx, key, decrement interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DecrBy", reflect.TypeOf((*MockUniversalClient)(nil).DecrBy), ctx, key, decrement)
}

// Del mocks base method.
func (m *MockUniversalClient) Del(ctx context.Context, keys ...string) *redis.IntCmd {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx}
	for _, a := range keys {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Del", varargs...)
	ret0, _ := ret[0].(*redis.IntCmd)
	return ret0
}

// Del indicates an expected call of Del.
func (mr *MockUniversalClientMockRecorder) Del(ctx interface{}, keys ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx}, keys...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Del", reflect.TypeOf((*MockUniversalClient)(nil).Del), varargs...)
}

// Do mocks base method.
func (m *MockUniversalClient) Do(ctx context.Context, args ...interface{}) *redis.Cmd {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Do", varargs...)
	ret0, _ := ret[0].(*redis.Cmd)
	return ret0
}

// Do indicates an expected call of Do.
func (mr *MockUniversalClientMockRecorder) Do(ctx interface{}, args ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx}, args...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Do", reflect.TypeOf((*MockUniversalClient)(nil).Do), varargs...)
}

// Dump mocks base method.
func (m *MockUniversalClient) Dump(ctx context.Context, key string) *redis.StringCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Dump", ctx, key)
	ret0, _ := ret[0].(*redis.StringCmd)
	return ret0
}

// Dump indicates an expected call of Dump.
func (mr *MockUniversalClientMockRecorder) Dump(ctx, key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Dump", reflect.TypeOf((*MockUniversalClient)(nil).Dump), ctx, key)
}

// Echo mocks base method.
func (m *MockUniversalClient) Echo(ctx context.Context, message interface{}) *redis.StringCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Echo", ctx, message)
	ret0, _ := ret[0].(*redis.StringCmd)
	return ret0
}

// Echo indicates an expected call of Echo.
func (mr *MockUniversalClientMockRecorder) Echo(ctx, message interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Echo", reflect.TypeOf((*MockUniversalClient)(nil).Echo), ctx, message)
}

// Eval mocks base method.
func (m *MockUniversalClient) Eval(ctx context.Context, script string, keys []string, args ...interface{}) *redis.Cmd {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, script, keys}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Eval", varargs...)
	ret0, _ := ret[0].(*redis.Cmd)
	return ret0
}

// Eval indicates an expected call of Eval.
func (mr *MockUniversalClientMockRecorder) Eval(ctx, script, keys interface{}, args ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, script, keys}, args...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Eval", reflect.TypeOf((*MockUniversalClient)(nil).Eval), varargs...)
}

// EvalSha mocks base method.
func (m *MockUniversalClient) EvalSha(ctx context.Context, sha1 string, keys []string, args ...interface{}) *redis.Cmd {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, sha1, keys}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "EvalSha", varargs...)
	ret0, _ := ret[0].(*redis.Cmd)
	return ret0
}

// EvalSha indicates an expected call of EvalSha.
func (mr *MockUniversalClientMockRecorder) EvalSha(ctx, sha1, keys interface{}, args ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, sha1, keys}, args...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EvalSha", reflect.TypeOf((*MockUniversalClient)(nil).EvalSha), varargs...)
}

// Exists mocks base method.
func (m *MockUniversalClient) Exists(ctx context.Context, keys ...string) *redis.IntCmd {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx}
	for _, a := range keys {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Exists", varargs...)
	ret0, _ := ret[0].(*redis.IntCmd)
	return ret0
}

// Exists indicates an expected call of Exists.
func (mr *MockUniversalClientMockRecorder) Exists(ctx interface{}, keys ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx}, keys...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Exists", reflect.TypeOf((*MockUniversalClient)(nil).Exists), varargs...)
}

// Expire mocks base method.
func (m *MockUniversalClient) Expire(ctx context.Context, key string, expiration time.Duration) *redis.BoolCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Expire", ctx, key, expiration)
	ret0, _ := ret[0].(*redis.BoolCmd)
	return ret0
}

// Expire indicates an expected call of Expire.
func (mr *MockUniversalClientMockRecorder) Expire(ctx, key, expiration interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Expire", reflect.TypeOf((*MockUniversalClient)(nil).Expire), ctx, key, expiration)
}

// ExpireAt mocks base method.
func (m *MockUniversalClient) ExpireAt(ctx context.Context, key string, tm time.Time) *redis.BoolCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ExpireAt", ctx, key, tm)
	ret0, _ := ret[0].(*redis.BoolCmd)
	return ret0
}

// ExpireAt indicates an expected call of ExpireAt.
func (mr *MockUniversalClientMockRecorder) ExpireAt(ctx, key, tm interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ExpireAt", reflect.TypeOf((*MockUniversalClient)(nil).ExpireAt), ctx, key, tm)
}

// ExpireGT mocks base method.
func (m *MockUniversalClient) ExpireGT(ctx context.Context, key string, expiration time.Duration) *redis.BoolCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ExpireGT", ctx, key, expiration)
	ret0, _ := ret[0].(*redis.BoolCmd)
	return ret0
}

// ExpireGT indicates an expected call of ExpireGT.
func (mr *MockUniversalClientMockRecorder) ExpireGT(ctx, key, expiration interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ExpireGT", reflect.TypeOf((*MockUniversalClient)(nil).ExpireGT), ctx, key, expiration)
}

// ExpireLT mocks base method.
func (m *MockUniversalClient) ExpireLT(ctx context.Context, key string, expiration time.Duration) *redis.BoolCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ExpireLT", ctx, key, expiration)
	ret0, _ := ret[0].(*redis.BoolCmd)
	return ret0
}

// ExpireLT indicates an expected call of ExpireLT.
func (mr *MockUniversalClientMockRecorder) ExpireLT(ctx, key, expiration interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ExpireLT", reflect.TypeOf((*MockUniversalClient)(nil).ExpireLT), ctx, key, expiration)
}

// ExpireNX mocks base method.
func (m *MockUniversalClient) ExpireNX(ctx context.Context, key string, expiration time.Duration) *redis.BoolCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ExpireNX", ctx, key, expiration)
	ret0, _ := ret[0].(*redis.BoolCmd)
	return ret0
}

// ExpireNX indicates an expected call of ExpireNX.
func (mr *MockUniversalClientMockRecorder) ExpireNX(ctx, key, expiration interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ExpireNX", reflect.TypeOf((*MockUniversalClient)(nil).ExpireNX), ctx, key, expiration)
}

// ExpireXX mocks base method.
func (m *MockUniversalClient) ExpireXX(ctx context.Context, key string, expiration time.Duration) *redis.BoolCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ExpireXX", ctx, key, expiration)
	ret0, _ := ret[0].(*redis.BoolCmd)
	return ret0
}

// ExpireXX indicates an expected call of ExpireXX.
func (mr *MockUniversalClientMockRecorder) ExpireXX(ctx, key, expiration interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ExpireXX", reflect.TypeOf((*MockUniversalClient)(nil).ExpireXX), ctx, key, expiration)
}

// FlushAll mocks base method.
func (m *MockUniversalClient) FlushAll(ctx context.Context) *redis.StatusCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FlushAll", ctx)
	ret0, _ := ret[0].(*redis.StatusCmd)
	return ret0
}

// FlushAll indicates an expected call of FlushAll.
func (mr *MockUniversalClientMockRecorder) FlushAll(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FlushAll", reflect.TypeOf((*MockUniversalClient)(nil).FlushAll), ctx)
}

// FlushAllAsync mocks base method.
func (m *MockUniversalClient) FlushAllAsync(ctx context.Context) *redis.StatusCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FlushAllAsync", ctx)
	ret0, _ := ret[0].(*redis.StatusCmd)
	return ret0
}

// FlushAllAsync indicates an expected call of FlushAllAsync.
func (mr *MockUniversalClientMockRecorder) FlushAllAsync(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FlushAllAsync", reflect.TypeOf((*MockUniversalClient)(nil).FlushAllAsync), ctx)
}

// FlushDB mocks base method.
func (m *MockUniversalClient) FlushDB(ctx context.Context) *redis.StatusCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FlushDB", ctx)
	ret0, _ := ret[0].(*redis.StatusCmd)
	return ret0
}

// FlushDB indicates an expected call of FlushDB.
func (mr *MockUniversalClientMockRecorder) FlushDB(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FlushDB", reflect.TypeOf((*MockUniversalClient)(nil).FlushDB), ctx)
}

// FlushDBAsync mocks base method.
func (m *MockUniversalClient) FlushDBAsync(ctx context.Context) *redis.StatusCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FlushDBAsync", ctx)
	ret0, _ := ret[0].(*redis.StatusCmd)
	return ret0
}

// FlushDBAsync indicates an expected call of FlushDBAsync.
func (mr *MockUniversalClientMockRecorder) FlushDBAsync(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FlushDBAsync", reflect.TypeOf((*MockUniversalClient)(nil).FlushDBAsync), ctx)
}

// GeoAdd mocks base method.
func (m *MockUniversalClient) GeoAdd(ctx context.Context, key string, geoLocation ...*redis.GeoLocation) *redis.IntCmd {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, key}
	for _, a := range geoLocation {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GeoAdd", varargs...)
	ret0, _ := ret[0].(*redis.IntCmd)
	return ret0
}

// GeoAdd indicates an expected call of GeoAdd.
func (mr *MockUniversalClientMockRecorder) GeoAdd(ctx, key interface{}, geoLocation ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, key}, geoLocation...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GeoAdd", reflect.TypeOf((*MockUniversalClient)(nil).GeoAdd), varargs...)
}

// GeoDist mocks base method.
func (m *MockUniversalClient) GeoDist(ctx context.Context, key, member1, member2, unit string) *redis.FloatCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GeoDist", ctx, key, member1, member2, unit)
	ret0, _ := ret[0].(*redis.FloatCmd)
	return ret0
}

// GeoDist indicates an expected call of GeoDist.
func (mr *MockUniversalClientMockRecorder) GeoDist(ctx, key, member1, member2, unit interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GeoDist", reflect.TypeOf((*MockUniversalClient)(nil).GeoDist), ctx, key, member1, member2, unit)
}

// GeoHash mocks base method.
func (m *MockUniversalClient) GeoHash(ctx context.Context, key string, members ...string) *redis.StringSliceCmd {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, key}
	for _, a := range members {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GeoHash", varargs...)
	ret0, _ := ret[0].(*redis.StringSliceCmd)
	return ret0
}

// GeoHash indicates an expected call of GeoHash.
func (mr *MockUniversalClientMockRecorder) GeoHash(ctx, key interface{}, members ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, key}, members...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GeoHash", reflect.TypeOf((*MockUniversalClient)(nil).GeoHash), varargs...)
}

// GeoPos mocks base method.
func (m *MockUniversalClient) GeoPos(ctx context.Context, key string, members ...string) *redis.GeoPosCmd {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, key}
	for _, a := range members {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GeoPos", varargs...)
	ret0, _ := ret[0].(*redis.GeoPosCmd)
	return ret0
}

// GeoPos indicates an expected call of GeoPos.
func (mr *MockUniversalClientMockRecorder) GeoPos(ctx, key interface{}, members ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, key}, members...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GeoPos", reflect.TypeOf((*MockUniversalClient)(nil).GeoPos), varargs...)
}

// GeoRadius mocks base method.
func (m *MockUniversalClient) GeoRadius(ctx context.Context, key string, longitude, latitude float64, query *redis.GeoRadiusQuery) *redis.GeoLocationCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GeoRadius", ctx, key, longitude, latitude, query)
	ret0, _ := ret[0].(*redis.GeoLocationCmd)
	return ret0
}

// GeoRadius indicates an expected call of GeoRadius.
func (mr *MockUniversalClientMockRecorder) GeoRadius(ctx, key, longitude, latitude, query interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GeoRadius", reflect.TypeOf((*MockUniversalClient)(nil).GeoRadius), ctx, key, longitude, latitude, query)
}

// GeoRadiusByMember mocks base method.
func (m *MockUniversalClient) GeoRadiusByMember(ctx context.Context, key, member string, query *redis.GeoRadiusQuery) *redis.GeoLocationCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GeoRadiusByMember", ctx, key, member, query)
	ret0, _ := ret[0].(*redis.GeoLocationCmd)
	return ret0
}

// GeoRadiusByMember indicates an expected call of GeoRadiusByMember.
func (mr *MockUniversalClientMockRecorder) GeoRadiusByMember(ctx, key, member, query interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GeoRadiusByMember", reflect.TypeOf((*MockUniversalClient)(nil).GeoRadiusByMember), ctx, key, member, query)
}

// GeoRadiusByMemberStore mocks base method.
func (m *MockUniversalClient) GeoRadiusByMemberStore(ctx context.Context, key, member string, query *redis.GeoRadiusQuery) *redis.IntCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GeoRadiusByMemberStore", ctx, key, member, query)
	ret0, _ := ret[0].(*redis.IntCmd)
	return ret0
}

// GeoRadiusByMemberStore indicates an expected call of GeoRadiusByMemberStore.
func (mr *MockUniversalClientMockRecorder) GeoRadiusByMemberStore(ctx, key, member, query interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GeoRadiusByMemberStore", reflect.TypeOf((*MockUniversalClient)(nil).GeoRadiusByMemberStore), ctx, key, member, query)
}

// GeoRadiusStore mocks base method.
func (m *MockUniversalClient) GeoRadiusStore(ctx context.Context, key string, longitude, latitude float64, query *redis.GeoRadiusQuery) *redis.IntCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GeoRadiusStore", ctx, key, longitude, latitude, query)
	ret0, _ := ret[0].(*redis.IntCmd)
	return ret0
}

// GeoRadiusStore indicates an expected call of GeoRadiusStore.
func (mr *MockUniversalClientMockRecorder) GeoRadiusStore(ctx, key, longitude, latitude, query interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GeoRadiusStore", reflect.TypeOf((*MockUniversalClient)(nil).GeoRadiusStore), ctx, key, longitude, latitude, query)
}

// GeoSearch mocks base method.
func (m *MockUniversalClient) GeoSearch(ctx context.Context, key string, q *redis.GeoSearchQuery) *redis.StringSliceCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GeoSearch", ctx, key, q)
	ret0, _ := ret[0].(*redis.StringSliceCmd)
	return ret0
}

// GeoSearch indicates an expected call of GeoSearch.
func (mr *MockUniversalClientMockRecorder) GeoSearch(ctx, key, q interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GeoSearch", reflect.TypeOf((*MockUniversalClient)(nil).GeoSearch), ctx, key, q)
}

// GeoSearchLocation mocks base method.
func (m *MockUniversalClient) GeoSearchLocation(ctx context.Context, key string, q *redis.GeoSearchLocationQuery) *redis.GeoSearchLocationCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GeoSearchLocation", ctx, key, q)
	ret0, _ := ret[0].(*redis.GeoSearchLocationCmd)
	return ret0
}

// GeoSearchLocation indicates an expected call of GeoSearchLocation.
func (mr *MockUniversalClientMockRecorder) GeoSearchLocation(ctx, key, q interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GeoSearchLocation", reflect.TypeOf((*MockUniversalClient)(nil).GeoSearchLocation), ctx, key, q)
}

// GeoSearchStore mocks base method.
func (m *MockUniversalClient) GeoSearchStore(ctx context.Context, key, store string, q *redis.GeoSearchStoreQuery) *redis.IntCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GeoSearchStore", ctx, key, store, q)
	ret0, _ := ret[0].(*redis.IntCmd)
	return ret0
}

// GeoSearchStore indicates an expected call of GeoSearchStore.
func (mr *MockUniversalClientMockRecorder) GeoSearchStore(ctx, key, store, q interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GeoSearchStore", reflect.TypeOf((*MockUniversalClient)(nil).GeoSearchStore), ctx, key, store, q)
}

// Get mocks base method.
func (m *MockUniversalClient) Get(ctx context.Context, key string) *redis.StringCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", ctx, key)
	ret0, _ := ret[0].(*redis.StringCmd)
	return ret0
}

// Get indicates an expected call of Get.
func (mr *MockUniversalClientMockRecorder) Get(ctx, key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockUniversalClient)(nil).Get), ctx, key)
}

// GetBit mocks base method.
func (m *MockUniversalClient) GetBit(ctx context.Context, key string, offset int64) *redis.IntCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBit", ctx, key, offset)
	ret0, _ := ret[0].(*redis.IntCmd)
	return ret0
}

// GetBit indicates an expected call of GetBit.
func (mr *MockUniversalClientMockRecorder) GetBit(ctx, key, offset interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBit", reflect.TypeOf((*MockUniversalClient)(nil).GetBit), ctx, key, offset)
}

// GetDel mocks base method.
func (m *MockUniversalClient) GetDel(ctx context.Context, key string) *redis.StringCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDel", ctx, key)
	ret0, _ := ret[0].(*redis.StringCmd)
	return ret0
}

// GetDel indicates an expected call of GetDel.
func (mr *MockUniversalClientMockRecorder) GetDel(ctx, key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDel", reflect.TypeOf((*MockUniversalClient)(nil).GetDel), ctx, key)
}

// GetEx mocks base method.
func (m *MockUniversalClient) GetEx(ctx context.Context, key string, expiration time.Duration) *redis.StringCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetEx", ctx, key, expiration)
	ret0, _ := ret[0].(*redis.StringCmd)
	return ret0
}

// GetEx indicates an expected call of GetEx.
func (mr *MockUniversalClientMockRecorder) GetEx(ctx, key, expiration interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetEx", reflect.TypeOf((*MockUniversalClient)(nil).GetEx), ctx, key, expiration)
}

// GetRange mocks base method.
func (m *MockUniversalClient) GetRange(ctx context.Context, key string, start, end int64) *redis.StringCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRange", ctx, key, start, end)
	ret0, _ := ret[0].(*redis.StringCmd)
	return ret0
}

// GetRange indicates an expected call of GetRange.
func (mr *MockUniversalClientMockRecorder) GetRange(ctx, key, start, end interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRange", reflect.TypeOf((*MockUniversalClient)(nil).GetRange), ctx, key, start, end)
}

// GetSet mocks base method.
func (m *MockUniversalClient) GetSet(ctx context.Context, key string, value interface{}) *redis.StringCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSet", ctx, key, value)
	ret0, _ := ret[0].(*redis.StringCmd)
	return ret0
}

// GetSet indicates an expected call of GetSet.
func (mr *MockUniversalClientMockRecorder) GetSet(ctx, key, value interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSet", reflect.TypeOf((*MockUniversalClient)(nil).GetSet), ctx, key, value)
}

// HDel mocks base method.
func (m *MockUniversalClient) HDel(ctx context.Context, key string, fields ...string) *redis.IntCmd {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, key}
	for _, a := range fields {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "HDel", varargs...)
	ret0, _ := ret[0].(*redis.IntCmd)
	return ret0
}

// HDel indicates an expected call of HDel.
func (mr *MockUniversalClientMockRecorder) HDel(ctx, key interface{}, fields ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, key}, fields...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HDel", reflect.TypeOf((*MockUniversalClient)(nil).HDel), varargs...)
}

// HExists mocks base method.
func (m *MockUniversalClient) HExists(ctx context.Context, key, field string) *redis.BoolCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HExists", ctx, key, field)
	ret0, _ := ret[0].(*redis.BoolCmd)
	return ret0
}

// HExists indicates an expected call of HExists.
func (mr *MockUniversalClientMockRecorder) HExists(ctx, key, field interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HExists", reflect.TypeOf((*MockUniversalClient)(nil).HExists), ctx, key, field)
}

// HGet mocks base method.
func (m *MockUniversalClient) HGet(ctx context.Context, key, field string) *redis.StringCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HGet", ctx, key, field)
	ret0, _ := ret[0].(*redis.StringCmd)
	return ret0
}

// HGet indicates an expected call of HGet.
func (mr *MockUniversalClientMockRecorder) HGet(ctx, key, field interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HGet", reflect.TypeOf((*MockUniversalClient)(nil).HGet), ctx, key, field)
}

// HGetAll mocks base method.
func (m *MockUniversalClient) HGetAll(ctx context.Context, key string) *redis.StringStringMapCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HGetAll", ctx, key)
	ret0, _ := ret[0].(*redis.StringStringMapCmd)
	return ret0
}

// HGetAll indicates an expected call of HGetAll.
func (mr *MockUniversalClientMockRecorder) HGetAll(ctx, key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HGetAll", reflect.TypeOf((*MockUniversalClient)(nil).HGetAll), ctx, key)
}

// HIncrBy mocks base method.
func (m *MockUniversalClient) HIncrBy(ctx context.Context, key, field string, incr int64) *redis.IntCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HIncrBy", ctx, key, field, incr)
	ret0, _ := ret[0].(*redis.IntCmd)
	return ret0
}

// HIncrBy indicates an expected call of HIncrBy.
func (mr *MockUniversalClientMockRecorder) HIncrBy(ctx, key, field, incr interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HIncrBy", reflect.TypeOf((*MockUniversalClient)(nil).HIncrBy), ctx, key, field, incr)
}

// HIncrByFloat mocks base method.
func (m *MockUniversalClient) HIncrByFloat(ctx context.Context, key, field string, incr float64) *redis.FloatCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HIncrByFloat", ctx, key, field, incr)
	ret0, _ := ret[0].(*redis.FloatCmd)
	return ret0
}

// HIncrByFloat indicates an expected call of HIncrByFloat.
func (mr *MockUniversalClientMockRecorder) HIncrByFloat(ctx, key, field, incr interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HIncrByFloat", reflect.TypeOf((*MockUniversalClient)(nil).HIncrByFloat), ctx, key, field, incr)
}

// HKeys mocks base method.
func (m *MockUniversalClient) HKeys(ctx context.Context, key string) *redis.StringSliceCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HKeys", ctx, key)
	ret0, _ := ret[0].(*redis.StringSliceCmd)
	return ret0
}

// HKeys indicates an expected call of HKeys.
func (mr *MockUniversalClientMockRecorder) HKeys(ctx, key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HKeys", reflect.TypeOf((*MockUniversalClient)(nil).HKeys), ctx, key)
}

// HLen mocks base method.
func (m *MockUniversalClient) HLen(ctx context.Context, key string) *redis.IntCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HLen", ctx, key)
	ret0, _ := ret[0].(*redis.IntCmd)
	return ret0
}

// HLen indicates an expected call of HLen.
func (mr *MockUniversalClientMockRecorder) HLen(ctx, key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HLen", reflect.TypeOf((*MockUniversalClient)(nil).HLen), ctx, key)
}

// HMGet mocks base method.
func (m *MockUniversalClient) HMGet(ctx context.Context, key string, fields ...string) *redis.SliceCmd {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, key}
	for _, a := range fields {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "HMGet", varargs...)
	ret0, _ := ret[0].(*redis.SliceCmd)
	return ret0
}

// HMGet indicates an expected call of HMGet.
func (mr *MockUniversalClientMockRecorder) HMGet(ctx, key interface{}, fields ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, key}, fields...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HMGet", reflect.TypeOf((*MockUniversalClient)(nil).HMGet), varargs...)
}

// HMSet mocks base method.
func (m *MockUniversalClient) HMSet(ctx context.Context, key string, values ...interface{}) *redis.BoolCmd {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, key}
	for _, a := range values {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "HMSet", varargs...)
	ret0, _ := ret[0].(*redis.BoolCmd)
	return ret0
}

// HMSet indicates an expected call of HMSet.
func (mr *MockUniversalClientMockRecorder) HMSet(ctx, key interface{}, values ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, key}, values...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HMSet", reflect.TypeOf((*MockUniversalClient)(nil).HMSet), varargs...)
}

// HRandField mocks base method.
func (m *MockUniversalClient) HRandField(ctx context.Context, key string, count int, withValues bool) *redis.StringSliceCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HRandField", ctx, key, count, withValues)
	ret0, _ := ret[0].(*redis.StringSliceCmd)
	return ret0
}

// HRandField indicates an expected call of HRandField.
func (mr *MockUniversalClientMockRecorder) HRandField(ctx, key, count, withValues interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HRandField", reflect.TypeOf((*MockUniversalClient)(nil).HRandField), ctx, key, count, withValues)
}

// HScan mocks base method.
func (m *MockUniversalClient) HScan(ctx context.Context, key string, cursor uint64, match string, count int64) *redis.ScanCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HScan", ctx, key, cursor, match, count)
	ret0, _ := ret[0].(*redis.ScanCmd)
	return ret0
}

// HScan indicates an expected call of HScan.
func (mr *MockUniversalClientMockRecorder) HScan(ctx, key, cursor, match, count interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HScan", reflect.TypeOf((*MockUniversalClient)(nil).HScan), ctx, key, cursor, match, count)
}

// HSet mocks base method.
func (m *MockUniversalClient) HSet(ctx context.Context, key string, values ...interface{}) *redis.IntCmd {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, key}
	for _, a := range values {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "HSet", varargs...)
	ret0, _ := ret[0].(*redis.IntCmd)
	return ret0
}

// HSet indicates an expected call of HSet.
func (mr *MockUniversalClientMockRecorder) HSet(ctx, key interface{}, values ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, key}, values...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HSet", reflect.TypeOf((*MockUniversalClient)(nil).HSet), varargs...)
}

// HSetNX mocks base method.
func (m *MockUniversalClient) HSetNX(ctx context.Context, key, field string, value interface{}) *redis.BoolCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HSetNX", ctx, key, field, value)
	ret0, _ := ret[0].(*redis.BoolCmd)
	return ret0
}

// HSetNX indicates an expected call of HSetNX.
func (mr *MockUniversalClientMockRecorder) HSetNX(ctx, key, field, value interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HSetNX", reflect.TypeOf((*MockUniversalClient)(nil).HSetNX), ctx, key, field, value)
}

// HVals mocks base method.
func (m *MockUniversalClient) HVals(ctx context.Context, key string) *redis.StringSliceCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HVals", ctx, key)
	ret0, _ := ret[0].(*redis.StringSliceCmd)
	return ret0
}

// HVals indicates an expected call of HVals.
func (mr *MockUniversalClientMockRecorder) HVals(ctx, key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HVals", reflect.TypeOf((*MockUniversalClient)(nil).HVals), ctx, key)
}

// Incr mocks base method.
func (m *MockUniversalClient) Incr(ctx context.Context, key string) *redis.IntCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Incr", ctx, key)
	ret0, _ := ret[0].(*redis.IntCmd)
	return ret0
}

// Incr indicates an expected call of Incr.
func (mr *MockUniversalClientMockRecorder) Incr(ctx, key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Incr", reflect.TypeOf((*MockUniversalClient)(nil).Incr), ctx, key)
}

// IncrBy mocks base method.
func (m *MockUniversalClient) IncrBy(ctx context.Context, key string, value int64) *redis.IntCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IncrBy", ctx, key, value)
	ret0, _ := ret[0].(*redis.IntCmd)
	return ret0
}

// IncrBy indicates an expected call of IncrBy.
func (mr *MockUniversalClientMockRecorder) IncrBy(ctx, key, value interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IncrBy", reflect.TypeOf((*MockUniversalClient)(nil).IncrBy), ctx, key, value)
}

// IncrByFloat mocks base method.
func (m *MockUniversalClient) IncrByFloat(ctx context.Context, key string, value float64) *redis.FloatCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IncrByFloat", ctx, key, value)
	ret0, _ := ret[0].(*redis.FloatCmd)
	return ret0
}

// IncrByFloat indicates an expected call of IncrByFloat.
func (mr *MockUniversalClientMockRecorder) IncrByFloat(ctx, key, value interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IncrByFloat", reflect.TypeOf((*MockUniversalClient)(nil).IncrByFloat), ctx, key, value)
}

// Info mocks base method.
func (m *MockUniversalClient) Info(ctx context.Context, section ...string) *redis.StringCmd {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx}
	for _, a := range section {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Info", varargs...)
	ret0, _ := ret[0].(*redis.StringCmd)
	return ret0
}

// Info indicates an expected call of Info.
func (mr *MockUniversalClientMockRecorder) Info(ctx interface{}, section ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx}, section...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Info", reflect.TypeOf((*MockUniversalClient)(nil).Info), varargs...)
}

// Keys mocks base method.
func (m *MockUniversalClient) Keys(ctx context.Context, pattern string) *redis.StringSliceCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Keys", ctx, pattern)
	ret0, _ := ret[0].(*redis.StringSliceCmd)
	return ret0
}

// Keys indicates an expected call of Keys.
func (mr *MockUniversalClientMockRecorder) Keys(ctx, pattern interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Keys", reflect.TypeOf((*MockUniversalClient)(nil).Keys), ctx, pattern)
}

// LIndex mocks base method.
func (m *MockUniversalClient) LIndex(ctx context.Context, key string, index int64) *redis.StringCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LIndex", ctx, key, index)
	ret0, _ := ret[0].(*redis.StringCmd)
	return ret0
}

// LIndex indicates an expected call of LIndex.
func (mr *MockUniversalClientMockRecorder) LIndex(ctx, key, index interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LIndex", reflect.TypeOf((*MockUniversalClient)(nil).LIndex), ctx, key, index)
}

// LInsert mocks base method.
func (m *MockUniversalClient) LInsert(ctx context.Context, key, op string, pivot, value interface{}) *redis.IntCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LInsert", ctx, key, op, pivot, value)
	ret0, _ := ret[0].(*redis.IntCmd)
	return ret0
}

// LInsert indicates an expected call of LInsert.
func (mr *MockUniversalClientMockRecorder) LInsert(ctx, key, op, pivot, value interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LInsert", reflect.TypeOf((*MockUniversalClient)(nil).LInsert), ctx, key, op, pivot, value)
}

// LInsertAfter mocks base method.
func (m *MockUniversalClient) LInsertAfter(ctx context.Context, key string, pivot, value interface{}) *redis.IntCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LInsertAfter", ctx, key, pivot, value)
	ret0, _ := ret[0].(*redis.IntCmd)
	return ret0
}

// LInsertAfter indicates an expected call of LInsertAfter.
func (mr *MockUniversalClientMockRecorder) LInsertAfter(ctx, key, pivot, value interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LInsertAfter", reflect.TypeOf((*MockUniversalClient)(nil).LInsertAfter), ctx, key, pivot, value)
}

// LInsertBefore mocks base method.
func (m *MockUniversalClient) LInsertBefore(ctx context.Context, key string, pivot, value interface{}) *redis.IntCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LInsertBefore", ctx, key, pivot, value)
	ret0, _ := ret[0].(*redis.IntCmd)
	return ret0
}

// LInsertBefore indicates an expected call of LInsertBefore.
func (mr *MockUniversalClientMockRecorder) LInsertBefore(ctx, key, pivot, value interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LInsertBefore", reflect.TypeOf((*MockUniversalClient)(nil).LInsertBefore), ctx, key, pivot, value)
}

// LLen mocks base method.
func (m *MockUniversalClient) LLen(ctx context.Context, key string) *redis.IntCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LLen", ctx, key)
	ret0, _ := ret[0].(*redis.IntCmd)
	return ret0
}

// LLen indicates an expected call of LLen.
func (mr *MockUniversalClientMockRecorder) LLen(ctx, key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LLen", reflect.TypeOf((*MockUniversalClient)(nil).LLen), ctx, key)
}

// LMove mocks base method.
func (m *MockUniversalClient) LMove(ctx context.Context, source, destination, srcpos, destpos string) *redis.StringCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LMove", ctx, source, destination, srcpos, destpos)
	ret0, _ := ret[0].(*redis.StringCmd)
	return ret0
}

// LMove indicates an expected call of LMove.
func (mr *MockUniversalClientMockRecorder) LMove(ctx, source, destination, srcpos, destpos interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LMove", reflect.TypeOf((*MockUniversalClient)(nil).LMove), ctx, source, destination, srcpos, destpos)
}

// LPop mocks base method.
func (m *MockUniversalClient) LPop(ctx context.Context, key string) *redis.StringCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LPop", ctx, key)
	ret0, _ := ret[0].(*redis.StringCmd)
	return ret0
}

// LPop indicates an expected call of LPop.
func (mr *MockUniversalClientMockRecorder) LPop(ctx, key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LPop", reflect.TypeOf((*MockUniversalClient)(nil).LPop), ctx, key)
}

// LPopCount mocks base method.
func (m *MockUniversalClient) LPopCount(ctx context.Context, key string, count int) *redis.StringSliceCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LPopCount", ctx, key, count)
	ret0, _ := ret[0].(*redis.StringSliceCmd)
	return ret0
}

// LPopCount indicates an expected call of LPopCount.
func (mr *MockUniversalClientMockRecorder) LPopCount(ctx, key, count interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LPopCount", reflect.TypeOf((*MockUniversalClient)(nil).LPopCount), ctx, key, count)
}

// LPos mocks base method.
func (m *MockUniversalClient) LPos(ctx context.Context, key, value string, args redis.LPosArgs) *redis.IntCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LPos", ctx, key, value, args)
	ret0, _ := ret[0].(*redis.IntCmd)
	return ret0
}

// LPos indicates an expected call of LPos.
func (mr *MockUniversalClientMockRecorder) LPos(ctx, key, value, args interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LPos", reflect.TypeOf((*MockUniversalClient)(nil).LPos), ctx, key, value, args)
}

// LPosCount mocks base method.
func (m *MockUniversalClient) LPosCount(ctx context.Context, key, value string, count int64, args redis.LPosArgs) *redis.IntSliceCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LPosCount", ctx, key, value, count, args)
	ret0, _ := ret[0].(*redis.IntSliceCmd)
	return ret0
}

// LPosCount indicates an expected call of LPosCount.
func (mr *MockUniversalClientMockRecorder) LPosCount(ctx, key, value, count, args interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LPosCount", reflect.TypeOf((*MockUniversalClient)(nil).LPosCount), ctx, key, value, count, args)
}

// LPush mocks base method.
func (m *MockUniversalClient) LPush(ctx context.Context, key string, values ...interface{}) *redis.IntCmd {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, key}
	for _, a := range values {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "LPush", varargs...)
	ret0, _ := ret[0].(*redis.IntCmd)
	return ret0
}

// LPush indicates an expected call of LPush.
func (mr *MockUniversalClientMockRecorder) LPush(ctx, key interface{}, values ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, key}, values...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LPush", reflect.TypeOf((*MockUniversalClient)(nil).LPush), varargs...)
}

// LPushX mocks base method.
func (m *MockUniversalClient) LPushX(ctx context.Context, key string, values ...interface{}) *redis.IntCmd {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, key}
	for _, a := range values {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "LPushX", varargs...)
	ret0, _ := ret[0].(*redis.IntCmd)
	return ret0
}

// LPushX indicates an expected call of LPushX.
func (mr *MockUniversalClientMockRecorder) LPushX(ctx, key interface{}, values ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, key}, values...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LPushX", reflect.TypeOf((*MockUniversalClient)(nil).LPushX), varargs...)
}

// LRange mocks base method.
func (m *MockUniversalClient) LRange(ctx context.Context, key string, start, stop int64) *redis.StringSliceCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LRange", ctx, key, start, stop)
	ret0, _ := ret[0].(*redis.StringSliceCmd)
	return ret0
}

// LRange indicates an expected call of LRange.
func (mr *MockUniversalClientMockRecorder) LRange(ctx, key, start, stop interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LRange", reflect.TypeOf((*MockUniversalClient)(nil).LRange), ctx, key, start, stop)
}

// LRem mocks base method.
func (m *MockUniversalClient) LRem(ctx context.Context, key string, count int64, value interface{}) *redis.IntCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LRem", ctx, key, count, value)
	ret0, _ := ret[0].(*redis.IntCmd)
	return ret0
}

// LRem indicates an expected call of LRem.
func (mr *MockUniversalClientMockRecorder) LRem(ctx, key, count, value interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LRem", reflect.TypeOf((*MockUniversalClient)(nil).LRem), ctx, key, count, value)
}

// LSet mocks base method.
func (m *MockUniversalClient) LSet(ctx context.Context, key string, index int64, value interface{}) *redis.StatusCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LSet", ctx, key, index, value)
	ret0, _ := ret[0].(*redis.StatusCmd)
	return ret0
}

// LSet indicates an expected call of LSet.
func (mr *MockUniversalClientMockRecorder) LSet(ctx, key, index, value interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LSet", reflect.TypeOf((*MockUniversalClient)(nil).LSet), ctx, key, index, value)
}

// LTrim mocks base method.
func (m *MockUniversalClient) LTrim(ctx context.Context, key string, start, stop int64) *redis.StatusCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LTrim", ctx, key, start, stop)
	ret0, _ := ret[0].(*redis.StatusCmd)
	return ret0
}

// LTrim indicates an expected call of LTrim.
func (mr *MockUniversalClientMockRecorder) LTrim(ctx, key, start, stop interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LTrim", reflect.TypeOf((*MockUniversalClient)(nil).LTrim), ctx, key, start, stop)
}

// LastSave mocks base method.
func (m *MockUniversalClient) LastSave(ctx context.Context) *redis.IntCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LastSave", ctx)
	ret0, _ := ret[0].(*redis.IntCmd)
	return ret0
}

// LastSave indicates an expected call of LastSave.
func (mr *MockUniversalClientMockRecorder) LastSave(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LastSave", reflect.TypeOf((*MockUniversalClient)(nil).LastSave), ctx)
}

// MGet mocks base method.
func (m *MockUniversalClient) MGet(ctx context.Context, keys ...string) *redis.SliceCmd {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx}
	for _, a := range keys {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "MGet", varargs...)
	ret0, _ := ret[0].(*redis.SliceCmd)
	return ret0
}

// MGet indicates an expected call of MGet.
func (mr *MockUniversalClientMockRecorder) MGet(ctx interface{}, keys ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx}, keys...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MGet", reflect.TypeOf((*MockUniversalClient)(nil).MGet), varargs...)
}

// MSet mocks base method.
func (m *MockUniversalClient) MSet(ctx context.Context, values ...interface{}) *redis.StatusCmd {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx}
	for _, a := range values {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "MSet", varargs...)
	ret0, _ := ret[0].(*redis.StatusCmd)
	return ret0
}

// MSet indicates an expected call of MSet.
func (mr *MockUniversalClientMockRecorder) MSet(ctx interface{}, values ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx}, values...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MSet", reflect.TypeOf((*MockUniversalClient)(nil).MSet), varargs...)
}

// MSetNX mocks base method.
func (m *MockUniversalClient) MSetNX(ctx context.Context, values ...interface{}) *redis.BoolCmd {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx}
	for _, a := range values {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "MSetNX", varargs...)
	ret0, _ := ret[0].(*redis.BoolCmd)
	return ret0
}

// MSetNX indicates an expected call of MSetNX.
func (mr *MockUniversalClientMockRecorder) MSetNX(ctx interface{}, values ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx}, values...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MSetNX", reflect.TypeOf((*MockUniversalClient)(nil).MSetNX), varargs...)
}

// MemoryUsage mocks base method.
func (m *MockUniversalClient) MemoryUsage(ctx context.Context, key string, samples ...int) *redis.IntCmd {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, key}
	for _, a := range samples {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "MemoryUsage", varargs...)
	ret0, _ := ret[0].(*redis.IntCmd)
	return ret0
}

// MemoryUsage indicates an expected call of MemoryUsage.
func (mr *MockUniversalClientMockRecorder) MemoryUsage(ctx, key interface{}, samples ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, key}, samples...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MemoryUsage", reflect.TypeOf((*MockUniversalClient)(nil).MemoryUsage), varargs...)
}

// Migrate mocks base method.
func (m *MockUniversalClient) Migrate(ctx context.Context, host, port, key string, db int, timeout time.Duration) *redis.StatusCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Migrate", ctx, host, port, key, db, timeout)
	ret0, _ := ret[0].(*redis.StatusCmd)
	return ret0
}

// Migrate indicates an expected call of Migrate.
func (mr *MockUniversalClientMockRecorder) Migrate(ctx, host, port, key, db, timeout interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Migrate", reflect.TypeOf((*MockUniversalClient)(nil).Migrate), ctx, host, port, key, db, timeout)
}

// Move mocks base method.
func (m *MockUniversalClient) Move(ctx context.Context, key string, db int) *redis.BoolCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Move", ctx, key, db)
	ret0, _ := ret[0].(*redis.BoolCmd)
	return ret0
}

// Move indicates an expected call of Move.
func (mr *MockUniversalClientMockRecorder) Move(ctx, key, db interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Move", reflect.TypeOf((*MockUniversalClient)(nil).Move), ctx, key, db)
}

// ObjectEncoding mocks base method.
func (m *MockUniversalClient) ObjectEncoding(ctx context.Context, key string) *redis.StringCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ObjectEncoding", ctx, key)
	ret0, _ := ret[0].(*redis.StringCmd)
	return ret0
}

// ObjectEncoding indicates an expected call of ObjectEncoding.
func (mr *MockUniversalClientMockRecorder) ObjectEncoding(ctx, key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ObjectEncoding", reflect.TypeOf((*MockUniversalClient)(nil).ObjectEncoding), ctx, key)
}

// ObjectIdleTime mocks base method.
func (m *MockUniversalClient) ObjectIdleTime(ctx context.Context, key string) *redis.DurationCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ObjectIdleTime", ctx, key)
	ret0, _ := ret[0].(*redis.DurationCmd)
	return ret0
}

// ObjectIdleTime indicates an expected call of ObjectIdleTime.
func (mr *MockUniversalClientMockRecorder) ObjectIdleTime(ctx, key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ObjectIdleTime", reflect.TypeOf((*MockUniversalClient)(nil).ObjectIdleTime), ctx, key)
}

// ObjectRefCount mocks base method.
func (m *MockUniversalClient) ObjectRefCount(ctx context.Context, key string) *redis.IntCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ObjectRefCount", ctx, key)
	ret0, _ := ret[0].(*redis.IntCmd)
	return ret0
}

// ObjectRefCount indicates an expected call of ObjectRefCount.
func (mr *MockUniversalClientMockRecorder) ObjectRefCount(ctx, key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ObjectRefCount", reflect.TypeOf((*MockUniversalClient)(nil).ObjectRefCount), ctx, key)
}

// PExpire mocks base method.
func (m *MockUniversalClient) PExpire(ctx context.Context, key string, expiration time.Duration) *redis.BoolCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PExpire", ctx, key, expiration)
	ret0, _ := ret[0].(*redis.BoolCmd)
	return ret0
}

// PExpire indicates an expected call of PExpire.
func (mr *MockUniversalClientMockRecorder) PExpire(ctx, key, expiration interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PExpire", reflect.TypeOf((*MockUniversalClient)(nil).PExpire), ctx, key, expiration)
}

// PExpireAt mocks base method.
func (m *MockUniversalClient) PExpireAt(ctx context.Context, key string, tm time.Time) *redis.BoolCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PExpireAt", ctx, key, tm)
	ret0, _ := ret[0].(*redis.BoolCmd)
	return ret0
}

// PExpireAt indicates an expected call of PExpireAt.
func (mr *MockUniversalClientMockRecorder) PExpireAt(ctx, key, tm interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PExpireAt", reflect.TypeOf((*MockUniversalClient)(nil).PExpireAt), ctx, key, tm)
}

// PFAdd mocks base method.
func (m *MockUniversalClient) PFAdd(ctx context.Context, key string, els ...interface{}) *redis.IntCmd {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, key}
	for _, a := range els {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "PFAdd", varargs...)
	ret0, _ := ret[0].(*redis.IntCmd)
	return ret0
}

// PFAdd indicates an expected call of PFAdd.
func (mr *MockUniversalClientMockRecorder) PFAdd(ctx, key interface{}, els ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, key}, els...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PFAdd", reflect.TypeOf((*MockUniversalClient)(nil).PFAdd), varargs...)
}

// PFCount mocks base method.
func (m *MockUniversalClient) PFCount(ctx context.Context, keys ...string) *redis.IntCmd {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx}
	for _, a := range keys {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "PFCount", varargs...)
	ret0, _ := ret[0].(*redis.IntCmd)
	return ret0
}

// PFCount indicates an expected call of PFCount.
func (mr *MockUniversalClientMockRecorder) PFCount(ctx interface{}, keys ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx}, keys...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PFCount", reflect.TypeOf((*MockUniversalClient)(nil).PFCount), varargs...)
}

// PFMerge mocks base method.
func (m *MockUniversalClient) PFMerge(ctx context.Context, dest string, keys ...string) *redis.StatusCmd {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, dest}
	for _, a := range keys {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "PFMerge", varargs...)
	ret0, _ := ret[0].(*redis.StatusCmd)
	return ret0
}

// PFMerge indicates an expected call of PFMerge.
func (mr *MockUniversalClientMockRecorder) PFMerge(ctx, dest interface{}, keys ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, dest}, keys...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PFMerge", reflect.TypeOf((*MockUniversalClient)(nil).PFMerge), varargs...)
}

// PSubscribe mocks base method.
func (m *MockUniversalClient) PSubscribe(ctx context.Context, channels ...string) *redis.PubSub {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx}
	for _, a := range channels {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "PSubscribe", varargs...)
	ret0, _ := ret[0].(*redis.PubSub)
	return ret0
}

// PSubscribe indicates an expected call of PSubscribe.
func (mr *MockUniversalClientMockRecorder) PSubscribe(ctx interface{}, channels ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx}, channels...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PSubscribe", reflect.TypeOf((*MockUniversalClient)(nil).PSubscribe), varargs...)
}

// PTTL mocks base method.
func (m *MockUniversalClient) PTTL(ctx context.Context, key string) *redis.DurationCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PTTL", ctx, key)
	ret0, _ := ret[0].(*redis.DurationCmd)
	return ret0
}

// PTTL indicates an expected call of PTTL.
func (mr *MockUniversalClientMockRecorder) PTTL(ctx, key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PTTL", reflect.TypeOf((*MockUniversalClient)(nil).PTTL), ctx, key)
}

// Persist mocks base method.
func (m *MockUniversalClient) Persist(ctx context.Context, key string) *redis.BoolCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Persist", ctx, key)
	ret0, _ := ret[0].(*redis.BoolCmd)
	return ret0
}

// Persist indicates an expected call of Persist.
func (mr *MockUniversalClientMockRecorder) Persist(ctx, key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Persist", reflect.TypeOf((*MockUniversalClient)(nil).Persist), ctx, key)
}

// Ping mocks base method.
func (m *MockUniversalClient) Ping(ctx context.Context) *redis.StatusCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Ping", ctx)
	ret0, _ := ret[0].(*redis.StatusCmd)
	return ret0
}

// Ping indicates an expected call of Ping.
func (mr *MockUniversalClientMockRecorder) Ping(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Ping", reflect.TypeOf((*MockUniversalClient)(nil).Ping), ctx)
}

// Pipeline mocks base method.
func (m *MockUniversalClient) Pipeline() redis.Pipeliner {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Pipeline")
	ret0, _ := ret[0].(redis.Pipeliner)
	return ret0
}

// Pipeline indicates an expected call of Pipeline.
func (mr *MockUniversalClientMockRecorder) Pipeline() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Pipeline", reflect.TypeOf((*MockUniversalClient)(nil).Pipeline))
}

// Pipelined mocks base method.
func (m *MockUniversalClient) Pipelined(ctx context.Context, fn func(redis.Pipeliner) error) ([]redis.Cmder, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Pipelined", ctx, fn)
	ret0, _ := ret[0].([]redis.Cmder)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Pipelined indicates an expected call of Pipelined.
func (mr *MockUniversalClientMockRecorder) Pipelined(ctx, fn interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Pipelined", reflect.TypeOf((*MockUniversalClient)(nil).Pipelined), ctx, fn)
}

// PoolStats mocks base method.
func (m *MockUniversalClient) PoolStats() *redis.PoolStats {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PoolStats")
	ret0, _ := ret[0].(*redis.PoolStats)
	return ret0
}

// PoolStats indicates an expected call of PoolStats.
func (mr *MockUniversalClientMockRecorder) PoolStats() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PoolStats", reflect.TypeOf((*MockUniversalClient)(nil).PoolStats))
}

// Process mocks base method.
func (m *MockUniversalClient) Process(ctx context.Context, cmd redis.Cmder) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Process", ctx, cmd)
	ret0, _ := ret[0].(error)
	return ret0
}

// Process indicates an expected call of Process.
func (mr *MockUniversalClientMockRecorder) Process(ctx, cmd interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Process", reflect.TypeOf((*MockUniversalClient)(nil).Process), ctx, cmd)
}

// PubSubChannels mocks base method.
func (m *MockUniversalClient) PubSubChannels(ctx context.Context, pattern string) *redis.StringSliceCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PubSubChannels", ctx, pattern)
	ret0, _ := ret[0].(*redis.StringSliceCmd)
	return ret0
}

// PubSubChannels indicates an expected call of PubSubChannels.
func (mr *MockUniversalClientMockRecorder) PubSubChannels(ctx, pattern interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PubSubChannels", reflect.TypeOf((*MockUniversalClient)(nil).PubSubChannels), ctx, pattern)
}

// PubSubNumPat mocks base method.
func (m *MockUniversalClient) PubSubNumPat(ctx context.Context) *redis.IntCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PubSubNumPat", ctx)
	ret0, _ := ret[0].(*redis.IntCmd)
	return ret0
}

// PubSubNumPat indicates an expected call of PubSubNumPat.
func (mr *MockUniversalClientMockRecorder) PubSubNumPat(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PubSubNumPat", reflect.TypeOf((*MockUniversalClient)(nil).PubSubNumPat), ctx)
}

// PubSubNumSub mocks base method.
func (m *MockUniversalClient) PubSubNumSub(ctx context.Context, channels ...string) *redis.StringIntMapCmd {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx}
	for _, a := range channels {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "PubSubNumSub", varargs...)
	ret0, _ := ret[0].(*redis.StringIntMapCmd)
	return ret0
}

// PubSubNumSub indicates an expected call of PubSubNumSub.
func (mr *MockUniversalClientMockRecorder) PubSubNumSub(ctx interface{}, channels ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx}, channels...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PubSubNumSub", reflect.TypeOf((*MockUniversalClient)(nil).PubSubNumSub), varargs...)
}

// Publish mocks base method.
func (m *MockUniversalClient) Publish(ctx context.Context, channel string, message interface{}) *redis.IntCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Publish", ctx, channel, message)
	ret0, _ := ret[0].(*redis.IntCmd)
	return ret0
}

// Publish indicates an expected call of Publish.
func (mr *MockUniversalClientMockRecorder) Publish(ctx, channel, message interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Publish", reflect.TypeOf((*MockUniversalClient)(nil).Publish), ctx, channel, message)
}

// Quit mocks base method.
func (m *MockUniversalClient) Quit(ctx context.Context) *redis.StatusCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Quit", ctx)
	ret0, _ := ret[0].(*redis.StatusCmd)
	return ret0
}

// Quit indicates an expected call of Quit.
func (mr *MockUniversalClientMockRecorder) Quit(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Quit", reflect.TypeOf((*MockUniversalClient)(nil).Quit), ctx)
}

// RPop mocks base method.
func (m *MockUniversalClient) RPop(ctx context.Context, key string) *redis.StringCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RPop", ctx, key)
	ret0, _ := ret[0].(*redis.StringCmd)
	return ret0
}

// RPop indicates an expected call of RPop.
func (mr *MockUniversalClientMockRecorder) RPop(ctx, key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RPop", reflect.TypeOf((*MockUniversalClient)(nil).RPop), ctx, key)
}

// RPopCount mocks base method.
func (m *MockUniversalClient) RPopCount(ctx context.Context, key string, count int) *redis.StringSliceCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RPopCount", ctx, key, count)
	ret0, _ := ret[0].(*redis.StringSliceCmd)
	return ret0
}

// RPopCount indicates an expected call of RPopCount.
func (mr *MockUniversalClientMockRecorder) RPopCount(ctx, key, count interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RPopCount", reflect.TypeOf((*MockUniversalClient)(nil).RPopCount), ctx, key, count)
}

// RPopLPush mocks base method.
func (m *MockUniversalClient) RPopLPush(ctx context.Context, source, destination string) *redis.StringCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RPopLPush", ctx, source, destination)
	ret0, _ := ret[0].(*redis.StringCmd)
	return ret0
}

// RPopLPush indicates an expected call of RPopLPush.
func (mr *MockUniversalClientMockRecorder) RPopLPush(ctx, source, destination interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RPopLPush", reflect.TypeOf((*MockUniversalClient)(nil).RPopLPush), ctx, source, destination)
}

// RPush mocks base method.
func (m *MockUniversalClient) RPush(ctx context.Context, key string, values ...interface{}) *redis.IntCmd {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, key}
	for _, a := range values {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "RPush", varargs...)
	ret0, _ := ret[0].(*redis.IntCmd)
	return ret0
}

// RPush indicates an expected call of RPush.
func (mr *MockUniversalClientMockRecorder) RPush(ctx, key interface{}, values ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, key}, values...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RPush", reflect.TypeOf((*MockUniversalClient)(nil).RPush), varargs...)
}

// RPushX mocks base method.
func (m *MockUniversalClient) RPushX(ctx context.Context, key string, values ...interface{}) *redis.IntCmd {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, key}
	for _, a := range values {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "RPushX", varargs...)
	ret0, _ := ret[0].(*redis.IntCmd)
	return ret0
}

// RPushX indicates an expected call of RPushX.
func (mr *MockUniversalClientMockRecorder) RPushX(ctx, key interface{}, values ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, key}, values...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RPushX", reflect.TypeOf((*MockUniversalClient)(nil).RPushX), varargs...)
}

// RandomKey mocks base method.
func (m *MockUniversalClient) RandomKey(ctx context.Context) *redis.StringCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RandomKey", ctx)
	ret0, _ := ret[0].(*redis.StringCmd)
	return ret0
}

// RandomKey indicates an expected call of RandomKey.
func (mr *MockUniversalClientMockRecorder) RandomKey(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RandomKey", reflect.TypeOf((*MockUniversalClient)(nil).RandomKey), ctx)
}

// ReadOnly mocks base method.
func (m *MockUniversalClient) ReadOnly(ctx context.Context) *redis.StatusCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReadOnly", ctx)
	ret0, _ := ret[0].(*redis.StatusCmd)
	return ret0
}

// ReadOnly indicates an expected call of ReadOnly.
func (mr *MockUniversalClientMockRecorder) ReadOnly(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadOnly", reflect.TypeOf((*MockUniversalClient)(nil).ReadOnly), ctx)
}

// ReadWrite mocks base method.
func (m *MockUniversalClient) ReadWrite(ctx context.Context) *redis.StatusCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReadWrite", ctx)
	ret0, _ := ret[0].(*redis.StatusCmd)
	return ret0
}

// ReadWrite indicates an expected call of ReadWrite.
func (mr *MockUniversalClientMockRecorder) ReadWrite(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadWrite", reflect.TypeOf((*MockUniversalClient)(nil).ReadWrite), ctx)
}

// Rename mocks base method.
func (m *MockUniversalClient) Rename(ctx context.Context, key, newkey string) *redis.StatusCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Rename", ctx, key, newkey)
	ret0, _ := ret[0].(*redis.StatusCmd)
	return ret0
}

// Rename indicates an expected call of Rename.
func (mr *MockUniversalClientMockRecorder) Rename(ctx, key, newkey interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Rename", reflect.TypeOf((*MockUniversalClient)(nil).Rename), ctx, key, newkey)
}

// RenameNX mocks base method.
func (m *MockUniversalClient) RenameNX(ctx context.Context, key, newkey string) *redis.BoolCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RenameNX", ctx, key, newkey)
	ret0, _ := ret[0].(*redis.BoolCmd)
	return ret0
}

// RenameNX indicates an expected call of RenameNX.
func (mr *MockUniversalClientMockRecorder) RenameNX(ctx, key, newkey interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RenameNX", reflect.TypeOf((*MockUniversalClient)(nil).RenameNX), ctx, key, newkey)
}

// Restore mocks base method.
func (m *MockUniversalClient) Restore(ctx context.Context, key string, ttl time.Duration, value string) *redis.StatusCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Restore", ctx, key, ttl, value)
	ret0, _ := ret[0].(*redis.StatusCmd)
	return ret0
}

// Restore indicates an expected call of Restore.
func (mr *MockUniversalClientMockRecorder) Restore(ctx, key, ttl, value interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Restore", reflect.TypeOf((*MockUniversalClient)(nil).Restore), ctx, key, ttl, value)
}

// RestoreReplace mocks base method.
func (m *MockUniversalClient) RestoreReplace(ctx context.Context, key string, ttl time.Duration, value string) *redis.StatusCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RestoreReplace", ctx, key, ttl, value)
	ret0, _ := ret[0].(*redis.StatusCmd)
	return ret0
}

// RestoreReplace indicates an expected call of RestoreReplace.
func (mr *MockUniversalClientMockRecorder) RestoreReplace(ctx, key, ttl, value interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RestoreReplace", reflect.TypeOf((*MockUniversalClient)(nil).RestoreReplace), ctx, key, ttl, value)
}

// SAdd mocks base method.
func (m *MockUniversalClient) SAdd(ctx context.Context, key string, members ...interface{}) *redis.IntCmd {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, key}
	for _, a := range members {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "SAdd", varargs...)
	ret0, _ := ret[0].(*redis.IntCmd)
	return ret0
}

// SAdd indicates an expected call of SAdd.
func (mr *MockUniversalClientMockRecorder) SAdd(ctx, key interface{}, members ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, key}, members...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SAdd", reflect.TypeOf((*MockUniversalClient)(nil).SAdd), varargs...)
}

// SCard mocks base method.
func (m *MockUniversalClient) SCard(ctx context.Context, key string) *redis.IntCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SCard", ctx, key)
	ret0, _ := ret[0].(*redis.IntCmd)
	return ret0
}

// SCard indicates an expected call of SCard.
func (mr *MockUniversalClientMockRecorder) SCard(ctx, key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SCard", reflect.TypeOf((*MockUniversalClient)(nil).SCard), ctx, key)
}

// SDiff mocks base method.
func (m *MockUniversalClient) SDiff(ctx context.Context, keys ...string) *redis.StringSliceCmd {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx}
	for _, a := range keys {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "SDiff", varargs...)
	ret0, _ := ret[0].(*redis.StringSliceCmd)
	return ret0
}

// SDiff indicates an expected call of SDiff.
func (mr *MockUniversalClientMockRecorder) SDiff(ctx interface{}, keys ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx}, keys...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SDiff", reflect.TypeOf((*MockUniversalClient)(nil).SDiff), varargs...)
}

// SDiffStore mocks base method.
func (m *MockUniversalClient) SDiffStore(ctx context.Context, destination string, keys ...string) *redis.IntCmd {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, destination}
	for _, a := range keys {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "SDiffStore", varargs...)
	ret0, _ := ret[0].(*redis.IntCmd)
	return ret0
}

// SDiffStore indicates an expected call of SDiffStore.
func (mr *MockUniversalClientMockRecorder) SDiffStore(ctx, destination interface{}, keys ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, destination}, keys...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SDiffStore", reflect.TypeOf((*MockUniversalClient)(nil).SDiffStore), varargs...)
}

// SInter mocks base method.
func (m *MockUniversalClient) SInter(ctx context.Context, keys ...string) *redis.StringSliceCmd {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx}
	for _, a := range keys {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "SInter", varargs...)
	ret0, _ := ret[0].(*redis.StringSliceCmd)
	return ret0
}

// SInter indicates an expected call of SInter.
func (mr *MockUniversalClientMockRecorder) SInter(ctx interface{}, keys ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx}, keys...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SInter", reflect.TypeOf((*MockUniversalClient)(nil).SInter), varargs...)
}

// SInterStore mocks base method.
func (m *MockUniversalClient) SInterStore(ctx context.Context, destination string, keys ...string) *redis.IntCmd {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, destination}
	for _, a := range keys {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "SInterStore", varargs...)
	ret0, _ := ret[0].(*redis.IntCmd)
	return ret0
}

// SInterStore indicates an expected call of SInterStore.
func (mr *MockUniversalClientMockRecorder) SInterStore(ctx, destination interface{}, keys ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, destination}, keys...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SInterStore", reflect.TypeOf((*MockUniversalClient)(nil).SInterStore), varargs...)
}

// SIsMember mocks base method.
func (m *MockUniversalClient) SIsMember(ctx context.Context, key string, member interface{}) *redis.BoolCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SIsMember", ctx, key, member)
	ret0, _ := ret[0].(*redis.BoolCmd)
	return ret0
}

// SIsMember indicates an expected call of SIsMember.
func (mr *MockUniversalClientMockRecorder) SIsMember(ctx, key, member interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SIsMember", reflect.TypeOf((*MockUniversalClient)(nil).SIsMember), ctx, key, member)
}

// SMIsMember mocks base method.
func (m *MockUniversalClient) SMIsMember(ctx context.Context, key string, members ...interface{}) *redis.BoolSliceCmd {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, key}
	for _, a := range members {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "SMIsMember", varargs...)
	ret0, _ := ret[0].(*redis.BoolSliceCmd)
	return ret0
}

// SMIsMember indicates an expected call of SMIsMember.
func (mr *MockUniversalClientMockRecorder) SMIsMember(ctx, key interface{}, members ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, key}, members...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SMIsMember", reflect.TypeOf((*MockUniversalClient)(nil).SMIsMember), varargs...)
}

// SMembers mocks base method.
func (m *MockUniversalClient) SMembers(ctx context.Context, key string) *redis.StringSliceCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SMembers", ctx, key)
	ret0, _ := ret[0].(*redis.StringSliceCmd)
	return ret0
}

// SMembers indicates an expected call of SMembers.
func (mr *MockUniversalClientMockRecorder) SMembers(ctx, key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SMembers", reflect.TypeOf((*MockUniversalClient)(nil).SMembers), ctx, key)
}

// SMembersMap mocks base method.
func (m *MockUniversalClient) SMembersMap(ctx context.Context, key string) *redis.StringStructMapCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SMembersMap", ctx, key)
	ret0, _ := ret[0].(*redis.StringStructMapCmd)
	return ret0
}

// SMembersMap indicates an expected call of SMembersMap.
func (mr *MockUniversalClientMockRecorder) SMembersMap(ctx, key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SMembersMap", reflect.TypeOf((*MockUniversalClient)(nil).SMembersMap), ctx, key)
}

// SMove mocks base method.
func (m *MockUniversalClient) SMove(ctx context.Context, source, destination string, member interface{}) *redis.BoolCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SMove", ctx, source, destination, member)
	ret0, _ := ret[0].(*redis.BoolCmd)
	return ret0
}

// SMove indicates an expected call of SMove.
func (mr *MockUniversalClientMockRecorder) SMove(ctx, source, destination, member interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SMove", reflect.TypeOf((*MockUniversalClient)(nil).SMove), ctx, source, destination, member)
}

// SPop mocks base method.
func (m *MockUniversalClient) SPop(ctx context.Context, key string) *redis.StringCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SPop", ctx, key)
	ret0, _ := ret[0].(*redis.StringCmd)
	return ret0
}

// SPop indicates an expected call of SPop.
func (mr *MockUniversalClientMockRecorder) SPop(ctx, key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SPop", reflect.TypeOf((*MockUniversalClient)(nil).SPop), ctx, key)
}

// SPopN mocks base method.
func (m *MockUniversalClient) SPopN(ctx context.Context, key string, count int64) *redis.StringSliceCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SPopN", ctx, key, count)
	ret0, _ := ret[0].(*redis.StringSliceCmd)
	return ret0
}

// SPopN indicates an expected call of SPopN.
func (mr *MockUniversalClientMockRecorder) SPopN(ctx, key, count interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SPopN", reflect.TypeOf((*MockUniversalClient)(nil).SPopN), ctx, key, count)
}

// SRandMember mocks base method.
func (m *MockUniversalClient) SRandMember(ctx context.Context, key string) *redis.StringCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SRandMember", ctx, key)
	ret0, _ := ret[0].(*redis.StringCmd)
	return ret0
}

// SRandMember indicates an expected call of SRandMember.
func (mr *MockUniversalClientMockRecorder) SRandMember(ctx, key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SRandMember", reflect.TypeOf((*MockUniversalClient)(nil).SRandMember), ctx, key)
}

// SRandMemberN mocks base method.
func (m *MockUniversalClient) SRandMemberN(ctx context.Context, key string, count int64) *redis.StringSliceCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SRandMemberN", ctx, key, count)
	ret0, _ := ret[0].(*redis.StringSliceCmd)
	return ret0
}

// SRandMemberN indicates an expected call of SRandMemberN.
func (mr *MockUniversalClientMockRecorder) SRandMemberN(ctx, key, count interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SRandMemberN", reflect.TypeOf((*MockUniversalClient)(nil).SRandMemberN), ctx, key, count)
}

// SRem mocks base method.
func (m *MockUniversalClient) SRem(ctx context.Context, key string, members ...interface{}) *redis.IntCmd {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, key}
	for _, a := range members {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "SRem", varargs...)
	ret0, _ := ret[0].(*redis.IntCmd)
	return ret0
}

// SRem indicates an expected call of SRem.
func (mr *MockUniversalClientMockRecorder) SRem(ctx, key interface{}, members ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, key}, members...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SRem", reflect.TypeOf((*MockUniversalClient)(nil).SRem), varargs...)
}

// SScan mocks base method.
func (m *MockUniversalClient) SScan(ctx context.Context, key string, cursor uint64, match string, count int64) *redis.ScanCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SScan", ctx, key, cursor, match, count)
	ret0, _ := ret[0].(*redis.ScanCmd)
	return ret0
}

// SScan indicates an expected call of SScan.
func (mr *MockUniversalClientMockRecorder) SScan(ctx, key, cursor, match, count interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SScan", reflect.TypeOf((*MockUniversalClient)(nil).SScan), ctx, key, cursor, match, count)
}

// SUnion mocks base method.
func (m *MockUniversalClient) SUnion(ctx context.Context, keys ...string) *redis.StringSliceCmd {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx}
	for _, a := range keys {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "SUnion", varargs...)
	ret0, _ := ret[0].(*redis.StringSliceCmd)
	return ret0
}

// SUnion indicates an expected call of SUnion.
func (mr *MockUniversalClientMockRecorder) SUnion(ctx interface{}, keys ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx}, keys...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SUnion", reflect.TypeOf((*MockUniversalClient)(nil).SUnion), varargs...)
}

// SUnionStore mocks base method.
func (m *MockUniversalClient) SUnionStore(ctx context.Context, destination string, keys ...string) *redis.IntCmd {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, destination}
	for _, a := range keys {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "SUnionStore", varargs...)
	ret0, _ := ret[0].(*redis.IntCmd)
	return ret0
}

// SUnionStore indicates an expected call of SUnionStore.
func (mr *MockUniversalClientMockRecorder) SUnionStore(ctx, destination interface{}, keys ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, destination}, keys...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SUnionStore", reflect.TypeOf((*MockUniversalClient)(nil).SUnionStore), varargs...)
}

// Save mocks base method.
func (m *MockUniversalClient) Save(ctx context.Context) *redis.StatusCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Save", ctx)
	ret0, _ := ret[0].(*redis.StatusCmd)
	return ret0
}

// Save indicates an expected call of Save.
func (mr *MockUniversalClientMockRecorder) Save(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Save", reflect.TypeOf((*MockUniversalClient)(nil).Save), ctx)
}

// Scan mocks base method.
func (m *MockUniversalClient) Scan(ctx context.Context, cursor uint64, match string, count int64) *redis.ScanCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Scan", ctx, cursor, match, count)
	ret0, _ := ret[0].(*redis.ScanCmd)
	return ret0
}

// Scan indicates an expected call of Scan.
func (mr *MockUniversalClientMockRecorder) Scan(ctx, cursor, match, count interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Scan", reflect.TypeOf((*MockUniversalClient)(nil).Scan), ctx, cursor, match, count)
}

// ScanType mocks base method.
func (m *MockUniversalClient) ScanType(ctx context.Context, cursor uint64, match string, count int64, keyType string) *redis.ScanCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ScanType", ctx, cursor, match, count, keyType)
	ret0, _ := ret[0].(*redis.ScanCmd)
	return ret0
}

// ScanType indicates an expected call of ScanType.
func (mr *MockUniversalClientMockRecorder) ScanType(ctx, cursor, match, count, keyType interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ScanType", reflect.TypeOf((*MockUniversalClient)(nil).ScanType), ctx, cursor, match, count, keyType)
}

// ScriptExists mocks base method.
func (m *MockUniversalClient) ScriptExists(ctx context.Context, hashes ...string) *redis.BoolSliceCmd {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx}
	for _, a := range hashes {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ScriptExists", varargs...)
	ret0, _ := ret[0].(*redis.BoolSliceCmd)
	return ret0
}

// ScriptExists indicates an expected call of ScriptExists.
func (mr *MockUniversalClientMockRecorder) ScriptExists(ctx interface{}, hashes ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx}, hashes...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ScriptExists", reflect.TypeOf((*MockUniversalClient)(nil).ScriptExists), varargs...)
}

// ScriptFlush mocks base method.
func (m *MockUniversalClient) ScriptFlush(ctx context.Context) *redis.StatusCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ScriptFlush", ctx)
	ret0, _ := ret[0].(*redis.StatusCmd)
	return ret0
}

// ScriptFlush indicates an expected call of ScriptFlush.
func (mr *MockUniversalClientMockRecorder) ScriptFlush(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ScriptFlush", reflect.TypeOf((*MockUniversalClient)(nil).ScriptFlush), ctx)
}

// ScriptKill mocks base method.
func (m *MockUniversalClient) ScriptKill(ctx context.Context) *redis.StatusCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ScriptKill", ctx)
	ret0, _ := ret[0].(*redis.StatusCmd)
	return ret0
}

// ScriptKill indicates an expected call of ScriptKill.
func (mr *MockUniversalClientMockRecorder) ScriptKill(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ScriptKill", reflect.TypeOf((*MockUniversalClient)(nil).ScriptKill), ctx)
}

// ScriptLoad mocks base method.
func (m *MockUniversalClient) ScriptLoad(ctx context.Context, script string) *redis.StringCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ScriptLoad", ctx, script)
	ret0, _ := ret[0].(*redis.StringCmd)
	return ret0
}

// ScriptLoad indicates an expected call of ScriptLoad.
func (mr *MockUniversalClientMockRecorder) ScriptLoad(ctx, script interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ScriptLoad", reflect.TypeOf((*MockUniversalClient)(nil).ScriptLoad), ctx, script)
}

// Set mocks base method.
func (m *MockUniversalClient) Set(ctx context.Context, key string, value interface{}, expiration time.Duration) *redis.StatusCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Set", ctx, key, value, expiration)
	ret0, _ := ret[0].(*redis.StatusCmd)
	return ret0
}

// Set indicates an expected call of Set.
func (mr *MockUniversalClientMockRecorder) Set(ctx, key, value, expiration interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Set", reflect.TypeOf((*MockUniversalClient)(nil).Set), ctx, key, value, expiration)
}

// SetArgs mocks base method.
func (m *MockUniversalClient) SetArgs(ctx context.Context, key string, value interface{}, a redis.SetArgs) *redis.StatusCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetArgs", ctx, key, value, a)
	ret0, _ := ret[0].(*redis.StatusCmd)
	return ret0
}

// SetArgs indicates an expected call of SetArgs.
func (mr *MockUniversalClientMockRecorder) SetArgs(ctx, key, value, a interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetArgs", reflect.TypeOf((*MockUniversalClient)(nil).SetArgs), ctx, key, value, a)
}

// SetBit mocks base method.
func (m *MockUniversalClient) SetBit(ctx context.Context, key string, offset int64, value int) *redis.IntCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetBit", ctx, key, offset, value)
	ret0, _ := ret[0].(*redis.IntCmd)
	return ret0
}

// SetBit indicates an expected call of SetBit.
func (mr *MockUniversalClientMockRecorder) SetBit(ctx, key, offset, value interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetBit", reflect.TypeOf((*MockUniversalClient)(nil).SetBit), ctx, key, offset, value)
}

// SetEX mocks base method.
func (m *MockUniversalClient) SetEX(ctx context.Context, key string, value interface{}, expiration time.Duration) *redis.StatusCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetEX", ctx, key, value, expiration)
	ret0, _ := ret[0].(*redis.StatusCmd)
	return ret0
}

// SetEX indicates an expected call of SetEX.
func (mr *MockUniversalClientMockRecorder) SetEX(ctx, key, value, expiration interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetEX", reflect.TypeOf((*MockUniversalClient)(nil).SetEX), ctx, key, value, expiration)
}

// SetNX mocks base method.
func (m *MockUniversalClient) SetNX(ctx context.Context, key string, value interface{}, expiration time.Duration) *redis.BoolCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetNX", ctx, key, value, expiration)
	ret0, _ := ret[0].(*redis.BoolCmd)
	return ret0
}

// SetNX indicates an expected call of SetNX.
func (mr *MockUniversalClientMockRecorder) SetNX(ctx, key, value, expiration interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetNX", reflect.TypeOf((*MockUniversalClient)(nil).SetNX), ctx, key, value, expiration)
}

// SetRange mocks base method.
func (m *MockUniversalClient) SetRange(ctx context.Context, key string, offset int64, value string) *redis.IntCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetRange", ctx, key, offset, value)
	ret0, _ := ret[0].(*redis.IntCmd)
	return ret0
}

// SetRange indicates an expected call of SetRange.
func (mr *MockUniversalClientMockRecorder) SetRange(ctx, key, offset, value interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetRange", reflect.TypeOf((*MockUniversalClient)(nil).SetRange), ctx, key, offset, value)
}

// SetXX mocks base method.
func (m *MockUniversalClient) SetXX(ctx context.Context, key string, value interface{}, expiration time.Duration) *redis.BoolCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetXX", ctx, key, value, expiration)
	ret0, _ := ret[0].(*redis.BoolCmd)
	return ret0
}

// SetXX indicates an expected call of SetXX.
func (mr *MockUniversalClientMockRecorder) SetXX(ctx, key, value, expiration interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetXX", reflect.TypeOf((*MockUniversalClient)(nil).SetXX), ctx, key, value, expiration)
}

// Shutdown mocks base method.
func (m *MockUniversalClient) Shutdown(ctx context.Context) *redis.StatusCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Shutdown", ctx)
	ret0, _ := ret[0].(*redis.StatusCmd)
	return ret0
}

// Shutdown indicates an expected call of Shutdown.
func (mr *MockUniversalClientMockRecorder) Shutdown(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Shutdown", reflect.TypeOf((*MockUniversalClient)(nil).Shutdown), ctx)
}

// ShutdownNoSave mocks base method.
func (m *MockUniversalClient) ShutdownNoSave(ctx context.Context) *redis.StatusCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ShutdownNoSave", ctx)
	ret0, _ := ret[0].(*redis.StatusCmd)
	return ret0
}

// ShutdownNoSave indicates an expected call of ShutdownNoSave.
func (mr *MockUniversalClientMockRecorder) ShutdownNoSave(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ShutdownNoSave", reflect.TypeOf((*MockUniversalClient)(nil).ShutdownNoSave), ctx)
}

// ShutdownSave mocks base method.
func (m *MockUniversalClient) ShutdownSave(ctx context.Context) *redis.StatusCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ShutdownSave", ctx)
	ret0, _ := ret[0].(*redis.StatusCmd)
	return ret0
}

// ShutdownSave indicates an expected call of ShutdownSave.
func (mr *MockUniversalClientMockRecorder) ShutdownSave(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ShutdownSave", reflect.TypeOf((*MockUniversalClient)(nil).ShutdownSave), ctx)
}

// SlaveOf mocks base method.
func (m *MockUniversalClient) SlaveOf(ctx context.Context, host, port string) *redis.StatusCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SlaveOf", ctx, host, port)
	ret0, _ := ret[0].(*redis.StatusCmd)
	return ret0
}

// SlaveOf indicates an expected call of SlaveOf.
func (mr *MockUniversalClientMockRecorder) SlaveOf(ctx, host, port interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SlaveOf", reflect.TypeOf((*MockUniversalClient)(nil).SlaveOf), ctx, host, port)
}

// Sort mocks base method.
func (m *MockUniversalClient) Sort(ctx context.Context, key string, sort *redis.Sort) *redis.StringSliceCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Sort", ctx, key, sort)
	ret0, _ := ret[0].(*redis.StringSliceCmd)
	return ret0
}

// Sort indicates an expected call of Sort.
func (mr *MockUniversalClientMockRecorder) Sort(ctx, key, sort interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Sort", reflect.TypeOf((*MockUniversalClient)(nil).Sort), ctx, key, sort)
}

// SortInterfaces mocks base method.
func (m *MockUniversalClient) SortInterfaces(ctx context.Context, key string, sort *redis.Sort) *redis.SliceCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SortInterfaces", ctx, key, sort)
	ret0, _ := ret[0].(*redis.SliceCmd)
	return ret0
}

// SortInterfaces indicates an expected call of SortInterfaces.
func (mr *MockUniversalClientMockRecorder) SortInterfaces(ctx, key, sort interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SortInterfaces", reflect.TypeOf((*MockUniversalClient)(nil).SortInterfaces), ctx, key, sort)
}

// SortStore mocks base method.
func (m *MockUniversalClient) SortStore(ctx context.Context, key, store string, sort *redis.Sort) *redis.IntCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SortStore", ctx, key, store, sort)
	ret0, _ := ret[0].(*redis.IntCmd)
	return ret0
}

// SortStore indicates an expected call of SortStore.
func (mr *MockUniversalClientMockRecorder) SortStore(ctx, key, store, sort interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SortStore", reflect.TypeOf((*MockUniversalClient)(nil).SortStore), ctx, key, store, sort)
}

// StrLen mocks base method.
func (m *MockUniversalClient) StrLen(ctx context.Context, key string) *redis.IntCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StrLen", ctx, key)
	ret0, _ := ret[0].(*redis.IntCmd)
	return ret0
}

// StrLen indicates an expected call of StrLen.
func (mr *MockUniversalClientMockRecorder) StrLen(ctx, key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StrLen", reflect.TypeOf((*MockUniversalClient)(nil).StrLen), ctx, key)
}

// Subscribe mocks base method.
func (m *MockUniversalClient) Subscribe(ctx context.Context, channels ...string) *redis.PubSub {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx}
	for _, a := range channels {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Subscribe", varargs...)
	ret0, _ := ret[0].(*redis.PubSub)
	return ret0
}

// Subscribe indicates an expected call of Subscribe.
func (mr *MockUniversalClientMockRecorder) Subscribe(ctx interface{}, channels ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx}, channels...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Subscribe", reflect.TypeOf((*MockUniversalClient)(nil).Subscribe), varargs...)
}

// TTL mocks base method.
func (m *MockUniversalClient) TTL(ctx context.Context, key string) *redis.DurationCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TTL", ctx, key)
	ret0, _ := ret[0].(*redis.DurationCmd)
	return ret0
}

// TTL indicates an expected call of TTL.
func (mr *MockUniversalClientMockRecorder) TTL(ctx, key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TTL", reflect.TypeOf((*MockUniversalClient)(nil).TTL), ctx, key)
}

// Time mocks base method.
func (m *MockUniversalClient) Time(ctx context.Context) *redis.TimeCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Time", ctx)
	ret0, _ := ret[0].(*redis.TimeCmd)
	return ret0
}

// Time indicates an expected call of Time.
func (mr *MockUniversalClientMockRecorder) Time(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Time", reflect.TypeOf((*MockUniversalClient)(nil).Time), ctx)
}

// Touch mocks base method.
func (m *MockUniversalClient) Touch(ctx context.Context, keys ...string) *redis.IntCmd {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx}
	for _, a := range keys {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Touch", varargs...)
	ret0, _ := ret[0].(*redis.IntCmd)
	return ret0
}

// Touch indicates an expected call of Touch.
func (mr *MockUniversalClientMockRecorder) Touch(ctx interface{}, keys ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx}, keys...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Touch", reflect.TypeOf((*MockUniversalClient)(nil).Touch), varargs...)
}

// TxPipeline mocks base method.
func (m *MockUniversalClient) TxPipeline() redis.Pipeliner {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TxPipeline")
	ret0, _ := ret[0].(redis.Pipeliner)
	return ret0
}

// TxPipeline indicates an expected call of TxPipeline.
func (mr *MockUniversalClientMockRecorder) TxPipeline() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TxPipeline", reflect.TypeOf((*MockUniversalClient)(nil).TxPipeline))
}

// TxPipelined mocks base method.
func (m *MockUniversalClient) TxPipelined(ctx context.Context, fn func(redis.Pipeliner) error) ([]redis.Cmder, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TxPipelined", ctx, fn)
	ret0, _ := ret[0].([]redis.Cmder)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// TxPipelined indicates an expected call of TxPipelined.
func (mr *MockUniversalClientMockRecorder) TxPipelined(ctx, fn interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TxPipelined", reflect.TypeOf((*MockUniversalClient)(nil).TxPipelined), ctx, fn)
}

// Type mocks base method.
func (m *MockUniversalClient) Type(ctx context.Context, key string) *redis.StatusCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Type", ctx, key)
	ret0, _ := ret[0].(*redis.StatusCmd)
	return ret0
}

// Type indicates an expected call of Type.
func (mr *MockUniversalClientMockRecorder) Type(ctx, key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Type", reflect.TypeOf((*MockUniversalClient)(nil).Type), ctx, key)
}

// Unlink mocks base method.
func (m *MockUniversalClient) Unlink(ctx context.Context, keys ...string) *redis.IntCmd {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx}
	for _, a := range keys {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Unlink", varargs...)
	ret0, _ := ret[0].(*redis.IntCmd)
	return ret0
}

// Unlink indicates an expected call of Unlink.
func (mr *MockUniversalClientMockRecorder) Unlink(ctx interface{}, keys ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx}, keys...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Unlink", reflect.TypeOf((*MockUniversalClient)(nil).Unlink), varargs...)
}

// Watch mocks base method.
func (m *MockUniversalClient) Watch(ctx context.Context, fn func(*redis.Tx) error, keys ...string) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, fn}
	for _, a := range keys {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Watch", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// Watch indicates an expected call of Watch.
func (mr *MockUniversalClientMockRecorder) Watch(ctx, fn interface{}, keys ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, fn}, keys...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Watch", reflect.TypeOf((*MockUniversalClient)(nil).Watch), varargs...)
}

// XAck mocks base method.
func (m *MockUniversalClient) XAck(ctx context.Context, stream, group string, ids ...string) *redis.IntCmd {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, stream, group}
	for _, a := range ids {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "XAck", varargs...)
	ret0, _ := ret[0].(*redis.IntCmd)
	return ret0
}

// XAck indicates an expected call of XAck.
func (mr *MockUniversalClientMockRecorder) XAck(ctx, stream, group interface{}, ids ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, stream, group}, ids...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "XAck", reflect.TypeOf((*MockUniversalClient)(nil).XAck), varargs...)
}

// XAdd mocks base method.
func (m *MockUniversalClient) XAdd(ctx context.Context, a *redis.XAddArgs) *redis.StringCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "XAdd", ctx, a)
	ret0, _ := ret[0].(*redis.StringCmd)
	return ret0
}

// XAdd indicates an expected call of XAdd.
func (mr *MockUniversalClientMockRecorder) XAdd(ctx, a interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "XAdd", reflect.TypeOf((*MockUniversalClient)(nil).XAdd), ctx, a)
}

// XAutoClaim mocks base method.
func (m *MockUniversalClient) XAutoClaim(ctx context.Context, a *redis.XAutoClaimArgs) *redis.XAutoClaimCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "XAutoClaim", ctx, a)
	ret0, _ := ret[0].(*redis.XAutoClaimCmd)
	return ret0
}

// XAutoClaim indicates an expected call of XAutoClaim.
func (mr *MockUniversalClientMockRecorder) XAutoClaim(ctx, a interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "XAutoClaim", reflect.TypeOf((*MockUniversalClient)(nil).XAutoClaim), ctx, a)
}

// XAutoClaimJustID mocks base method.
func (m *MockUniversalClient) XAutoClaimJustID(ctx context.Context, a *redis.XAutoClaimArgs) *redis.XAutoClaimJustIDCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "XAutoClaimJustID", ctx, a)
	ret0, _ := ret[0].(*redis.XAutoClaimJustIDCmd)
	return ret0
}

// XAutoClaimJustID indicates an expected call of XAutoClaimJustID.
func (mr *MockUniversalClientMockRecorder) XAutoClaimJustID(ctx, a interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "XAutoClaimJustID", reflect.TypeOf((*MockUniversalClient)(nil).XAutoClaimJustID), ctx, a)
}

// XClaim mocks base method.
func (m *MockUniversalClient) XClaim(ctx context.Context, a *redis.XClaimArgs) *redis.XMessageSliceCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "XClaim", ctx, a)
	ret0, _ := ret[0].(*redis.XMessageSliceCmd)
	return ret0
}

// XClaim indicates an expected call of XClaim.
func (mr *MockUniversalClientMockRecorder) XClaim(ctx, a interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "XClaim", reflect.TypeOf((*MockUniversalClient)(nil).XClaim), ctx, a)
}

// XClaimJustID mocks base method.
func (m *MockUniversalClient) XClaimJustID(ctx context.Context, a *redis.XClaimArgs) *redis.StringSliceCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "XClaimJustID", ctx, a)
	ret0, _ := ret[0].(*redis.StringSliceCmd)
	return ret0
}

// XClaimJustID indicates an expected call of XClaimJustID.
func (mr *MockUniversalClientMockRecorder) XClaimJustID(ctx, a interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "XClaimJustID", reflect.TypeOf((*MockUniversalClient)(nil).XClaimJustID), ctx, a)
}

// XDel mocks base method.
func (m *MockUniversalClient) XDel(ctx context.Context, stream string, ids ...string) *redis.IntCmd {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, stream}
	for _, a := range ids {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "XDel", varargs...)
	ret0, _ := ret[0].(*redis.IntCmd)
	return ret0
}

// XDel indicates an expected call of XDel.
func (mr *MockUniversalClientMockRecorder) XDel(ctx, stream interface{}, ids ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, stream}, ids...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "XDel", reflect.TypeOf((*MockUniversalClient)(nil).XDel), varargs...)
}

// XGroupCreate mocks base method.
func (m *MockUniversalClient) XGroupCreate(ctx context.Context, stream, group, start string) *redis.StatusCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "XGroupCreate", ctx, stream, group, start)
	ret0, _ := ret[0].(*redis.StatusCmd)
	return ret0
}

// XGroupCreate indicates an expected call of XGroupCreate.
func (mr *MockUniversalClientMockRecorder) XGroupCreate(ctx, stream, group, start interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "XGroupCreate", reflect.TypeOf((*MockUniversalClient)(nil).XGroupCreate), ctx, stream, group, start)
}

// XGroupCreateConsumer mocks base method.
func (m *MockUniversalClient) XGroupCreateConsumer(ctx context.Context, stream, group, consumer string) *redis.IntCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "XGroupCreateConsumer", ctx, stream, group, consumer)
	ret0, _ := ret[0].(*redis.IntCmd)
	return ret0
}

// XGroupCreateConsumer indicates an expected call of XGroupCreateConsumer.
func (mr *MockUniversalClientMockRecorder) XGroupCreateConsumer(ctx, stream, group, consumer interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "XGroupCreateConsumer", reflect.TypeOf((*MockUniversalClient)(nil).XGroupCreateConsumer), ctx, stream, group, consumer)
}

// XGroupCreateMkStream mocks base method.
func (m *MockUniversalClient) XGroupCreateMkStream(ctx context.Context, stream, group, start string) *redis.StatusCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "XGroupCreateMkStream", ctx, stream, group, start)
	ret0, _ := ret[0].(*redis.StatusCmd)
	return ret0
}

// XGroupCreateMkStream indicates an expected call of XGroupCreateMkStream.
func (mr *MockUniversalClientMockRecorder) XGroupCreateMkStream(ctx, stream, group, start interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "XGroupCreateMkStream", reflect.TypeOf((*MockUniversalClient)(nil).XGroupCreateMkStream), ctx, stream, group, start)
}

// XGroupDelConsumer mocks base method.
func (m *MockUniversalClient) XGroupDelConsumer(ctx context.Context, stream, group, consumer string) *redis.IntCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "XGroupDelConsumer", ctx, stream, group, consumer)
	ret0, _ := ret[0].(*redis.IntCmd)
	return ret0
}

// XGroupDelConsumer indicates an expected call of XGroupDelConsumer.
func (mr *MockUniversalClientMockRecorder) XGroupDelConsumer(ctx, stream, group, consumer interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "XGroupDelConsumer", reflect.TypeOf((*MockUniversalClient)(nil).XGroupDelConsumer), ctx, stream, group, consumer)
}

// XGroupDestroy mocks base method.
func (m *MockUniversalClient) XGroupDestroy(ctx context.Context, stream, group string) *redis.IntCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "XGroupDestroy", ctx, stream, group)
	ret0, _ := ret[0].(*redis.IntCmd)
	return ret0
}

// XGroupDestroy indicates an expected call of XGroupDestroy.
func (mr *MockUniversalClientMockRecorder) XGroupDestroy(ctx, stream, group interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "XGroupDestroy", reflect.TypeOf((*MockUniversalClient)(nil).XGroupDestroy), ctx, stream, group)
}

// XGroupSetID mocks base method.
func (m *MockUniversalClient) XGroupSetID(ctx context.Context, stream, group, start string) *redis.StatusCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "XGroupSetID", ctx, stream, group, start)
	ret0, _ := ret[0].(*redis.StatusCmd)
	return ret0
}

// XGroupSetID indicates an expected call of XGroupSetID.
func (mr *MockUniversalClientMockRecorder) XGroupSetID(ctx, stream, group, start interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "XGroupSetID", reflect.TypeOf((*MockUniversalClient)(nil).XGroupSetID), ctx, stream, group, start)
}

// XInfoConsumers mocks base method.
func (m *MockUniversalClient) XInfoConsumers(ctx context.Context, key, group string) *redis.XInfoConsumersCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "XInfoConsumers", ctx, key, group)
	ret0, _ := ret[0].(*redis.XInfoConsumersCmd)
	return ret0
}

// XInfoConsumers indicates an expected call of XInfoConsumers.
func (mr *MockUniversalClientMockRecorder) XInfoConsumers(ctx, key, group interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "XInfoConsumers", reflect.TypeOf((*MockUniversalClient)(nil).XInfoConsumers), ctx, key, group)
}

// XInfoGroups mocks base method.
func (m *MockUniversalClient) XInfoGroups(ctx context.Context, key string) *redis.XInfoGroupsCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "XInfoGroups", ctx, key)
	ret0, _ := ret[0].(*redis.XInfoGroupsCmd)
	return ret0
}

// XInfoGroups indicates an expected call of XInfoGroups.
func (mr *MockUniversalClientMockRecorder) XInfoGroups(ctx, key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "XInfoGroups", reflect.TypeOf((*MockUniversalClient)(nil).XInfoGroups), ctx, key)
}

// XInfoStream mocks base method.
func (m *MockUniversalClient) XInfoStream(ctx context.Context, key string) *redis.XInfoStreamCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "XInfoStream", ctx, key)
	ret0, _ := ret[0].(*redis.XInfoStreamCmd)
	return ret0
}

// XInfoStream indicates an expected call of XInfoStream.
func (mr *MockUniversalClientMockRecorder) XInfoStream(ctx, key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "XInfoStream", reflect.TypeOf((*MockUniversalClient)(nil).XInfoStream), ctx, key)
}

// XInfoStreamFull mocks base method.
func (m *MockUniversalClient) XInfoStreamFull(ctx context.Context, key string, count int) *redis.XInfoStreamFullCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "XInfoStreamFull", ctx, key, count)
	ret0, _ := ret[0].(*redis.XInfoStreamFullCmd)
	return ret0
}

// XInfoStreamFull indicates an expected call of XInfoStreamFull.
func (mr *MockUniversalClientMockRecorder) XInfoStreamFull(ctx, key, count interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "XInfoStreamFull", reflect.TypeOf((*MockUniversalClient)(nil).XInfoStreamFull), ctx, key, count)
}

// XLen mocks base method.
func (m *MockUniversalClient) XLen(ctx context.Context, stream string) *redis.IntCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "XLen", ctx, stream)
	ret0, _ := ret[0].(*redis.IntCmd)
	return ret0
}

// XLen indicates an expected call of XLen.
func (mr *MockUniversalClientMockRecorder) XLen(ctx, stream interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "XLen", reflect.TypeOf((*MockUniversalClient)(nil).XLen), ctx, stream)
}

// XPending mocks base method.
func (m *MockUniversalClient) XPending(ctx context.Context, stream, group string) *redis.XPendingCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "XPending", ctx, stream, group)
	ret0, _ := ret[0].(*redis.XPendingCmd)
	return ret0
}

// XPending indicates an expected call of XPending.
func (mr *MockUniversalClientMockRecorder) XPending(ctx, stream, group interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "XPending", reflect.TypeOf((*MockUniversalClient)(nil).XPending), ctx, stream, group)
}

// XPendingExt mocks base method.
func (m *MockUniversalClient) XPendingExt(ctx context.Context, a *redis.XPendingExtArgs) *redis.XPendingExtCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "XPendingExt", ctx, a)
	ret0, _ := ret[0].(*redis.XPendingExtCmd)
	return ret0
}

// XPendingExt indicates an expected call of XPendingExt.
func (mr *MockUniversalClientMockRecorder) XPendingExt(ctx, a interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "XPendingExt", reflect.TypeOf((*MockUniversalClient)(nil).XPendingExt), ctx, a)
}

// XRange mocks base method.
func (m *MockUniversalClient) XRange(ctx context.Context, stream, start, stop string) *redis.XMessageSliceCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "XRange", ctx, stream, start, stop)
	ret0, _ := ret[0].(*redis.XMessageSliceCmd)
	return ret0
}

// XRange indicates an expected call of XRange.
func (mr *MockUniversalClientMockRecorder) XRange(ctx, stream, start, stop interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "XRange", reflect.TypeOf((*MockUniversalClient)(nil).XRange), ctx, stream, start, stop)
}

// XRangeN mocks base method.
func (m *MockUniversalClient) XRangeN(ctx context.Context, stream, start, stop string, count int64) *redis.XMessageSliceCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "XRangeN", ctx, stream, start, stop, count)
	ret0, _ := ret[0].(*redis.XMessageSliceCmd)
	return ret0
}

// XRangeN indicates an expected call of XRangeN.
func (mr *MockUniversalClientMockRecorder) XRangeN(ctx, stream, start, stop, count interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "XRangeN", reflect.TypeOf((*MockUniversalClient)(nil).XRangeN), ctx, stream, start, stop, count)
}

// XRead mocks base method.
func (m *MockUniversalClient) XRead(ctx context.Context, a *redis.XReadArgs) *redis.XStreamSliceCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "XRead", ctx, a)
	ret0, _ := ret[0].(*redis.XStreamSliceCmd)
	return ret0
}

// XRead indicates an expected call of XRead.
func (mr *MockUniversalClientMockRecorder) XRead(ctx, a interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "XRead", reflect.TypeOf((*MockUniversalClient)(nil).XRead), ctx, a)
}

// XReadGroup mocks base method.
func (m *MockUniversalClient) XReadGroup(ctx context.Context, a *redis.XReadGroupArgs) *redis.XStreamSliceCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "XReadGroup", ctx, a)
	ret0, _ := ret[0].(*redis.XStreamSliceCmd)
	return ret0
}

// XReadGroup indicates an expected call of XReadGroup.
func (mr *MockUniversalClientMockRecorder) XReadGroup(ctx, a interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "XReadGroup", reflect.TypeOf((*MockUniversalClient)(nil).XReadGroup), ctx, a)
}

// XReadStreams mocks base method.
func (m *MockUniversalClient) XReadStreams(ctx context.Context, streams ...string) *redis.XStreamSliceCmd {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx}
	for _, a := range streams {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "XReadStreams", varargs...)
	ret0, _ := ret[0].(*redis.XStreamSliceCmd)
	return ret0
}

// XReadStreams indicates an expected call of XReadStreams.
func (mr *MockUniversalClientMockRecorder) XReadStreams(ctx interface{}, streams ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx}, streams...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "XReadStreams", reflect.TypeOf((*MockUniversalClient)(nil).XReadStreams), varargs...)
}

// XRevRange mocks base method.
func (m *MockUniversalClient) XRevRange(ctx context.Context, stream, start, stop string) *redis.XMessageSliceCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "XRevRange", ctx, stream, start, stop)
	ret0, _ := ret[0].(*redis.XMessageSliceCmd)
	return ret0
}

// XRevRange indicates an expected call of XRevRange.
func (mr *MockUniversalClientMockRecorder) XRevRange(ctx, stream, start, stop interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "XRevRange", reflect.TypeOf((*MockUniversalClient)(nil).XRevRange), ctx, stream, start, stop)
}

// XRevRangeN mocks base method.
func (m *MockUniversalClient) XRevRangeN(ctx context.Context, stream, start, stop string, count int64) *redis.XMessageSliceCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "XRevRangeN", ctx, stream, start, stop, count)
	ret0, _ := ret[0].(*redis.XMessageSliceCmd)
	return ret0
}

// XRevRangeN indicates an expected call of XRevRangeN.
func (mr *MockUniversalClientMockRecorder) XRevRangeN(ctx, stream, start, stop, count interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "XRevRangeN", reflect.TypeOf((*MockUniversalClient)(nil).XRevRangeN), ctx, stream, start, stop, count)
}

// XTrim mocks base method.
func (m *MockUniversalClient) XTrim(ctx context.Context, key string, maxLen int64) *redis.IntCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "XTrim", ctx, key, maxLen)
	ret0, _ := ret[0].(*redis.IntCmd)
	return ret0
}

// XTrim indicates an expected call of XTrim.
func (mr *MockUniversalClientMockRecorder) XTrim(ctx, key, maxLen interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "XTrim", reflect.TypeOf((*MockUniversalClient)(nil).XTrim), ctx, key, maxLen)
}

// XTrimApprox mocks base method.
func (m *MockUniversalClient) XTrimApprox(ctx context.Context, key string, maxLen int64) *redis.IntCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "XTrimApprox", ctx, key, maxLen)
	ret0, _ := ret[0].(*redis.IntCmd)
	return ret0
}

// XTrimApprox indicates an expected call of XTrimApprox.
func (mr *MockUniversalClientMockRecorder) XTrimApprox(ctx, key, maxLen interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "XTrimApprox", reflect.TypeOf((*MockUniversalClient)(nil).XTrimApprox), ctx, key, maxLen)
}

// XTrimMaxLen mocks base method.
func (m *MockUniversalClient) XTrimMaxLen(ctx context.Context, key string, maxLen int64) *redis.IntCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "XTrimMaxLen", ctx, key, maxLen)
	ret0, _ := ret[0].(*redis.IntCmd)
	return ret0
}

// XTrimMaxLen indicates an expected call of XTrimMaxLen.
func (mr *MockUniversalClientMockRecorder) XTrimMaxLen(ctx, key, maxLen interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "XTrimMaxLen", reflect.TypeOf((*MockUniversalClient)(nil).XTrimMaxLen), ctx, key, maxLen)
}

// XTrimMaxLenApprox mocks base method.
func (m *MockUniversalClient) XTrimMaxLenApprox(ctx context.Context, key string, maxLen, limit int64) *redis.IntCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "XTrimMaxLenApprox", ctx, key, maxLen, limit)
	ret0, _ := ret[0].(*redis.IntCmd)
	return ret0
}

// XTrimMaxLenApprox indicates an expected call of XTrimMaxLenApprox.
func (mr *MockUniversalClientMockRecorder) XTrimMaxLenApprox(ctx, key, maxLen, limit interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "XTrimMaxLenApprox", reflect.TypeOf((*MockUniversalClient)(nil).XTrimMaxLenApprox), ctx, key, maxLen, limit)
}

// XTrimMinID mocks base method.
func (m *MockUniversalClient) XTrimMinID(ctx context.Context, key, minID string) *redis.IntCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "XTrimMinID", ctx, key, minID)
	ret0, _ := ret[0].(*redis.IntCmd)
	return ret0
}

// XTrimMinID indicates an expected call of XTrimMinID.
func (mr *MockUniversalClientMockRecorder) XTrimMinID(ctx, key, minID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "XTrimMinID", reflect.TypeOf((*MockUniversalClient)(nil).XTrimMinID), ctx, key, minID)
}

// XTrimMinIDApprox mocks base method.
func (m *MockUniversalClient) XTrimMinIDApprox(ctx context.Context, key, minID string, limit int64) *redis.IntCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "XTrimMinIDApprox", ctx, key, minID, limit)
	ret0, _ := ret[0].(*redis.IntCmd)
	return ret0
}

// XTrimMinIDApprox indicates an expected call of XTrimMinIDApprox.
func (mr *MockUniversalClientMockRecorder) XTrimMinIDApprox(ctx, key, minID, limit interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "XTrimMinIDApprox", reflect.TypeOf((*MockUniversalClient)(nil).XTrimMinIDApprox), ctx, key, minID, limit)
}

// ZAdd mocks base method.
func (m *MockUniversalClient) ZAdd(ctx context.Context, key string, members ...*redis.Z) *redis.IntCmd {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, key}
	for _, a := range members {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ZAdd", varargs...)
	ret0, _ := ret[0].(*redis.IntCmd)
	return ret0
}

// ZAdd indicates an expected call of ZAdd.
func (mr *MockUniversalClientMockRecorder) ZAdd(ctx, key interface{}, members ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, key}, members...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ZAdd", reflect.TypeOf((*MockUniversalClient)(nil).ZAdd), varargs...)
}

// ZAddArgsIncr mocks base method.
func (m *MockUniversalClient) ZAddArgsIncr(ctx context.Context, key string, args redis.ZAddArgs) *redis.FloatCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ZAddArgsIncr", ctx, key, args)
	ret0, _ := ret[0].(*redis.FloatCmd)
	return ret0
}

// ZAddArgsIncr indicates an expected call of ZAddArgsIncr.
func (mr *MockUniversalClientMockRecorder) ZAddArgsIncr(ctx, key, args interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ZAddArgsIncr", reflect.TypeOf((*MockUniversalClient)(nil).ZAddArgsIncr), ctx, key, args)
}

// ZAddCh mocks base method.
func (m *MockUniversalClient) ZAddCh(ctx context.Context, key string, members ...*redis.Z) *redis.IntCmd {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, key}
	for _, a := range members {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ZAddCh", varargs...)
	ret0, _ := ret[0].(*redis.IntCmd)
	return ret0
}

// ZAddCh indicates an expected call of ZAddCh.
func (mr *MockUniversalClientMockRecorder) ZAddCh(ctx, key interface{}, members ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, key}, members...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ZAddCh", reflect.TypeOf((*MockUniversalClient)(nil).ZAddCh), varargs...)
}

// ZAddNX mocks base method.
func (m *MockUniversalClient) ZAddNX(ctx context.Context, key string, members ...*redis.Z) *redis.IntCmd {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, key}
	for _, a := range members {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ZAddNX", varargs...)
	ret0, _ := ret[0].(*redis.IntCmd)
	return ret0
}

// ZAddNX indicates an expected call of ZAddNX.
func (mr *MockUniversalClientMockRecorder) ZAddNX(ctx, key interface{}, members ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, key}, members...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ZAddNX", reflect.TypeOf((*MockUniversalClient)(nil).ZAddNX), varargs...)
}

// ZAddNXCh mocks base method.
func (m *MockUniversalClient) ZAddNXCh(ctx context.Context, key string, members ...*redis.Z) *redis.IntCmd {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, key}
	for _, a := range members {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ZAddNXCh", varargs...)
	ret0, _ := ret[0].(*redis.IntCmd)
	return ret0
}

// ZAddNXCh indicates an expected call of ZAddNXCh.
func (mr *MockUniversalClientMockRecorder) ZAddNXCh(ctx, key interface{}, members ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, key}, members...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ZAddNXCh", reflect.TypeOf((*MockUniversalClient)(nil).ZAddNXCh), varargs...)
}

// ZAddXX mocks base method.
func (m *MockUniversalClient) ZAddXX(ctx context.Context, key string, members ...*redis.Z) *redis.IntCmd {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, key}
	for _, a := range members {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ZAddXX", varargs...)
	ret0, _ := ret[0].(*redis.IntCmd)
	return ret0
}

// ZAddXX indicates an expected call of ZAddXX.
func (mr *MockUniversalClientMockRecorder) ZAddXX(ctx, key interface{}, members ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, key}, members...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ZAddXX", reflect.TypeOf((*MockUniversalClient)(nil).ZAddXX), varargs...)
}

// ZAddXXCh mocks base method.
func (m *MockUniversalClient) ZAddXXCh(ctx context.Context, key string, members ...*redis.Z) *redis.IntCmd {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, key}
	for _, a := range members {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ZAddXXCh", varargs...)
	ret0, _ := ret[0].(*redis.IntCmd)
	return ret0
}

// ZAddXXCh indicates an expected call of ZAddXXCh.
func (mr *MockUniversalClientMockRecorder) ZAddXXCh(ctx, key interface{}, members ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, key}, members...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ZAddXXCh", reflect.TypeOf((*MockUniversalClient)(nil).ZAddXXCh), varargs...)
}

// ZCard mocks base method.
func (m *MockUniversalClient) ZCard(ctx context.Context, key string) *redis.IntCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ZCard", ctx, key)
	ret0, _ := ret[0].(*redis.IntCmd)
	return ret0
}

// ZCard indicates an expected call of ZCard.
func (mr *MockUniversalClientMockRecorder) ZCard(ctx, key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ZCard", reflect.TypeOf((*MockUniversalClient)(nil).ZCard), ctx, key)
}

// ZCount mocks base method.
func (m *MockUniversalClient) ZCount(ctx context.Context, key, min, max string) *redis.IntCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ZCount", ctx, key, min, max)
	ret0, _ := ret[0].(*redis.IntCmd)
	return ret0
}

// ZCount indicates an expected call of ZCount.
func (mr *MockUniversalClientMockRecorder) ZCount(ctx, key, min, max interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ZCount", reflect.TypeOf((*MockUniversalClient)(nil).ZCount), ctx, key, min, max)
}

// ZDiff mocks base method.
func (m *MockUniversalClient) ZDiff(ctx context.Context, keys ...string) *redis.StringSliceCmd {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx}
	for _, a := range keys {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ZDiff", varargs...)
	ret0, _ := ret[0].(*redis.StringSliceCmd)
	return ret0
}

// ZDiff indicates an expected call of ZDiff.
func (mr *MockUniversalClientMockRecorder) ZDiff(ctx interface{}, keys ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx}, keys...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ZDiff", reflect.TypeOf((*MockUniversalClient)(nil).ZDiff), varargs...)
}

// ZDiffStore mocks base method.
func (m *MockUniversalClient) ZDiffStore(ctx context.Context, destination string, keys ...string) *redis.IntCmd {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, destination}
	for _, a := range keys {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ZDiffStore", varargs...)
	ret0, _ := ret[0].(*redis.IntCmd)
	return ret0
}

// ZDiffStore indicates an expected call of ZDiffStore.
func (mr *MockUniversalClientMockRecorder) ZDiffStore(ctx, destination interface{}, keys ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, destination}, keys...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ZDiffStore", reflect.TypeOf((*MockUniversalClient)(nil).ZDiffStore), varargs...)
}

// ZDiffWithScores mocks base method.
func (m *MockUniversalClient) ZDiffWithScores(ctx context.Context, keys ...string) *redis.ZSliceCmd {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx}
	for _, a := range keys {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ZDiffWithScores", varargs...)
	ret0, _ := ret[0].(*redis.ZSliceCmd)
	return ret0
}

// ZDiffWithScores indicates an expected call of ZDiffWithScores.
func (mr *MockUniversalClientMockRecorder) ZDiffWithScores(ctx interface{}, keys ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx}, keys...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ZDiffWithScores", reflect.TypeOf((*MockUniversalClient)(nil).ZDiffWithScores), varargs...)
}

// ZIncr mocks base method.
func (m *MockUniversalClient) ZIncr(ctx context.Context, key string, member *redis.Z) *redis.FloatCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ZIncr", ctx, key, member)
	ret0, _ := ret[0].(*redis.FloatCmd)
	return ret0
}

// ZIncr indicates an expected call of ZIncr.
func (mr *MockUniversalClientMockRecorder) ZIncr(ctx, key, member interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ZIncr", reflect.TypeOf((*MockUniversalClient)(nil).ZIncr), ctx, key, member)
}

// ZIncrBy mocks base method.
func (m *MockUniversalClient) ZIncrBy(ctx context.Context, key string, increment float64, member string) *redis.FloatCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ZIncrBy", ctx, key, increment, member)
	ret0, _ := ret[0].(*redis.FloatCmd)
	return ret0
}

// ZIncrBy indicates an expected call of ZIncrBy.
func (mr *MockUniversalClientMockRecorder) ZIncrBy(ctx, key, increment, member interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ZIncrBy", reflect.TypeOf((*MockUniversalClient)(nil).ZIncrBy), ctx, key, increment, member)
}

// ZIncrNX mocks base method.
func (m *MockUniversalClient) ZIncrNX(ctx context.Context, key string, member *redis.Z) *redis.FloatCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ZIncrNX", ctx, key, member)
	ret0, _ := ret[0].(*redis.FloatCmd)
	return ret0
}

// ZIncrNX indicates an expected call of ZIncrNX.
func (mr *MockUniversalClientMockRecorder) ZIncrNX(ctx, key, member interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ZIncrNX", reflect.TypeOf((*MockUniversalClient)(nil).ZIncrNX), ctx, key, member)
}

// ZIncrXX mocks base method.
func (m *MockUniversalClient) ZIncrXX(ctx context.Context, key string, member *redis.Z) *redis.FloatCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ZIncrXX", ctx, key, member)
	ret0, _ := ret[0].(*redis.FloatCmd)
	return ret0
}

// ZIncrXX indicates an expected call of ZIncrXX.
func (mr *MockUniversalClientMockRecorder) ZIncrXX(ctx, key, member interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ZIncrXX", reflect.TypeOf((*MockUniversalClient)(nil).ZIncrXX), ctx, key, member)
}

// ZInter mocks base method.
func (m *MockUniversalClient) ZInter(ctx context.Context, store *redis.ZStore) *redis.StringSliceCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ZInter", ctx, store)
	ret0, _ := ret[0].(*redis.StringSliceCmd)
	return ret0
}

// ZInter indicates an expected call of ZInter.
func (mr *MockUniversalClientMockRecorder) ZInter(ctx, store interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ZInter", reflect.TypeOf((*MockUniversalClient)(nil).ZInter), ctx, store)
}

// ZInterStore mocks base method.
func (m *MockUniversalClient) ZInterStore(ctx context.Context, destination string, store *redis.ZStore) *redis.IntCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ZInterStore", ctx, destination, store)
	ret0, _ := ret[0].(*redis.IntCmd)
	return ret0
}

// ZInterStore indicates an expected call of ZInterStore.
func (mr *MockUniversalClientMockRecorder) ZInterStore(ctx, destination, store interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ZInterStore", reflect.TypeOf((*MockUniversalClient)(nil).ZInterStore), ctx, destination, store)
}

// ZInterWithScores mocks base method.
func (m *MockUniversalClient) ZInterWithScores(ctx context.Context, store *redis.ZStore) *redis.ZSliceCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ZInterWithScores", ctx, store)
	ret0, _ := ret[0].(*redis.ZSliceCmd)
	return ret0
}

// ZInterWithScores indicates an expected call of ZInterWithScores.
func (mr *MockUniversalClientMockRecorder) ZInterWithScores(ctx, store interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ZInterWithScores", reflect.TypeOf((*MockUniversalClient)(nil).ZInterWithScores), ctx, store)
}

// ZLexCount mocks base method.
func (m *MockUniversalClient) ZLexCount(ctx context.Context, key, min, max string) *redis.IntCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ZLexCount", ctx, key, min, max)
	ret0, _ := ret[0].(*redis.IntCmd)
	return ret0
}

// ZLexCount indicates an expected call of ZLexCount.
func (mr *MockUniversalClientMockRecorder) ZLexCount(ctx, key, min, max interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ZLexCount", reflect.TypeOf((*MockUniversalClient)(nil).ZLexCount), ctx, key, min, max)
}

// ZMScore mocks base method.
func (m *MockUniversalClient) ZMScore(ctx context.Context, key string, members ...string) *redis.FloatSliceCmd {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, key}
	for _, a := range members {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ZMScore", varargs...)
	ret0, _ := ret[0].(*redis.FloatSliceCmd)
	return ret0
}

// ZMScore indicates an expected call of ZMScore.
func (mr *MockUniversalClientMockRecorder) ZMScore(ctx, key interface{}, members ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, key}, members...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ZMScore", reflect.TypeOf((*MockUniversalClient)(nil).ZMScore), varargs...)
}

// ZPopMax mocks base method.
func (m *MockUniversalClient) ZPopMax(ctx context.Context, key string, count ...int64) *redis.ZSliceCmd {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, key}
	for _, a := range count {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ZPopMax", varargs...)
	ret0, _ := ret[0].(*redis.ZSliceCmd)
	return ret0
}

// ZPopMax indicates an expected call of ZPopMax.
func (mr *MockUniversalClientMockRecorder) ZPopMax(ctx, key interface{}, count ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, key}, count...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ZPopMax", reflect.TypeOf((*MockUniversalClient)(nil).ZPopMax), varargs...)
}

// ZPopMin mocks base method.
func (m *MockUniversalClient) ZPopMin(ctx context.Context, key string, count ...int64) *redis.ZSliceCmd {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, key}
	for _, a := range count {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ZPopMin", varargs...)
	ret0, _ := ret[0].(*redis.ZSliceCmd)
	return ret0
}

// ZPopMin indicates an expected call of ZPopMin.
func (mr *MockUniversalClientMockRecorder) ZPopMin(ctx, key interface{}, count ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, key}, count...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ZPopMin", reflect.TypeOf((*MockUniversalClient)(nil).ZPopMin), varargs...)
}

// ZRandMember mocks base method.
func (m *MockUniversalClient) ZRandMember(ctx context.Context, key string, count int, withScores bool) *redis.StringSliceCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ZRandMember", ctx, key, count, withScores)
	ret0, _ := ret[0].(*redis.StringSliceCmd)
	return ret0
}

// ZRandMember indicates an expected call of ZRandMember.
func (mr *MockUniversalClientMockRecorder) ZRandMember(ctx, key, count, withScores interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ZRandMember", reflect.TypeOf((*MockUniversalClient)(nil).ZRandMember), ctx, key, count, withScores)
}

// ZRange mocks base method.
func (m *MockUniversalClient) ZRange(ctx context.Context, key string, start, stop int64) *redis.StringSliceCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ZRange", ctx, key, start, stop)
	ret0, _ := ret[0].(*redis.StringSliceCmd)
	return ret0
}

// ZRange indicates an expected call of ZRange.
func (mr *MockUniversalClientMockRecorder) ZRange(ctx, key, start, stop interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ZRange", reflect.TypeOf((*MockUniversalClient)(nil).ZRange), ctx, key, start, stop)
}

// ZRangeArgs mocks base method.
func (m *MockUniversalClient) ZRangeArgs(ctx context.Context, z redis.ZRangeArgs) *redis.StringSliceCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ZRangeArgs", ctx, z)
	ret0, _ := ret[0].(*redis.StringSliceCmd)
	return ret0
}

// ZRangeArgs indicates an expected call of ZRangeArgs.
func (mr *MockUniversalClientMockRecorder) ZRangeArgs(ctx, z interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ZRangeArgs", reflect.TypeOf((*MockUniversalClient)(nil).ZRangeArgs), ctx, z)
}

// ZRangeArgsWithScores mocks base method.
func (m *MockUniversalClient) ZRangeArgsWithScores(ctx context.Context, z redis.ZRangeArgs) *redis.ZSliceCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ZRangeArgsWithScores", ctx, z)
	ret0, _ := ret[0].(*redis.ZSliceCmd)
	return ret0
}

// ZRangeArgsWithScores indicates an expected call of ZRangeArgsWithScores.
func (mr *MockUniversalClientMockRecorder) ZRangeArgsWithScores(ctx, z interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ZRangeArgsWithScores", reflect.TypeOf((*MockUniversalClient)(nil).ZRangeArgsWithScores), ctx, z)
}

// ZRangeByLex mocks base method.
func (m *MockUniversalClient) ZRangeByLex(ctx context.Context, key string, opt *redis.ZRangeBy) *redis.StringSliceCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ZRangeByLex", ctx, key, opt)
	ret0, _ := ret[0].(*redis.StringSliceCmd)
	return ret0
}

// ZRangeByLex indicates an expected call of ZRangeByLex.
func (mr *MockUniversalClientMockRecorder) ZRangeByLex(ctx, key, opt interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ZRangeByLex", reflect.TypeOf((*MockUniversalClient)(nil).ZRangeByLex), ctx, key, opt)
}

// ZRangeByScore mocks base method.
func (m *MockUniversalClient) ZRangeByScore(ctx context.Context, key string, opt *redis.ZRangeBy) *redis.StringSliceCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ZRangeByScore", ctx, key, opt)
	ret0, _ := ret[0].(*redis.StringSliceCmd)
	return ret0
}

// ZRangeByScore indicates an expected call of ZRangeByScore.
func (mr *MockUniversalClientMockRecorder) ZRangeByScore(ctx, key, opt interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ZRangeByScore", reflect.TypeOf((*MockUniversalClient)(nil).ZRangeByScore), ctx, key, opt)
}

// ZRangeByScoreWithScores mocks base method.
func (m *MockUniversalClient) ZRangeByScoreWithScores(ctx context.Context, key string, opt *redis.ZRangeBy) *redis.ZSliceCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ZRangeByScoreWithScores", ctx, key, opt)
	ret0, _ := ret[0].(*redis.ZSliceCmd)
	return ret0
}

// ZRangeByScoreWithScores indicates an expected call of ZRangeByScoreWithScores.
func (mr *MockUniversalClientMockRecorder) ZRangeByScoreWithScores(ctx, key, opt interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ZRangeByScoreWithScores", reflect.TypeOf((*MockUniversalClient)(nil).ZRangeByScoreWithScores), ctx, key, opt)
}

// ZRangeStore mocks base method.
func (m *MockUniversalClient) ZRangeStore(ctx context.Context, dst string, z redis.ZRangeArgs) *redis.IntCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ZRangeStore", ctx, dst, z)
	ret0, _ := ret[0].(*redis.IntCmd)
	return ret0
}

// ZRangeStore indicates an expected call of ZRangeStore.
func (mr *MockUniversalClientMockRecorder) ZRangeStore(ctx, dst, z interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ZRangeStore", reflect.TypeOf((*MockUniversalClient)(nil).ZRangeStore), ctx, dst, z)
}

// ZRangeWithScores mocks base method.
func (m *MockUniversalClient) ZRangeWithScores(ctx context.Context, key string, start, stop int64) *redis.ZSliceCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ZRangeWithScores", ctx, key, start, stop)
	ret0, _ := ret[0].(*redis.ZSliceCmd)
	return ret0
}

// ZRangeWithScores indicates an expected call of ZRangeWithScores.
func (mr *MockUniversalClientMockRecorder) ZRangeWithScores(ctx, key, start, stop interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ZRangeWithScores", reflect.TypeOf((*MockUniversalClient)(nil).ZRangeWithScores), ctx, key, start, stop)
}

// ZRank mocks base method.
func (m *MockUniversalClient) ZRank(ctx context.Context, key, member string) *redis.IntCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ZRank", ctx, key, member)
	ret0, _ := ret[0].(*redis.IntCmd)
	return ret0
}

// ZRank indicates an expected call of ZRank.
func (mr *MockUniversalClientMockRecorder) ZRank(ctx, key, member interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ZRank", reflect.TypeOf((*MockUniversalClient)(nil).ZRank), ctx, key, member)
}

// ZRem mocks base method.
func (m *MockUniversalClient) ZRem(ctx context.Context, key string, members ...interface{}) *redis.IntCmd {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, key}
	for _, a := range members {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ZRem", varargs...)
	ret0, _ := ret[0].(*redis.IntCmd)
	return ret0
}

// ZRem indicates an expected call of ZRem.
func (mr *MockUniversalClientMockRecorder) ZRem(ctx, key interface{}, members ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, key}, members...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ZRem", reflect.TypeOf((*MockUniversalClient)(nil).ZRem), varargs...)
}

// ZRemRangeByLex mocks base method.
func (m *MockUniversalClient) ZRemRangeByLex(ctx context.Context, key, min, max string) *redis.IntCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ZRemRangeByLex", ctx, key, min, max)
	ret0, _ := ret[0].(*redis.IntCmd)
	return ret0
}

// ZRemRangeByLex indicates an expected call of ZRemRangeByLex.
func (mr *MockUniversalClientMockRecorder) ZRemRangeByLex(ctx, key, min, max interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ZRemRangeByLex", reflect.TypeOf((*MockUniversalClient)(nil).ZRemRangeByLex), ctx, key, min, max)
}

// ZRemRangeByRank mocks base method.
func (m *MockUniversalClient) ZRemRangeByRank(ctx context.Context, key string, start, stop int64) *redis.IntCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ZRemRangeByRank", ctx, key, start, stop)
	ret0, _ := ret[0].(*redis.IntCmd)
	return ret0
}

// ZRemRangeByRank indicates an expected call of ZRemRangeByRank.
func (mr *MockUniversalClientMockRecorder) ZRemRangeByRank(ctx, key, start, stop interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ZRemRangeByRank", reflect.TypeOf((*MockUniversalClient)(nil).ZRemRangeByRank), ctx, key, start, stop)
}

// ZRemRangeByScore mocks base method.
func (m *MockUniversalClient) ZRemRangeByScore(ctx context.Context, key, min, max string) *redis.IntCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ZRemRangeByScore", ctx, key, min, max)
	ret0, _ := ret[0].(*redis.IntCmd)
	return ret0
}

// ZRemRangeByScore indicates an expected call of ZRemRangeByScore.
func (mr *MockUniversalClientMockRecorder) ZRemRangeByScore(ctx, key, min, max interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ZRemRangeByScore", reflect.TypeOf((*MockUniversalClient)(nil).ZRemRangeByScore), ctx, key, min, max)
}

// ZRevRange mocks base method.
func (m *MockUniversalClient) ZRevRange(ctx context.Context, key string, start, stop int64) *redis.StringSliceCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ZRevRange", ctx, key, start, stop)
	ret0, _ := ret[0].(*redis.StringSliceCmd)
	return ret0
}

// ZRevRange indicates an expected call of ZRevRange.
func (mr *MockUniversalClientMockRecorder) ZRevRange(ctx, key, start, stop interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ZRevRange", reflect.TypeOf((*MockUniversalClient)(nil).ZRevRange), ctx, key, start, stop)
}

// ZRevRangeByLex mocks base method.
func (m *MockUniversalClient) ZRevRangeByLex(ctx context.Context, key string, opt *redis.ZRangeBy) *redis.StringSliceCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ZRevRangeByLex", ctx, key, opt)
	ret0, _ := ret[0].(*redis.StringSliceCmd)
	return ret0
}

// ZRevRangeByLex indicates an expected call of ZRevRangeByLex.
func (mr *MockUniversalClientMockRecorder) ZRevRangeByLex(ctx, key, opt interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ZRevRangeByLex", reflect.TypeOf((*MockUniversalClient)(nil).ZRevRangeByLex), ctx, key, opt)
}

// ZRevRangeByScore mocks base method.
func (m *MockUniversalClient) ZRevRangeByScore(ctx context.Context, key string, opt *redis.ZRangeBy) *redis.StringSliceCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ZRevRangeByScore", ctx, key, opt)
	ret0, _ := ret[0].(*redis.StringSliceCmd)
	return ret0
}

// ZRevRangeByScore indicates an expected call of ZRevRangeByScore.
func (mr *MockUniversalClientMockRecorder) ZRevRangeByScore(ctx, key, opt interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ZRevRangeByScore", reflect.TypeOf((*MockUniversalClient)(nil).ZRevRangeByScore), ctx, key, opt)
}

// ZRevRangeByScoreWithScores mocks base method.
func (m *MockUniversalClient) ZRevRangeByScoreWithScores(ctx context.Context, key string, opt *redis.ZRangeBy) *redis.ZSliceCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ZRevRangeByScoreWithScores", ctx, key, opt)
	ret0, _ := ret[0].(*redis.ZSliceCmd)
	return ret0
}

// ZRevRangeByScoreWithScores indicates an expected call of ZRevRangeByScoreWithScores.
func (mr *MockUniversalClientMockRecorder) ZRevRangeByScoreWithScores(ctx, key, opt interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ZRevRangeByScoreWithScores", reflect.TypeOf((*MockUniversalClient)(nil).ZRevRangeByScoreWithScores), ctx, key, opt)
}

// ZRevRangeWithScores mocks base method.
func (m *MockUniversalClient) ZRevRangeWithScores(ctx context.Context, key string, start, stop int64) *redis.ZSliceCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ZRevRangeWithScores", ctx, key, start, stop)
	ret0, _ := ret[0].(*redis.ZSliceCmd)
	return ret0
}

// ZRevRangeWithScores indicates an expected call of ZRevRangeWithScores.
func (mr *MockUniversalClientMockRecorder) ZRevRangeWithScores(ctx, key, start, stop interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ZRevRangeWithScores", reflect.TypeOf((*MockUniversalClient)(nil).ZRevRangeWithScores), ctx, key, start, stop)
}

// ZRevRank mocks base method.
func (m *MockUniversalClient) ZRevRank(ctx context.Context, key, member string) *redis.IntCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ZRevRank", ctx, key, member)
	ret0, _ := ret[0].(*redis.IntCmd)
	return ret0
}

// ZRevRank indicates an expected call of ZRevRank.
func (mr *MockUniversalClientMockRecorder) ZRevRank(ctx, key, member interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ZRevRank", reflect.TypeOf((*MockUniversalClient)(nil).ZRevRank), ctx, key, member)
}

// ZScan mocks base method.
func (m *MockUniversalClient) ZScan(ctx context.Context, key string, cursor uint64, match string, count int64) *redis.ScanCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ZScan", ctx, key, cursor, match, count)
	ret0, _ := ret[0].(*redis.ScanCmd)
	return ret0
}

// ZScan indicates an expected call of ZScan.
func (mr *MockUniversalClientMockRecorder) ZScan(ctx, key, cursor, match, count interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ZScan", reflect.TypeOf((*MockUniversalClient)(nil).ZScan), ctx, key, cursor, match, count)
}

// ZScore mocks base method.
func (m *MockUniversalClient) ZScore(ctx context.Context, key, member string) *redis.FloatCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ZScore", ctx, key, member)
	ret0, _ := ret[0].(*redis.FloatCmd)
	return ret0
}

// ZScore indicates an expected call of ZScore.
func (mr *MockUniversalClientMockRecorder) ZScore(ctx, key, member interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ZScore", reflect.TypeOf((*MockUniversalClient)(nil).ZScore), ctx, key, member)
}

// ZUnion mocks base method.
func (m *MockUniversalClient) ZUnion(ctx context.Context, store redis.ZStore) *redis.StringSliceCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ZUnion", ctx, store)
	ret0, _ := ret[0].(*redis.StringSliceCmd)
	return ret0
}

// ZUnion indicates an expected call of ZUnion.
func (mr *MockUniversalClientMockRecorder) ZUnion(ctx, store interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ZUnion", reflect.TypeOf((*MockUniversalClient)(nil).ZUnion), ctx, store)
}

// ZUnionStore mocks base method.
func (m *MockUniversalClient) ZUnionStore(ctx context.Context, dest string, store *redis.ZStore) *redis.IntCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ZUnionStore", ctx, dest, store)
	ret0, _ := ret[0].(*redis.IntCmd)
	return ret0
}

// ZUnionStore indicates an expected call of ZUnionStore.
func (mr *MockUniversalClientMockRecorder) ZUnionStore(ctx, dest, store interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ZUnionStore", reflect.TypeOf((*MockUniversalClient)(nil).ZUnionStore), ctx, dest, store)
}

// ZUnionWithScores mocks base method.
func (m *MockUniversalClient) ZUnionWithScores(ctx context.Context, store redis.ZStore) *redis.ZSliceCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ZUnionWithScores", ctx, store)
	ret0, _ := ret[0].(*redis.ZSliceCmd)
	return ret0
}

// ZUnionWithScores indicates an expected call of ZUnionWithScores.
func (mr *MockUniversalClientMockRecorder) ZUnionWithScores(ctx, store interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ZUnionWithScores", reflect.TypeOf((*MockUniversalClient)(nil).ZUnionWithScores), ctx, store)
}
