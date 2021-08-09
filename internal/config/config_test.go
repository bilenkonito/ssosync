package config_test

import (
	"testing"

	. "github.com/awslabs/ssosync/internal/config"

	"github.com/stretchr/testify/assert"
)

func TestConfig(t *testing.T) {
	a := assert.New(t)

	cfg := New()

	a.NotNil(cfg)

	a.Equal(cfg.LogLevel, DefaultLogLevel)
	a.Equal(cfg.LogFormat, DefaultLogFormat)
	a.Equal(cfg.Debug, DefaultDebug)
	a.Equal(cfg.GoogleCredentials, DefaultGoogleCredentials)
	a.Equal(cfg.GoogleCustomer, DefaultGoogleCustomer)
}
