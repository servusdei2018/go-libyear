// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/nieomylnieja/go-libyear (interfaces: ModulesRepo,VersionsGetter)
//
// Generated by this command:
//
//	mockgen -destination internal/mocks/mocks.go -package mocks -typed . ModulesRepo,VersionsGetter
//
// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	semver "github.com/Masterminds/semver"
	gomock "go.uber.org/mock/gomock"

	internal "github.com/nieomylnieja/go-libyear/internal"
)

// MockModulesRepo is a mock of ModulesRepo interface.
type MockModulesRepo struct {
	ctrl     *gomock.Controller
	recorder *MockModulesRepoMockRecorder
}

// MockModulesRepoMockRecorder is the mock recorder for MockModulesRepo.
type MockModulesRepoMockRecorder struct {
	mock *MockModulesRepo
}

// NewMockModulesRepo creates a new mock instance.
func NewMockModulesRepo(ctrl *gomock.Controller) *MockModulesRepo {
	mock := &MockModulesRepo{ctrl: ctrl}
	mock.recorder = &MockModulesRepoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockModulesRepo) EXPECT() *MockModulesRepoMockRecorder {
	return m.recorder
}

// GetInfo mocks base method.
func (m *MockModulesRepo) GetInfo(arg0 string, arg1 *semver.Version) (*internal.Module, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetInfo", arg0, arg1)
	ret0, _ := ret[0].(*internal.Module)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetInfo indicates an expected call of GetInfo.
func (mr *MockModulesRepoMockRecorder) GetInfo(arg0, arg1 any) *ModulesRepoGetInfoCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetInfo", reflect.TypeOf((*MockModulesRepo)(nil).GetInfo), arg0, arg1)
	return &ModulesRepoGetInfoCall{Call: call}
}

// ModulesRepoGetInfoCall wrap *gomock.Call
type ModulesRepoGetInfoCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *ModulesRepoGetInfoCall) Return(arg0 *internal.Module, arg1 error) *ModulesRepoGetInfoCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *ModulesRepoGetInfoCall) Do(f func(string, *semver.Version) (*internal.Module, error)) *ModulesRepoGetInfoCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *ModulesRepoGetInfoCall) DoAndReturn(f func(string, *semver.Version) (*internal.Module, error)) *ModulesRepoGetInfoCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// GetLatestInfo mocks base method.
func (m *MockModulesRepo) GetLatestInfo(arg0 string) (*internal.Module, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetLatestInfo", arg0)
	ret0, _ := ret[0].(*internal.Module)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetLatestInfo indicates an expected call of GetLatestInfo.
func (mr *MockModulesRepoMockRecorder) GetLatestInfo(arg0 any) *ModulesRepoGetLatestInfoCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLatestInfo", reflect.TypeOf((*MockModulesRepo)(nil).GetLatestInfo), arg0)
	return &ModulesRepoGetLatestInfoCall{Call: call}
}

// ModulesRepoGetLatestInfoCall wrap *gomock.Call
type ModulesRepoGetLatestInfoCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *ModulesRepoGetLatestInfoCall) Return(arg0 *internal.Module, arg1 error) *ModulesRepoGetLatestInfoCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *ModulesRepoGetLatestInfoCall) Do(f func(string) (*internal.Module, error)) *ModulesRepoGetLatestInfoCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *ModulesRepoGetLatestInfoCall) DoAndReturn(f func(string) (*internal.Module, error)) *ModulesRepoGetLatestInfoCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// GetModFile mocks base method.
func (m *MockModulesRepo) GetModFile(arg0 string, arg1 *semver.Version) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetModFile", arg0, arg1)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetModFile indicates an expected call of GetModFile.
func (mr *MockModulesRepoMockRecorder) GetModFile(arg0, arg1 any) *ModulesRepoGetModFileCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetModFile", reflect.TypeOf((*MockModulesRepo)(nil).GetModFile), arg0, arg1)
	return &ModulesRepoGetModFileCall{Call: call}
}

// ModulesRepoGetModFileCall wrap *gomock.Call
type ModulesRepoGetModFileCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *ModulesRepoGetModFileCall) Return(arg0 []byte, arg1 error) *ModulesRepoGetModFileCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *ModulesRepoGetModFileCall) Do(f func(string, *semver.Version) ([]byte, error)) *ModulesRepoGetModFileCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *ModulesRepoGetModFileCall) DoAndReturn(f func(string, *semver.Version) ([]byte, error)) *ModulesRepoGetModFileCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// GetVersions mocks base method.
func (m *MockModulesRepo) GetVersions(arg0 string) ([]*semver.Version, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetVersions", arg0)
	ret0, _ := ret[0].([]*semver.Version)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetVersions indicates an expected call of GetVersions.
func (mr *MockModulesRepoMockRecorder) GetVersions(arg0 any) *ModulesRepoGetVersionsCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetVersions", reflect.TypeOf((*MockModulesRepo)(nil).GetVersions), arg0)
	return &ModulesRepoGetVersionsCall{Call: call}
}

// ModulesRepoGetVersionsCall wrap *gomock.Call
type ModulesRepoGetVersionsCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *ModulesRepoGetVersionsCall) Return(arg0 []*semver.Version, arg1 error) *ModulesRepoGetVersionsCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *ModulesRepoGetVersionsCall) Do(f func(string) ([]*semver.Version, error)) *ModulesRepoGetVersionsCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *ModulesRepoGetVersionsCall) DoAndReturn(f func(string) ([]*semver.Version, error)) *ModulesRepoGetVersionsCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// MockVersionsGetter is a mock of VersionsGetter interface.
type MockVersionsGetter struct {
	ctrl     *gomock.Controller
	recorder *MockVersionsGetterMockRecorder
}

// MockVersionsGetterMockRecorder is the mock recorder for MockVersionsGetter.
type MockVersionsGetterMockRecorder struct {
	mock *MockVersionsGetter
}

// NewMockVersionsGetter creates a new mock instance.
func NewMockVersionsGetter(ctrl *gomock.Controller) *MockVersionsGetter {
	mock := &MockVersionsGetter{ctrl: ctrl}
	mock.recorder = &MockVersionsGetterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockVersionsGetter) EXPECT() *MockVersionsGetterMockRecorder {
	return m.recorder
}

// GetVersions mocks base method.
func (m *MockVersionsGetter) GetVersions(arg0 string) ([]*semver.Version, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetVersions", arg0)
	ret0, _ := ret[0].([]*semver.Version)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetVersions indicates an expected call of GetVersions.
func (mr *MockVersionsGetterMockRecorder) GetVersions(arg0 any) *VersionsGetterGetVersionsCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetVersions", reflect.TypeOf((*MockVersionsGetter)(nil).GetVersions), arg0)
	return &VersionsGetterGetVersionsCall{Call: call}
}

// VersionsGetterGetVersionsCall wrap *gomock.Call
type VersionsGetterGetVersionsCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *VersionsGetterGetVersionsCall) Return(arg0 []*semver.Version, arg1 error) *VersionsGetterGetVersionsCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *VersionsGetterGetVersionsCall) Do(f func(string) ([]*semver.Version, error)) *VersionsGetterGetVersionsCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *VersionsGetterGetVersionsCall) DoAndReturn(f func(string) ([]*semver.Version, error)) *VersionsGetterGetVersionsCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}