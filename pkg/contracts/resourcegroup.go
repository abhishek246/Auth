package contracts

type ResourceGroupCreateRequest struct {
	GroupName        string `json:"group_name" validate:"required"`
	GroupDescription string `json:"group_description"`
}

type ResourceGroupUpdateRequest struct {
	GroupName        string `json:"group_name" validate:"required"`
	GroupDescription string `json:"group_description"`
}
