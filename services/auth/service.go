package auth

import (
    "github.com/templateOfService/models"
)

type Service struct {
    repo *Repo
}

func NewService() *Service {
    return &Service{
        repo: NewRepo(),
    }
}

func (s *Service) Signup(user models.User) string {
    err := s.repo.InsertUser(user)
    if err != nil {
        return ErrInternal
    }
    return SuccessMessage
}
