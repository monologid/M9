package serviceprovider

type (
	// AccessTokenSchema ...
	AccessTokenSchema struct {
		AccessToken string `json:"access_token"`
		TokenType   string `json:"token_type"`
		Scope       string `json:"scope"`
		ExpiresIn   int32  `json:"expires_in"`
	}

	// ProfileSchema ...
	ProfileSchema struct {
		ID        string `json:"id"`
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
		Email     string `json:"email"`
		Birthday  string `json:"birthday"`
	}
)
