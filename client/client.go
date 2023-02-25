package client

import (
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/spf13/viper"
)

// ------------------------------------------------- --------------------------------------------------------------------

type Config struct {
	FooCount     int `json:"foo-count" yaml:"foo-count" mapstructure:"foo-count"`
	BarCount     int `json:"bar-count" yaml:"bar-count" mapstructure:"bar-count"`
	SleepSeconds int `json:"sleep-seconds" yaml:"sleep-seconds" mapstructure:"sleep-seconds"`
}

// ------------------------------------------------- --------------------------------------------------------------------

type Client struct {
	Config *Config
}

func NewClient(config *viper.Viper) (*Client, *schema.Diagnostics) {
	c := Config{}
	err := config.Unmarshal(&c)
	if err != nil {
		return nil, schema.NewDiagnostics().AddErrorMsg("config unmarshal error: %s", err.Error())
	}

	diagnostics := schema.NewDiagnostics()
	if c.SleepSeconds < 0 {
		diagnostics.AddErrorMsg("sleep-seconds must >= 0")
	}
	if c.FooCount <= 0 {
		diagnostics.AddErrorMsg("foo-count must > 0")
	}
	if c.BarCount <= 0 {
		diagnostics.AddErrorMsg("bar-count must > 0")
	}

	return &Client{
		Config: &c,
	}, diagnostics
}

// ------------------------------------------------- --------------------------------------------------------------------
