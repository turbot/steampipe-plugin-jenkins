package jenkins

import (
	"context"
	"strings"

	"github.com/turbot/steampipe-plugin-sdk/v4/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v4/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v4/plugin/transform"
)

//// TABLE DEFINITION

func tableJenkinsNode() *plugin.Table {
	return &plugin.Table{
		Name:        "jenkins_node",
		Description: "A node is a runnable entity on Jenkins.",
		List: &plugin.ListConfig{
			Hydrate: listJenkinsNodes,
			KeyColumns: plugin.KeyColumnSlice{
				{Name: "id", Require: plugin.Optional},
			},
		},

		Columns: []*plugin.Column{
			{Name: "actions", Type: proto.ColumnType_JSON, Transform: transform.FromField("Actions"), Description: ""},
			{Name: "display_name", Type: proto.ColumnType_STRING, Transform: transform.FromField("DisplayName"), Description: ""},
			{Name: "executors", Type: proto.ColumnType_JSON, Transform: transform.FromField("Executors"), Description: ""},
			{Name: "icon", Type: proto.ColumnType_STRING, Transform: transform.FromField("Icon"), Description: ""},
			{Name: "icon_class_name", Type: proto.ColumnType_STRING, Transform: transform.FromField("IconClassName"), Description: ""},
			{Name: "id", Type: proto.ColumnType_STRING, Transform: transform.FromField("DisplayName"), Description: "Unique key for node."},
			{Name: "idle", Type: proto.ColumnType_BOOL, Transform: transform.FromField("Idle"), Description: ""},
			{Name: "jnlp_agent", Type: proto.ColumnType_BOOL, Transform: transform.FromField("JnlpAgent"), Description: ""},
			{Name: "launch_supported", Type: proto.ColumnType_BOOL, Transform: transform.FromField("LaunchSupported"), Description: ""},
			{Name: "load_statistics", Type: proto.ColumnType_JSON, Transform: transform.FromField("LoadStatistics"), Description: ""},
			{Name: "manual_launch_allowed", Type: proto.ColumnType_BOOL, Transform: transform.FromField("ManualLaunchAllowed"), Description: ""},
			{Name: "monitor_data", Type: proto.ColumnType_JSON, Transform: transform.FromField("MonitorData"), Description: ""},
			{Name: "num_executors", Type: proto.ColumnType_INT, Transform: transform.FromField("NumExecutors"), Description: ""},
			{Name: "offline", Type: proto.ColumnType_BOOL, Transform: transform.FromField("Offline"), Description: ""},
			{Name: "offline_cause", Type: proto.ColumnType_JSON, Transform: transform.FromField("OfflineCause"), Description: ""},
			{Name: "offline_cause_reason", Type: proto.ColumnType_STRING, Transform: transform.FromField("OfflineCauseReason"), Description: ""},
			{Name: "one_off_executors", Type: proto.ColumnType_JSON, Transform: transform.FromField("OneOffExecutors"), Description: ""},
			{Name: "temporarily_offline", Type: proto.ColumnType_BOOL, Transform: transform.FromField("TemporarilyOffline"), Description: ""},
		},
	}
}

//// LIST FUNCTION

func listJenkinsNodes(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)
	client, err := Connect(ctx, d)
	if err != nil {
		logger.Error("listJenkinsNodes", "connect_error", err)
		return nil, err
	}

	nodes, err := client.GetAllNodes(ctx)
	if err != nil {
		logger.Error("listJenkinsNodes", "list_nodes_error", err)
		if strings.Contains(err.Error(), "Not found") {
			return nil, nil
		}
		return nil, err
	}

	for _, node := range nodes {
		d.StreamListItem(ctx, node.Raw)

		// Context can be cancelled due to manual cancellation or the limit has been hit
		if d.QueryStatus.RowsRemaining(ctx) == 0 {
			return nil, nil
		}
	}

	return nil, err
}
