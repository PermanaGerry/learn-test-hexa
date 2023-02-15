package notifrepords

import (
	"context"
	"encoding/json"
	"learn-test-hexa/internal/core/domain"
	"os"

	"github.com/go-redis/redis"
	"github.com/rs/zerolog/log"
)

var (
	list  domain.ShowNotif
	table = "notif"
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

func (repo *RepositoryNotif) Send(data domain.PushNotif) domain.PushNotif {
	notif := domain.NewNotif(data.ID, data.Message, data.Type)

	return notif
}

func (repo *RepositoryNotif) Show() domain.ShowNotif {

	val, err := repo.db.HGet("notif", "storage").Bytes()

	if err != nil {
		log.Error().Msg("error get redis")
		os.Exit(0)
	}

	err = json.Unmarshal(val, &list.Lists)
	if err != nil {
		log.Error().Msg("error unmarshall data redis")
		os.Exit(0)
	}

	return list
}
