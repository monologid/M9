package config

import "time"

// Schema ...
type Schema struct {
	Application ApplicationSchema `yaml:"application"`
	Facebook    FacebookSchema    `yaml:"facebook"`
	Google      GoogleSchema      `yaml:"Google"`
}

// ApplicationSchema ...
type ApplicationSchema struct {
	Host                    string        `yaml:"host"`
	Port                    string        `yaml:"port"`
	ReadTimeout             time.Duration `yaml:"readTimeout"`
	WriteTimeout            time.Duration `yaml:"writeTimeout"`
	GracefulShutdownTimeout time.Duration `yaml:"gracefulShutdownTimeout"`
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
