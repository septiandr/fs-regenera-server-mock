package services

import (
	"encoding/json"
	"errors"
	"fs-regenera/src/model"
	"os"
	"sort"
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

	// 2.5 Sorting
	if query.SortBy != "" && query.SortType != "" {
		sort.Slice(filtered, func(i, j int) bool {
			var less bool
			switch query.SortBy {
			case model.BookingListSortBookingDate:
				less = filtered[i].BookingDate < filtered[j].BookingDate
			case model.BookingListSortCustomerName:
				less = filtered[i].CustomerName < filtered[j].CustomerName
			case model.BookingListSortOutletName:
				less = filtered[i].OutletName < filtered[j].OutletName
			case model.BookingListSortCreatedAt:
				less = filtered[i].CreatedAt < filtered[j].CreatedAt
			default:
				less = true
			}

			if query.SortType == model.SortDESC {
				return !less
			}
			return less
		})
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

func GetDetailBookingService() (
	model.Booking,
	error,
) {
	file, err := os.ReadFile("src/data/booking_list.json")
	if err != nil {
		return model.Booking{}, err
	}

	var bookings []model.Booking
	if err := json.Unmarshal(file, &bookings); err != nil {
		return model.Booking{}, err
	}

	// for _, b := range bookings {
	// 	if query.BookingUUID != "" && b.UUID == query.BookingUUID {
	// 		return b, nil
	// 	}

	// }

	return model.Booking{}, errors.New("booking not found")
}

func GetListLogBookingService() (
	model.BookingListLog,
	error,
) {
	file, err := os.ReadFile("src/data/bookings_log.json")
	if err != nil {
		return model.BookingListLog{}, err
	}

	var bookings []model.Booking
	if err := json.Unmarshal(file, &bookings); err != nil {
		return model.BookingListLog{}, err
	}

	return model.BookingListLog{}, errors.New("booking not found")

}

func GetBookingByUUIDService() (
	model.BookingDetail,
	error,
) {
	file, err := os.ReadFile("src/data/bookings_log.json")
	if err != nil {
		return model.BookingDetail{}, err
	}

	var bookings []model.Booking
	if err := json.Unmarshal(file, &bookings); err != nil {
		return model.BookingDetail{}, err
	}

	return model.BookingDetail{}, errors.New("booking not found")

}
