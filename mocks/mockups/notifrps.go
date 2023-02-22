package mockups

import (
	domain "learn-test-hexa/internal/core/domain"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

/**
 * ===============
 * mock repository
 * ===============
 */

type MockNotifRepository struct {
	ctrl     *gomock.Controller
	recorder *MockNotifRepositoryMockRecorder
}

type MockNotifRepositoryMockRecorder struct {
	mock *MockNotifRepository
}

func NewMockNotifRepository(ctrl *gomock.Controller) *MockNotifRepository {
	mock := &MockNotifRepository{ctrl: ctrl}
	mock.recorder = &MockNotifRepositoryMockRecorder{mock}

	return mock
}

func (m *MockNotifRepository) EXPECT() *MockNotifRepositoryMockRecorder {
	return m.recorder
}

func (m *MockNotifRepository) Send(arg0 domain.PushNotif) (domain.PushNotif, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Send", arg0)
	ret0, _ := ret[0].(domain.PushNotif)
	ret1, _ := ret[1].(error)

	return ret0, ret1
}

func (mr *MockNotifRepositoryMockRecorder) Send(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Send", reflect.TypeOf((*MockNotifRepository)(nil).Send), arg0)
}

func (m *MockNotifRepository) Show() (domain.ShowNotif, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Show")
	ret0, _ := ret[0].(domain.ShowNotif)
	ret1, _ := ret[1].(error)

	return ret0, ret1
}

func (mr *MockNotifRepositoryMockRecorder) Show() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Show", reflect.TypeOf((*MockNotifRepository)(nil).Show))
}

/**
 * ===============
 * mock service
 * ===============
 */

type MockNotifService struct {
	ctrl     *gomock.Controller
	recorder *MockNotifServiceRecorder
}

type MockNotifServiceRecorder struct {
	mock *MockNotifService
}

func NewMockNotifService(ctrl *gomock.Controller) *MockNotifService {
	mock := &MockNotifService{ctrl: ctrl}
	mock.recorder = &MockNotifServiceRecorder{mock}

	return mock
}

func (m *MockNotifService) EXPECT() *MockNotifServiceRecorder {
	return m.recorder
}

func (m *MockNotifService) Send(arg0 domain.PushNotif) (domain.PushNotif, error) {
	m.ctrl.T.Helper()

	ret := m.ctrl.Call(m, "Send", arg0)
	ret0, _ := ret[0].(domain.PushNotif)
	ret1, _ := ret[1].(error)

	return ret0, ret1
}

func (mr *MockNotifServiceRecorder) Send(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Send", reflect.TypeOf((*MockNotifService)(nil).Send), arg0)
}

func (m *MockNotifService) List() (domain.ShowNotif, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Show")
	ret0, _ := ret[0].(domain.ShowNotif)
	ret1, _ := ret[1].(error)

	return ret0, ret1
}

func (mr *MockNotifServiceRecorder) List() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Show", reflect.TypeOf((*MockNotifRepository)(nil).Show))
}

/**
 * ===============
 * mock handler
 * ===============
 */

type MockNotifHandler struct {
	ctrl     *gomock.Controller
	recorder *MockNotifHandlerRecover
}

type MockNotifHandlerRecover struct {
	mock *MockNotifHandler
}

func NewMockNotifHandler(ctrl *gomock.Controller) *MockNotifHandler {
	mock := &MockNotifHandler{ctrl: ctrl}
	mock.recorder = &MockNotifHandlerRecover{mock}

	return mock
}

func (m *MockNotifHandler) EXPECT() *MockNotifHandlerRecover {
	return m.recorder
}

func (m *MockNotifHandler) Send(arg0 domain.PushNotif) (domain.PushNotif, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Send", arg0)
	ret0 := ret[0].(domain.PushNotif)
	ret1 := ret[1].(error)

	return ret0, ret1
}

func (mr *MockNotifHandlerRecover) Send(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Send", reflect.TypeOf((*MockNotifHandler)(nil).Send), arg0)
}

func (m *MockNotifHandler) Show() (domain.ShowNotif, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Show")
	ret0, _ := ret[0].(domain.ShowNotif)
	ret1, _ := ret[1].(error)

	return ret0, ret1
}

func (mr *MockNotifHandlerRecover) Show() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Show", reflect.TypeOf((*MockNotifHandler)(nil).Show))
}
