package services

import (
	"context"
	"fs-regenera/src/model"
)

func LoginService(ctx context.Context, req model.LoginRequest) (model.LoginResponse, error) {

	response := model.LoginResponse{
        "uuid": "c28f392b-337c-41eb-8277-03bc0246d89b",
        "name": "Super Admin",
        "email": "superadmin@mailinator.com",
        "is_super_admin": true,
        "role_id": null,
        "role_name": null,
        "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6InN1cGVyYWRtaW5AbWFpbGluYXRvci5jb20iLCJleHAiOjE3Njc4NTkyODAsImZlYXR1cmVzIjpudWxsLCJpZCI6MSwiaW1hZ2UiOiIiLCJpc19zdXBlcl9hZG1pbiI6dHJ1ZSwibmFtZSI6IlN1cGVyIEFkbWluIiwicm9sZV9pZCI6IiIsInJvbGVfbmFtZSI6IiIsInV1aWQiOiJjMjhmMzkyYi0zMzdjLTQxZWItODI3Ny0wM2JjMDI0NmQ4OWIifQ.DO3JpoHlORLo9v8MfZ3TYBkMpibPlYyEdf0YP6GmTLU",
        "features": null,
    },

	return response, nil
}
