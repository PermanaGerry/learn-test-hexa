package notifhdl

import (
	"learn-test-hexa/internal/core/domain"
	"learn-test-hexa/internal/core/ports"

	"github.com/gin-gonic/gin"
)

type HTTPHandler struct {
	notifService ports.PushNotificationService
}

func NewHTTPHadler(notifService ports.PushNotificationService) *HTTPHandler {
	return &HTTPHandler{
		notifService: notifService,
	}
}

func (hdl *HTTPHandler) Send(c *gin.Context) {

	notif := hdl.notifService.Send(domain.PushNotif{
		ID:      "uui-exx-aa",
		Message: "test-message",
		Type:    1,
	})

	c.JSON(200, notif)
}
