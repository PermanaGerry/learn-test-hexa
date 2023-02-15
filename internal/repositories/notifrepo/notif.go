package notifrepo

import (
	"learn-test-hexa/internal/core/domain"
)

type memNotif struct {
	notif string
}

func NewMemNotif() *memNotif {
	return &memNotif{
		notif: "test",
	}
}

func (repo *memNotif) Send(data domain.PushNotif) domain.PushNotif {
	notif := domain.NewNotif(data.ID, data.Message, data.Type)

	return notif
}

func (repo *memNotif) Show() domain.ShowNotif {
	list := domain.ShowNotif{
		Lists: []domain.PushNotif{
			{
				ID:      "test-test",
				Message: "test",
				Type:    1,
			},
			{
				ID:      "test-test-1",
				Message: "test",
				Type:    1,
			},
		},
	}

	return list
}
