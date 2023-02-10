package notifrepo

import (
	"learn-test-hexa/internal/core/domain"

	"github.com/google/uuid"
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

	id := uuid.New().String()
	notif := domain.NewNotif(id, "test pesan", 1)

	return notif
}
