package service

import (
	"github.com/dankru/practice-task1/internal/domain"
)

type UserRepository interface {
	GetAll() ([]domain.User, error)
	Create(user domain.User) error
	Update(id int64, user domain.User) error
	Delete(id int64) error
}

type Service struct {
	repository UserRepository
}

func NewService(repository UserRepository) *Service {
	return &Service{repository: repository}
}

func (s *Service) GetAll() ([]domain.User, error) {
	users, err := s.repository.GetAll()
	return users, err
}

func (s *Service) Create(user domain.User) error {
	err := s.repository.Create(user)
	return err
}

func (s *Service) Update(id int64, user domain.User) error {
	err := s.repository.Update(id, user)
	return err
}

func (s *Service) Delete(id int64) error {
	err := s.repository.Delete(id)
	return err
}
