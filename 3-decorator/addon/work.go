package addon

import (
	"3decorator/interfaces"
	"3decorator/model"
)

type WorkAddOn struct {
	User interfaces.UserInterface
}

func NewWorkAddOn(user interfaces.UserInterface) interfaces.UserInterface {
	return &WorkAddOn{
		User: user,
	}
}

func (a *WorkAddOn) GetData() model.User {
	user := a.User.GetData()
	user.Work = &model.WorkData{
		JobTitle:       "Wallfacer",
		CompanyName:    "United Nations",
		CompanyAddress: "760 United Nations Plaza, Manhattan, New York, U.S.",
		Industry:       "Government",
	}
	return user
}
