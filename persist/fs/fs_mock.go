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
// Source: github.com/m3db/m3db/persist/fs (interfaces: FileSetWriter,FileSetReader,FileSetSeeker)

package fs

import (
	"reflect"
	"time"

	"github.com/m3db/m3x/checked"
	"github.com/m3db/m3x/ident"
	time0 "github.com/m3db/m3x/time"

	"github.com/golang/mock/gomock"
)

// MockFileSetWriter is a mock of FileSetWriter interface
type MockFileSetWriter struct {
	ctrl     *gomock.Controller
	recorder *MockFileSetWriterMockRecorder
}

// MockFileSetWriterMockRecorder is the mock recorder for MockFileSetWriter
type MockFileSetWriterMockRecorder struct {
	mock *MockFileSetWriter
}

// NewMockFileSetWriter creates a new mock instance
func NewMockFileSetWriter(ctrl *gomock.Controller) *MockFileSetWriter {
	mock := &MockFileSetWriter{ctrl: ctrl}
	mock.recorder = &MockFileSetWriterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (_m *MockFileSetWriter) EXPECT() *MockFileSetWriterMockRecorder {
	return _m.recorder
}

// Close mocks base method
func (_m *MockFileSetWriter) Close() error {
	ret := _m.ctrl.Call(_m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close
func (_mr *MockFileSetWriterMockRecorder) Close() *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "Close", reflect.TypeOf((*MockFileSetWriter)(nil).Close))
}

// Open mocks base method
func (_m *MockFileSetWriter) Open(_param0 ident.ID, _param1 time.Duration, _param2 uint32, _param3 time.Time) error {
	ret := _m.ctrl.Call(_m, "Open", _param0, _param1, _param2, _param3)
	ret0, _ := ret[0].(error)
	return ret0
}

// Open indicates an expected call of Open
func (_mr *MockFileSetWriterMockRecorder) Open(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "Open", reflect.TypeOf((*MockFileSetWriter)(nil).Open), arg0, arg1, arg2, arg3)
}

// Write mocks base method
func (_m *MockFileSetWriter) Write(_param0 ident.ID, _param1 checked.Bytes, _param2 uint32) error {
	ret := _m.ctrl.Call(_m, "Write", _param0, _param1, _param2)
	ret0, _ := ret[0].(error)
	return ret0
}

// Write indicates an expected call of Write
func (_mr *MockFileSetWriterMockRecorder) Write(arg0, arg1, arg2 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "Write", reflect.TypeOf((*MockFileSetWriter)(nil).Write), arg0, arg1, arg2)
}

// WriteAll mocks base method
func (_m *MockFileSetWriter) WriteAll(_param0 ident.ID, _param1 []checked.Bytes, _param2 uint32) error {
	ret := _m.ctrl.Call(_m, "WriteAll", _param0, _param1, _param2)
	ret0, _ := ret[0].(error)
	return ret0
}

// WriteAll indicates an expected call of WriteAll
func (_mr *MockFileSetWriterMockRecorder) WriteAll(arg0, arg1, arg2 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "WriteAll", reflect.TypeOf((*MockFileSetWriter)(nil).WriteAll), arg0, arg1, arg2)
}

// MockFileSetReader is a mock of FileSetReader interface
type MockFileSetReader struct {
	ctrl     *gomock.Controller
	recorder *MockFileSetReaderMockRecorder
}

// MockFileSetReaderMockRecorder is the mock recorder for MockFileSetReader
type MockFileSetReaderMockRecorder struct {
	mock *MockFileSetReader
}

// NewMockFileSetReader creates a new mock instance
func NewMockFileSetReader(ctrl *gomock.Controller) *MockFileSetReader {
	mock := &MockFileSetReader{ctrl: ctrl}
	mock.recorder = &MockFileSetReaderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (_m *MockFileSetReader) EXPECT() *MockFileSetReaderMockRecorder {
	return _m.recorder
}

// Close mocks base method
func (_m *MockFileSetReader) Close() error {
	ret := _m.ctrl.Call(_m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close
func (_mr *MockFileSetReaderMockRecorder) Close() *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "Close", reflect.TypeOf((*MockFileSetReader)(nil).Close))
}

// Entries mocks base method
func (_m *MockFileSetReader) Entries() int {
	ret := _m.ctrl.Call(_m, "Entries")
	ret0, _ := ret[0].(int)
	return ret0
}

// Entries indicates an expected call of Entries
func (_mr *MockFileSetReaderMockRecorder) Entries() *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "Entries", reflect.TypeOf((*MockFileSetReader)(nil).Entries))
}

// EntriesRead mocks base method
func (_m *MockFileSetReader) EntriesRead() int {
	ret := _m.ctrl.Call(_m, "EntriesRead")
	ret0, _ := ret[0].(int)
	return ret0
}

// EntriesRead indicates an expected call of EntriesRead
func (_mr *MockFileSetReaderMockRecorder) EntriesRead() *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "EntriesRead", reflect.TypeOf((*MockFileSetReader)(nil).EntriesRead))
}

// MetadataRead mocks base method
func (_m *MockFileSetReader) MetadataRead() int {
	ret := _m.ctrl.Call(_m, "MetadataRead")
	ret0, _ := ret[0].(int)
	return ret0
}

// MetadataRead indicates an expected call of MetadataRead
func (_mr *MockFileSetReaderMockRecorder) MetadataRead() *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "MetadataRead", reflect.TypeOf((*MockFileSetReader)(nil).MetadataRead))
}

// Open mocks base method
func (_m *MockFileSetReader) Open(_param0 ident.ID, _param1 uint32, _param2 time.Time) error {
	ret := _m.ctrl.Call(_m, "Open", _param0, _param1, _param2)
	ret0, _ := ret[0].(error)
	return ret0
}

// Open indicates an expected call of Open
func (_mr *MockFileSetReaderMockRecorder) Open(arg0, arg1, arg2 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "Open", reflect.TypeOf((*MockFileSetReader)(nil).Open), arg0, arg1, arg2)
}

// Range mocks base method
func (_m *MockFileSetReader) Range() time0.Range {
	ret := _m.ctrl.Call(_m, "Range")
	ret0, _ := ret[0].(time0.Range)
	return ret0
}

// Range indicates an expected call of Range
func (_mr *MockFileSetReaderMockRecorder) Range() *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "Range", reflect.TypeOf((*MockFileSetReader)(nil).Range))
}

// Read mocks base method
func (_m *MockFileSetReader) Read() (ident.ID, checked.Bytes, uint32, error) {
	ret := _m.ctrl.Call(_m, "Read")
	ret0, _ := ret[0].(ident.ID)
	ret1, _ := ret[1].(checked.Bytes)
	ret2, _ := ret[2].(uint32)
	ret3, _ := ret[3].(error)
	return ret0, ret1, ret2, ret3
}

// Read indicates an expected call of Read
func (_mr *MockFileSetReaderMockRecorder) Read() *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "Read", reflect.TypeOf((*MockFileSetReader)(nil).Read))
}

// ReadBloomFilter mocks base method
func (_m *MockFileSetReader) ReadBloomFilter() (*ManagedConcurrentBloomFilter, error) {
	ret := _m.ctrl.Call(_m, "ReadBloomFilter")
	ret0, _ := ret[0].(*ManagedConcurrentBloomFilter)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReadBloomFilter indicates an expected call of ReadBloomFilter
func (_mr *MockFileSetReaderMockRecorder) ReadBloomFilter() *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "ReadBloomFilter", reflect.TypeOf((*MockFileSetReader)(nil).ReadBloomFilter))
}

// ReadMetadata mocks base method
func (_m *MockFileSetReader) ReadMetadata() (ident.ID, int, uint32, error) {
	ret := _m.ctrl.Call(_m, "ReadMetadata")
	ret0, _ := ret[0].(ident.ID)
	ret1, _ := ret[1].(int)
	ret2, _ := ret[2].(uint32)
	ret3, _ := ret[3].(error)
	return ret0, ret1, ret2, ret3
}

// ReadMetadata indicates an expected call of ReadMetadata
func (_mr *MockFileSetReaderMockRecorder) ReadMetadata() *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "ReadMetadata", reflect.TypeOf((*MockFileSetReader)(nil).ReadMetadata))
}

// Status mocks base method
func (_m *MockFileSetReader) Status() FileSetReaderStatus {
	ret := _m.ctrl.Call(_m, "Status")
	ret0, _ := ret[0].(FileSetReaderStatus)
	return ret0
}

// Status indicates an expected call of Status
func (_mr *MockFileSetReaderMockRecorder) Status() *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "Status", reflect.TypeOf((*MockFileSetReader)(nil).Status))
}

// Validate mocks base method
func (_m *MockFileSetReader) Validate() error {
	ret := _m.ctrl.Call(_m, "Validate")
	ret0, _ := ret[0].(error)
	return ret0
}

// Validate indicates an expected call of Validate
func (_mr *MockFileSetReaderMockRecorder) Validate() *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "Validate", reflect.TypeOf((*MockFileSetReader)(nil).Validate))
}

// ValidateData mocks base method
func (_m *MockFileSetReader) ValidateData() error {
	ret := _m.ctrl.Call(_m, "ValidateData")
	ret0, _ := ret[0].(error)
	return ret0
}

// ValidateData indicates an expected call of ValidateData
func (_mr *MockFileSetReaderMockRecorder) ValidateData() *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "ValidateData", reflect.TypeOf((*MockFileSetReader)(nil).ValidateData))
}

// ValidateMetadata mocks base method
func (_m *MockFileSetReader) ValidateMetadata() error {
	ret := _m.ctrl.Call(_m, "ValidateMetadata")
	ret0, _ := ret[0].(error)
	return ret0
}

// ValidateMetadata indicates an expected call of ValidateMetadata
func (_mr *MockFileSetReaderMockRecorder) ValidateMetadata() *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "ValidateMetadata", reflect.TypeOf((*MockFileSetReader)(nil).ValidateMetadata))
}

// MockFileSetSeeker is a mock of FileSetSeeker interface
type MockFileSetSeeker struct {
	ctrl     *gomock.Controller
	recorder *MockFileSetSeekerMockRecorder
}

// MockFileSetSeekerMockRecorder is the mock recorder for MockFileSetSeeker
type MockFileSetSeekerMockRecorder struct {
	mock *MockFileSetSeeker
}

// NewMockFileSetSeeker creates a new mock instance
func NewMockFileSetSeeker(ctrl *gomock.Controller) *MockFileSetSeeker {
	mock := &MockFileSetSeeker{ctrl: ctrl}
	mock.recorder = &MockFileSetSeekerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (_m *MockFileSetSeeker) EXPECT() *MockFileSetSeekerMockRecorder {
	return _m.recorder
}

// Close mocks base method
func (_m *MockFileSetSeeker) Close() error {
	ret := _m.ctrl.Call(_m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close
func (_mr *MockFileSetSeekerMockRecorder) Close() *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "Close", reflect.TypeOf((*MockFileSetSeeker)(nil).Close))
}

// ConcurrentClone mocks base method
func (_m *MockFileSetSeeker) ConcurrentClone() (ConcurrentFileSetSeeker, error) {
	ret := _m.ctrl.Call(_m, "ConcurrentClone")
	ret0, _ := ret[0].(ConcurrentFileSetSeeker)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ConcurrentClone indicates an expected call of ConcurrentClone
func (_mr *MockFileSetSeekerMockRecorder) ConcurrentClone() *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "ConcurrentClone", reflect.TypeOf((*MockFileSetSeeker)(nil).ConcurrentClone))
}

// ConcurrentIDBloomFilter mocks base method
func (_m *MockFileSetSeeker) ConcurrentIDBloomFilter() *ManagedConcurrentBloomFilter {
	ret := _m.ctrl.Call(_m, "ConcurrentIDBloomFilter")
	ret0, _ := ret[0].(*ManagedConcurrentBloomFilter)
	return ret0
}

// ConcurrentIDBloomFilter indicates an expected call of ConcurrentIDBloomFilter
func (_mr *MockFileSetSeekerMockRecorder) ConcurrentIDBloomFilter() *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "ConcurrentIDBloomFilter", reflect.TypeOf((*MockFileSetSeeker)(nil).ConcurrentIDBloomFilter))
}

// Entries mocks base method
func (_m *MockFileSetSeeker) Entries() int {
	ret := _m.ctrl.Call(_m, "Entries")
	ret0, _ := ret[0].(int)
	return ret0
}

// Entries indicates an expected call of Entries
func (_mr *MockFileSetSeekerMockRecorder) Entries() *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "Entries", reflect.TypeOf((*MockFileSetSeeker)(nil).Entries))
}

// Open mocks base method
func (_m *MockFileSetSeeker) Open(_param0 ident.ID, _param1 uint32, _param2 time.Time) error {
	ret := _m.ctrl.Call(_m, "Open", _param0, _param1, _param2)
	ret0, _ := ret[0].(error)
	return ret0
}

// Open indicates an expected call of Open
func (_mr *MockFileSetSeekerMockRecorder) Open(arg0, arg1, arg2 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "Open", reflect.TypeOf((*MockFileSetSeeker)(nil).Open), arg0, arg1, arg2)
}

// Range mocks base method
func (_m *MockFileSetSeeker) Range() time0.Range {
	ret := _m.ctrl.Call(_m, "Range")
	ret0, _ := ret[0].(time0.Range)
	return ret0
}

// Range indicates an expected call of Range
func (_mr *MockFileSetSeekerMockRecorder) Range() *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "Range", reflect.TypeOf((*MockFileSetSeeker)(nil).Range))
}

// SeekByID mocks base method
func (_m *MockFileSetSeeker) SeekByID(_param0 ident.ID) (checked.Bytes, error) {
	ret := _m.ctrl.Call(_m, "SeekByID", _param0)
	ret0, _ := ret[0].(checked.Bytes)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SeekByID indicates an expected call of SeekByID
func (_mr *MockFileSetSeekerMockRecorder) SeekByID(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "SeekByID", reflect.TypeOf((*MockFileSetSeeker)(nil).SeekByID), arg0)
}

// SeekByIndexEntry mocks base method
func (_m *MockFileSetSeeker) SeekByIndexEntry(_param0 IndexEntry) (checked.Bytes, error) {
	ret := _m.ctrl.Call(_m, "SeekByIndexEntry", _param0)
	ret0, _ := ret[0].(checked.Bytes)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SeekByIndexEntry indicates an expected call of SeekByIndexEntry
func (_mr *MockFileSetSeekerMockRecorder) SeekByIndexEntry(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "SeekByIndexEntry", reflect.TypeOf((*MockFileSetSeeker)(nil).SeekByIndexEntry), arg0)
}

// SeekIndexEntry mocks base method
func (_m *MockFileSetSeeker) SeekIndexEntry(_param0 ident.ID) (IndexEntry, error) {
	ret := _m.ctrl.Call(_m, "SeekIndexEntry", _param0)
	ret0, _ := ret[0].(IndexEntry)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SeekIndexEntry indicates an expected call of SeekIndexEntry
func (_mr *MockFileSetSeekerMockRecorder) SeekIndexEntry(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "SeekIndexEntry", reflect.TypeOf((*MockFileSetSeeker)(nil).SeekIndexEntry), arg0)
}
