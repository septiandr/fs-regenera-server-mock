package services

import (
	"context"
	"encoding/json"
	"fs-regenera/src/model"
	"os"
	"sort"
	"strings"
)

func GetOutletListServices(ctx context.Context, params model.OutletListParams) ([]model.OutletResponse, int, error) {
	//dummy data
	filePath := "src/data/outlets.json"

	bytes, err := os.ReadFile(filePath)
	if err != nil {
		return nil, 0, err
	}

	var outlets []model.OutletResponse
	err = json.Unmarshal(bytes, &outlets)
	if err != nil {
		return nil, 0, err
	}

	filteredOutlets := make([]model.OutletResponse, 0)
	for _, outlet := range outlets {
		//filter status
		if params.Status != "" && outlet.Status != string(params.Status) {
			continue
		}

		//filter search ""
		if params.Search != "" {
			if !strings.Contains(strings.ToLower(outlet.Name), strings.ToLower(params.Search)) {
				continue
			}
		}
		filteredOutlets = append(filteredOutlets, outlet)
	}

	total := len(filteredOutlets)

	//=====Sort===

	if params.SortBy != "" && params.SortType != "" {
		sort.Slice(filteredOutlets, func(i, j int) bool {
			var less bool
			switch params.SortBy {
			case model.SortByCode:
				less = filteredOutlets[i].Code < filteredOutlets[j].Code
			case model.SortByName:
				less = filteredOutlets[i].Name < filteredOutlets[j].Name
			case model.SortByCreatedAt:
				less = filteredOutlets[i].CreatedAt.Before(filteredOutlets[j].CreatedAt)
			default:
				less = true
			}
			if params.SortType == model.SortDESC {
				return !less
			}
			return less
		})
	}

	// =========================
	// PAGINATION
	// =========================
	start := (params.Page - 1) * params.Limit
	end := start + params.Limit

	if start > total {
		return []model.OutletResponse{}, total, nil
	}

	if end > total {
		end = total
	}

	paged := filteredOutlets[start:end]

	return paged, total, nil
}
