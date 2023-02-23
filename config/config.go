package config

import "os"

type Config struct {
	SERVER_ADDRESS string
	DB_USERNAME    string
	DB_PASSWORD    string
	DB_PORT        string
	DB_HOST        string
	DB_NAME        string
	JWT_KEY        string
	LOC            string
}

func InitConfiguration() Config {

	return Config{
		SERVER_ADDRESS: GetOrDefault("SERVER_ADDRESS", "0.0.0.0:1323"),
		DB_USERNAME:    GetOrDefault("DB_USERNAME", "root"),
		DB_PASSWORD:    GetOrDefault("DB_PASSWORD", "T3Csk2IF3bW4R44EsFhM"),
		DB_NAME:        GetOrDefault("DB_NAME", "railway"),
		DB_PORT:        GetOrDefault("DB_PORT", "7938"),
		DB_HOST:        GetOrDefault("DB_HOST", "containers-us-west-129.railway.app"),
		JWT_KEY:        GetOrDefault("JWT_KEY", "sungguhRahasia"),
		LOC:            GetOrDefault("LOC", "Asia%2FJakarta"),
	}
}

func GetOrDefault(envKey, defaultValue string) string {
	// cek env
	if val, exist := os.LookupEnv(envKey); exist {
		return val
	}

	return defaultValue
}
