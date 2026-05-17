package models

import "context"

// This file will contains many DTOs for only HelloWorld Usecase
// Start Domain can be used by many usecase, each usecase has its own DTOs
// Models and filename in usecase are specific to that usecase rather than domain
type HelloWorldUsecaseRequest struct {
	Ctx context.Context
}
type HelloWorldUsecaseResponse struct {
	Message string
}
