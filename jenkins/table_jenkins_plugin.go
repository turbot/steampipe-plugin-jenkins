package jenkins

import (
	"context"
	"strings"

	"github.com/turbot/steampipe-plugin-sdk/v4/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v4/plugin/transform"

	"github.com/turbot/steampipe-plugin-sdk/v4/plugin"
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
			{Name: "active", Type: proto.ColumnType_BOOL, Description: ""},
			{Name: "backup_version", Type: proto.ColumnType_JSON, Description: ""},
			{Name: "bundled", Type: proto.ColumnType_BOOL, Description: ""},
			{Name: "deleted", Type: proto.ColumnType_BOOL, Description: ""},
			{Name: "dependencies", Type: proto.ColumnType_JSON, Description: ""},
			{Name: "downgradable", Type: proto.ColumnType_BOOL, Description: ""},
			{Name: "enabled", Type: proto.ColumnType_BOOL, Description: ""},
			{Name: "has_update", Type: proto.ColumnType_BOOL, Description: ""},
			{Name: "long_name", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "pinned", Type: proto.ColumnType_BOOL, Description: ""},
			{Name: "short_name", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "supports_dynamic_load", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "title", Type: proto.ColumnType_STRING, Transform: transform.FromField("LongName"), Description: ""},
			{Name: "url", Type: proto.ColumnType_STRING, Transform: transform.FromField("URL"), Description: ""},
			{Name: "version", Type: proto.ColumnType_STRING, Description: ""},
		},
	}
}

//// LIST FUNCTION

func listJenkinsPlugins(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)
	client, err := Connect(ctx, d)
	if err != nil {
		logger.Error("listJenkinsPlugins", "connect_error", err)
		return nil, err
	}

	plugins, err := client.GetPlugins(ctx, 1)
	if err != nil {
		logger.Error("listJenkinsPlugins", "list_plugins_error", err)
		if strings.Contains(err.Error(), "Not found") {
			return nil, nil
		}
		return nil, err
	}

	for _, plugin := range plugins.Raw.Plugins {
		d.StreamListItem(ctx, plugin)

		// Context can be cancelled due to manual cancellation or the limit has been hit
		if d.QueryStatus.RowsRemaining(ctx) == 0 {
			return nil, nil
		}
	}

	return nil, err
}
