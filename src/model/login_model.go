package model

type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type LoginResponse struct {
	UUID         string   `json:"uuid"`
	Name         string   `json:"name"`
	Email        string   `json:"email"`
	IsSuperAdmin bool     `json:"is_super_admin"`
	RoleID       *int     `json:"role_id"`
	RoleName     *string  `json:"role_name"`
	Token        string   `json:"token"`
	Features     []string `json:"features"`
}
