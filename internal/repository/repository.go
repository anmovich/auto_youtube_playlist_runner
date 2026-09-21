package repository

import (
	"ypp/internal/repository/postgres"
	"ypp/models"

	"github.com/jmoiron/sqlx"
)

type Authorization interface{
	CreateUser(models.User, []byte) (int, error)
	GetPass(string) (*string, *string, error)
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
