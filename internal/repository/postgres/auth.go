package postgres

import (
	"fmt"
	"log"
	"ypp/models"

	"github.com/jmoiron/sqlx"
)

type Auth struct{
	db *sqlx.DB
}

func NewAuth(db *sqlx.DB) *Auth {
	return &Auth{
		db: db,
	}
}

func (r *Auth)CreateUser(input models.User, salt []byte) (int, error){
	var id int
	query := "INSERT INTO users (name, username, password_hash, password_salt, email) values ($1, $2, $3, $4, $5) RETURNING id;"
	row := r.db.QueryRow(query, input.Name, input.Username, input.Password,fmt.Sprintf("%x", salt), input.Email)
	if err := row.Err(); err != nil{
		return 0, err
	}
	if err := row.Scan(&id); err != nil{
		return 0, err
	}
	log.Println("INSERT db done")
	return id, nil
}
/*
------------------------DOPISAT--------------------------
func (r *Auth)GetUser(username string) (models.User, error) {
	query := "SELECT * FROM users WHERE username=$1"
	var user models.User
	err := r.db.QueryRow(query, username).Scan(
		&uster.ID,
		&user.Name,
			
	)
} */

func (r *Auth) GetPass(username string) (*string, *string, error){
	var pass, salt string
	query := "SELECT (password_hash, password_salt) WHERE username=$1)"
	err := r.db.QueryRow(query, username).Scan(
		&pass,
		&salt,
	)
	if err != nil{
		return nil, nil, err
	}
	return pass, salt, nil
}
