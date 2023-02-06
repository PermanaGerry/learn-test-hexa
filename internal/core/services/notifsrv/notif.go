package notifsrv

import (
	"learn-test-hexa/internal/core/domain"
	"learn-test-hexa/internal/core/ports"
)

type service struct {
	pushNotificationRepository ports.PushNotificationRepository
}

func New(pushNotificationRepository ports.PushNotificationRepository) *service {
	return &service{
		pushNotificationRepository: pushNotificationRepository,
	}
}

func (srv *service) Send(data domain.PushNotif) domain.PushNotif {

	notif := srv.pushNotificationRepository.Send(data)

	return notif
}
