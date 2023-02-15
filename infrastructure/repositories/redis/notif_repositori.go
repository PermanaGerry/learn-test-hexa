package notifrepords

import (
	"context"
	"encoding/json"
	"errors"
	"learn-test-hexa/internal/core/domain"

	"github.com/go-redis/redis"
)

var (
	list  domain.ShowNotif
	table = "notif"
	key   = "storage"
	ctx   = context.Background()
)

type (
	RepositoryNotif struct {
		db *redis.Client
	}
)

func NewRepositoryNotif(db *redis.Client) *RepositoryNotif {
	return &RepositoryNotif{
		db: db,
	}
}

func (repo *RepositoryNotif) Send(data domain.PushNotif) (domain.PushNotif, error) {
	notif := domain.NewNotif(data.ID, data.Message, data.Type)

	val, err := repo.db.HGet(table, key).Bytes()
	if err != nil {
		return notif, errors.New("error get redis")
	}

	err = json.Unmarshal(val, &list.Lists)
	if err != nil {
		return notif, errors.New("error unmarshall data redis")
	}

	list.Lists = append(list.Lists, notif)

	save, err := json.Marshal(list.Lists)
	if err != nil {
		return notif, errors.New("error marshal save data redis")
	}

	cmd := repo.db.HSet(table, key, save)
	if cmd.Err() != nil {
		return notif, errors.New("error save data redis")
	}

	return notif, nil
}

func (repo *RepositoryNotif) Show() (domain.ShowNotif, error) {

	val, err := repo.db.HGet(table, key).Bytes()

	if err != nil {
		return list, errors.New("error get redis")
	}

	err = json.Unmarshal(val, &list.Lists)
	if err != nil {
		return list, errors.New("error unmarshall data redis")
	}

	return list, nil
}
