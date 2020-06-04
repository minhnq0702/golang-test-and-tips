package models

// Partner be using for customer/supplier
type Partner struct {
	Model
	Name        string
	Code        string
	PartnerType string
	Title       string
	JobPosition string
	Street      string
	Street2     string
	District    string
	City        string
	Country     string
	Zip         string
	IsCompany   bool
	Active      bool
}
