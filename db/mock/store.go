// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/shawgichan/tourist/db/sqlc (interfaces: Store)
//
// Generated by this command:
//
//	mockgen -package mockdb -destination db/mock/store.go github.com/shawgichan/tourist/db/sqlc Store
//

// Package mockdb is a generated GoMock package.
package mockdb

import (
	context "context"
	reflect "reflect"

	db "github.com/shawgichan/tourist/db/sqlc"
	gomock "go.uber.org/mock/gomock"
)

// MockStore is a mock of Store interface.
type MockStore struct {
	ctrl     *gomock.Controller
	recorder *MockStoreMockRecorder
}

// MockStoreMockRecorder is the mock recorder for MockStore.
type MockStoreMockRecorder struct {
	mock *MockStore
}

// NewMockStore creates a new mock instance.
func NewMockStore(ctrl *gomock.Controller) *MockStore {
	mock := &MockStore{ctrl: ctrl}
	mock.recorder = &MockStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStore) EXPECT() *MockStoreMockRecorder {
	return m.recorder
}

// CheckUsernameAndEmail mocks base method.
func (m *MockStore) CheckUsernameAndEmail(arg0 context.Context, arg1 db.CheckUsernameAndEmailParams) (db.CheckUsernameAndEmailRow, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckUsernameAndEmail", arg0, arg1)
	ret0, _ := ret[0].(db.CheckUsernameAndEmailRow)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CheckUsernameAndEmail indicates an expected call of CheckUsernameAndEmail.
func (mr *MockStoreMockRecorder) CheckUsernameAndEmail(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckUsernameAndEmail", reflect.TypeOf((*MockStore)(nil).CheckUsernameAndEmail), arg0, arg1)
}

// CreateLocation mocks base method.
func (m *MockStore) CreateLocation(arg0 context.Context, arg1 db.CreateLocationParams) (db.Location, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateLocation", arg0, arg1)
	ret0, _ := ret[0].(db.Location)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateLocation indicates an expected call of CreateLocation.
func (mr *MockStoreMockRecorder) CreateLocation(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateLocation", reflect.TypeOf((*MockStore)(nil).CreateLocation), arg0, arg1)
}

// CreatePlace mocks base method.
func (m *MockStore) CreatePlace(arg0 context.Context, arg1 db.CreatePlaceParams) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreatePlace", arg0, arg1)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreatePlace indicates an expected call of CreatePlace.
func (mr *MockStoreMockRecorder) CreatePlace(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreatePlace", reflect.TypeOf((*MockStore)(nil).CreatePlace), arg0, arg1)
}

// CreateProfile mocks base method.
func (m *MockStore) CreateProfile(arg0 context.Context, arg1 db.CreateProfileParams) (db.Profile, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateProfile", arg0, arg1)
	ret0, _ := ret[0].(db.Profile)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateProfile indicates an expected call of CreateProfile.
func (mr *MockStoreMockRecorder) CreateProfile(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateProfile", reflect.TypeOf((*MockStore)(nil).CreateProfile), arg0, arg1)
}

// CreateUser mocks base method.
func (m *MockStore) CreateUser(arg0 context.Context, arg1 db.CreateUserParams) (db.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUser", arg0, arg1)
	ret0, _ := ret[0].(db.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateUser indicates an expected call of CreateUser.
func (mr *MockStoreMockRecorder) CreateUser(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUser", reflect.TypeOf((*MockStore)(nil).CreateUser), arg0, arg1)
}

// GetPlace mocks base method.
func (m *MockStore) GetPlace(arg0 context.Context, arg1 int64) (db.Place, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPlace", arg0, arg1)
	ret0, _ := ret[0].(db.Place)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPlace indicates an expected call of GetPlace.
func (mr *MockStoreMockRecorder) GetPlace(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPlace", reflect.TypeOf((*MockStore)(nil).GetPlace), arg0, arg1)
}

// GetPlaces mocks base method.
func (m *MockStore) GetPlaces(arg0 context.Context) ([]db.Place, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPlaces", arg0)
	ret0, _ := ret[0].([]db.Place)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPlaces indicates an expected call of GetPlaces.
func (mr *MockStoreMockRecorder) GetPlaces(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPlaces", reflect.TypeOf((*MockStore)(nil).GetPlaces), arg0)
}

// GetUserByName mocks base method.
func (m *MockStore) GetUserByName(arg0 context.Context, arg1 string) (db.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserByName", arg0, arg1)
	ret0, _ := ret[0].(db.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserByName indicates an expected call of GetUserByName.
func (mr *MockStoreMockRecorder) GetUserByName(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserByName", reflect.TypeOf((*MockStore)(nil).GetUserByName), arg0, arg1)
}
