package models

type User struct{
	ID int `json:"id"`
	Name string `json:"name"`
	Username string `json:"username"`
	Password string `json:"password" `
	Email string `json:"email"`
}

type UserSingInUsername struct{
	Username string `json:"username"`
	Password string `json:"password"`
}
