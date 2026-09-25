package service

import (
	"ypp/internal/repository"
	"ypp/models"
)

type Authorization interface{
	CreateUser(*models.User) (int, error)
	SignIN(usr *models.UserSingInUsername) (error) 
	ClearLocked(username string)
}



type Service struct{
	Authorization
}

func NewService(repo *repository.Repository) *Service{
	return &Service{
		Authorization:  NewAuth(repo),
	}
}
