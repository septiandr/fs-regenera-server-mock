package services

import (
	"encoding/json"
	"fs-regenera/src/model"
	"os"
	"strings"
	"time"
)

func GetListCustomersService(
	query model.CustomerCheckQuery,
) (
	result []model.Customer,
	total int,
	err error,
) {

	// Read JSON file
	file, err := os.ReadFile("src/data/customer_list.json")
	if err != nil {
		return nil, 0, err
	}

	var customers []model.Customer
	if err := json.Unmarshal(file, &customers); err != nil {
		return nil, 0, err
	}

	filtered := make([]model.Customer, 0)

	for _, c := range customers {

		// 🔍 global search (name, phone, legacy)
		if query.Search != "" {
			s := strings.ToLower(query.Search)
			if !strings.Contains(strings.ToLower(c.Name), s) &&
				!strings.Contains(c.Phone, s) &&
				!strings.Contains(strings.ToLower(c.LegacyNoCustomer), s) {
				continue
			}
		}

		// 👤 name
		if query.Name != "" &&
			!strings.Contains(strings.ToLower(c.Name), strings.ToLower(query.Name)) {
			continue
		}

		// 📞 phone code
		if query.PhoneCode != "" && c.PhoneCode != query.PhoneCode {
			continue
		}

		// 📱 phone
		if query.Phone != "" && !strings.Contains(c.Phone, query.Phone) {
			continue
		}

		// 🆔 legacy number
		if query.LegacyNoCustomer != "" &&
			!strings.Contains(strings.ToLower(c.LegacyNoCustomer), strings.ToLower(query.LegacyNoCustomer)) {
			continue
		}

		// ⚧ gender
		if query.Gender != "" && c.Gender != query.Gender {
			continue
		}

		// 🏥 registration outlet
		if query.RegistrationOutletUUID != "" &&
			c.RegistrationOutletUUID != query.RegistrationOutletUUID {
			continue
		}

		// 🎂 birth date (optional, if exists in data)
		if query.BirthDate != "" {
			birth, err1 := time.Parse("2006-01-02", c.BirthDate)
			filter, err2 := time.Parse("2006-01-02", query.BirthDate)
			if err1 == nil && err2 == nil && !birth.Equal(filter) {
				continue
			}
		}

		filtered = append(filtered, c)
	}

	total = len(filtered)

	// Pagination
	start := (query.Page - 1) * query.Limit
	end := start + query.Limit

	if start > total {
		return []model.Customer{}, total, nil
	}
	if end > total {
		end = total
	}

	result = filtered[start:end]
	return result, total, nil
}
