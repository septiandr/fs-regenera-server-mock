package model

import "time"

type DoctorListResponse struct {
	UUID         string         `json:"uuid"`
	Name         string         `json:"name"`
	ProfilePhoto string         `json:"profile_photo"`
	SIPNumber    string         `json:"sip_number"`
	RegisteredAt time.Time      `json:"registered_at"`
	Type         string         `json:"type"` // e.g. "REGULAR"
	Outlets      []DoctorOutlet `json:"outlets"`
	Age          int            `json:"age"`
	Status       string         `json:"status"` // e.g. "ACTIVE"
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
}

type DoctorOutlet struct {
	UUID string `json:"uuid"`
	Name string `json:"name"`
}

type DoctorListParams struct {
	Page       int    `form:"page" binding:"omitempty,min=1"`
	Limit      int    `form:"limit" binding:"omitempty,min=1"`
	Search     string `form:"search" binding:"omitempty,min=3"`
	OutletUUID string `form:"outlet_uuid" binding:"omitempty,uuid"`
	Type       string `form:"type" binding:"omitempty,oneof=REGULAR GUEST"`
	Status     string `form:"status" binding:"omitempty,oneof=ACTIVE INACTIVE"`
	SortBy     string `form:"sort_by" binding:"omitempty,oneof=name created_at sip_number registered_at outlet age"`
	SortType   string `form:"sort_type" binding:"omitempty,oneof=ASC DESC"`
}
