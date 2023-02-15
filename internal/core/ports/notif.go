package ports

import (
	"learn-test-hexa/internal/core/domain"
)

type PushNotificationRepository interface {
	Send(notif domain.PushNotif) domain.PushNotif
	Show() domain.ShowNotif
}

type PushNotificationService interface {
	Send(notif domain.PushNotif) domain.PushNotif
	List() domain.ShowNotif
}
