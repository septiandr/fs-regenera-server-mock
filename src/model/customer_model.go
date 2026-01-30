package model

type CustomerCheckQuery struct {
	Page                   int    `form:"page"`
	Limit                  int    `form:"limit"`
	PhoneCode              string `form:"phone_code"`
	Phone                  string `form:"phone"`
	Name                   string `form:"name"`
	Search                 string `form:"search"`
	LegacyNoCustomer       string `form:"legacy_no_customer"`
	Gender                 string `form:"gender"`
	BirthDate              string `form:"birth_date"`
	RegistrationOutletUUID string `form:"registration_outlet_uuid"`
}

type RegistrationOutlet struct {
	UUID string `json:"uuid"`
	Name string `json:"name"`
	Code string `json:"code"`
}

type Customer struct {
	UUID                   string `json:"uuid"`
	PhoneCode              string `json:"phone_code"`
	Phone                  string `json:"phone"`
	Email                  string `json:"email"`
	IdentityNumber         string `json:"identity_number"`
	Name                   string `json:"name"`
	LegacyNoCustomer       string `json:"legacy_no_customer"`
	Gender                 string `json:"gender"`
	BirthDate              string `json:"birth_date"`
	RegistrationOutletUUID string `json:"registration_outlet_uuid"`
}
