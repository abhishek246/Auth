package contracts

type ChangePassword struct {
	Email    string `json:"email_id" validate:"required"`
	Password string `json:"password" validate:"required"`
	ClientID string `json:"client_id" validate:"required"`
}
