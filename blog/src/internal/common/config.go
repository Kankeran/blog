package common

type EnvGetter func(string) string

type Config struct {
	getEnv EnvGetter
}

func NewConfig(getEnv EnvGetter) *Config {
	return &Config{getEnv: getEnv}
}

func (c *Config) GetEnv(key string) string {
	return c.getEnv(key)
}
