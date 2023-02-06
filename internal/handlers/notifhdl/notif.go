package notifhdl

import (
	"learn-test-hexa/internal/core/ports"

	"github.com/gin-gonic/gin"
)

type HTTPHandler struct {
	notifService ports.PushNotificationService
}

func NewHTTPHadler(notifService ports.PushNotificationRepository) *HTTPHandler {
	return &NewHTTPHadler{
		notifService: notifService,
	}
}

func (hdl *HTTPHandler) Send(c *gin.Context) {
	notif := hdl.notifService.Send()

	c.JSON(200, notif)
}
