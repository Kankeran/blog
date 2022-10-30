package common

import (
	"os"
	"sync"
)

var (
	envBlogWorkdir     = "BLOG_WORKDIR"
	envPostgresHost    = "POSTGRES_HOST"
	envPostgresPort    = "POSTGRES_PORT"
	envPostgresUser    = "POSTGRES_USER"
	envPostgresPass    = "POSTGRES_PASS"
	envPostgresDb      = "POSTGRES_DB"
	envPostgresSslmode = "POSTGRES_SSLMODE"
)

var (
	configService   = NewConfig(ProvideEnvGetter())
	databaseService = NewDatabase(
		ProvideConfig().GetEnv(envPostgresHost),
		ProvideConfig().GetEnv(envPostgresPort),
		ProvideConfig().GetEnv(envPostgresUser),
		ProvideConfig().GetEnv(envPostgresPass),
		ProvideConfig().GetEnv(envPostgresDb),
		ProvideConfig().GetEnv(envPostgresSslmode),
		&sync.Once{},
	)
)

func ProvideEnvGetter() EnvGetter {
	return os.Getenv
}

func ProvideConfig() *Config {
	return configService
}

func ProvideWorkdir() string {
	return ProvideConfig().GetEnv(envBlogWorkdir) + "/blog"
}

func ProvideDatabase() *Database {
	return databaseService
}
