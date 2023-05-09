// /*
// Copyright 2023 The KEDA Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
// */
//

// Code generated by MockGen. DO NOT EDIT.
// Source: operator/generated/clientset/versioned/typed/http/v1alpha1/httpscaledobject.go

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	v1alpha1 "github.com/kedacore/http-add-on/operator/apis/http/v1alpha1"
	v1alpha10 "github.com/kedacore/http-add-on/operator/generated/clientset/versioned/typed/http/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
)

// MockHTTPScaledObjectsGetter is a mock of HTTPScaledObjectsGetter interface.
type MockHTTPScaledObjectsGetter struct {
	ctrl     *gomock.Controller
	recorder *MockHTTPScaledObjectsGetterMockRecorder
}

// MockHTTPScaledObjectsGetterMockRecorder is the mock recorder for MockHTTPScaledObjectsGetter.
type MockHTTPScaledObjectsGetterMockRecorder struct {
	mock *MockHTTPScaledObjectsGetter
}

// NewMockHTTPScaledObjectsGetter creates a new mock instance.
func NewMockHTTPScaledObjectsGetter(ctrl *gomock.Controller) *MockHTTPScaledObjectsGetter {
	mock := &MockHTTPScaledObjectsGetter{ctrl: ctrl}
	mock.recorder = &MockHTTPScaledObjectsGetterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockHTTPScaledObjectsGetter) EXPECT() *MockHTTPScaledObjectsGetterMockRecorder {
	return m.recorder
}

// HTTPScaledObjects mocks base method.
func (m *MockHTTPScaledObjectsGetter) HTTPScaledObjects(namespace string) v1alpha10.HTTPScaledObjectInterface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HTTPScaledObjects", namespace)
	ret0, _ := ret[0].(v1alpha10.HTTPScaledObjectInterface)
	return ret0
}

// HTTPScaledObjects indicates an expected call of HTTPScaledObjects.
func (mr *MockHTTPScaledObjectsGetterMockRecorder) HTTPScaledObjects(namespace interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HTTPScaledObjects", reflect.TypeOf((*MockHTTPScaledObjectsGetter)(nil).HTTPScaledObjects), namespace)
}

// MockHTTPScaledObjectInterface is a mock of HTTPScaledObjectInterface interface.
type MockHTTPScaledObjectInterface struct {
	ctrl     *gomock.Controller
	recorder *MockHTTPScaledObjectInterfaceMockRecorder
}

// MockHTTPScaledObjectInterfaceMockRecorder is the mock recorder for MockHTTPScaledObjectInterface.
type MockHTTPScaledObjectInterfaceMockRecorder struct {
	mock *MockHTTPScaledObjectInterface
}

// NewMockHTTPScaledObjectInterface creates a new mock instance.
func NewMockHTTPScaledObjectInterface(ctrl *gomock.Controller) *MockHTTPScaledObjectInterface {
	mock := &MockHTTPScaledObjectInterface{ctrl: ctrl}
	mock.recorder = &MockHTTPScaledObjectInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockHTTPScaledObjectInterface) EXPECT() *MockHTTPScaledObjectInterfaceMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockHTTPScaledObjectInterface) Create(ctx context.Context, hTTPScaledObject *v1alpha1.HTTPScaledObject, opts v1.CreateOptions) (*v1alpha1.HTTPScaledObject, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", ctx, hTTPScaledObject, opts)
	ret0, _ := ret[0].(*v1alpha1.HTTPScaledObject)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockHTTPScaledObjectInterfaceMockRecorder) Create(ctx, hTTPScaledObject, opts interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockHTTPScaledObjectInterface)(nil).Create), ctx, hTTPScaledObject, opts)
}

// Delete mocks base method.
func (m *MockHTTPScaledObjectInterface) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", ctx, name, opts)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockHTTPScaledObjectInterfaceMockRecorder) Delete(ctx, name, opts interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockHTTPScaledObjectInterface)(nil).Delete), ctx, name, opts)
}

// DeleteCollection mocks base method.
func (m *MockHTTPScaledObjectInterface) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteCollection", ctx, opts, listOpts)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteCollection indicates an expected call of DeleteCollection.
func (mr *MockHTTPScaledObjectInterfaceMockRecorder) DeleteCollection(ctx, opts, listOpts interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteCollection", reflect.TypeOf((*MockHTTPScaledObjectInterface)(nil).DeleteCollection), ctx, opts, listOpts)
}

// Get mocks base method.
func (m *MockHTTPScaledObjectInterface) Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.HTTPScaledObject, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", ctx, name, opts)
	ret0, _ := ret[0].(*v1alpha1.HTTPScaledObject)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockHTTPScaledObjectInterfaceMockRecorder) Get(ctx, name, opts interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockHTTPScaledObjectInterface)(nil).Get), ctx, name, opts)
}

// List mocks base method.
func (m *MockHTTPScaledObjectInterface) List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.HTTPScaledObjectList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", ctx, opts)
	ret0, _ := ret[0].(*v1alpha1.HTTPScaledObjectList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *MockHTTPScaledObjectInterfaceMockRecorder) List(ctx, opts interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockHTTPScaledObjectInterface)(nil).List), ctx, opts)
}

// Patch mocks base method.
func (m *MockHTTPScaledObjectInterface) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (*v1alpha1.HTTPScaledObject, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, name, pt, data, opts}
	for _, a := range subresources {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Patch", varargs...)
	ret0, _ := ret[0].(*v1alpha1.HTTPScaledObject)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Patch indicates an expected call of Patch.
func (mr *MockHTTPScaledObjectInterfaceMockRecorder) Patch(ctx, name, pt, data, opts interface{}, subresources ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, name, pt, data, opts}, subresources...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Patch", reflect.TypeOf((*MockHTTPScaledObjectInterface)(nil).Patch), varargs...)
}

// Update mocks base method.
func (m *MockHTTPScaledObjectInterface) Update(ctx context.Context, hTTPScaledObject *v1alpha1.HTTPScaledObject, opts v1.UpdateOptions) (*v1alpha1.HTTPScaledObject, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", ctx, hTTPScaledObject, opts)
	ret0, _ := ret[0].(*v1alpha1.HTTPScaledObject)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockHTTPScaledObjectInterfaceMockRecorder) Update(ctx, hTTPScaledObject, opts interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockHTTPScaledObjectInterface)(nil).Update), ctx, hTTPScaledObject, opts)
}

// UpdateStatus mocks base method.
func (m *MockHTTPScaledObjectInterface) UpdateStatus(ctx context.Context, hTTPScaledObject *v1alpha1.HTTPScaledObject, opts v1.UpdateOptions) (*v1alpha1.HTTPScaledObject, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateStatus", ctx, hTTPScaledObject, opts)
	ret0, _ := ret[0].(*v1alpha1.HTTPScaledObject)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateStatus indicates an expected call of UpdateStatus.
func (mr *MockHTTPScaledObjectInterfaceMockRecorder) UpdateStatus(ctx, hTTPScaledObject, opts interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateStatus", reflect.TypeOf((*MockHTTPScaledObjectInterface)(nil).UpdateStatus), ctx, hTTPScaledObject, opts)
}

// Watch mocks base method.
func (m *MockHTTPScaledObjectInterface) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Watch", ctx, opts)
	ret0, _ := ret[0].(watch.Interface)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Watch indicates an expected call of Watch.
func (mr *MockHTTPScaledObjectInterfaceMockRecorder) Watch(ctx, opts interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Watch", reflect.TypeOf((*MockHTTPScaledObjectInterface)(nil).Watch), ctx, opts)
}
