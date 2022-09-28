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
		Description: "A plugin is a runnable entity on Jenkins.",
		List: &plugin.ListConfig{
			Hydrate: listJenkinsPlugins,
			KeyColumns: plugin.KeyColumnSlice{
				{Name: "id", Require: plugin.Optional},
			},
		},

		Columns: []*plugin.Column{
			{Name: "active", Type: proto.ColumnType_BOOL, Transform: transform.FromField("Active"), Description: ""},
			{Name: "bundled", Type: proto.ColumnType_BOOL, Transform: transform.FromField("Bundled"), Description: ""},
			{Name: "deleted", Type: proto.ColumnType_BOOL, Transform: transform.FromField("Deleted"), Description: ""},
			{Name: "downgradable", Type: proto.ColumnType_BOOL, Transform: transform.FromField("Downgradable"), Description: ""},
			{Name: "enabled", Type: proto.ColumnType_BOOL, Transform: transform.FromField("Enabled"), Description: ""},
			{Name: "hasUpdate", Type: proto.ColumnType_BOOL, Transform: transform.FromField("HasUpdate"), Description: ""},
			{Name: "id", Type: proto.ColumnType_STRING, Transform: transform.FromField("ShortName"), Description: ""},
			{Name: "longName", Type: proto.ColumnType_STRING, Transform: transform.FromField("LongName"), Description: ""},
			{Name: "pinned", Type: proto.ColumnType_BOOL, Transform: transform.FromField("Pinned"), Description: ""},
			{Name: "shortName", Type: proto.ColumnType_STRING, Transform: transform.FromField("ShortName"), Description: ""},
			{Name: "supportsDynamicLoad", Type: proto.ColumnType_STRING, Transform: transform.FromField("SupportsDynamicLoad"), Description: ""},
			{Name: "url", Type: proto.ColumnType_STRING, Transform: transform.FromField("URL"), Description: ""},
			{Name: "version", Type: proto.ColumnType_STRING, Transform: transform.FromField("Version"), Description: ""},
			{Name: "backupVersion", Type: proto.ColumnType_JSON, Transform: transform.FromField("BackupVersion"), Description: ""},
			{Name: "dependencies", Type: proto.ColumnType_JSON, Transform: transform.FromField("Dependencies"), Description: ""},
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
