package entities

type GoogleOAuthRequest struct {
	Code string `json:"code" validate:"required"`
}

type GoogleOAuthResponse struct {
	Name    string `json:"name"`
	Email   string `json:"email"`
	Picture string `json:"picture"`
}
