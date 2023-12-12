package jenkins

import (
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

type jenkinsConfig struct {
	ServerURL *string `hcl:"server_url"`
	Username  *string `hcl:"username"`
	Password  *string `hcl:"password"`
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
