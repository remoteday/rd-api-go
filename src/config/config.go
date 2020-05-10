package config

import "time"

// Auth -
type Auth struct {
	Auth0ClientID     string `envconfig:"AUTH0_CLIENT_ID"`
	Auth0ClientSecret string `envconfig:"AUTH0_CLIENT_SECRET"`
	Auth0Subdomain    string `envconfig:"AUTH0_SUBDOMAIN"`
}

// Database -
type Database struct {
	DBHost string `default:"localhost" envconfig:"DB_HOST"`
	DBPort string `default:"5432" envconfig:"DB_PORT"`
	DBUser string `default:"master" envconfig:"DB_USER"`
	DBPass string `default:"Episource123" envconfig:"DB_PASS"`
	DBName string `default:"epia_test" envconfig:"DB_NAME"`
}

// Web -
type Web struct {
	APIPort         string        `default:"3000" envconfig:"PORT"`
	DebugPort       string        `default:"4000" envconfig:"DEBUG_PORT"`
	APIHost         string        `default:"0.0.0.0" envconfig:"API_HOST"`
	DebugHost       string        `default:"0.0.0.0" envconfig:"DEBUG_HOST"`
	ReadTimeout     time.Duration `default:"5s" envconfig:"READ_TIMEOUT"`
	WriteTimeout    time.Duration `default:"5s" envconfig:"WRITE_TIMEOUT"`
	ShutdownTimeout time.Duration `default:"5s" envconfig:"SHUTDOWN_TIMEOUT"`
}
