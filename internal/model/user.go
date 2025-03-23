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

type LoginInput struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

type UpdatePasswordInput struct {
	Password     string
	UserSalt     string
	SecretAnswer string
}

type UpdatePasswordOutput struct {
	Id uint
}
