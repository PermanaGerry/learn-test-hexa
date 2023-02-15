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
	notif, err := hdl.notifService.Send(domain.PushNotif{
		ID:      id,
		Message: "test-message",
		Type:    1,
	})

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
