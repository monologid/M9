package config

import "time"

// Schema is a base schema for configuration
type Schema struct {
	Application ApplicationSchema `yaml:"application"`
	Database    DatabaseSchema    `yaml:"database"`
	Security    SecuritySchema    `yaml:"security"`
	Facebook    FacebookSchema    `yaml:"facebook"`
	Google      GoogleSchema      `yaml:"Google"`

	Module ModuleSchema `yaml:"module"`
}

// ApplicationSchema is a schema related with application configuration
type ApplicationSchema struct {
	Host                    string        `yaml:"host"`
	Port                    string        `yaml:"port"`
	ReadTimeout             time.Duration `yaml:"readTimeout"`
	WriteTimeout            time.Duration `yaml:"writeTimeout"`
	GracefulShutdownTimeout time.Duration `yaml:"gracefulShutdownTimeout"`
}

// DatabaseSchema is a schema related with database configuration
type DatabaseSchema struct {
	Engine string `yaml:"engine"`
	URL    string `yaml:"url"`
}

// SecuritySchema is a schema related with security configuration
type SecuritySchema struct {
	JWT JWTSchema `yaml:"jwt"`
}

// JWTSchema is a schema related with JWT configuration
type JWTSchema struct {
	Secret string `yaml:"secret"`
}

// FacebookSchema is a schema related with Facebook configuration
type FacebookSchema struct {
	ClientID     string              `yaml:"clientId"`
	ClientSecret string              `yaml:"clientSecret"`
	GraphURL     string              `yaml:"graphUrl"`
	Oauth        FacebookOauthSchema `yaml:"oauth"`
}

// FacebookOauthSchema is a schema related with Facebook OAuth configuration
type FacebookOauthSchema struct {
	URL         string `yaml:"url"`
	RedirectURI string `yaml:"redirectUri"`
	Scope       string `yaml:"scope"`
	Fields      string `yaml:"fields"`
}

// GoogleSchema is a schema related with Google configuration
type GoogleSchema struct {
	ClientID     string            `yaml:"clientId"`
	ClientSecret string            `yaml:"clientSecret"`
	APIURL       string            `yaml:"apiUrl"`
	Oauth        GoogleOauthSchema `yaml:"oauth"`
}

// GoogleOauthSchema is a schema related with Google OAuth configuration
type GoogleOauthSchema struct {
	URL         string `yaml:"url"`
	TokenURL    string `yaml:"tokenUrl"`
	RedirectURI string `yaml:"redirectUri"`
	Scope       string `yaml:"scope"`
}

// ModuleSchema is a schema related with module configuration
type ModuleSchema struct {
	Login ModuleLoginSchema `yaml:"login"`
}

// ModuleLoginSchema is a schema related with login module configuration
type ModuleLoginSchema struct {
	IsRedirect  bool   `yaml:"isRedirect"`
	RedirectURI string `yaml:"redirectUri"`
}
