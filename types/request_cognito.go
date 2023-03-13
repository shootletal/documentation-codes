package types

// AuthParametersRequest struct auth parameters for TokenRequest.
type AuthParametersRequest struct {
	Username string `json:"USERNAME"`
	Password string `json:"PASSWORD"`
}

// TokenRequest struct for get token form cognito.
type TokenRequest struct {
	AuthParameters AuthParametersRequest `json:"AuthParameters"`
	AuthFlow       string                `json:"AuthFlow"`
	ClientId       string                `json:"ClientId"`
}
