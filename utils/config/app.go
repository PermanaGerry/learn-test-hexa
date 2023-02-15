package config

func LoadConfig() {
	LoadEnvVar()
	OpenRedisPool()
}
