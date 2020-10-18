package serviceprovider

// AccessTokenSchema is a schema for parsing token object
// from the service provider
type AccessTokenSchema struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	Scope       string `json:"scope"`
	ExpiresIn   int32  `json:"expires_in"`
}
