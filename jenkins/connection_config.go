package jenkins

import (
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/schema"
)

type jenkinsConfig struct {
	Domain   *string `cty:"url"`
	UserId   *string `cty:"user_id"`
	ApiToken *string `cty:"api_token"`
}

var ConfigSchema = map[string]*schema.Attribute{
	"url": {
		Type: schema.TypeString,
	},
	"user_id": {
		Type: schema.TypeString,
	},
	"api_token": {
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
