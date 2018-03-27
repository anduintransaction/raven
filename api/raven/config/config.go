package config

import (
	"io/ioutil"
	"os"

	"github.com/palantir/stacktrace"
	"gopkg.in/yaml.v2"
)

// Config .
type Config struct {
	Logging    *LoggingConfig          `yaml:"logging"`
	Database   *DatabaseConfig         `yaml:"database"`
	Admin      *AdminAPIServerConfig   `yaml:"admin"`
	Mailgun    *MailgunAPIServerConfig `yaml:"mailgun"`
	SMTPServer *SMTPServerConfig       `yaml:"smtp_server"`
}

// LoggingConfig .
type LoggingConfig struct {
	Level  string `yaml:"level"`
	Output string `yaml:"output"`
}

// DatabaseConfig .
type DatabaseConfig struct {
	Driver           string `yaml:"driver"`
	ConnectionString string `yaml:"connection_string"`
}

// AdminAPIServerConfig .
type AdminAPIServerConfig struct {
	ListenAddress string `yaml:"listen_address"`
}

// MailgunAPIServerConfig .
type MailgunAPIServerConfig struct {
	ListenAddress string `yaml:"listen_address"`
}

// SMTPServerConfig .
type SMTPServerConfig struct {
	ListenAddress string `yaml:"listen_address"`
}

// ParseConfig .
func ParseConfig(configFile string) (*Config, error) {
	content, err := ioutil.ReadFile(configFile)
	if err != nil {
		return nil, stacktrace.Propagate(err, "cannot read config file %q", configFile)
	}
	contentStr := os.ExpandEnv(string(content))
	config := &Config{}
	err = yaml.Unmarshal([]byte(contentStr), config)
	if err != nil {
		return nil, stacktrace.Propagate(err, "cannot parse config file %q", configFile)
	}
	return config, nil
}
