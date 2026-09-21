package postgres

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	 _ "github.com/lib/pq"
)

type Config struct{
	Host string
	Port string
	Username string
	Password string
	DBName string
	SSlMode string
}

func NewDB(cfg Config) (*sqlx.DB, error){
	db, err := sqlx.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
				cfg.Host, cfg.Port, cfg.Username, cfg.DBName, cfg.Password, cfg.SSlMode))
	if err != nil{
		return nil, err
	}
	if err := db.Ping(); err != nil{
		return nil, err
	}
	return db, nil
}
