package config

type Config struct {
	Addr string
	Assets string
}

func New() Config {
	return Config{
		Addr: "localhost:8080",
		Assets: "web/public/assets",
	}
}