package config

type (
	Config struct {
		App      `yaml:"app"`
		HTTP     `yaml:"HTTP"`
		Postgres `yaml:"postgres"`
	}
	App struct {
		Name    string `yaml:"name" env:"APP_NAME"`
		Version string `yaml:"version" env:"APP_VERSION"`
	}
	Postgres struct {
		Host     string
		Port     string
		Username string
		Password string
		DbName   string
		Timeout  int
	}
	HTTP struct {
		Port string `yaml:"port" env:"HTTP_PORT"`
	}
)
