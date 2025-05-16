package addon

import (
	"3decorator/interfaces"
	"3decorator/model"
)

type PersonalAddOn struct {
	User interfaces.UserInterface
}

func NewPersonalAddOn(user interfaces.UserInterface) interfaces.UserInterface {
	return &PersonalAddOn{
		User: user,
	}
}

func (a *PersonalAddOn) GetData() model.User {
	user := a.User.GetData()
	user.Personal = &model.PersonalData{
		IDNumber:   "234567890",
		Name:       "Luo Ji",
		BirthDate:  "1975-03-03",
		BirthPlace: "Shanghai",
	}
	return user
}
