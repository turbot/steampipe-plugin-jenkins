package jenkins

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

const pluginName = "steampipe-plugin-jenkins"

// Plugin creates this (jenkins) plugin
func Plugin(ctx context.Context) *plugin.Plugin {
	p := &plugin.Plugin{
		Name:             pluginName,
		DefaultTransform: transform.FromCamel(),
		ConnectionConfigSchema: &plugin.ConnectionConfigSchema{
			NewInstance: ConfigInstance,
			Schema:      ConfigSchema,
		},
		TableMap: map[string]*plugin.Table{
			"jenkins_build":     tableJenkinsBuild(),
			"jenkins_folder":    tableJenkinsFolder(),
			"jenkins_freestyle": tableJenkinsFreestyle(),
			"jenkins_job":       tableJenkinsJob(),
			"jenkins_node":      tableJenkinsNode(),
			"jenkins_pipeline":  tableJenkinsPipeline(),
			"jenkins_plugin":    tableJenkinsPlugin(),
		},
	}

	return p
}
