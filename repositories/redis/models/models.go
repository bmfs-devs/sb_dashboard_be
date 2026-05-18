package models

import "context"

type RedisMessageSetRequest struct {
	Ctx   context.Context
	Key   string
	Value string
	TTL   int // Time to live in seconds, 0 for no expiration
}

type RedisMessageHSetRequest struct {
	Ctx   context.Context
	Key   string
	Field map[string]interface{}
	TTL   int // Time to live in seconds, 0 for no expiration

}

type RedisMessageHGetRequest struct {
	Ctx   context.Context
	Key   string
	Field string
}
