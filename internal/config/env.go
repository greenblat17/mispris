package config

import "os"

var (
	ConfigPath = Getter("CONFIG_PATH", ".\\config.yaml")
)

func Getter(key, defaultValue string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return defaultValue
}
