package notifsrv_test

import (
	// "fmt"
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"learn-test-hexa/internal/core/domain"
	"learn-test-hexa/internal/core/services/notifsrv"
	"learn-test-hexa/mocks/mockups"
)

type mocks struct {
	notifRepository *mockups.MockNotifRepository
}

func TestSend(t *testing.T) {

	assert := assert.New(t)

	notif := domain.PushNotif{
		ID:      "test-test",
		Message: "test",
		Type:    1,
	}

	type args struct {
		id      string
		message string
		tipe    int
	}

	type want struct {
		result domain.PushNotif
		err    error
	}

	test := []struct {
		name  string
		args  args
		want  want
		mocks func(m mocks)
	}{
		{
			name: "[positif] Should send notif successfully",
			args: args{id: "test-test", message: "test", tipe: 1},
			want: want{result: notif},
			mocks: func(m mocks) {
				m.notifRepository.EXPECT().
					Send(notif).
					Return(notif, nil)
			},
		},
		{
			name: "[negatif] Should return error - error get redis",
			args: args{id: "test-test", message: "test", tipe: 1},
			want: want{err: errors.New("error get redis")},
			mocks: func(m mocks) {
				m.notifRepository.EXPECT().
					Send(notif).
					Return(domain.PushNotif{}, errors.New("error get redis"))
			},
		},
		{
			name: "[negatif] Should return error - error unmarshall data redis",
			args: args{id: "test-test", message: "test", tipe: 1},
			want: want{err: errors.New("error unmarshall data redis")},
			mocks: func(m mocks) {
				m.notifRepository.EXPECT().
					Send(notif).
					Return(domain.PushNotif{}, errors.New("error unmarshall data redis"))
			},
		},
		{
			name: "[negatif] Should return error - error marshal save data redis",
			args: args{id: "test-test", message: "test", tipe: 1},
			want: want{err: errors.New("error marshal save data redis")},
			mocks: func(m mocks) {
				m.notifRepository.EXPECT().
					Send(notif).
					Return(domain.PushNotif{}, errors.New("error marshal save data redis"))
			},
		},
		{
			name: "[negatif] Should return error - error save data redis",
			args: args{id: "test-test", message: "test", tipe: 1},
			want: want{err: errors.New("error save data redis")},
			mocks: func(m mocks) {
				m.notifRepository.EXPECT().
					Send(notif).
					Return(domain.PushNotif{}, errors.New("error save data redis"))
			},
		},
	}

	for _, tt := range test {
		tt := tt

		m := mocks{
			notifRepository: mockups.NewMockNotifRepository(gomock.NewController(t)),
		}

		tt.mocks(m)
		service := notifsrv.New(m.notifRepository)

		result, err := service.Send(domain.NewNotif(tt.args.id, tt.args.message, tt.args.tipe))

		if tt.want.err != nil && err != nil {
			assert.Equal(tt.want.err.Error(), err.Error())
		}

		assert.Equal(tt.want.result, result, "they should be equal")
	}
}

func TestList(t *testing.T) {
	assert := assert.New(t)

	list := domain.ShowNotif{
		Lists: []domain.PushNotif{
			{
				ID:      "test-test-test-1",
				Message: "test",
				Type:    1,
			},
			{
				ID:      "test-test-test-2",
				Message: "test 2",
				Type:    2,
			},
		},
	}

	type want struct {
		result domain.ShowNotif
		err    error
	}

	test := []struct {
		name  string
		want  want
		mocks func(m mocks)
	}{
		{
			name: "[positif] Should show list successfully",
			want: want{result: list},
			mocks: func(m mocks) {
				m.notifRepository.EXPECT().
					Show().
					Return(list, nil)
			},
		},
		{
			name: "[negatif] Should return error - error get redis",
			want: want{err: errors.New("error get redis")},
			mocks: func(m mocks) {
				m.notifRepository.EXPECT().
					Show().
					Return(domain.ShowNotif{}, errors.New("error get redis"))
			},
		},
		{
			name: "[negatif] Should return error - error get redis",
			want: want{err: errors.New("error unmarshall data redis")},
			mocks: func(m mocks) {
				m.notifRepository.EXPECT().
					Show().
					Return(domain.ShowNotif{}, errors.New("error unmarshall data redis"))
			},
		},
	}

	for _, tt := range test {
		tt := tt

		m := mocks{
			notifRepository: mockups.NewMockNotifRepository(gomock.NewController(t)),
		}

		tt.mocks(m)
		service := notifsrv.New(m.notifRepository)

		result, err := service.List()
		if tt.want.err != nil && err != nil {
			assert.Equal(tt.want.err.Error(), err.Error())
		}

		assert.Equal(tt.want.result, result, "they should be equal")
	}
}
