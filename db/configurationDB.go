package config

type EnvDbConvig struct {
	DBPort     string
	DBHost     string
	DBUser     string
	DBPassword string
	DBName     string
}

func NewEnvDbConfig() *EnvDbConvig {
	return &EnvDbConvig{
		DBPort:     "DB_PORT",
		DBHost:     "DB_HOST",
		DBUser:     "DB_USER",
		DBPassword: "DB_PASSWORD",
		DBName:     "DB_NAME",
	}
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
