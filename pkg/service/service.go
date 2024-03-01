package service

import "library-app/pkg/repository"

type Authorization interface {
}

type BookList interface {
}

type Book interface {
}

type Service struct {
	Authorization
	BookList
	Book
}

func NewService(repos *repository.Repository) *Service {
	return &Service{}
}
