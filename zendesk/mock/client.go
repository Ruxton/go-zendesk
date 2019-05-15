// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/nukosuke/go-zendesk/v0.3/zendesk (interfaces: API)

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	zendesk "github.com/nukosuke/go-zendesk/v0.3/zendesk"
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
func (m *Client) CreateBrand(arg0 context.Context, arg1 zendesk.Brand) (zendesk.Brand, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateBrand", arg0, arg1)
	ret0, _ := ret[0].(zendesk.Brand)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateBrand indicates an expected call of CreateBrand
func (mr *ClientMockRecorder) CreateBrand(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateBrand", reflect.TypeOf((*Client)(nil).CreateBrand), arg0, arg1)
}

// CreateDynamicContentItem mocks base method
func (m *Client) CreateDynamicContentItem(arg0 context.Context, arg1 zendesk.DynamicContentItem) (zendesk.DynamicContentItem, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateDynamicContentItem", arg0, arg1)
	ret0, _ := ret[0].(zendesk.DynamicContentItem)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateDynamicContentItem indicates an expected call of CreateDynamicContentItem
func (mr *ClientMockRecorder) CreateDynamicContentItem(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateDynamicContentItem", reflect.TypeOf((*Client)(nil).CreateDynamicContentItem), arg0, arg1)
}

// CreateGroup mocks base method
func (m *Client) CreateGroup(arg0 context.Context, arg1 zendesk.Group) (zendesk.Group, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateGroup", arg0, arg1)
	ret0, _ := ret[0].(zendesk.Group)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateGroup indicates an expected call of CreateGroup
func (mr *ClientMockRecorder) CreateGroup(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateGroup", reflect.TypeOf((*Client)(nil).CreateGroup), arg0, arg1)
}

// CreateTicketField mocks base method
func (m *Client) CreateTicketField(arg0 context.Context, arg1 zendesk.TicketField) (zendesk.TicketField, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateTicketField", arg0, arg1)
	ret0, _ := ret[0].(zendesk.TicketField)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateTicketField indicates an expected call of CreateTicketField
func (mr *ClientMockRecorder) CreateTicketField(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateTicketField", reflect.TypeOf((*Client)(nil).CreateTicketField), arg0, arg1)
}

// CreateTicketForm mocks base method
func (m *Client) CreateTicketForm(arg0 context.Context, arg1 zendesk.TicketForm) (zendesk.TicketForm, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateTicketForm", arg0, arg1)
	ret0, _ := ret[0].(zendesk.TicketForm)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateTicketForm indicates an expected call of CreateTicketForm
func (mr *ClientMockRecorder) CreateTicketForm(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateTicketForm", reflect.TypeOf((*Client)(nil).CreateTicketForm), arg0, arg1)
}

// CreateTrigger mocks base method
func (m *Client) CreateTrigger(arg0 context.Context, arg1 zendesk.Trigger) (zendesk.Trigger, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateTrigger", arg0, arg1)
	ret0, _ := ret[0].(zendesk.Trigger)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateTrigger indicates an expected call of CreateTrigger
func (mr *ClientMockRecorder) CreateTrigger(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateTrigger", reflect.TypeOf((*Client)(nil).CreateTrigger), arg0, arg1)
}

// CreateUser mocks base method
func (m *Client) CreateUser(arg0 context.Context, arg1 zendesk.User) (zendesk.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUser", arg0, arg1)
	ret0, _ := ret[0].(zendesk.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateUser indicates an expected call of CreateUser
func (mr *ClientMockRecorder) CreateUser(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUser", reflect.TypeOf((*Client)(nil).CreateUser), arg0, arg1)
}

// DeleteBrand mocks base method
func (m *Client) DeleteBrand(arg0 context.Context, arg1 int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteBrand", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteBrand indicates an expected call of DeleteBrand
func (mr *ClientMockRecorder) DeleteBrand(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteBrand", reflect.TypeOf((*Client)(nil).DeleteBrand), arg0, arg1)
}

// DeleteGroup mocks base method
func (m *Client) DeleteGroup(arg0 context.Context, arg1 int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteGroup", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteGroup indicates an expected call of DeleteGroup
func (mr *ClientMockRecorder) DeleteGroup(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteGroup", reflect.TypeOf((*Client)(nil).DeleteGroup), arg0, arg1)
}

// DeleteTicketField mocks base method
func (m *Client) DeleteTicketField(arg0 context.Context, arg1 int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteTicketField", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteTicketField indicates an expected call of DeleteTicketField
func (mr *ClientMockRecorder) DeleteTicketField(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteTicketField", reflect.TypeOf((*Client)(nil).DeleteTicketField), arg0, arg1)
}

// DeleteTicketForm mocks base method
func (m *Client) DeleteTicketForm(arg0 context.Context, arg1 int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteTicketForm", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteTicketForm indicates an expected call of DeleteTicketForm
func (mr *ClientMockRecorder) DeleteTicketForm(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteTicketForm", reflect.TypeOf((*Client)(nil).DeleteTicketForm), arg0, arg1)
}

// DeleteTrigger mocks base method
func (m *Client) DeleteTrigger(arg0 context.Context, arg1 int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteTrigger", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteTrigger indicates an expected call of DeleteTrigger
func (mr *ClientMockRecorder) DeleteTrigger(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteTrigger", reflect.TypeOf((*Client)(nil).DeleteTrigger), arg0, arg1)
}

// DeleteUpload mocks base method
func (m *Client) DeleteUpload(arg0 context.Context, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteUpload", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteUpload indicates an expected call of DeleteUpload
func (mr *ClientMockRecorder) DeleteUpload(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteUpload", reflect.TypeOf((*Client)(nil).DeleteUpload), arg0, arg1)
}

// GetAttachment mocks base method
func (m *Client) GetAttachment(arg0 context.Context, arg1 int64) (zendesk.Attachment, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAttachment", arg0, arg1)
	ret0, _ := ret[0].(zendesk.Attachment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAttachment indicates an expected call of GetAttachment
func (mr *ClientMockRecorder) GetAttachment(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAttachment", reflect.TypeOf((*Client)(nil).GetAttachment), arg0, arg1)
}

// GetBrand mocks base method
func (m *Client) GetBrand(arg0 context.Context, arg1 int64) (zendesk.Brand, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBrand", arg0, arg1)
	ret0, _ := ret[0].(zendesk.Brand)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBrand indicates an expected call of GetBrand
func (mr *ClientMockRecorder) GetBrand(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBrand", reflect.TypeOf((*Client)(nil).GetBrand), arg0, arg1)
}

// GetDynamicContentItems mocks base method
func (m *Client) GetDynamicContentItems(arg0 context.Context) ([]zendesk.DynamicContentItem, zendesk.Page, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDynamicContentItems", arg0)
	ret0, _ := ret[0].([]zendesk.DynamicContentItem)
	ret1, _ := ret[1].(zendesk.Page)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetDynamicContentItems indicates an expected call of GetDynamicContentItems
func (mr *ClientMockRecorder) GetDynamicContentItems(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDynamicContentItems", reflect.TypeOf((*Client)(nil).GetDynamicContentItems), arg0)
}

// GetGroup mocks base method
func (m *Client) GetGroup(arg0 context.Context, arg1 int64) (zendesk.Group, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetGroup", arg0, arg1)
	ret0, _ := ret[0].(zendesk.Group)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetGroup indicates an expected call of GetGroup
func (mr *ClientMockRecorder) GetGroup(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetGroup", reflect.TypeOf((*Client)(nil).GetGroup), arg0, arg1)
}

// GetGroups mocks base method
func (m *Client) GetGroups(arg0 context.Context) ([]zendesk.Group, zendesk.Page, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetGroups", arg0)
	ret0, _ := ret[0].([]zendesk.Group)
	ret1, _ := ret[1].(zendesk.Page)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetGroups indicates an expected call of GetGroups
func (mr *ClientMockRecorder) GetGroups(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetGroups", reflect.TypeOf((*Client)(nil).GetGroups), arg0)
}

// GetLocales mocks base method
func (m *Client) GetLocales(arg0 context.Context) ([]zendesk.Locale, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetLocales", arg0)
	ret0, _ := ret[0].([]zendesk.Locale)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetLocales indicates an expected call of GetLocales
func (mr *ClientMockRecorder) GetLocales(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLocales", reflect.TypeOf((*Client)(nil).GetLocales), arg0)
}

// GetTicketField mocks base method
func (m *Client) GetTicketField(arg0 context.Context, arg1 int64) (zendesk.TicketField, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTicketField", arg0, arg1)
	ret0, _ := ret[0].(zendesk.TicketField)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTicketField indicates an expected call of GetTicketField
func (mr *ClientMockRecorder) GetTicketField(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTicketField", reflect.TypeOf((*Client)(nil).GetTicketField), arg0, arg1)
}

// GetTicketFields mocks base method
func (m *Client) GetTicketFields(arg0 context.Context) ([]zendesk.TicketField, zendesk.Page, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTicketFields", arg0)
	ret0, _ := ret[0].([]zendesk.TicketField)
	ret1, _ := ret[1].(zendesk.Page)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetTicketFields indicates an expected call of GetTicketFields
func (mr *ClientMockRecorder) GetTicketFields(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTicketFields", reflect.TypeOf((*Client)(nil).GetTicketFields), arg0)
}

// GetTicketForm mocks base method
func (m *Client) GetTicketForm(arg0 context.Context, arg1 int64) (zendesk.TicketForm, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTicketForm", arg0, arg1)
	ret0, _ := ret[0].(zendesk.TicketForm)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTicketForm indicates an expected call of GetTicketForm
func (mr *ClientMockRecorder) GetTicketForm(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTicketForm", reflect.TypeOf((*Client)(nil).GetTicketForm), arg0, arg1)
}

// GetTicketForms mocks base method
func (m *Client) GetTicketForms(arg0 context.Context) ([]zendesk.TicketForm, zendesk.Page, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTicketForms", arg0)
	ret0, _ := ret[0].([]zendesk.TicketForm)
	ret1, _ := ret[1].(zendesk.Page)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetTicketForms indicates an expected call of GetTicketForms
func (mr *ClientMockRecorder) GetTicketForms(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTicketForms", reflect.TypeOf((*Client)(nil).GetTicketForms), arg0)
}

// GetTrigger mocks base method
func (m *Client) GetTrigger(arg0 context.Context, arg1 int64) (zendesk.Trigger, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTrigger", arg0, arg1)
	ret0, _ := ret[0].(zendesk.Trigger)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTrigger indicates an expected call of GetTrigger
func (mr *ClientMockRecorder) GetTrigger(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTrigger", reflect.TypeOf((*Client)(nil).GetTrigger), arg0, arg1)
}

// GetTriggers mocks base method
func (m *Client) GetTriggers(arg0 context.Context) ([]zendesk.Trigger, zendesk.Page, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTriggers", arg0)
	ret0, _ := ret[0].([]zendesk.Trigger)
	ret1, _ := ret[1].(zendesk.Page)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetTriggers indicates an expected call of GetTriggers
func (mr *ClientMockRecorder) GetTriggers(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTriggers", reflect.TypeOf((*Client)(nil).GetTriggers), arg0)
}

// GetUsers mocks base method
func (m *Client) GetUsers(arg0 context.Context) ([]zendesk.User, zendesk.Page, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUsers", arg0)
	ret0, _ := ret[0].([]zendesk.User)
	ret1, _ := ret[1].(zendesk.Page)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetUsers indicates an expected call of GetUsers
func (mr *ClientMockRecorder) GetUsers(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUsers", reflect.TypeOf((*Client)(nil).GetUsers), arg0)
}

// UpdateBrand mocks base method
func (m *Client) UpdateBrand(arg0 context.Context, arg1 int64, arg2 zendesk.Brand) (zendesk.Brand, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateBrand", arg0, arg1, arg2)
	ret0, _ := ret[0].(zendesk.Brand)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateBrand indicates an expected call of UpdateBrand
func (mr *ClientMockRecorder) UpdateBrand(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateBrand", reflect.TypeOf((*Client)(nil).UpdateBrand), arg0, arg1, arg2)
}

// UpdateGroup mocks base method
func (m *Client) UpdateGroup(arg0 context.Context, arg1 int64, arg2 zendesk.Group) (zendesk.Group, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateGroup", arg0, arg1, arg2)
	ret0, _ := ret[0].(zendesk.Group)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateGroup indicates an expected call of UpdateGroup
func (mr *ClientMockRecorder) UpdateGroup(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateGroup", reflect.TypeOf((*Client)(nil).UpdateGroup), arg0, arg1, arg2)
}

// UpdateTicketField mocks base method
func (m *Client) UpdateTicketField(arg0 context.Context, arg1 int64, arg2 zendesk.TicketField) (zendesk.TicketField, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateTicketField", arg0, arg1, arg2)
	ret0, _ := ret[0].(zendesk.TicketField)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateTicketField indicates an expected call of UpdateTicketField
func (mr *ClientMockRecorder) UpdateTicketField(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateTicketField", reflect.TypeOf((*Client)(nil).UpdateTicketField), arg0, arg1, arg2)
}

// UpdateTicketForm mocks base method
func (m *Client) UpdateTicketForm(arg0 context.Context, arg1 int64, arg2 zendesk.TicketForm) (zendesk.TicketForm, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateTicketForm", arg0, arg1, arg2)
	ret0, _ := ret[0].(zendesk.TicketForm)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateTicketForm indicates an expected call of UpdateTicketForm
func (mr *ClientMockRecorder) UpdateTicketForm(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateTicketForm", reflect.TypeOf((*Client)(nil).UpdateTicketForm), arg0, arg1, arg2)
}

// UpdateTrigger mocks base method
func (m *Client) UpdateTrigger(arg0 context.Context, arg1 int64, arg2 zendesk.Trigger) (zendesk.Trigger, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateTrigger", arg0, arg1, arg2)
	ret0, _ := ret[0].(zendesk.Trigger)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateTrigger indicates an expected call of UpdateTrigger
func (mr *ClientMockRecorder) UpdateTrigger(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateTrigger", reflect.TypeOf((*Client)(nil).UpdateTrigger), arg0, arg1, arg2)
}

// UploadAttachment mocks base method
func (m *Client) UploadAttachment(arg0 context.Context, arg1, arg2 string) zendesk.UploadWriter {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UploadAttachment", arg0, arg1, arg2)
	ret0, _ := ret[0].(zendesk.UploadWriter)
	return ret0
}

// UploadAttachment indicates an expected call of UploadAttachment
func (mr *ClientMockRecorder) UploadAttachment(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UploadAttachment", reflect.TypeOf((*Client)(nil).UploadAttachment), arg0, arg1, arg2)
}
