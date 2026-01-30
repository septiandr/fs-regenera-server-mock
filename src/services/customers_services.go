package services

import (
	"encoding/json"
	"fmt"
	"fs-regenera/src/model"
	"os"
	"regexp"
)

var nonDigitRegex = regexp.MustCompile(`\D`)

func normalizePhoneCode(v string) string {
	return nonDigitRegex.ReplaceAllString(v, "")
}

func GetListCustomersService(
	query model.CustomerCheckQuery,
) (
	result []model.Customer,
	total int,
	err error,
) {

	// ⚠️ ABSOLUTE PATH (BIAR TIDAK SALAH FILE)
	file, err := os.ReadFile("src/data/customer_list.json")
	if err != nil {
		fmt.Println("READ FILE ERROR:", err)
		return nil, 0, err
	}

	fmt.Println("RAW JSON SIZE:", len(file))

	var customers []model.Customer
	if err := json.Unmarshal(file, &customers); err != nil {
		fmt.Println("UNMARSHAL ERROR:", err)
		return nil, 0, err
	}

	fmt.Println("TOTAL CUSTOMER RAW:", len(customers))

	filtered := make([]model.Customer, 0)

	for _, c := range customers {

		fmt.Println("CHECK CUSTOMER:", c.Name, c.PhoneCode)

		// phone code
		if query.PhoneCode != "" {
			qc := normalizePhoneCode(query.PhoneCode)
			cc := normalizePhoneCode(c.PhoneCode)

			if qc != cc {
				continue
			}
		}

		filtered = append(filtered, c)
	}

	fmt.Println("FILTERED COUNT:", len(filtered))

	total = len(filtered)

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
