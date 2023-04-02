package dto

type Auth struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type UpdateAuth struct {
	Username    string `json:"username" validate:"required"`
	OldPassword string `json:"old_password" validate:"omitempty"`
	NewPassword string `json:"new_password" validate:"required_with=OldPassword"`
}
