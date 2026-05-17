package main

import (
	hive "github.com/bmfs-devs/sb_dashboard_be/dal/hive"
	redis "github.com/bmfs-devs/sb_dashboard_be/dal/redis"
	resource "github.com/bmfs-devs/sb_dashboard_be/dal/resources"
	"github.com/bmfs-devs/sb_dashboard_be/routes"
)

func main() {
	// Init Client
	hive.InitHiveMQ()
	redis.InitRedisClient()

	// Init Repository and Usecase
	resource.InitNewRepo()
	resource.InitNewUsecase()
	routes.InitFiber()
}
