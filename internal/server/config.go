package server

// Config ...
type Config struct {
	Port 			string `toml:"port"`
	LogLevel 		string `toml:"log_level"`
	Static 			string `toml:"static"`
	Templates 		string `toml:"templates"`
	sqlURL		 	string `toml:"sql"`
	tokenPassword   string `toml:"token_password"`
}

// NewConfig ...
func NewConfig() *Config {
	return &Config{
		Port: 			":8080",
		LogLevel:		"debug",
		Static: 		"web/static",
		Templates: 		"web/",
		sqlURL: 		"user=postgres password=12345 dbname=web port=5432 sslmode=disable",
		tokenPassword: 	"supersecretpass",
	}
}