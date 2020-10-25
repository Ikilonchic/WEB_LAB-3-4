package server

// Config ...
type Config struct {
	Port 			string `toml:"port"`
	LogLevel 		string `toml:"log_level"`
	Static 			string `toml:"static"`
	Templates 		string `toml:"templates"`
	sqlURL		 	string `toml:"sql"`
	redisURL		string `toml:"redis"`
}

// NewConfig ...
func NewConfig() *Config {
	return &Config{
		Port: ":8080",
		LogLevel: "debug",
		Static: "web/static",
		Templates: "web/",
		sqlURL: "",
		redisURL: "",
	}
}