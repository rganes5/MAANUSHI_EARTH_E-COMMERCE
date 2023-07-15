// Code generated by MockGen. DO NOT EDIT.
// Source: pkg/repository/interface/admin.go

// Package mockRepo is a generated GoMock package.
package mockRepo

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	domain "github.com/rganes5/maanushi_earth_e-commerce/pkg/domain"
	utils "github.com/rganes5/maanushi_earth_e-commerce/pkg/utils"
)

// MockAdminRepository is a mock of AdminRepository interface.
type MockAdminRepository struct {
	ctrl     *gomock.Controller
	recorder *MockAdminRepositoryMockRecorder
}

// MockAdminRepositoryMockRecorder is the mock recorder for MockAdminRepository.
type MockAdminRepositoryMockRecorder struct {
	mock *MockAdminRepository
}

// NewMockAdminRepository creates a new mock instance.
func NewMockAdminRepository(ctrl *gomock.Controller) *MockAdminRepository {
	mock := &MockAdminRepository{ctrl: ctrl}
	mock.recorder = &MockAdminRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAdminRepository) EXPECT() *MockAdminRepositoryMockRecorder {
	return m.recorder
}

// AccessHandler mocks base method.
func (m *MockAdminRepository) AccessHandler(ctx context.Context, id string, access bool) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AccessHandler", ctx, id, access)
	ret0, _ := ret[0].(error)
	return ret0
}

// AccessHandler indicates an expected call of AccessHandler.
func (mr *MockAdminRepositoryMockRecorder) AccessHandler(ctx, id, access interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AccessHandler", reflect.TypeOf((*MockAdminRepository)(nil).AccessHandler), ctx, id, access)
}

// AddCoupon mocks base method.
func (m *MockAdminRepository) AddCoupon(ctx context.Context, coupon domain.Coupon) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddCoupon", ctx, coupon)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddCoupon indicates an expected call of AddCoupon.
func (mr *MockAdminRepositoryMockRecorder) AddCoupon(ctx, coupon interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddCoupon", reflect.TypeOf((*MockAdminRepository)(nil).AddCoupon), ctx, coupon)
}

// Dashboard mocks base method.
func (m *MockAdminRepository) Dashboard(ctx context.Context) (utils.ResponseWidgets, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Dashboard", ctx)
	ret0, _ := ret[0].(utils.ResponseWidgets)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Dashboard indicates an expected call of Dashboard.
func (mr *MockAdminRepositoryMockRecorder) Dashboard(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Dashboard", reflect.TypeOf((*MockAdminRepository)(nil).Dashboard), ctx)
}

// DeleteCoupon mocks base method.
func (m *MockAdminRepository) DeleteCoupon(ctx context.Context, couponId string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteCoupon", ctx, couponId)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteCoupon indicates an expected call of DeleteCoupon.
func (mr *MockAdminRepositoryMockRecorder) DeleteCoupon(ctx, couponId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteCoupon", reflect.TypeOf((*MockAdminRepository)(nil).DeleteCoupon), ctx, couponId)
}

// FindByEmail mocks base method.
func (m *MockAdminRepository) FindByEmail(ctx context.Context, Email string) (domain.Admin, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByEmail", ctx, Email)
	ret0, _ := ret[0].(domain.Admin)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByEmail indicates an expected call of FindByEmail.
func (mr *MockAdminRepositoryMockRecorder) FindByEmail(ctx, Email interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByEmail", reflect.TypeOf((*MockAdminRepository)(nil).FindByEmail), ctx, Email)
}

// GetAllCoupons mocks base method.
func (m *MockAdminRepository) GetAllCoupons(ctx context.Context, pagination utils.Pagination) ([]domain.Coupon, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllCoupons", ctx, pagination)
	ret0, _ := ret[0].([]domain.Coupon)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllCoupons indicates an expected call of GetAllCoupons.
func (mr *MockAdminRepositoryMockRecorder) GetAllCoupons(ctx, pagination interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllCoupons", reflect.TypeOf((*MockAdminRepository)(nil).GetAllCoupons), ctx, pagination)
}

// ListUsers mocks base method.
func (m *MockAdminRepository) ListUsers(ctx context.Context, pagination utils.Pagination) ([]utils.ResponseUsers, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListUsers", ctx, pagination)
	ret0, _ := ret[0].([]utils.ResponseUsers)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListUsers indicates an expected call of ListUsers.
func (mr *MockAdminRepositoryMockRecorder) ListUsers(ctx, pagination interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListUsers", reflect.TypeOf((*MockAdminRepository)(nil).ListUsers), ctx, pagination)
}

// SalesReport mocks base method.
func (m *MockAdminRepository) SalesReport(arg0 utils.SalesReport) ([]utils.ResponseSalesReport, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SalesReport", arg0)
	ret0, _ := ret[0].([]utils.ResponseSalesReport)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SalesReport indicates an expected call of SalesReport.
func (mr *MockAdminRepositoryMockRecorder) SalesReport(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SalesReport", reflect.TypeOf((*MockAdminRepository)(nil).SalesReport), arg0)
}

// SignUpAdmin mocks base method.
func (m *MockAdminRepository) SignUpAdmin(ctx context.Context, body utils.AdminSignUp) (domain.Admin, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SignUpAdmin", ctx, body)
	ret0, _ := ret[0].(domain.Admin)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SignUpAdmin indicates an expected call of SignUpAdmin.
func (mr *MockAdminRepositoryMockRecorder) SignUpAdmin(ctx, body interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SignUpAdmin", reflect.TypeOf((*MockAdminRepository)(nil).SignUpAdmin), ctx, body)
}

// UpdateCoupon mocks base method.
func (m *MockAdminRepository) UpdateCoupon(ctx context.Context, coupon domain.Coupon, couponId string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateCoupon", ctx, coupon, couponId)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateCoupon indicates an expected call of UpdateCoupon.
func (mr *MockAdminRepositoryMockRecorder) UpdateCoupon(ctx, coupon, couponId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateCoupon", reflect.TypeOf((*MockAdminRepository)(nil).UpdateCoupon), ctx, coupon, couponId)
}