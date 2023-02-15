package ports

import (
	"learn-test-hexa/internal/core/domain"
)

type PushNotificationRepository interface {
	Send(notif domain.PushNotif) (domain.PushNotif, error)
	Show() (domain.ShowNotif, error)
}

type PushNotificationService interface {
	Send(notif domain.PushNotif) (domain.PushNotif, error)
	List() (domain.ShowNotif, error)
}
