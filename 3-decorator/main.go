package main

import (
	"3decorator/addon"
	"3decorator/base"
	"3decorator/interfaces"
	"encoding/json"
	"fmt"
)

func main() {
	var jcart []byte
	var user interfaces.UserInterface

	fmt.Println("All data:")
	user = base.NewUser()
	user = addon.NewPersonalAddOn(user)
	user = addon.NewWorkAddOn(user)
	user = addon.NewResidenceAddOn(user)

	jcart, _ = json.MarshalIndent(user.GetData(), "", "\t")
	fmt.Println(string(jcart))

	fmt.Println("Only Personal and Work Data:")
	user = base.NewUser()
	user = addon.NewPersonalAddOn(user)
	user = addon.NewWorkAddOn(user)

	jcart, _ = json.MarshalIndent(user.GetData(), "", "\t")
	fmt.Println(string(jcart))

	fmt.Println("Only Personal Data:")
	user = base.NewUser()
	user = addon.NewPersonalAddOn(user)

	jcart, _ = json.MarshalIndent(user.GetData(), "", "\t")
	fmt.Println(string(jcart))
}
