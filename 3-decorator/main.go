package main

import (
	"3decorator/addon"
	"3decorator/base"
	"encoding/json"
	"fmt"
)

func main() {
	var jcart []byte

	fmt.Println("All data:")
	luoJi := base.NewUser()
	personalDecorator := addon.NewPersonalAddOn(luoJi)
	workDecorator := addon.NewWorkAddOn(personalDecorator)
	residenceDecorator := addon.NewResidenceAddOn(workDecorator)

	jcart, _ = json.MarshalIndent(residenceDecorator.GetData(), "", "\t")
	fmt.Println(string(jcart))

	fmt.Println("Only Personal and Work Data:")
	luoJi_two := base.NewUser()
	personalDecorator_two := addon.NewPersonalAddOn(luoJi_two)
	workDecorator_two := addon.NewWorkAddOn(personalDecorator_two)

	jcart, _ = json.MarshalIndent(workDecorator_two.GetData(), "", "\t")
	fmt.Println(string(jcart))
}
