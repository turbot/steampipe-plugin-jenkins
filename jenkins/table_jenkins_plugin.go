package jenkins

import (
	"context"
	"strings"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"

	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

//// TABLE DEFINITION

func tableJenkinsPlugin() *plugin.Table {
	return &plugin.Table{
		Name:        "jenkins_plugin",
		Description: "An extension to Jenkins functionality provided separately from Jenkins Core.",
		List: &plugin.ListConfig{
			Hydrate: listJenkinsPlugins,
		},

		Columns: []*plugin.Column{
			{Name: "active", Type: proto.ColumnType_BOOL, Description: "Boolean to indicate whether the plugin is active."},
			{Name: "backup_version", Type: proto.ColumnType_JSON, Description: "The backup version of the plugin available for downgrade."},
			{Name: "bundled", Type: proto.ColumnType_BOOL, Description: "Boolean to indicate whether the plugin is bundled with Jenkins in the WAR file."},
			{Name: "deleted", Type: proto.ColumnType_BOOL, Description: "Boolean to indicate whether the plugin has been deleted."},
			{Name: "dependencies", Type: proto.ColumnType_JSON, Description: "A list of other plugins this depends on."},
			{Name: "downgradable", Type: proto.ColumnType_BOOL, Description: "Boolean to indicate whether a downgrade can be performed on the plugin."},
			{Name: "enabled", Type: proto.ColumnType_BOOL, Description: "Boolean to indicate whether the plugin is enable."},
			{Name: "has_update", Type: proto.ColumnType_BOOL, Description: "Boolean to indicate when an update is available."},
			{Name: "long_name", Type: proto.ColumnType_STRING, Description: "A human-readable full name of the plugin."},
			{Name: "pinned", Type: proto.ColumnType_BOOL, Description: "Boolean to indicate whether the plugin is pinned on UI."},
			{Name: "short_name", Type: proto.ColumnType_STRING, Description: "Unique key for the plugin."},
			{Name: "supports_dynamic_load", Type: proto.ColumnType_STRING, Description: "Boolean to indicate whether the plugin can be dynamically loaded."},
			{Name: "title", Type: proto.ColumnType_STRING, Transform: transform.FromField("LongName"), Description: titleDescription},
			{Name: "url", Type: proto.ColumnType_STRING, Transform: transform.FromField("URL"), Description: "Full URL to the installed plugin."},
			{Name: "version", Type: proto.ColumnType_STRING, Description: "Current installed version."},
		},
	}
}

//// LIST FUNCTION

func listJenkinsPlugins(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)
	client, err := Connect(ctx, d)
	if err != nil {
		logger.Error("jenkins_plugin.listJenkinsPlugins", "connect_error", err)
		return nil, err
	}

	plugins, err := client.GetPlugins(ctx, 10)
	if err != nil {
		logger.Error("jenkins_plugin.listJenkinsPlugins", "query_error", err)
		if strings.Contains(err.Error(), "Not found") {
			return nil, nil
		}
		return nil, err
	}

	for _, plugin := range plugins.Raw.Plugins {
		d.StreamListItem(ctx, plugin)

		// Context can be cancelled due to manual cancellation or the limit has been hit
		if d.RowsRemaining(ctx) == 0 {
			return nil, nil
		}
	}

	return nil, err
}
