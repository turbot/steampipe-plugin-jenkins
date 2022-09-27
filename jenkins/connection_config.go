package jenkins

import (
	"github.com/turbot/steampipe-plugin-sdk/v4/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v4/plugin/schema"
)

type jenkinsConfig struct {
	Domain   *string `cty:"domain"`
	User     *string `cty:"user"`
	Password *string `cty:"password"`
}

var ConfigSchema = map[string]*schema.Attribute{
	"domain": {
		Type: schema.TypeString,
	},
	"user": {
		Type: schema.TypeString,
	},
	"password": {
		Type: schema.TypeString,
	},
}

func ConfigInstance() interface{} {
	return &jenkinsConfig{}
}

// GetConfig :: retrieve and cast connection config from query data
func GetConfig(connection *plugin.Connection) jenkinsConfig {
	if connection == nil || connection.Config == nil {
		return jenkinsConfig{}
	}
	config, _ := connection.Config.(jenkinsConfig)
	return config
}
