package notifhdl_test

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
			name: "Should send notif successfully - positif",
			args: args{id: "test-test", message: "test", tipe: 1},
			want: want{result: notif},
			mocks: func(m mocks) {
				m.notifRepository.EXPECT().
					Send(notif).
					Return(notif, nil)
			},
		},
		{
			name: "Should return error - error get redis",
			args: args{id: "test-test", message: "test", tipe: 1},
			want: want{err: errors.New("error get redis")},
			mocks: func(m mocks) {
				m.notifRepository.EXPECT().
					Send(notif).
					Return(domain.PushNotif{}, errors.New("error get redis"))
			},
		},
		{
			name: "Should return error - error unmarshall data redis",
			args: args{id: "test-test", message: "test", tipe: 1},
			want: want{err: errors.New("error unmarshall data redis")},
			mocks: func(m mocks) {
				m.notifRepository.EXPECT().
					Send(notif).
					Return(domain.PushNotif{}, errors.New("error unmarshall data redis"))
			},
		},
		{
			name: "Should return error - error marshal save data redis",
			args: args{id: "test-test", message: "test", tipe: 1},
			want: want{err: errors.New("error marshal save data redis")},
			mocks: func(m mocks) {
				m.notifRepository.EXPECT().
					Send(notif).
					Return(domain.PushNotif{}, errors.New("error marshal save data redis"))
			},
		},
		{
			name: "Should return error - error save data redis",
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
