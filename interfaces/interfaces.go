package interfaces

import (
	mHive "github.com/bmfs-devs/sb_dashboard_be/repositories/queue/hive/models"
	mStart "github.com/bmfs-devs/sb_dashboard_be/repositories/start/models"
	mUsecase "github.com/bmfs-devs/sb_dashboard_be/usecases/models"
)

type Usecase interface {
	GetHelloWorld(params mUsecase.HelloWorldUsecaseRequest) mUsecase.HelloWorldUsecaseResponse
	OnOffAC(params mUsecase.ACRemoteRequest) (mUsecase.ACRemoteResponse, error)
}

type HelloWorldRepository interface {
	GetHelloWorld(params mStart.HelloWorldRequest) mStart.HelloWorldResponse
}

type RedisRepository interface{}

type HiveRepository interface {
	Publish(message mHive.HiveMessage) error
}
