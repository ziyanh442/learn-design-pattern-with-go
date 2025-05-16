package addon

import (
	"3decorator/interfaces"
	"3decorator/model"
)

type ResidenceAddOn struct {
	User interfaces.UserInterface
}

func NewResidenceAddOn(user interfaces.UserInterface) interfaces.UserInterface {
	return &ResidenceAddOn{
		User: user,
	}
}

func (a *ResidenceAddOn) GetData() model.User {
	user := a.User.GetData()
	user.Residence = &model.ResidenceData{
		Address:       "Somewhere in the mountain where the snow never cease",
		OwnershipType: "Sole Ownership",
	}
	return user
}
