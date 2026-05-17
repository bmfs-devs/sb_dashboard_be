package redis

import "github.com/redis/go-redis/v9"

type Repository struct {
	Client *redis.Client
}

func NewRepository(
	client *redis.Client,
) *Repository {
	return &Repository{
		Client: client,
	}
}
