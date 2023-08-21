package contracts

type UserPermissionRequest struct {
	UserId           string `json:"user_id" validate:"required"`
	PermissionName   string `json:"permission_name"`
	PermissionAction string `json:"action"`
}
