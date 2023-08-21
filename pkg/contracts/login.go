package contracts

type LoginRequest struct {
	EmailID  string `json:"email_id"`
	Password string `json:"password"`
	ClientId string `json:"client_id"`
}

type LoginResponse struct {
	Token       string `json:"token"`
	RedirectURI string `json:"redirect_uri"`
}
