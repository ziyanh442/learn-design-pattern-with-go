package model

type User struct {
	Personal  *PersonalData
	Work      *WorkData
	Residence *ResidenceData
}

type PersonalData struct {
	IDNumber   string
	Name       string
	BirthDate  string
	BirthPlace string
}

type WorkData struct {
	JobTitle       string
	CompanyName    string
	CompanyAddress string
	Industry       string
}

type ResidenceData struct {
	Address       string
	OwnershipType string
}
