package services

import (
	"encoding/json"
	"fmt"
	"os"
	"regexp"
	"strings"

	"fs-regenera/src/model"
)

var nonDigitRegex = regexp.MustCompile(`\D`)

func normalizePhoneCode(v string) string {
	return nonDigitRegex.ReplaceAllString(v, "")
}

func normalizeString(v string) string {
	return strings.ToLower(strings.TrimSpace(v))
}

func GetListCustomersService(
	query model.CustomerCheckQuery,
) (
	result []model.Customer,
	total int,
	err error,
) {

	// read json file
	file, err := os.ReadFile("src/data/customer_list.json")
	if err != nil {
		fmt.Println("READ FILE ERROR:", err)
		return nil, 0, err
	}

	var customers []model.Customer
	if err := json.Unmarshal(file, &customers); err != nil {
		fmt.Println("UNMARSHAL ERROR:", err)
		return nil, 0, err
	}

	filtered := make([]model.Customer, 0)

	for _, c := range customers {

		// filter phone code (exact)
		if query.PhoneCode != "" {
			qc := normalizePhoneCode(query.PhoneCode)
			cc := normalizePhoneCode(c.PhoneCode)

			if qc != cc {
				continue
			}
		}
		// filter birth date (exact)
		if query.BirthDate != "" {
			if c.BirthDate != query.BirthDate {
				continue
			}
		}

		// filter name (include / contains)
		if query.Name != "" {
			qn := normalizeString(query.Name)
			cn := normalizeString(c.Name)

			if !strings.Contains(cn, qn) {
				continue
			}
		}

		// filter legacy no customer (exact)
		if query.LegacyNoCustomer != "" {
			if c.LegacyNoCustomer != query.LegacyNoCustomer {
				continue
			}
		}

		filtered = append(filtered, c)
	}

	total = len(filtered)

	// pagination
	if query.Page <= 0 {
		query.Page = 1
	}
	if query.Limit <= 0 {
		query.Limit = 10
	}

	start := (query.Page - 1) * query.Limit
	end := start + query.Limit

	if start >= total {
		return []model.Customer{}, total, nil
	}

	if end > total {
		end = total
	}

	return filtered[start:end], total, nil
}
