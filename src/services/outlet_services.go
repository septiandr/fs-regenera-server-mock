package services

import (
	"context"
	"encoding/json"
	"fs-regenera/src/model"
	"os"
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
	return outlets, len(outlets), nil
}
