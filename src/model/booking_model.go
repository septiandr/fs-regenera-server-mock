package model

type BookingRequest struct {
	CustomerUUID    string `json:"customer_uuid"`
	MerchantUUID    string `json:"merchant_uuid"`
	OutletUUID      string `json:"outlet_uuid"`
	DoctorUUID      string `json:"doctor_uuid"`
	DoctorSessionID int    `json:"doctor_session_id"`
	BookingDate     string `json:"booking_date"` // Format: YYYY-MM-DD
}

type BookingSummaryResponse struct {
	Today       int `json:"today"`
	Scheduled   int `json:"scheduled"`
	Rescheduled int `json:"rescheduled"`
	Upcoming    int `json:"upcoming"`
	Canceled    int `json:"canceled"`
	Finished    int `json:"finished"`
}

type BookingListQuery struct {
	Page             int    `form:"page" json:"page"`
	Limit            int    `form:"limit" json:"limit"`
	Search           string `form:"search" json:"search"`
	Status           string `form:"status" json:"status"`
	BookingStartDate string `form:"booking_start_date" json:"booking_start_date"`
	BookingEndDate   string `form:"booking_end_date" json:"booking_end_date"`
	OutletUUID       string `form:"outlet_uuid" json:"outlet_uuid"`
	DoctorUUID       string `form:"doctor_uuid" json:"doctor_uuid"`
}

type Booking struct {
	UUID             string `json:"uuid"`
	Code             string `json:"code"`
	CustomerName     string `json:"customer_name"`
	LegacyNoCustomer string `json:"legacy_no_customer"`
	Status           string `json:"status"`
	BookingDate      string `json:"booking_date"`
	BookingStart     string `json:"booking_start"`
	BookingEnd       string `json:"booking_end"`
	Timezone         string `json:"timezone"`
	MerchantUUID     string `json:"merchant_uuid"`
	MerchantName     string `json:"merchant_name"`
	OutletUUID       string `json:"outlet_uuid"`
	OutletName       string `json:"outlet_name"`
	DoctorUUID       string `json:"doctor_uuid"`
	DoctorName       string `json:"doctor_name"`
	CreatedAt        string `json:"created_at"`
	UpdatedAt        string `json:"updated_at"`
}

type BookingDetailQueryParams struct {
	BookingUUID string `form:"booking_uuid" json:"booking_uuid"`
}
