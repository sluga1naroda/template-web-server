package config

// Config is a config.
type Config struct {
	Application ApplicationConfig
	HTTPServer  HTTPServerConfig
	UserInfo    UserInfo
}

// UserInfo is a userinfo for application.
type UserInfo struct {
	Name                   string
	Status                 string
	Email                  string
	Education              string
	EducationAdd           string
	EducationDirector      string
	EducationDirectorPhone string
	Friend1Name            string
	Friend1Email           string
	Friend1ImageUrl        string
	Friend2Name            string
	Friend2Email           string
	Friend2ImageUrl        string
}

// ApplicationConfig is a config for application.
type ApplicationConfig struct {
	Name     string
	Debug    bool
	LogLevel string
}

// HTTPServerConfig is a config for http settings.
type HTTPServerConfig struct {
	Host string
	Port int
}
