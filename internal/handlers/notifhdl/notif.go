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

	// id := uuid.New().String()
	notif := domain.PushNotif{
		ID:      "test-test",
		Message: "test-message",
		Type:    1,
	}

	notif, err := hdl.notifService.Send(notif)

	if err != nil {
		c.JSON(200, domain.ErrorMessage{
			Message: err.Error(),
		})
	}

	c.JSON(200, notif)
}

func (hdl *HTTPHandler) List(c *gin.Context) {
	show, err := hdl.notifService.List()

	if err != nil {
		c.JSON(200, domain.ErrorMessage{
			Message: err.Error(),
		})
	}

	c.JSON(200, show.Lists)
}
