package models

import "context"

// This file will contains many DTOs for Start Repository
// Start Domain can have many repo, each repo has its own DTOs
type HelloWorldRequest struct {
	Ctx context.Context
}
type HelloWorldResponse struct {
	Message string
}
