package types

type AuthResult struct {
	AccessToken  string `json:"AccessToken"`
	ExpiresIn    int    `json:"ExpiresIn"`
	IdToken      string `json:"IdToken"`
	TokenType    string `json:"TokenType"`
	RefreshToken string `json:"RefreshToken"`
}

type ResponseCognito struct {
	AuthenticationResult AuthResult `json:"AuthenticationResult,omitempty"`
	Message              string     `json:"message,omitempty"`
	Type                 string     `json:"__type,omitempty"`
}
