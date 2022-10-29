package common

import (
	"os"
)

var (
	envBlogWorkdir = "BLOG_WORKDIR"
)

var (
	configService = NewConfig(ProvideEnvGetter())
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
