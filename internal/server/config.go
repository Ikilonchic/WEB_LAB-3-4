package server

// Config ...
type Config struct {
	Port 			string `toml:"port"`
	LogLevel 		string `toml:"log_level"`
	Static 			string `toml:"static"`
	Templates 		string `toml:"templates"`
	DatabaseURL 	string `toml:"db"`
	MailHost 		string `toml:"mail_host"`
	Email			string `toml:"email"`
	Password 		string `toml:"password"`
}

// NewConfig ...
func NewConfig() *Config {
	return &Config{
		Port: ":8080",
		LogLevel: "debug",
		Static: "web/static",
		Templates: "web/",
		DatabaseURL: "",
		MailHost: "",
		Email: "",
		Password: "",
	}
}