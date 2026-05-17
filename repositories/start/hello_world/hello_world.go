package helloworld

import (
	mStart "github.com/bmfs-devs/sb_dashboard_be/repositories/start/models"
)

type Repository struct {
}

func NewRepository() *Repository {
	return &Repository{}
}

// Repositories acts as gateway or access to data source (external and internal)
// Repository must be stateless and thread-safe and general so any func can use this
func (r *Repository) GetHelloWorld(params mStart.HelloWorldRequest) mStart.HelloWorldResponse {
	return mStart.HelloWorldResponse{
		Message: "Hello World",
	}
}
