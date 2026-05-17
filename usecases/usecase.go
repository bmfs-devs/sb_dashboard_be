package usecases

import (
	domainInterfaces "github.com/bmfs-devs/sb_dashboard_be/interfaces"
	mStart "github.com/bmfs-devs/sb_dashboard_be/repositories/start/models"
	mUsecase "github.com/bmfs-devs/sb_dashboard_be/usecases/models"
)

type Usecase struct {
	HelloWorldRepo domainInterfaces.HelloWorldRepository
	HiveRepo       domainInterfaces.HiveRepository
}

func NewUsecase(
	HelloWorldRepo domainInterfaces.HelloWorldRepository,
	HiveRepo domainInterfaces.HiveRepository,
) *Usecase {
	return &Usecase{
		HelloWorldRepo: HelloWorldRepo,
		HiveRepo:       HiveRepo,
	}
}

// Usecase only for Business Logic. Business Logic is forbidden in handler and repository
func (u *Usecase) GetHelloWorld(params mUsecase.HelloWorldUsecaseRequest) mUsecase.HelloWorldUsecaseResponse {
	r := u.HelloWorldRepo.GetHelloWorld(mStart.HelloWorldRequest{
		Ctx: params.Ctx,
	})
	return mUsecase.HelloWorldUsecaseResponse{
		Message: r.Message,
	}
}
