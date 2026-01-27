package model

type BookingRequest struct {
	CustomerUUID    string `json:"customer_uuid"`
	MerchantUUID    string `json:"merchant_uuid"`
	OutletUUID      string `json:"outlet_uuid"`
	DoctorUUID      string `json:"doctor_uuid"`
	DoctorSessionID int    `json:"doctor_session_id"`
	BookingDate     string `json:"booking_date"` // Format: YYYY-MM-DD
}
