package contracts

type ResourcePermissionCreateRequest struct {
	PermissionName    string `json:"name"`
	PermissionAction  string `json:"action"`
	ResourceGroupName string `json:"resource_group_name"`
}
