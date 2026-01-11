package config

import "os"

type EnvDbConvig struct {
	DBPort     string
	DBHost     string
	DBUser     string
	DBPassword string
	DBName     string
}

func NewEnvDbConfig() *EnvDbConvig {
	return &EnvDbConvig{
		DBPort:     getEnv("DB_PORT", "5432"),
		DBHost:     getEnv("DB_HOST", "localhost"),
		DBUser:     getEnv("DB_USER", "postgres"),
		DBPassword: getEnv("DB_PASSWORD", "password"),
		DBName:     getEnv("DB_NAME", "videocall_db"),
	}
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

func (c *EnvDbConvig) GetUsernameDb() string {
	return c.DBUser
}

func (c *EnvDbConvig) GetHostDb() string {
	return c.DBHost
}

func (c *EnvDbConvig) GetPasswordDb() string {
	return c.DBPassword
}

func (c *EnvDbConvig) GetPortDb() string {
	return c.DBPort
}

func (c *EnvDbConvig) GetNameDb() string {
	return c.DBName
}
