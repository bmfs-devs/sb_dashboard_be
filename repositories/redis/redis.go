package redis

import (
	"errors"
	"fmt"
	"time"

	mRedis "github.com/bmfs-devs/sb_dashboard_be/repositories/redis/models"
	"github.com/redis/go-redis/v9"
)

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

func (r *Repository) Set(params mRedis.RedisMessageSetRequest) error {
	var ttl time.Duration
	if params.TTL == 0 {
		ttl = 365 * 24 * time.Hour
	} else {
		ttl = time.Duration(params.TTL) * time.Second
	}
	return r.Client.Set(params.Ctx, params.Key, params.Value, ttl).Err()
}

func (r *Repository) Delete(params mRedis.RedisMessageSetRequest) error {
	return r.Client.Del(params.Ctx, params.Key).Err()
}

func (r *Repository) Get(params mRedis.RedisMessageSetRequest) (string, error) {
	res, err := r.Client.Get(params.Ctx, params.Key).Result()
	if err != nil {
		// Check if the error is specifically because the key doesn't exist
		if errors.Is(err, redis.Nil) {
			return "", nil // Or return a custom app error like ErrNotFound
		}

		// This is a real error (e.g., network down, connection timeout)
		return "", err
	}
	return res, nil
}

func (r *Repository) HSet(params mRedis.RedisMessageHSetRequest) error {
	err := r.Client.HSet(params.Ctx, params.Key, params.Field).Err()
	if err != nil {
		return fmt.Errorf("failed to HSet key %s: %w", params.Key, err)
	}

	var ttl time.Duration
	if params.TTL == 0 {
		ttl = 365 * 24 * time.Hour
	} else {
		ttl = time.Duration(params.TTL) * time.Second
	}
	err = r.Client.Expire(params.Ctx, params.Key, ttl).Err()
	if err != nil {
		return fmt.Errorf("failed to set TTL for key %s: %w", params.Key, err)
	}
	return nil
}

func (r *Repository) HGet(params mRedis.RedisMessageHGetRequest) (string, error) {
	val, err := r.Client.HGet(params.Ctx, params.Key, params.Field).Result()

	if err != nil {
		// Check if the overall key or the specific field doesn't exist
		if errors.Is(err, redis.Nil) {
			return "", nil // Or return a custom domain error like ErrFieldNotFound
		}
		return "", fmt.Errorf("redis error: %w", err)
	}

	return val, nil
}

func (r *Repository) HDel(params mRedis.RedisMessageHGetRequest) error {
	return r.Client.HDel(params.Ctx, params.Key, params.Field).Err()
}

func (r *Repository) HGetAll(params mRedis.RedisMessageHGetRequest) (map[string]string, error) {
	res, err := r.Client.HGetAll(params.Ctx, params.Key).Result()
	if err != nil {
		return nil, fmt.Errorf("redis error: %w", err)
	}
	return res, nil
}
