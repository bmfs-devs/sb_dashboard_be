package interfaces

import (
	mHive "github.com/bmfs-devs/sb_dashboard_be/repositories/queue/hive/models"
	mRedis "github.com/bmfs-devs/sb_dashboard_be/repositories/redis/models"
	mStart "github.com/bmfs-devs/sb_dashboard_be/repositories/start/models"
	mUsecase "github.com/bmfs-devs/sb_dashboard_be/usecases/models"
)

type Usecase interface {
	GetHelloWorld(params mUsecase.HelloWorldUsecaseRequest) mUsecase.HelloWorldUsecaseResponse
	OnOffAC(params mUsecase.ACRemoteRequest) (mUsecase.ACRemoteResponse, error)
	SetTemperature(params mUsecase.ACRemoteRequest) (mUsecase.ACRemoteResponse, error)
	GetTemperature(params mUsecase.GetTemperatureRequest) (mUsecase.ACRemoteResponse, error)
}

type HelloWorldRepository interface {
	GetHelloWorld(params mStart.HelloWorldRequest) mStart.HelloWorldResponse
}

type RedisRepository interface {
	Set(params mRedis.RedisMessageSetRequest) error
	Delete(params mRedis.RedisMessageSetRequest) error
	Get(params mRedis.RedisMessageSetRequest) (string, error)
	HSet(params mRedis.RedisMessageHSetRequest) error
	HGet(params mRedis.RedisMessageHGetRequest) (string, error)
	HDel(params mRedis.RedisMessageHGetRequest) error
	HGetAll(params mRedis.RedisMessageHGetRequest) (map[string]string, error)
}

type HiveRepository interface {
	Publish(message mHive.HiveMessage) error
}
