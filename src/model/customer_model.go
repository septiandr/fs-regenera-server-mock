package model

type CustomerCheckQuery struct {
	Page                   int    `form:"page" json:"page"`
	Limit                  int    `form:"limit" json:"limit"`
	Search                 string `form:"search" json:"search"`
	Gender                 string `form:"gender" json:"gender"`
	RegistrationOutletUUID string `form:"registration_outlet_uuid" json:"registration_outlet_uuid"`
	Name                   string `form:"name" json:"name"`
	PhoneCode              string `form:"phone_code" json:"phone_code"`
	Phone                  string `form:"phone" json:"phone"`
	LegacyNoCustomer       string `form:"legacy_no_customer" json:"legacy_no_customer"`
	BirthDate              string `form:"birth_date" json:"birth_date"` // YYYY-MM-DD
}

type RegistrationOutlet struct {
	UUID string `json:"uuid"`
	Name string `json:"name"`
	Code string `json:"code"`
}

type Customer struct {
	UUID                   string             `json:"uuid"`
	PhoneCode              string             `json:"phone_code"`
	Phone                  string             `json:"phone"`
	Name                   string             `json:"name"`
	LegacyNoCustomer       string             `json:"legacy_no_customer"`
	Gender                 string             `json:"gender"`
	BirthDate              string             `json:"birth_date"` // YYYY-MM-DD
	RegistrationOutletUUID string             `json:"registration_outlet_uuid"`
	RegistrationOutlet     RegistrationOutlet `json:"registration_outlet"`
}
