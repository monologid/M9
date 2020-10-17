package login

type (
	// FacebookAccessTokenSchema ...
	FacebookAccessTokenSchema struct {
		AccessToken string `json:"access_token"`
		TokenType   string `json:"token_type"`
		ExpiresIn   int32  `json:"expires_in"`
	}

	// FacebookProfileSchema ...
	FacebookProfileSchema struct {
		ID        string `json:"id"`
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
		Email     string `json:"email"`
		Birthday  string `json:"birthday"`
	}
)
