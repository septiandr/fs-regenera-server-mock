package model

import "time"

type Status string

const (
	StatusInactive Status = "INACTIVE"
	StatusActive   Status = "ACTIVE"
)

type SortBy string

const (
	SortByCode      SortBy = "code"
	SortByName      SortBy = "name"
	SortByCreatedAt SortBy = "created_at"
)

type SortType string

const (
	SortASC  SortType = "ASC"
	SortDESC SortType = "DESC"
)

type OutletListParams struct {
	Page     int      `form:"page" binding:"omitempty,min=1"`
	Limit    int      `form:"limit" binding:"omitempty,min=1"`
	Search   string   `form:"search" binding:"omitempty"`
	Status   Status   `form:"status" binding:"omitempty,oneof=ACTIVE INACTIVE"`
	SortBy   SortBy   `form:"sort_by" binding:"omitempty,oneof=code name created_at,required_with=SortType"`
	SortType SortType `form:"sort_type" binding:"omitempty,oneof=ASC DESC,required_with=SortBy"`
}

type OutletResponse struct {
	UUID      string    `json:"uuid"`
	Code      string    `json:"code"`
	Name      string    `json:"name"`
	Address   string    `json:"address"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
