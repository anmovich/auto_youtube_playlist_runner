package service

import (
	"crypto/rand"
	"crypto/sha1"
	"fmt"
	"ypp/internal/repository"
	"ypp/models"
)

type Auth struct {
	repo repository.Authorization
}

func NewAuth(repo repository.Authorization) *Auth{
	return &Auth{
		repo: repo,
	}
}

type Passwd struct{
	Salt []byte
	Password_hash string
}

func (s *Auth)CreateUser(usr *models.User) (int, error) {
	
	pwd, err := generatePasswordHash(usr.Password)
	usr.Password = pwd.Password_hash
	if err != nil{
		return 0, err
	}
	id, err := s.repo.CreateUser(*usr, pwd.Salt)
	if err != nil{
		return 0, err
	}
	return id, nil
}

func (s *Auth) SignIN(usr *models.User) (string, error) {
	return "", nil
}

func generatePasswordHash(password string) (*Passwd, error) {
	var ret Passwd
	var err error
	hash := sha1.New()
	ret.Salt, err = generateSalt(16)
	if err != nil{
		return nil, err
	}
	hash.Write([]byte(password))
	ret.Password_hash = fmt.Sprintf("%x", hash.Sum(ret.Salt))
	return &ret, nil
}

func generateSalt(lenght int) ([]byte, error){
	petersalt := make([]byte, lenght)
	if _, err := rand.Read(petersalt); err != nil{
		return nil, err
	}
	return petersalt, nil
}

