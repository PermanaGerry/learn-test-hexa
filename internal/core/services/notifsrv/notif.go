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

func (srv *service) Send(data domain.PushNotif) (domain.PushNotif, error) {

	notif, err := srv.pushNotificationRepository.Send(data)

	if err != nil {
		return notif, err
	}

	return notif, nil
}

func (srv *service) List() (domain.ShowNotif, error) {
	show, err := srv.pushNotificationRepository.Show()

	if err != nil {
		return show, err
	}

	return show, nil
}
