package config

type Postgres struct {
	URL string `env:"POSTGRES_URL" env-required:"true"`
}
