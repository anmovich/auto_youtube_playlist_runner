package postgres

import (
	"errors"
	"log"
	"time"
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

func (r *Auth)CreateUser(input models.User) (int, error){
	var id int
	query := "INSERT INTO users (name, username, password_hash, email) values ($1, $2, $3, $4) RETURNING id;"
	row := r.db.QueryRow(query, input.Name, input.Username, input.Password, input.Email)
	if err := row.Err(); err != nil{
		return 0, err
	}
	if err := row.Scan(&id); err != nil{
		return 0, err
	}
	log.Println("INSERT db done")
	return id, nil
}

func (r *Auth) GetUserPass(username string) (string, error){
	var pass string
	query := "SELECT password_hash FROM users WHERE username=$1"
	err := r.db.QueryRow(query, username).Scan(
		&pass,
	)
	if err != nil{
		log.Println(err)
		return "", errors.New("No user with this username")
	}
	return pass, nil
}

func (r *Auth) FailedLogin(username string){
	r.addOneToFailedLogin(username)
	attempts, err := r.checkLoginAttempts(username)
	if err != nil {
		log.Println("Fail login ", err, username)
		return
	}
	if attempts >= 3 {
		log.Println("Login locked for ", username)
		r.setLockedUntil(username)
	}

}

//Return true if sing in locked and if locked_until before time now return false and null locked_until
func (r *Auth) CheckLockedUntil(username string) bool{
	query := "SELECT locked_until FROM users WHERE username = $1"
	var t *time.Time
	r.db.QueryRow(query, username).Scan(
		&t,
	)
	
	if t != nil{
		if t.Before(time.Now()){
			log.Println("time now before until", username)
			r.setNULLLockedUntil(username)
			r.clearLoginAttemps(username)
			return false
		}
		log.Println("Login locked", username)
		return  true
	}
	log.Println("Login not locked", username)
	return false
}


func (r *Auth) setNULLLockedUntil (username string){
	query := "UPDATE users SET locked_until = NULL WHERE username = $1"
	if  _, err := r.db.Query(query, username); err != nil{
		log.Println(err)
	}
}

func (r *Auth) setLockedUntil(username string) {
	t := time.Now()
	t = t.Add(5 * time.Minute)
	log.Println(t)
	query := "UPDATE users SET locked_until = $1 WHERE username = $2"
	r.db.Query(query, t, username)
}

func (r *Auth) checkLoginAttempts(username string) (int, error){
	query := "SELECT failed_login_attempts FROM users WHERE username = $1"
	var attempts int
	err := r.db.QueryRow(query, username).Scan(
		&attempts,
	)
	if err != nil{
		return 0, err
	}
	return attempts, nil
}

func (r *Auth) clearLoginAttemps(username string) {
	query := "UPDATE users SET failed_login_attempts = 0 WHERE username = $1"
	r.db.Query(query, username)
}

func (r *Auth) addOneToFailedLogin(username string) {
	query := "UPDATE users SET failed_login_attempts = failed_login_attempts + 1 WHERE username = $1"
	_, err := r.db.Query(query, username)
	log.Println("add one:", err)
}
