package base

import (
	"3decorator/interfaces"
	"3decorator/model"
)

type Base struct {
	User model.User
}

func NewUser() interfaces.UserInterface {
	return &Base{
		User: model.User{},
	}
}

func (p *Base) GetData() model.User {
	return p.User
}
