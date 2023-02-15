package notifhdl

import (
	"learn-test-hexa/internal/core/domain"
	"learn-test-hexa/internal/core/ports"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
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

	id := uuid.New().String()
	notif := hdl.notifService.Send(domain.PushNotif{
		ID:      id,
		Message: "test-message",
		Type:    1,
	})

	c.JSON(200, notif)
}

func (hdl *HTTPHandler) List(c *gin.Context) {
	show := hdl.notifService.List()

	c.JSON(200, show.Lists)
}
