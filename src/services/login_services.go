package services

import (
	"context"
	"fs-regenera/src/model"
)

func LoginService(ctx context.Context, req model.LoginRequest) (string, error) {
	return "token-12345", nil
}
