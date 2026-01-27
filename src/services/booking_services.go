package services

import (
	"encoding/json"
	"fs-regenera/src/model"
	"os"
	"strings"
	"time"
)

func GetListBookingService(query model.BookingListQuery) (
	result []model.Booking,
	total int,
	err error,
) {
	// 1. Read JSON file
	file, err := os.ReadFile("src/data/booking_list.json")
	if err != nil {
		return nil, 0, err
	}

	var bookings []model.Booking
	if err := json.Unmarshal(file, &bookings); err != nil {
		return nil, 0, err
	}

	// 2. Filtering
	filtered := make([]model.Booking, 0)

	for _, b := range bookings {

		// 🔍 Search
		if query.Search != "" {
			search := strings.ToLower(query.Search)
			if !strings.Contains(strings.ToLower(b.CustomerName), search) &&
				!strings.Contains(strings.ToLower(b.LegacyNoCustomer), search) &&
				!strings.Contains(strings.ToLower(b.Code), search) {
				continue
			}
		}

		// 📌 Status
		if query.Status != "" && b.Status != query.Status {
			continue
		}

		// 🏥 Outlet
		if query.OutletUUID != "" && b.OutletUUID != query.OutletUUID {
			continue
		}

		// 👨‍⚕️ Doctor
		if query.DoctorUUID != "" && b.DoctorUUID != query.DoctorUUID {
			continue
		}

		// 📅 Booking date range
		if query.BookingStartDate != "" || query.BookingEndDate != "" {
			bookingDate, err := time.Parse("2006-01-02", b.BookingDate)
			if err != nil {
				continue
			}

			if query.BookingStartDate != "" {
				start, err := time.Parse("2006-01-02", query.BookingStartDate)
				if err == nil && bookingDate.Before(start) {
					continue
				}
			}

			if query.BookingEndDate != "" {
				end, err := time.Parse("2006-01-02", query.BookingEndDate)
				if err == nil && bookingDate.After(end) {
					continue
				}
			}
		}

		filtered = append(filtered, b)
	}

	total = len(filtered)

	// 3. Pagination
	start := (query.Page - 1) * query.Limit
	end := start + query.Limit

	if start > total {
		return []model.Booking{}, total, nil
	}

	if end > total {
		end = total
	}

	result = filtered[start:end]

	return result, total, nil
}
