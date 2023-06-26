package jenkins

import (
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/schema"
)

type jenkinsConfig struct {
	ServerURL *string `cty:"server_url"`
	Username  *string `cty:"username"`
	Password  *string `cty:"password"`
}

var ConfigSchema = map[string]*schema.Attribute{
	"server_url": {
		Type: schema.TypeString,
	},
	"username": {
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
