package model

import "github.com/gogf/gf/v2/frame/g"

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

type UserInfoBase struct {
	g.Meta `orm:"table:user_info"`
	Id     uint
	Name   string
	Avatar string
	Sex    uint8
	Sign   string
	Status uint8
}
