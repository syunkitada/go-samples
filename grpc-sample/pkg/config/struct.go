package config

import (
	"path/filepath"
)

type Config struct {
	Default      DefaultConfig
	SimpleServer SimpleServerConfig
}

type DefaultConfig struct {
	ConfigDir string
}

type SimpleServerConfig struct {
	Grpc GrpcConfig
}

type GrpcConfig struct {
	Listen             string
	CertFile           string
	KeyFile            string
	CaFile             string
	ServerHostOverride string
	Timeout            int
	Targets            []string
}

func newConfig(configDir string) *Config {
	defaultConfig := &Config{
		Default: DefaultConfig{
			ConfigDir: configDir,
		},
		SimpleServer: SimpleServerConfig{
			Grpc: GrpcConfig{
				Listen:             "localhost:23300",
				CertFile:           "server.pem",
				KeyFile:            "server.key",
				CaFile:             "ca.pem",
				ServerHostOverride: "x.test.youtube.com",
				Timeout:            2,
				Targets: []string{
					"localhost:23300",
				},
			},
		},
	}
	return defaultConfig
}

func (conf *Config) Path(path string) string {
	return filepath.Join(conf.Default.ConfigDir, path)
}
