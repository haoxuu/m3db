// Copyright (c) 2018 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/m3db/m3db/serialize (interfaces: Encoder,EncoderPool,Decoder,DecoderPool)

package serialize

import (
	"reflect"

	"github.com/m3db/m3x/checked"
	"github.com/m3db/m3x/ident"

	"github.com/golang/mock/gomock"
)

// MockEncoder is a mock of Encoder interface
type MockEncoder struct {
	ctrl     *gomock.Controller
	recorder *MockEncoderMockRecorder
}

// MockEncoderMockRecorder is the mock recorder for MockEncoder
type MockEncoderMockRecorder struct {
	mock *MockEncoder
}

// NewMockEncoder creates a new mock instance
func NewMockEncoder(ctrl *gomock.Controller) *MockEncoder {
	mock := &MockEncoder{ctrl: ctrl}
	mock.recorder = &MockEncoderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (_m *MockEncoder) EXPECT() *MockEncoderMockRecorder {
	return _m.recorder
}

// Data mocks base method
func (_m *MockEncoder) Data() (checked.Bytes, bool) {
	ret := _m.ctrl.Call(_m, "Data")
	ret0, _ := ret[0].(checked.Bytes)
	ret1, _ := ret[1].(bool)
	return ret0, ret1
}

// Data indicates an expected call of Data
func (_mr *MockEncoderMockRecorder) Data() *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "Data", reflect.TypeOf((*MockEncoder)(nil).Data))
}

// Encode mocks base method
func (_m *MockEncoder) Encode(_param0 ident.TagIterator) error {
	ret := _m.ctrl.Call(_m, "Encode", _param0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Encode indicates an expected call of Encode
func (_mr *MockEncoderMockRecorder) Encode(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "Encode", reflect.TypeOf((*MockEncoder)(nil).Encode), arg0)
}

// Finalize mocks base method
func (_m *MockEncoder) Finalize() {
	_m.ctrl.Call(_m, "Finalize")
}

// Finalize indicates an expected call of Finalize
func (_mr *MockEncoderMockRecorder) Finalize() *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "Finalize", reflect.TypeOf((*MockEncoder)(nil).Finalize))
}

// Reset mocks base method
func (_m *MockEncoder) Reset() {
	_m.ctrl.Call(_m, "Reset")
}

// Reset indicates an expected call of Reset
func (_mr *MockEncoderMockRecorder) Reset() *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "Reset", reflect.TypeOf((*MockEncoder)(nil).Reset))
}

// MockEncoderPool is a mock of EncoderPool interface
type MockEncoderPool struct {
	ctrl     *gomock.Controller
	recorder *MockEncoderPoolMockRecorder
}

// MockEncoderPoolMockRecorder is the mock recorder for MockEncoderPool
type MockEncoderPoolMockRecorder struct {
	mock *MockEncoderPool
}

// NewMockEncoderPool creates a new mock instance
func NewMockEncoderPool(ctrl *gomock.Controller) *MockEncoderPool {
	mock := &MockEncoderPool{ctrl: ctrl}
	mock.recorder = &MockEncoderPoolMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (_m *MockEncoderPool) EXPECT() *MockEncoderPoolMockRecorder {
	return _m.recorder
}

// Get mocks base method
func (_m *MockEncoderPool) Get() Encoder {
	ret := _m.ctrl.Call(_m, "Get")
	ret0, _ := ret[0].(Encoder)
	return ret0
}

// Get indicates an expected call of Get
func (_mr *MockEncoderPoolMockRecorder) Get() *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "Get", reflect.TypeOf((*MockEncoderPool)(nil).Get))
}

// Init mocks base method
func (_m *MockEncoderPool) Init() {
	_m.ctrl.Call(_m, "Init")
}

// Init indicates an expected call of Init
func (_mr *MockEncoderPoolMockRecorder) Init() *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "Init", reflect.TypeOf((*MockEncoderPool)(nil).Init))
}

// Put mocks base method
func (_m *MockEncoderPool) Put(_param0 Encoder) {
	_m.ctrl.Call(_m, "Put", _param0)
}

// Put indicates an expected call of Put
func (_mr *MockEncoderPoolMockRecorder) Put(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "Put", reflect.TypeOf((*MockEncoderPool)(nil).Put), arg0)
}

// MockDecoder is a mock of Decoder interface
type MockDecoder struct {
	ctrl     *gomock.Controller
	recorder *MockDecoderMockRecorder
}

// MockDecoderMockRecorder is the mock recorder for MockDecoder
type MockDecoderMockRecorder struct {
	mock *MockDecoder
}

// NewMockDecoder creates a new mock instance
func NewMockDecoder(ctrl *gomock.Controller) *MockDecoder {
	mock := &MockDecoder{ctrl: ctrl}
	mock.recorder = &MockDecoderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (_m *MockDecoder) EXPECT() *MockDecoderMockRecorder {
	return _m.recorder
}

// Close mocks base method
func (_m *MockDecoder) Close() {
	_m.ctrl.Call(_m, "Close")
}

// Close indicates an expected call of Close
func (_mr *MockDecoderMockRecorder) Close() *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "Close", reflect.TypeOf((*MockDecoder)(nil).Close))
}

// Current mocks base method
func (_m *MockDecoder) Current() ident.Tag {
	ret := _m.ctrl.Call(_m, "Current")
	ret0, _ := ret[0].(ident.Tag)
	return ret0
}

// Current indicates an expected call of Current
func (_mr *MockDecoderMockRecorder) Current() *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "Current", reflect.TypeOf((*MockDecoder)(nil).Current))
}

// Duplicate mocks base method
func (_m *MockDecoder) Duplicate() ident.TagIterator {
	ret := _m.ctrl.Call(_m, "Duplicate")
	ret0, _ := ret[0].(ident.TagIterator)
	return ret0
}

// Duplicate indicates an expected call of Duplicate
func (_mr *MockDecoderMockRecorder) Duplicate() *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "Duplicate", reflect.TypeOf((*MockDecoder)(nil).Duplicate))
}

// Err mocks base method
func (_m *MockDecoder) Err() error {
	ret := _m.ctrl.Call(_m, "Err")
	ret0, _ := ret[0].(error)
	return ret0
}

// Err indicates an expected call of Err
func (_mr *MockDecoderMockRecorder) Err() *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "Err", reflect.TypeOf((*MockDecoder)(nil).Err))
}

// Finalize mocks base method
func (_m *MockDecoder) Finalize() {
	_m.ctrl.Call(_m, "Finalize")
}

// Finalize indicates an expected call of Finalize
func (_mr *MockDecoderMockRecorder) Finalize() *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "Finalize", reflect.TypeOf((*MockDecoder)(nil).Finalize))
}

// Next mocks base method
func (_m *MockDecoder) Next() bool {
	ret := _m.ctrl.Call(_m, "Next")
	ret0, _ := ret[0].(bool)
	return ret0
}

// Next indicates an expected call of Next
func (_mr *MockDecoderMockRecorder) Next() *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "Next", reflect.TypeOf((*MockDecoder)(nil).Next))
}

// Remaining mocks base method
func (_m *MockDecoder) Remaining() int {
	ret := _m.ctrl.Call(_m, "Remaining")
	ret0, _ := ret[0].(int)
	return ret0
}

// Remaining indicates an expected call of Remaining
func (_mr *MockDecoderMockRecorder) Remaining() *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "Remaining", reflect.TypeOf((*MockDecoder)(nil).Remaining))
}

// Reset mocks base method
func (_m *MockDecoder) Reset(_param0 checked.Bytes) {
	_m.ctrl.Call(_m, "Reset", _param0)
}

// Reset indicates an expected call of Reset
func (_mr *MockDecoderMockRecorder) Reset(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "Reset", reflect.TypeOf((*MockDecoder)(nil).Reset), arg0)
}

// MockDecoderPool is a mock of DecoderPool interface
type MockDecoderPool struct {
	ctrl     *gomock.Controller
	recorder *MockDecoderPoolMockRecorder
}

// MockDecoderPoolMockRecorder is the mock recorder for MockDecoderPool
type MockDecoderPoolMockRecorder struct {
	mock *MockDecoderPool
}

// NewMockDecoderPool creates a new mock instance
func NewMockDecoderPool(ctrl *gomock.Controller) *MockDecoderPool {
	mock := &MockDecoderPool{ctrl: ctrl}
	mock.recorder = &MockDecoderPoolMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (_m *MockDecoderPool) EXPECT() *MockDecoderPoolMockRecorder {
	return _m.recorder
}

// Get mocks base method
func (_m *MockDecoderPool) Get() Decoder {
	ret := _m.ctrl.Call(_m, "Get")
	ret0, _ := ret[0].(Decoder)
	return ret0
}

// Get indicates an expected call of Get
func (_mr *MockDecoderPoolMockRecorder) Get() *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "Get", reflect.TypeOf((*MockDecoderPool)(nil).Get))
}

// Init mocks base method
func (_m *MockDecoderPool) Init() {
	_m.ctrl.Call(_m, "Init")
}

// Init indicates an expected call of Init
func (_mr *MockDecoderPoolMockRecorder) Init() *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "Init", reflect.TypeOf((*MockDecoderPool)(nil).Init))
}

// Put mocks base method
func (_m *MockDecoderPool) Put(_param0 Decoder) {
	_m.ctrl.Call(_m, "Put", _param0)
}

// Put indicates an expected call of Put
func (_mr *MockDecoderPoolMockRecorder) Put(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "Put", reflect.TypeOf((*MockDecoderPool)(nil).Put), arg0)
}
