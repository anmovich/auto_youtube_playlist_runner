package service

import (
	"errors"
	"log"
	"ypp/internal/repository"
	"ypp/models"

	"golang.org/x/crypto/bcrypt"
)

type Auth struct {
	repo repository.Authorization
}

func NewAuth(repo repository.Authorization) *Auth{
	return &Auth{
		repo: repo,
	}
}

func (s *Auth)CreateUser(usr *models.User) (int, error) {
	var err error
	usr.Password, err = generatePasswordHash(usr.Password)
	if err != nil{
		return 0, err
	}
	id, err := s.repo.CreateUser(*usr)
	if err != nil{
		return 0, err
	}
	return id, nil
}

func (s *Auth) SignIN(usr *models.UserSingInUsername) (error) {
	pass_hash, err := s.repo.GetUserPass(usr.Username)	
	if err != nil{
		return errors.New("Smt wrong")
	}
	if s.repo.CheckLockedUntil(usr.Username){
		return errors.New("You cant log in now")
	}
	if err := checkPasswordHash(usr.Password, pass_hash); err != nil{
		s.repo.FailedLogin(usr.Username)
		log.Println("service: ", err)
		return  errors.New("Fail log in")
	}
	s.repo.ClearLocked(usr.Username)
	
	return nil
}


func checkPasswordHash(passwordInput string, passwordDB string) (error) {
	err := bcrypt.CompareHashAndPassword([]byte(passwordDB), []byte(passwordInput))
	if err != nil{
		return err
	}
	return nil
}

func generatePasswordHash(password string) (string, error) {
	passwd := []byte(password)
	ret, err := bcrypt.GenerateFromPassword(passwd, bcrypt.DefaultCost)
	if err != nil{
		return "", err
	}
	return string(ret), nil
}

