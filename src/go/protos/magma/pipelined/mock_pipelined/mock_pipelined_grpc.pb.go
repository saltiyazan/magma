// Code generated by MockGen. DO NOT EDIT.
// Source: magma/pipelined/pipelined_grpc.pb.go

// Package mock_pipelined is a generated GoMock package.
package mock_pipelined

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	orc8r "github.com/magma/magma/src/go/protos/magma/orc8r"
	pipelined "github.com/magma/magma/src/go/protos/magma/pipelined"
	session_manager "github.com/magma/magma/src/go/protos/magma/session_manager"
	grpc "google.golang.org/grpc"
)

// MockPipelinedClient is a mock of PipelinedClient interface.
type MockPipelinedClient struct {
	ctrl     *gomock.Controller
	recorder *MockPipelinedClientMockRecorder
}

// MockPipelinedClientMockRecorder is the mock recorder for MockPipelinedClient.
type MockPipelinedClientMockRecorder struct {
	mock *MockPipelinedClient
}

// NewMockPipelinedClient creates a new mock instance.
func NewMockPipelinedClient(ctrl *gomock.Controller) *MockPipelinedClient {
	mock := &MockPipelinedClient{ctrl: ctrl}
	mock.recorder = &MockPipelinedClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPipelinedClient) EXPECT() *MockPipelinedClientMockRecorder {
	return m.recorder
}

// ActivateFlows mocks base method.
func (m *MockPipelinedClient) ActivateFlows(ctx context.Context, in *pipelined.ActivateFlowsRequest, opts ...grpc.CallOption) (*pipelined.ActivateFlowsResult, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ActivateFlows", varargs...)
	ret0, _ := ret[0].(*pipelined.ActivateFlowsResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ActivateFlows indicates an expected call of ActivateFlows.
func (mr *MockPipelinedClientMockRecorder) ActivateFlows(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ActivateFlows", reflect.TypeOf((*MockPipelinedClient)(nil).ActivateFlows), varargs...)
}

// AddUEMacFlow mocks base method.
func (m *MockPipelinedClient) AddUEMacFlow(ctx context.Context, in *pipelined.UEMacFlowRequest, opts ...grpc.CallOption) (*pipelined.FlowResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "AddUEMacFlow", varargs...)
	ret0, _ := ret[0].(*pipelined.FlowResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddUEMacFlow indicates an expected call of AddUEMacFlow.
func (mr *MockPipelinedClientMockRecorder) AddUEMacFlow(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddUEMacFlow", reflect.TypeOf((*MockPipelinedClient)(nil).AddUEMacFlow), varargs...)
}

// CreateFlow mocks base method.
func (m *MockPipelinedClient) CreateFlow(ctx context.Context, in *pipelined.FlowRequest, opts ...grpc.CallOption) (*pipelined.FlowResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CreateFlow", varargs...)
	ret0, _ := ret[0].(*pipelined.FlowResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateFlow indicates an expected call of CreateFlow.
func (mr *MockPipelinedClientMockRecorder) CreateFlow(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateFlow", reflect.TypeOf((*MockPipelinedClient)(nil).CreateFlow), varargs...)
}

// DeactivateFlows mocks base method.
func (m *MockPipelinedClient) DeactivateFlows(ctx context.Context, in *pipelined.DeactivateFlowsRequest, opts ...grpc.CallOption) (*pipelined.DeactivateFlowsResult, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DeactivateFlows", varargs...)
	ret0, _ := ret[0].(*pipelined.DeactivateFlowsResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeactivateFlows indicates an expected call of DeactivateFlows.
func (mr *MockPipelinedClientMockRecorder) DeactivateFlows(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeactivateFlows", reflect.TypeOf((*MockPipelinedClient)(nil).DeactivateFlows), varargs...)
}

// DeleteUEMacFlow mocks base method.
func (m *MockPipelinedClient) DeleteUEMacFlow(ctx context.Context, in *pipelined.UEMacFlowRequest, opts ...grpc.CallOption) (*pipelined.FlowResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DeleteUEMacFlow", varargs...)
	ret0, _ := ret[0].(*pipelined.FlowResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteUEMacFlow indicates an expected call of DeleteUEMacFlow.
func (mr *MockPipelinedClientMockRecorder) DeleteUEMacFlow(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteUEMacFlow", reflect.TypeOf((*MockPipelinedClient)(nil).DeleteUEMacFlow), varargs...)
}

// GetAllTableAssignments mocks base method.
func (m *MockPipelinedClient) GetAllTableAssignments(ctx context.Context, in *orc8r.Void, opts ...grpc.CallOption) (*pipelined.AllTableAssignments, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetAllTableAssignments", varargs...)
	ret0, _ := ret[0].(*pipelined.AllTableAssignments)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllTableAssignments indicates an expected call of GetAllTableAssignments.
func (mr *MockPipelinedClientMockRecorder) GetAllTableAssignments(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllTableAssignments", reflect.TypeOf((*MockPipelinedClient)(nil).GetAllTableAssignments), varargs...)
}

// GetPolicyUsage mocks base method.
func (m *MockPipelinedClient) GetPolicyUsage(ctx context.Context, in *orc8r.Void, opts ...grpc.CallOption) (*session_manager.RuleRecordTable, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetPolicyUsage", varargs...)
	ret0, _ := ret[0].(*session_manager.RuleRecordTable)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPolicyUsage indicates an expected call of GetPolicyUsage.
func (mr *MockPipelinedClientMockRecorder) GetPolicyUsage(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPolicyUsage", reflect.TypeOf((*MockPipelinedClient)(nil).GetPolicyUsage), varargs...)
}

// GetStats mocks base method.
func (m *MockPipelinedClient) GetStats(ctx context.Context, in *pipelined.GetStatsRequest, opts ...grpc.CallOption) (*session_manager.RuleRecordTable, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetStats", varargs...)
	ret0, _ := ret[0].(*session_manager.RuleRecordTable)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetStats indicates an expected call of GetStats.
func (mr *MockPipelinedClientMockRecorder) GetStats(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetStats", reflect.TypeOf((*MockPipelinedClient)(nil).GetStats), varargs...)
}

// RemoveFlow mocks base method.
func (m *MockPipelinedClient) RemoveFlow(ctx context.Context, in *pipelined.FlowRequest, opts ...grpc.CallOption) (*pipelined.FlowResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "RemoveFlow", varargs...)
	ret0, _ := ret[0].(*pipelined.FlowResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RemoveFlow indicates an expected call of RemoveFlow.
func (mr *MockPipelinedClientMockRecorder) RemoveFlow(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveFlow", reflect.TypeOf((*MockPipelinedClient)(nil).RemoveFlow), varargs...)
}

// SetSMFSessions mocks base method.
func (m *MockPipelinedClient) SetSMFSessions(ctx context.Context, in *pipelined.SessionSet, opts ...grpc.CallOption) (*pipelined.UPFSessionContextState, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "SetSMFSessions", varargs...)
	ret0, _ := ret[0].(*pipelined.UPFSessionContextState)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SetSMFSessions indicates an expected call of SetSMFSessions.
func (mr *MockPipelinedClientMockRecorder) SetSMFSessions(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetSMFSessions", reflect.TypeOf((*MockPipelinedClient)(nil).SetSMFSessions), varargs...)
}

// SetupDefaultControllers mocks base method.
func (m *MockPipelinedClient) SetupDefaultControllers(ctx context.Context, in *pipelined.SetupDefaultRequest, opts ...grpc.CallOption) (*pipelined.SetupFlowsResult, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "SetupDefaultControllers", varargs...)
	ret0, _ := ret[0].(*pipelined.SetupFlowsResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SetupDefaultControllers indicates an expected call of SetupDefaultControllers.
func (mr *MockPipelinedClientMockRecorder) SetupDefaultControllers(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetupDefaultControllers", reflect.TypeOf((*MockPipelinedClient)(nil).SetupDefaultControllers), varargs...)
}

// SetupPolicyFlows mocks base method.
func (m *MockPipelinedClient) SetupPolicyFlows(ctx context.Context, in *pipelined.SetupPolicyRequest, opts ...grpc.CallOption) (*pipelined.SetupFlowsResult, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "SetupPolicyFlows", varargs...)
	ret0, _ := ret[0].(*pipelined.SetupFlowsResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SetupPolicyFlows indicates an expected call of SetupPolicyFlows.
func (mr *MockPipelinedClientMockRecorder) SetupPolicyFlows(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetupPolicyFlows", reflect.TypeOf((*MockPipelinedClient)(nil).SetupPolicyFlows), varargs...)
}

// SetupQuotaFlows mocks base method.
func (m *MockPipelinedClient) SetupQuotaFlows(ctx context.Context, in *pipelined.SetupQuotaRequest, opts ...grpc.CallOption) (*pipelined.SetupFlowsResult, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "SetupQuotaFlows", varargs...)
	ret0, _ := ret[0].(*pipelined.SetupFlowsResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SetupQuotaFlows indicates an expected call of SetupQuotaFlows.
func (mr *MockPipelinedClientMockRecorder) SetupQuotaFlows(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetupQuotaFlows", reflect.TypeOf((*MockPipelinedClient)(nil).SetupQuotaFlows), varargs...)
}

// SetupUEMacFlows mocks base method.
func (m *MockPipelinedClient) SetupUEMacFlows(ctx context.Context, in *pipelined.SetupUEMacRequest, opts ...grpc.CallOption) (*pipelined.SetupFlowsResult, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "SetupUEMacFlows", varargs...)
	ret0, _ := ret[0].(*pipelined.SetupFlowsResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SetupUEMacFlows indicates an expected call of SetupUEMacFlows.
func (mr *MockPipelinedClientMockRecorder) SetupUEMacFlows(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetupUEMacFlows", reflect.TypeOf((*MockPipelinedClient)(nil).SetupUEMacFlows), varargs...)
}

// UpdateFlowStats mocks base method.
func (m *MockPipelinedClient) UpdateFlowStats(ctx context.Context, in *pipelined.FlowRequest, opts ...grpc.CallOption) (*pipelined.FlowResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UpdateFlowStats", varargs...)
	ret0, _ := ret[0].(*pipelined.FlowResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateFlowStats indicates an expected call of UpdateFlowStats.
func (mr *MockPipelinedClientMockRecorder) UpdateFlowStats(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateFlowStats", reflect.TypeOf((*MockPipelinedClient)(nil).UpdateFlowStats), varargs...)
}

// UpdateIPFIXFlow mocks base method.
func (m *MockPipelinedClient) UpdateIPFIXFlow(ctx context.Context, in *pipelined.UEMacFlowRequest, opts ...grpc.CallOption) (*pipelined.FlowResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UpdateIPFIXFlow", varargs...)
	ret0, _ := ret[0].(*pipelined.FlowResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateIPFIXFlow indicates an expected call of UpdateIPFIXFlow.
func (mr *MockPipelinedClientMockRecorder) UpdateIPFIXFlow(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateIPFIXFlow", reflect.TypeOf((*MockPipelinedClient)(nil).UpdateIPFIXFlow), varargs...)
}

// UpdateSubscriberQuotaState mocks base method.
func (m *MockPipelinedClient) UpdateSubscriberQuotaState(ctx context.Context, in *pipelined.UpdateSubscriberQuotaStateRequest, opts ...grpc.CallOption) (*pipelined.FlowResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UpdateSubscriberQuotaState", varargs...)
	ret0, _ := ret[0].(*pipelined.FlowResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateSubscriberQuotaState indicates an expected call of UpdateSubscriberQuotaState.
func (mr *MockPipelinedClientMockRecorder) UpdateSubscriberQuotaState(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateSubscriberQuotaState", reflect.TypeOf((*MockPipelinedClient)(nil).UpdateSubscriberQuotaState), varargs...)
}

// UpdateUEState mocks base method.
func (m *MockPipelinedClient) UpdateUEState(ctx context.Context, in *pipelined.UESessionSet, opts ...grpc.CallOption) (*pipelined.UESessionContextResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UpdateUEState", varargs...)
	ret0, _ := ret[0].(*pipelined.UESessionContextResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateUEState indicates an expected call of UpdateUEState.
func (mr *MockPipelinedClientMockRecorder) UpdateUEState(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateUEState", reflect.TypeOf((*MockPipelinedClient)(nil).UpdateUEState), varargs...)
}

// MockPipelinedServer is a mock of PipelinedServer interface.
type MockPipelinedServer struct {
	ctrl     *gomock.Controller
	recorder *MockPipelinedServerMockRecorder
}

// MockPipelinedServerMockRecorder is the mock recorder for MockPipelinedServer.
type MockPipelinedServerMockRecorder struct {
	mock *MockPipelinedServer
}

// NewMockPipelinedServer creates a new mock instance.
func NewMockPipelinedServer(ctrl *gomock.Controller) *MockPipelinedServer {
	mock := &MockPipelinedServer{ctrl: ctrl}
	mock.recorder = &MockPipelinedServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPipelinedServer) EXPECT() *MockPipelinedServerMockRecorder {
	return m.recorder
}

// ActivateFlows mocks base method.
func (m *MockPipelinedServer) ActivateFlows(arg0 context.Context, arg1 *pipelined.ActivateFlowsRequest) (*pipelined.ActivateFlowsResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ActivateFlows", arg0, arg1)
	ret0, _ := ret[0].(*pipelined.ActivateFlowsResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ActivateFlows indicates an expected call of ActivateFlows.
func (mr *MockPipelinedServerMockRecorder) ActivateFlows(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ActivateFlows", reflect.TypeOf((*MockPipelinedServer)(nil).ActivateFlows), arg0, arg1)
}

// AddUEMacFlow mocks base method.
func (m *MockPipelinedServer) AddUEMacFlow(arg0 context.Context, arg1 *pipelined.UEMacFlowRequest) (*pipelined.FlowResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddUEMacFlow", arg0, arg1)
	ret0, _ := ret[0].(*pipelined.FlowResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddUEMacFlow indicates an expected call of AddUEMacFlow.
func (mr *MockPipelinedServerMockRecorder) AddUEMacFlow(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddUEMacFlow", reflect.TypeOf((*MockPipelinedServer)(nil).AddUEMacFlow), arg0, arg1)
}

// CreateFlow mocks base method.
func (m *MockPipelinedServer) CreateFlow(arg0 context.Context, arg1 *pipelined.FlowRequest) (*pipelined.FlowResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateFlow", arg0, arg1)
	ret0, _ := ret[0].(*pipelined.FlowResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateFlow indicates an expected call of CreateFlow.
func (mr *MockPipelinedServerMockRecorder) CreateFlow(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateFlow", reflect.TypeOf((*MockPipelinedServer)(nil).CreateFlow), arg0, arg1)
}

// DeactivateFlows mocks base method.
func (m *MockPipelinedServer) DeactivateFlows(arg0 context.Context, arg1 *pipelined.DeactivateFlowsRequest) (*pipelined.DeactivateFlowsResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeactivateFlows", arg0, arg1)
	ret0, _ := ret[0].(*pipelined.DeactivateFlowsResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeactivateFlows indicates an expected call of DeactivateFlows.
func (mr *MockPipelinedServerMockRecorder) DeactivateFlows(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeactivateFlows", reflect.TypeOf((*MockPipelinedServer)(nil).DeactivateFlows), arg0, arg1)
}

// DeleteUEMacFlow mocks base method.
func (m *MockPipelinedServer) DeleteUEMacFlow(arg0 context.Context, arg1 *pipelined.UEMacFlowRequest) (*pipelined.FlowResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteUEMacFlow", arg0, arg1)
	ret0, _ := ret[0].(*pipelined.FlowResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteUEMacFlow indicates an expected call of DeleteUEMacFlow.
func (mr *MockPipelinedServerMockRecorder) DeleteUEMacFlow(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteUEMacFlow", reflect.TypeOf((*MockPipelinedServer)(nil).DeleteUEMacFlow), arg0, arg1)
}

// GetAllTableAssignments mocks base method.
func (m *MockPipelinedServer) GetAllTableAssignments(arg0 context.Context, arg1 *orc8r.Void) (*pipelined.AllTableAssignments, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllTableAssignments", arg0, arg1)
	ret0, _ := ret[0].(*pipelined.AllTableAssignments)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllTableAssignments indicates an expected call of GetAllTableAssignments.
func (mr *MockPipelinedServerMockRecorder) GetAllTableAssignments(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllTableAssignments", reflect.TypeOf((*MockPipelinedServer)(nil).GetAllTableAssignments), arg0, arg1)
}

// GetPolicyUsage mocks base method.
func (m *MockPipelinedServer) GetPolicyUsage(arg0 context.Context, arg1 *orc8r.Void) (*session_manager.RuleRecordTable, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPolicyUsage", arg0, arg1)
	ret0, _ := ret[0].(*session_manager.RuleRecordTable)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPolicyUsage indicates an expected call of GetPolicyUsage.
func (mr *MockPipelinedServerMockRecorder) GetPolicyUsage(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPolicyUsage", reflect.TypeOf((*MockPipelinedServer)(nil).GetPolicyUsage), arg0, arg1)
}

// GetStats mocks base method.
func (m *MockPipelinedServer) GetStats(arg0 context.Context, arg1 *pipelined.GetStatsRequest) (*session_manager.RuleRecordTable, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetStats", arg0, arg1)
	ret0, _ := ret[0].(*session_manager.RuleRecordTable)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetStats indicates an expected call of GetStats.
func (mr *MockPipelinedServerMockRecorder) GetStats(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetStats", reflect.TypeOf((*MockPipelinedServer)(nil).GetStats), arg0, arg1)
}

// RemoveFlow mocks base method.
func (m *MockPipelinedServer) RemoveFlow(arg0 context.Context, arg1 *pipelined.FlowRequest) (*pipelined.FlowResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveFlow", arg0, arg1)
	ret0, _ := ret[0].(*pipelined.FlowResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RemoveFlow indicates an expected call of RemoveFlow.
func (mr *MockPipelinedServerMockRecorder) RemoveFlow(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveFlow", reflect.TypeOf((*MockPipelinedServer)(nil).RemoveFlow), arg0, arg1)
}

// SetSMFSessions mocks base method.
func (m *MockPipelinedServer) SetSMFSessions(arg0 context.Context, arg1 *pipelined.SessionSet) (*pipelined.UPFSessionContextState, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetSMFSessions", arg0, arg1)
	ret0, _ := ret[0].(*pipelined.UPFSessionContextState)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SetSMFSessions indicates an expected call of SetSMFSessions.
func (mr *MockPipelinedServerMockRecorder) SetSMFSessions(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetSMFSessions", reflect.TypeOf((*MockPipelinedServer)(nil).SetSMFSessions), arg0, arg1)
}

// SetupDefaultControllers mocks base method.
func (m *MockPipelinedServer) SetupDefaultControllers(arg0 context.Context, arg1 *pipelined.SetupDefaultRequest) (*pipelined.SetupFlowsResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetupDefaultControllers", arg0, arg1)
	ret0, _ := ret[0].(*pipelined.SetupFlowsResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SetupDefaultControllers indicates an expected call of SetupDefaultControllers.
func (mr *MockPipelinedServerMockRecorder) SetupDefaultControllers(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetupDefaultControllers", reflect.TypeOf((*MockPipelinedServer)(nil).SetupDefaultControllers), arg0, arg1)
}

// SetupPolicyFlows mocks base method.
func (m *MockPipelinedServer) SetupPolicyFlows(arg0 context.Context, arg1 *pipelined.SetupPolicyRequest) (*pipelined.SetupFlowsResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetupPolicyFlows", arg0, arg1)
	ret0, _ := ret[0].(*pipelined.SetupFlowsResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SetupPolicyFlows indicates an expected call of SetupPolicyFlows.
func (mr *MockPipelinedServerMockRecorder) SetupPolicyFlows(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetupPolicyFlows", reflect.TypeOf((*MockPipelinedServer)(nil).SetupPolicyFlows), arg0, arg1)
}

// SetupQuotaFlows mocks base method.
func (m *MockPipelinedServer) SetupQuotaFlows(arg0 context.Context, arg1 *pipelined.SetupQuotaRequest) (*pipelined.SetupFlowsResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetupQuotaFlows", arg0, arg1)
	ret0, _ := ret[0].(*pipelined.SetupFlowsResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SetupQuotaFlows indicates an expected call of SetupQuotaFlows.
func (mr *MockPipelinedServerMockRecorder) SetupQuotaFlows(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetupQuotaFlows", reflect.TypeOf((*MockPipelinedServer)(nil).SetupQuotaFlows), arg0, arg1)
}

// SetupUEMacFlows mocks base method.
func (m *MockPipelinedServer) SetupUEMacFlows(arg0 context.Context, arg1 *pipelined.SetupUEMacRequest) (*pipelined.SetupFlowsResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetupUEMacFlows", arg0, arg1)
	ret0, _ := ret[0].(*pipelined.SetupFlowsResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SetupUEMacFlows indicates an expected call of SetupUEMacFlows.
func (mr *MockPipelinedServerMockRecorder) SetupUEMacFlows(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetupUEMacFlows", reflect.TypeOf((*MockPipelinedServer)(nil).SetupUEMacFlows), arg0, arg1)
}

// UpdateFlowStats mocks base method.
func (m *MockPipelinedServer) UpdateFlowStats(arg0 context.Context, arg1 *pipelined.FlowRequest) (*pipelined.FlowResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateFlowStats", arg0, arg1)
	ret0, _ := ret[0].(*pipelined.FlowResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateFlowStats indicates an expected call of UpdateFlowStats.
func (mr *MockPipelinedServerMockRecorder) UpdateFlowStats(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateFlowStats", reflect.TypeOf((*MockPipelinedServer)(nil).UpdateFlowStats), arg0, arg1)
}

// UpdateIPFIXFlow mocks base method.
func (m *MockPipelinedServer) UpdateIPFIXFlow(arg0 context.Context, arg1 *pipelined.UEMacFlowRequest) (*pipelined.FlowResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateIPFIXFlow", arg0, arg1)
	ret0, _ := ret[0].(*pipelined.FlowResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateIPFIXFlow indicates an expected call of UpdateIPFIXFlow.
func (mr *MockPipelinedServerMockRecorder) UpdateIPFIXFlow(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateIPFIXFlow", reflect.TypeOf((*MockPipelinedServer)(nil).UpdateIPFIXFlow), arg0, arg1)
}

// UpdateSubscriberQuotaState mocks base method.
func (m *MockPipelinedServer) UpdateSubscriberQuotaState(arg0 context.Context, arg1 *pipelined.UpdateSubscriberQuotaStateRequest) (*pipelined.FlowResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateSubscriberQuotaState", arg0, arg1)
	ret0, _ := ret[0].(*pipelined.FlowResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateSubscriberQuotaState indicates an expected call of UpdateSubscriberQuotaState.
func (mr *MockPipelinedServerMockRecorder) UpdateSubscriberQuotaState(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateSubscriberQuotaState", reflect.TypeOf((*MockPipelinedServer)(nil).UpdateSubscriberQuotaState), arg0, arg1)
}

// UpdateUEState mocks base method.
func (m *MockPipelinedServer) UpdateUEState(arg0 context.Context, arg1 *pipelined.UESessionSet) (*pipelined.UESessionContextResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateUEState", arg0, arg1)
	ret0, _ := ret[0].(*pipelined.UESessionContextResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateUEState indicates an expected call of UpdateUEState.
func (mr *MockPipelinedServerMockRecorder) UpdateUEState(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateUEState", reflect.TypeOf((*MockPipelinedServer)(nil).UpdateUEState), arg0, arg1)
}

// mustEmbedUnimplementedPipelinedServer mocks base method.
func (m *MockPipelinedServer) mustEmbedUnimplementedPipelinedServer() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "mustEmbedUnimplementedPipelinedServer")
}

// mustEmbedUnimplementedPipelinedServer indicates an expected call of mustEmbedUnimplementedPipelinedServer.
func (mr *MockPipelinedServerMockRecorder) mustEmbedUnimplementedPipelinedServer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "mustEmbedUnimplementedPipelinedServer", reflect.TypeOf((*MockPipelinedServer)(nil).mustEmbedUnimplementedPipelinedServer))
}

// MockUnsafePipelinedServer is a mock of UnsafePipelinedServer interface.
type MockUnsafePipelinedServer struct {
	ctrl     *gomock.Controller
	recorder *MockUnsafePipelinedServerMockRecorder
}

// MockUnsafePipelinedServerMockRecorder is the mock recorder for MockUnsafePipelinedServer.
type MockUnsafePipelinedServerMockRecorder struct {
	mock *MockUnsafePipelinedServer
}

// NewMockUnsafePipelinedServer creates a new mock instance.
func NewMockUnsafePipelinedServer(ctrl *gomock.Controller) *MockUnsafePipelinedServer {
	mock := &MockUnsafePipelinedServer{ctrl: ctrl}
	mock.recorder = &MockUnsafePipelinedServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUnsafePipelinedServer) EXPECT() *MockUnsafePipelinedServerMockRecorder {
	return m.recorder
}

// mustEmbedUnimplementedPipelinedServer mocks base method.
func (m *MockUnsafePipelinedServer) mustEmbedUnimplementedPipelinedServer() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "mustEmbedUnimplementedPipelinedServer")
}

// mustEmbedUnimplementedPipelinedServer indicates an expected call of mustEmbedUnimplementedPipelinedServer.
func (mr *MockUnsafePipelinedServerMockRecorder) mustEmbedUnimplementedPipelinedServer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "mustEmbedUnimplementedPipelinedServer", reflect.TypeOf((*MockUnsafePipelinedServer)(nil).mustEmbedUnimplementedPipelinedServer))
}