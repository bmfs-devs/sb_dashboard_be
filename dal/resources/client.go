package resource

import (
	"github.com/bmfs-devs/sb_dashboard_be/dal/hive"
	redis "github.com/bmfs-devs/sb_dashboard_be/dal/redis"
	domainInterfaces "github.com/bmfs-devs/sb_dashboard_be/interfaces"
	domainHiveRepo "github.com/bmfs-devs/sb_dashboard_be/repositories/queue/hive"
	domainRedisRepo "github.com/bmfs-devs/sb_dashboard_be/repositories/redis"
	domainHelloWorldRepo "github.com/bmfs-devs/sb_dashboard_be/repositories/start/hello_world"
	"github.com/bmfs-devs/sb_dashboard_be/usecases"
)

// All Repositories
var (
	HelloWorldRepository domainInterfaces.HelloWorldRepository
	HiveRepository       domainInterfaces.HiveRepository
	RedisRepository      domainInterfaces.RedisRepository
)

// All Usecase
var (
	Usecase domainInterfaces.Usecase
)

// Init All Repo
func InitNewRepo() {
	InitHelloWorldRepository()
	InitHiveRepository()
}

func InitHelloWorldRepository() {
	HelloWorldRepository = domainHelloWorldRepo.NewRepository()
}

func InitHiveRepository() {
	HiveRepository = domainHiveRepo.NewRepository(hive.MQTTClient)
}

func InitRedisRepository() {
	RedisRepository = domainRedisRepo.NewRepository(redis.RedisClient)
}

// Init All Usecase
func InitNewUsecase() {
	Usecase = usecases.NewUsecase(HelloWorldRepository, HiveRepository)
}
