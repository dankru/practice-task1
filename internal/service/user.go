package service

import (
	"github.com/dankru/practice-task1/internal/domain"
)

type UserRepository interface {
	GetUsers() ([]domain.User, error)
}

type Service struct {
	repository UserRepository
}

func NewService(repository UserRepository) *Service {
	return &Service{repository: repository}
}

func (s *Service) GetUsers() ([]domain.User, error) {
	users, err := s.repository.GetUsers()
	return users, err
}
