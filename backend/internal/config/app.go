package config

type App struct {
	Name string `envconfig:"APP_NAME" required:"true"`
	Port string `envconfig:"APP_PORT" required:"true" default:"8080"`
}
