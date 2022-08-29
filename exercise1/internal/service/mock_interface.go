// Code generated by MockGen. DO NOT EDIT.
// Source: interface.go

// Package service is a generated GoMock package.
package service

import (
	reflect "reflect"
	models "trainig/exercise1/internal/models"

	gofr "developer.zopsmart.com/go/gofr/pkg/gofr"
	gomock "github.com/golang/mock/gomock"
)

// MockProductStore is a mock of ProductStore interface.
type MockProductStore struct {
	ctrl     *gomock.Controller
	recorder *MockProductStoreMockRecorder
}

// MockProductStoreMockRecorder is the mock recorder for MockProductStore.
type MockProductStoreMockRecorder struct {
	mock *MockProductStore
}

// NewMockProductStore creates a new mock instance.
func NewMockProductStore(ctrl *gomock.Controller) *MockProductStore {
	mock := &MockProductStore{ctrl: ctrl}
	mock.recorder = &MockProductStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockProductStore) EXPECT() *MockProductStoreMockRecorder {
	return m.recorder
}

// AddProduct mocks base method.
func (m *MockProductStore) AddProduct(ctx *gofr.Context, product *models.Product) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddProduct", ctx, product)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddProduct indicates an expected call of AddProduct.
func (mr *MockProductStoreMockRecorder) AddProduct(ctx, product interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddProduct", reflect.TypeOf((*MockProductStore)(nil).AddProduct), ctx, product)
}

// AddVariant mocks base method.
func (m *MockProductStore) AddVariant(ctx *gofr.Context, variant *models.Variant) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddVariant", ctx, variant)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddVariant indicates an expected call of AddVariant.
func (mr *MockProductStoreMockRecorder) AddVariant(ctx, variant interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddVariant", reflect.TypeOf((*MockProductStore)(nil).AddVariant), ctx, variant)
}

// GetProduct mocks base method.
func (m *MockProductStore) GetProduct(ctx *gofr.Context, product *models.Product, id int) models.Product {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetProduct", ctx, product, id)
	ret0, _ := ret[0].(models.Product)
	return ret0
}

// GetProduct indicates an expected call of GetProduct.
func (mr *MockProductStoreMockRecorder) GetProduct(ctx, product, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetProduct", reflect.TypeOf((*MockProductStore)(nil).GetProduct), ctx, product, id)
}

// GetProductVariants mocks base method.
func (m *MockProductStore) GetProductVariants(ctx *gofr.Context, pid int) []models.Variant {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetProductVariants", ctx, pid)
	ret0, _ := ret[0].([]models.Variant)
	return ret0
}

// GetProductVariants indicates an expected call of GetProductVariants.
func (mr *MockProductStoreMockRecorder) GetProductVariants(ctx, pid interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetProductVariants", reflect.TypeOf((*MockProductStore)(nil).GetProductVariants), ctx, pid)
}

// GetVariant mocks base method.
func (m *MockProductStore) GetVariant(ctx *gofr.Context, productID, variantID string) *models.Variant {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetVariant", ctx, productID, variantID)
	ret0, _ := ret[0].(*models.Variant)
	return ret0
}

// GetVariant indicates an expected call of GetVariant.
func (mr *MockProductStoreMockRecorder) GetVariant(ctx, productID, variantID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetVariant", reflect.TypeOf((*MockProductStore)(nil).GetVariant), ctx, productID, variantID)
}
