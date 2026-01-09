package services

import (
	"context"
	"encoding/json"
	"fs-regenera/src/model"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

func getFirstOutletName(outlets []model.DoctorOutlet) string {
	if len(outlets) == 0 {
		return ""
	}
	return outlets[0].Name
}

func GetDoctorListServices(c context.Context, params model.DoctorListParams) ([]model.DoctorListResponse, int, error) {
	filePath := "src/data/doctors.json"

	bytes, err := os.ReadFile(filePath)
	if err != nil {
		return nil, 0, err
	}

	var doctors []model.DoctorListResponse
	err = json.Unmarshal(bytes, &doctors)
	if err != nil {
		return nil, 0, err
	}

	filtered := make([]model.DoctorListResponse, 0)

	for _, d := range doctors {

		// filter search (by name)
		if params.Search != "" {
			if !strings.Contains(
				strings.ToLower(d.Name),
				strings.ToLower(params.Search),
			) {
				continue
			}
		}

		// filter outlet_uuid
		if params.OutletUUID != "" {
			found := false
			for _, o := range d.Outlets {
				if o.UUID == params.OutletUUID {
					found = true
					break
				}
			}
			if !found {
				continue
			}
		}

		// filter type
		if params.Type != "" && d.Type != params.Type {

			continue
		}

		// filter status
		if params.Status != "" && d.Status != params.Status {
			continue
		}

		filtered = append(filtered, d)
	}

	total := len(filtered)

	// =========================
	// SORTING
	// =========================
	if params.SortBy != "" && params.SortType != "" {
		sort.Slice(filtered, func(i, j int) bool {

			var less bool

			switch params.SortBy {
			case "name":
				less = filtered[i].Name < filtered[j].Name
			case "created_at":
				less = filtered[i].CreatedAt.Before(filtered[j].CreatedAt)
			case "sip_number":
				less = filtered[i].SIPNumber < filtered[j].SIPNumber
			case "registered_at":
				less = filtered[i].RegisteredAt.Before(filtered[j].RegisteredAt)
			case "outlet":
				less = getFirstOutletName(filtered[i].Outlets) <
					getFirstOutletName(filtered[j].Outlets)
			case "age":
				less = filtered[i].Age < filtered[j].Age
			default:
				less = true
			}

			if params.SortType == "DESC" {
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

	if start >= total {
		return []model.DoctorListResponse{}, total, nil
	}

	if end > total {
		end = total
	}

	paged := filtered[start:end]

	return paged, total, nil

}

func GetDoctorSessionsServices(
	ctx context.Context,
	params model.DoctorSessionsParams,
) ([]model.DoctorSessionResponse, error) {
	filePath := "src/data/doctor_sessions.json"

	bytes, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	var sessions []model.DoctorSessionResponse
	if err := json.Unmarshal(bytes, &sessions); err != nil {
		return nil, err
	}

	filtered := make([]model.DoctorSessionResponse, 0)

	for _, s := range sessions {

		// filter outlet_uuid
		if params.OutletUUID != "" && s.OutletUUID != params.OutletUUID {
			continue
		}

		// filter date
		if params.Date != "" {
			if s.Date == nil || *s.Date != params.Date {
				continue
			}
		}

		filtered = append(filtered, s)
	}

	return filtered, nil
}

func GetListDoctorBookedService(
	ctx context.Context,
	doctorUUID string,
	date string,
) ([]model.DoctorBookedResponse, error) {

	// 1. read sessions
	sessionBytes, err := os.ReadFile(filepath.Join("src", "data", "doctor_session.json"))
	if err != nil {
		return nil, err
	}

	var sessions []model.DoctorSessionResponse
	if err := json.Unmarshal(sessionBytes, &sessions); err != nil {
		return nil, err
	}

	// filter sessions sesuai doctor + date
	sessionIDs := make(map[int]bool)
	for _, s := range sessions {
		if s.DoctorUUID != doctorUUID {
			continue
		}
		if date != "" && (s.Date == nil || *s.Date != date) {
			continue
		}
		sessionIDs[s.ID] = true
	}

	if len(sessionIDs) == 0 {
		return []model.DoctorBookedResponse{}, nil
	}

	// 2. read booked
	bookedBytes, err := os.ReadFile(filepath.Join("src", "data", "doctor_booked.json"))
	if err != nil {
		return nil, err
	}

	var booked []model.DoctorBookedResponse
	if err := json.Unmarshal(bookedBytes, &booked); err != nil {
		return nil, err
	}

	// 3. filter booked sesuai sessionIDs
	filteredBooked := make([]model.DoctorBookedResponse, 0)
	for _, b := range booked {
		if sessionIDs[b.DoctorSessionID] {
			filteredBooked = append(filteredBooked, b)
		}
	}

	return filteredBooked, nil
}
