package mockups

import (
	domain "learn-test-hexa/internal/core/domain"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

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
