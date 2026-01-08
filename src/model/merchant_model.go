package model

type MerchantQueryParams struct {
	Page   int    `form:"page"`
	Limit  int    `form:"limit"`
	Status string `form:"status"`
}

type Merchant struct {
	Code             string           `json:"code"`
	CreatedAt        string           `json:"created_at"`
	Description      string           `json:"description"`
	IsProvider       bool             `json:"is_provider"`
	IsRedemption     bool             `json:"is_redemption"`
	MerchantCategory MerchantCategory `json:"merchant_category"`
	Name             string           `json:"name"`
	PriorityNumber   int              `json:"priority_number"`
	Status           string           `json:"status"`
	UpdatedAt        string           `json:"updated_at"`
	UUID             string           `json:"uuid"`
}

type MerchantCategory struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}
