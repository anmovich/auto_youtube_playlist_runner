package repository

import (
	"ypp/internal/repository/postgres"
	"ypp/models"

	"github.com/jmoiron/sqlx"
)

type Authorization interface{
	CreateUser(models.User) (int, error)
	GetUserPass(string) (string, error)
	FailedLogin(username string)
	CheckLockedUntil(username string) bool
	// DEBUG -----------------------------------------------
	ClearLocked(username string) 	
}

type ChannelLink interface{

}

type VideoLink interface{

}

type Repository struct{
	Authorization
	ChannelLink
	VideoLink
}

func New(db *sqlx.DB) *Repository{
	return &Repository{
		Authorization: postgres.NewAuth(db),
	}
}
