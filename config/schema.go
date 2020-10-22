package config

import "time"

// Schema ...
type Schema struct {
	Application ApplicationSchema `yaml:"application"`
	Database    DatabaseSchema    `yaml:"database"`
	Security    SecuritySchema    `yaml:"security"`
	Facebook    FacebookSchema    `yaml:"facebook"`
	Google      GoogleSchema      `yaml:"Google"`

	Module ModuleSchema `yaml:"module"`
}

// ApplicationSchema ...
type ApplicationSchema struct {
	Host                    string        `yaml:"host"`
	Port                    string        `yaml:"port"`
	ReadTimeout             time.Duration `yaml:"readTimeout"`
	WriteTimeout            time.Duration `yaml:"writeTimeout"`
	GracefulShutdownTimeout time.Duration `yaml:"gracefulShutdownTimeout"`
}

// DatabaseSchema ...
type DatabaseSchema struct {
	Engine string `yaml:"engine"`
	URL    string `yaml:"url"`
}

// SecuritySchema ...
type SecuritySchema struct {
	JWT JWTSchema `yaml:"jwt"`
}

// JWTSchema ...
type JWTSchema struct {
	Secret string `yaml:"secret"`
}

// FacebookSchema ...
type FacebookSchema struct {
	ClientID     string              `yaml:"clientId"`
	ClientSecret string              `yaml:"clientSecret"`
	GraphURL     string              `yaml:"graphUrl"`
	Oauth        FacebookOauthSchema `yaml:"oauth"`
}

// FacebookOauthSchema ...
type FacebookOauthSchema struct {
	URL         string `yaml:"url"`
	RedirectURI string `yaml:"redirectUri"`
	Scope       string `yaml:"scope"`
	Fields      string `yaml:"fields"`
}

// GoogleSchema ...
type GoogleSchema struct {
	ClientID     string            `yaml:"clientId"`
	ClientSecret string            `yaml:"clientSecret"`
	APIURL       string            `yaml:"apiUrl"`
	Oauth        GoogleOauthSchema `yaml:"oauth"`
}

// GoogleOauthSchema ...
type GoogleOauthSchema struct {
	URL         string `yaml:"url"`
	TokenURL    string `yaml:"tokenUrl"`
	RedirectURI string `yaml:"redirectUri"`
	Scope       string `yaml:"scope"`
}

// ModuleSchema ...
type ModuleSchema struct {
	Login ModuleLoginSchema `yaml:"login"`
}

// ModuleLoginSchema ...
type ModuleLoginSchema struct {
	IsRedirect  bool   `yaml:"isRedirect"`
	RedirectURI string `yaml:"redirectUri"`
}
