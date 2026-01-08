package services

import (
	"context"
	"encoding/json"
	"fs-regenera/src/model"
	"os"
)

func GetMerchantsListServices(
	ctx context.Context,
	params model.MerchantQueryParams,
) ([]model.Merchant, int, error) {

	filePath := "src/data/merchants.json"

	bytes, err := os.ReadFile(filePath)
	if err != nil {
		return nil, 0, err
	}

	var merchants []model.Merchant
	if err := json.Unmarshal(bytes, &merchants); err != nil {
		return nil, 0, err
	}

	// =====================
	// FILTER
	// =====================
	var filtered []model.Merchant
	for _, merchant := range merchants {
		if params.Status != "" && merchant.Status != params.Status {
			continue
		}
		filtered = append(filtered, merchant)
	}

	total := len(filtered)

	// =====================
	// PAGINATION
	// =====================
	page := params.Page
	limit := params.Limit

	if page <= 0 {
		page = 1
	}
	if limit <= 0 {
		limit = 10
	}

	start := (page - 1) * limit
	end := start + limit

	if start >= total {
		return []model.Merchant{}, total, nil
	}
	if end > total {
		end = total
	}

	return filtered[start:end], total, nil
}
