// Code generated by MockGen. DO NOT EDIT.
// Source: measurement.go

// Package stats is a generated GoMock package.
package stats

import (
	reflect "reflect"
	time "time"

	gomock "github.com/golang/mock/gomock"
)

// MockCounter is a mock of Counter interface.
type MockCounter struct {
	ctrl     *gomock.Controller
	recorder *MockCounterMockRecorder
}

// MockCounterMockRecorder is the mock recorder for MockCounter.
type MockCounterMockRecorder struct {
	mock *MockCounter
}

// NewMockCounter creates a new mock instance.
func NewMockCounter(ctrl *gomock.Controller) *MockCounter {
	mock := &MockCounter{ctrl: ctrl}
	mock.recorder = &MockCounterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCounter) EXPECT() *MockCounterMockRecorder {
	return m.recorder
}

// Count mocks base method.
func (m *MockCounter) Count(n int) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Count", n)
}

// Count indicates an expected call of Count.
func (mr *MockCounterMockRecorder) Count(n interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Count", reflect.TypeOf((*MockCounter)(nil).Count), n)
}

// Increment mocks base method.
func (m *MockCounter) Increment() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Increment")
}

// Increment indicates an expected call of Increment.
func (mr *MockCounterMockRecorder) Increment() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Increment", reflect.TypeOf((*MockCounter)(nil).Increment))
}

// MockGauge is a mock of Gauge interface.
type MockGauge struct {
	ctrl     *gomock.Controller
	recorder *MockGaugeMockRecorder
}

// MockGaugeMockRecorder is the mock recorder for MockGauge.
type MockGaugeMockRecorder struct {
	mock *MockGauge
}

// NewMockGauge creates a new mock instance.
func NewMockGauge(ctrl *gomock.Controller) *MockGauge {
	mock := &MockGauge{ctrl: ctrl}
	mock.recorder = &MockGaugeMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockGauge) EXPECT() *MockGaugeMockRecorder {
	return m.recorder
}

// Gauge mocks base method.
func (m *MockGauge) Gauge(value interface{}) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Gauge", value)
}

// Gauge indicates an expected call of Gauge.
func (mr *MockGaugeMockRecorder) Gauge(value interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Gauge", reflect.TypeOf((*MockGauge)(nil).Gauge), value)
}

// MockHistogram is a mock of Histogram interface.
type MockHistogram struct {
	ctrl     *gomock.Controller
	recorder *MockHistogramMockRecorder
}

// MockHistogramMockRecorder is the mock recorder for MockHistogram.
type MockHistogramMockRecorder struct {
	mock *MockHistogram
}

// NewMockHistogram creates a new mock instance.
func NewMockHistogram(ctrl *gomock.Controller) *MockHistogram {
	mock := &MockHistogram{ctrl: ctrl}
	mock.recorder = &MockHistogramMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockHistogram) EXPECT() *MockHistogramMockRecorder {
	return m.recorder
}

// Observe mocks base method.
func (m *MockHistogram) Observe(value float64) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Observe", value)
}

// Observe indicates an expected call of Observe.
func (mr *MockHistogramMockRecorder) Observe(value interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Observe", reflect.TypeOf((*MockHistogram)(nil).Observe), value)
}

// MockTimer is a mock of Timer interface.
type MockTimer struct {
	ctrl     *gomock.Controller
	recorder *MockTimerMockRecorder
}

// MockTimerMockRecorder is the mock recorder for MockTimer.
type MockTimerMockRecorder struct {
	mock *MockTimer
}

// NewMockTimer creates a new mock instance.
func NewMockTimer(ctrl *gomock.Controller) *MockTimer {
	mock := &MockTimer{ctrl: ctrl}
	mock.recorder = &MockTimerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTimer) EXPECT() *MockTimerMockRecorder {
	return m.recorder
}

// RecordDuration mocks base method.
func (m *MockTimer) RecordDuration() func() {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RecordDuration")
	ret0, _ := ret[0].(func())
	return ret0
}

// RecordDuration indicates an expected call of RecordDuration.
func (mr *MockTimerMockRecorder) RecordDuration() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RecordDuration", reflect.TypeOf((*MockTimer)(nil).RecordDuration))
}

// SendTiming mocks base method.
func (m *MockTimer) SendTiming(duration time.Duration) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SendTiming", duration)
}

// SendTiming indicates an expected call of SendTiming.
func (mr *MockTimerMockRecorder) SendTiming(duration interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendTiming", reflect.TypeOf((*MockTimer)(nil).SendTiming), duration)
}

// Since mocks base method.
func (m *MockTimer) Since(start time.Time) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Since", start)
}

// Since indicates an expected call of Since.
func (mr *MockTimerMockRecorder) Since(start interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Since", reflect.TypeOf((*MockTimer)(nil).Since), start)
}

// MockMeasurement is a mock of Measurement interface.
type MockMeasurement struct {
	ctrl     *gomock.Controller
	recorder *MockMeasurementMockRecorder
}

// MockMeasurementMockRecorder is the mock recorder for MockMeasurement.
type MockMeasurementMockRecorder struct {
	mock *MockMeasurement
}

// NewMockMeasurement creates a new mock instance.
func NewMockMeasurement(ctrl *gomock.Controller) *MockMeasurement {
	mock := &MockMeasurement{ctrl: ctrl}
	mock.recorder = &MockMeasurementMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMeasurement) EXPECT() *MockMeasurementMockRecorder {
	return m.recorder
}

// Count mocks base method.
func (m *MockMeasurement) Count(n int) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Count", n)
}

// Count indicates an expected call of Count.
func (mr *MockMeasurementMockRecorder) Count(n interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Count", reflect.TypeOf((*MockMeasurement)(nil).Count), n)
}

// Gauge mocks base method.
func (m *MockMeasurement) Gauge(value interface{}) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Gauge", value)
}

// Gauge indicates an expected call of Gauge.
func (mr *MockMeasurementMockRecorder) Gauge(value interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Gauge", reflect.TypeOf((*MockMeasurement)(nil).Gauge), value)
}

// Increment mocks base method.
func (m *MockMeasurement) Increment() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Increment")
}

// Increment indicates an expected call of Increment.
func (mr *MockMeasurementMockRecorder) Increment() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Increment", reflect.TypeOf((*MockMeasurement)(nil).Increment))
}

// Observe mocks base method.
func (m *MockMeasurement) Observe(value float64) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Observe", value)
}

// Observe indicates an expected call of Observe.
func (mr *MockMeasurementMockRecorder) Observe(value interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Observe", reflect.TypeOf((*MockMeasurement)(nil).Observe), value)
}

// RecordDuration mocks base method.
func (m *MockMeasurement) RecordDuration() func() {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RecordDuration")
	ret0, _ := ret[0].(func())
	return ret0
}

// RecordDuration indicates an expected call of RecordDuration.
func (mr *MockMeasurementMockRecorder) RecordDuration() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RecordDuration", reflect.TypeOf((*MockMeasurement)(nil).RecordDuration))
}

// SendTiming mocks base method.
func (m *MockMeasurement) SendTiming(duration time.Duration) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SendTiming", duration)
}

// SendTiming indicates an expected call of SendTiming.
func (mr *MockMeasurementMockRecorder) SendTiming(duration interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendTiming", reflect.TypeOf((*MockMeasurement)(nil).SendTiming), duration)
}

// Since mocks base method.
func (m *MockMeasurement) Since(start time.Time) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Since", start)
}

// Since indicates an expected call of Since.
func (mr *MockMeasurementMockRecorder) Since(start interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Since", reflect.TypeOf((*MockMeasurement)(nil).Since), start)
}
