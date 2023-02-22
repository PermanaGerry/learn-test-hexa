package notifhdl_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	// "github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"learn-test-hexa/internal/core/domain"
	"learn-test-hexa/internal/handlers/notifhdl"
	"learn-test-hexa/mocks/mockups"
	"net/http"
	"net/http/httptest"
	"testing"
)

type mocks struct {
	notifService *mockups.MockNotifService
}

func SetUpRouter() *gin.Engine {
	router := gin.Default()
	return router
}

func TestSend(t *testing.T) {

	assert := assert.New(t)

	r := SetUpRouter()

	// id, _ := uuid.NewRandom()
	notif := domain.PushNotif{
		ID:      "test-test",
		Message: "test-message",
		Type:    1,
	}

	type args struct {
		data domain.PushNotif
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
			name: "Should send notif successfully",
			args: args{data: notif},
			want: want{result: notif},
			mocks: func(m mocks) {
				m.notifService.EXPECT().
					Send(gomock.Eq(notif)).
					Return(domain.PushNotif{}, nil)
			},
		},
		// {
		// 	name: "[negatif] Should return error - error get redis",
		// 	args: args{data: notif},
		// 	want: want{err: errors.New("error get redis")},
		// 	mocks: func(m mocks) {
		// 		m.notifRepository.EXPECT().
		// 			Send(notif).
		// 			Return(domain.PushNotif{}, errors.New("error get redis"))

		// 		m.notifService.EXPECT().
		// 			Send(notif).
		// 			Return(domain.PushNotif{}, domain.)
		// 	},
		// },
		// {
		// 	name: "[negatif] Should return error - error unmarshall data redis",
		// 	args: args{data: notif},
		// 	want: want{err: errors.New("error unmarshall data redis")},
		// 	mocks: func(m mocks) {
		// 		m.notifService.EXPECT().
		// 			Send(notif).
		// 			Return(domain.PushNotif{}, errors.New("error unmarshall data redis"))
		// 	},
		// },
		// {
		// 	name: "[negatif] Should return error - error marshal save data redis",
		// 	args: args{data: notif},
		// 	want: want{err: errors.New("error marshal save data redis")},
		// 	mocks: func(m mocks) {
		// 		m.notifService.EXPECT().
		// 			Send(notif).
		// 			Return(domain.PushNotif{}, errors.New("error marshal save data redis"))
		// 	},
		// },
		// {
		// 	name: "[negatif] Should return error - error save data redis",
		// 	args: args{data: notif},
		// 	want: want{err: errors.New("error save data redis")},
		// 	mocks: func(m mocks) {
		// 		m.notifService.EXPECT().
		// 			Send(notif).
		// 			Return(domain.PushNotif{}, errors.New("error save data redis"))
		// 	},
		// },
	}

	for _, tt := range test {
		tt := tt

		m := mocks{
			notifService: mockups.NewMockNotifService(gomock.NewController(t)),
		}

		tt.mocks(m)
		handler := notifhdl.NewHTTPHadler(m.notifService)

		r.GET("/send", handler.Send)

		jsonValue, _ := json.Marshal(tt.args.data)
		req, _ := http.NewRequest("GET", "/send", bytes.NewBuffer(jsonValue))

		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		res := w.Result()
		defer res.Body.Close()

		result := new(domain.PushNotif)
		json.NewDecoder(w.Body).Decode(result)

		fmt.Println(result, err)

		assert.Equal(jsonValue, w.Body)
	}
}

func TestShow(*testing.T) {

}
