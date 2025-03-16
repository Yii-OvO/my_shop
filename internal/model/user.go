package model

type RegisterInput struct {
	Name         string `json:"name"`
	Avatar       string `json:"avatar"`
	Password     string `json:"password"`
	UserSalt     string `json:"userSalt"`
	Sex          int    `json:"sex"`
	Status       int    `json:"status"`
	Sign         string `json:"sign"`
	SecretAnswer string `json:"secretAnswer"`
}

type RegisterOutput struct {
	Id uint
}
