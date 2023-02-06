package ports

import (
	"learn-test-hexa/internal/core/domain"
)

type PushNotificationRepository interface {
	Send(notif domain.PushNotif) domain.PushNotif
}

type PushNotificationService interface {
	Send() domain.PushNotif
}
