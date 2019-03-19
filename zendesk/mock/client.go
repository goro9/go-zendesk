// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/nukosuke/go-zendesk/zendesk (interfaces: API)

// Package mock is a generated GoMock package.
package mock

import (
	gomock "github.com/golang/mock/gomock"
	zendesk "github.com/nukosuke/go-zendesk/zendesk"
	reflect "reflect"
)

// Client is a mock of API interface
type Client struct {
	ctrl     *gomock.Controller
	recorder *ClientMockRecorder
}

// ClientMockRecorder is the mock recorder for Client
type ClientMockRecorder struct {
	mock *Client
}

// NewClient creates a new mock instance
func NewClient(ctrl *gomock.Controller) *Client {
	mock := &Client{ctrl: ctrl}
	mock.recorder = &ClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *Client) EXPECT() *ClientMockRecorder {
	return m.recorder
}

// CreateBrand mocks base method
func (m *Client) CreateBrand(arg0 zendesk.Brand) (zendesk.Brand, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateBrand", arg0)
	ret0, _ := ret[0].(zendesk.Brand)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateBrand indicates an expected call of CreateBrand
func (mr *ClientMockRecorder) CreateBrand(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateBrand", reflect.TypeOf((*Client)(nil).CreateBrand), arg0)
}

// CreateGroup mocks base method
func (m *Client) CreateGroup(arg0 zendesk.Group) (zendesk.Group, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateGroup", arg0)
	ret0, _ := ret[0].(zendesk.Group)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateGroup indicates an expected call of CreateGroup
func (mr *ClientMockRecorder) CreateGroup(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateGroup", reflect.TypeOf((*Client)(nil).CreateGroup), arg0)
}

// CreateTicketField mocks base method
func (m *Client) CreateTicketField(arg0 zendesk.TicketField) (zendesk.TicketField, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateTicketField", arg0)
	ret0, _ := ret[0].(zendesk.TicketField)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateTicketField indicates an expected call of CreateTicketField
func (mr *ClientMockRecorder) CreateTicketField(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateTicketField", reflect.TypeOf((*Client)(nil).CreateTicketField), arg0)
}

// CreateTicketForm mocks base method
func (m *Client) CreateTicketForm(arg0 zendesk.TicketForm) (zendesk.TicketForm, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateTicketForm", arg0)
	ret0, _ := ret[0].(zendesk.TicketForm)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateTicketForm indicates an expected call of CreateTicketForm
func (mr *ClientMockRecorder) CreateTicketForm(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateTicketForm", reflect.TypeOf((*Client)(nil).CreateTicketForm), arg0)
}

// CreateTrigger mocks base method
func (m *Client) CreateTrigger(arg0 zendesk.Trigger) (zendesk.Trigger, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateTrigger", arg0)
	ret0, _ := ret[0].(zendesk.Trigger)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateTrigger indicates an expected call of CreateTrigger
func (mr *ClientMockRecorder) CreateTrigger(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateTrigger", reflect.TypeOf((*Client)(nil).CreateTrigger), arg0)
}

// CreateUser mocks base method
func (m *Client) CreateUser(arg0 zendesk.User) (zendesk.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUser", arg0)
	ret0, _ := ret[0].(zendesk.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateUser indicates an expected call of CreateUser
func (mr *ClientMockRecorder) CreateUser(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUser", reflect.TypeOf((*Client)(nil).CreateUser), arg0)
}

// DeleteBrand mocks base method
func (m *Client) DeleteBrand(arg0 int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteBrand", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteBrand indicates an expected call of DeleteBrand
func (mr *ClientMockRecorder) DeleteBrand(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteBrand", reflect.TypeOf((*Client)(nil).DeleteBrand), arg0)
}

// DeleteGroup mocks base method
func (m *Client) DeleteGroup(arg0 int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteGroup", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteGroup indicates an expected call of DeleteGroup
func (mr *ClientMockRecorder) DeleteGroup(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteGroup", reflect.TypeOf((*Client)(nil).DeleteGroup), arg0)
}

// DeleteTicketField mocks base method
func (m *Client) DeleteTicketField(arg0 int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteTicketField", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteTicketField indicates an expected call of DeleteTicketField
func (mr *ClientMockRecorder) DeleteTicketField(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteTicketField", reflect.TypeOf((*Client)(nil).DeleteTicketField), arg0)
}

// GetBrand mocks base method
func (m *Client) GetBrand(arg0 int64) (zendesk.Brand, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBrand", arg0)
	ret0, _ := ret[0].(zendesk.Brand)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBrand indicates an expected call of GetBrand
func (mr *ClientMockRecorder) GetBrand(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBrand", reflect.TypeOf((*Client)(nil).GetBrand), arg0)
}

// GetGroup mocks base method
func (m *Client) GetGroup(arg0 int64) (zendesk.Group, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetGroup", arg0)
	ret0, _ := ret[0].(zendesk.Group)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetGroup indicates an expected call of GetGroup
func (mr *ClientMockRecorder) GetGroup(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetGroup", reflect.TypeOf((*Client)(nil).GetGroup), arg0)
}

// GetGroups mocks base method
func (m *Client) GetGroups() ([]zendesk.Group, zendesk.Page, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetGroups")
	ret0, _ := ret[0].([]zendesk.Group)
	ret1, _ := ret[1].(zendesk.Page)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetGroups indicates an expected call of GetGroups
func (mr *ClientMockRecorder) GetGroups() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetGroups", reflect.TypeOf((*Client)(nil).GetGroups))
}

// GetLocales mocks base method
func (m *Client) GetLocales() ([]zendesk.Locale, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetLocales")
	ret0, _ := ret[0].([]zendesk.Locale)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetLocales indicates an expected call of GetLocales
func (mr *ClientMockRecorder) GetLocales() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLocales", reflect.TypeOf((*Client)(nil).GetLocales))
}

// GetTicketField mocks base method
func (m *Client) GetTicketField(arg0 int64) (zendesk.TicketField, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTicketField", arg0)
	ret0, _ := ret[0].(zendesk.TicketField)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTicketField indicates an expected call of GetTicketField
func (mr *ClientMockRecorder) GetTicketField(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTicketField", reflect.TypeOf((*Client)(nil).GetTicketField), arg0)
}

// GetTicketFields mocks base method
func (m *Client) GetTicketFields() ([]zendesk.TicketField, zendesk.Page, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTicketFields")
	ret0, _ := ret[0].([]zendesk.TicketField)
	ret1, _ := ret[1].(zendesk.Page)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetTicketFields indicates an expected call of GetTicketFields
func (mr *ClientMockRecorder) GetTicketFields() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTicketFields", reflect.TypeOf((*Client)(nil).GetTicketFields))
}

// GetTicketForms mocks base method
func (m *Client) GetTicketForms() ([]zendesk.TicketForm, zendesk.Page, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTicketForms")
	ret0, _ := ret[0].([]zendesk.TicketForm)
	ret1, _ := ret[1].(zendesk.Page)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetTicketForms indicates an expected call of GetTicketForms
func (mr *ClientMockRecorder) GetTicketForms() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTicketForms", reflect.TypeOf((*Client)(nil).GetTicketForms))
}

// GetTriggers mocks base method
func (m *Client) GetTriggers() ([]zendesk.Trigger, zendesk.Page, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTriggers")
	ret0, _ := ret[0].([]zendesk.Trigger)
	ret1, _ := ret[1].(zendesk.Page)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetTriggers indicates an expected call of GetTriggers
func (mr *ClientMockRecorder) GetTriggers() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTriggers", reflect.TypeOf((*Client)(nil).GetTriggers))
}

// GetUsers mocks base method
func (m *Client) GetUsers() ([]zendesk.User, zendesk.Page, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUsers")
	ret0, _ := ret[0].([]zendesk.User)
	ret1, _ := ret[1].(zendesk.Page)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetUsers indicates an expected call of GetUsers
func (mr *ClientMockRecorder) GetUsers() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUsers", reflect.TypeOf((*Client)(nil).GetUsers))
}

// UpdateBrand mocks base method
func (m *Client) UpdateBrand(arg0 int64, arg1 zendesk.Brand) (zendesk.Brand, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateBrand", arg0, arg1)
	ret0, _ := ret[0].(zendesk.Brand)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateBrand indicates an expected call of UpdateBrand
func (mr *ClientMockRecorder) UpdateBrand(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateBrand", reflect.TypeOf((*Client)(nil).UpdateBrand), arg0, arg1)
}

// UpdateGroup mocks base method
func (m *Client) UpdateGroup(arg0 int64, arg1 zendesk.Group) (zendesk.Group, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateGroup", arg0, arg1)
	ret0, _ := ret[0].(zendesk.Group)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateGroup indicates an expected call of UpdateGroup
func (mr *ClientMockRecorder) UpdateGroup(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateGroup", reflect.TypeOf((*Client)(nil).UpdateGroup), arg0, arg1)
}

// UpdateTicketField mocks base method
func (m *Client) UpdateTicketField(arg0 int64, arg1 zendesk.TicketField) (zendesk.TicketField, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateTicketField", arg0, arg1)
	ret0, _ := ret[0].(zendesk.TicketField)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateTicketField indicates an expected call of UpdateTicketField
func (mr *ClientMockRecorder) UpdateTicketField(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateTicketField", reflect.TypeOf((*Client)(nil).UpdateTicketField), arg0, arg1)
}

// UploadAttachment mocks base method
func (m *Client) UploadAttachment(arg0, arg1 string) zendesk.UploadWriter {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UploadAttachment", arg0, arg1)
	ret0, _ := ret[0].(zendesk.UploadWriter)
	return ret0
}

// UploadAttachment indicates an expected call of UploadAttachment
func (mr *ClientMockRecorder) UploadAttachment(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UploadAttachment", reflect.TypeOf((*Client)(nil).UploadAttachment), arg0, arg1)
}
