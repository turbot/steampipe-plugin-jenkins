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
		Description: "A machine which is part of the Jenkins environment and capable of executing Pipelines or jobs.",
		Get: &plugin.GetConfig{
			Hydrate:    getJenkinsNode,
			KeyColumns: plugin.SingleColumn("display_name"),
		},
		List: &plugin.ListConfig{
			Hydrate: listJenkinsNodes,
		},

		Columns: []*plugin.Column{
			{Name: "actions", Type: proto.ColumnType_JSON, Description: ""},
			{Name: "display_name", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "executors", Type: proto.ColumnType_JSON, Description: ""},
			{Name: "icon_class_name", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "icon", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "idle", Type: proto.ColumnType_BOOL, Description: ""},
			{Name: "jnlp_agent", Type: proto.ColumnType_BOOL, Description: ""},
			{Name: "launch_supported", Type: proto.ColumnType_BOOL, Description: ""},
			{Name: "load_statistics", Type: proto.ColumnType_JSON, Description: ""},
			{Name: "manual_launch_allowed", Type: proto.ColumnType_BOOL, Description: ""},
			{Name: "monitor_data", Type: proto.ColumnType_JSON, Description: ""},
			{Name: "num_executors", Type: proto.ColumnType_INT, Description: ""},
			{Name: "offline_cause_reason", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "offline_cause", Type: proto.ColumnType_JSON, Description: ""},
			{Name: "offline", Type: proto.ColumnType_BOOL, Description: ""},
			{Name: "one_off_executors", Type: proto.ColumnType_JSON, Description: ""},
			{Name: "temporarily_offline", Type: proto.ColumnType_BOOL, Description: ""},
			{Name: "title", Type: proto.ColumnType_STRING, Transform: transform.FromField("DisplayName"), Description: ""},
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

//// HYDRATE FUNCTION

func getJenkinsNode(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)
	logger.Trace("getJenkinsNode")
	nodeName := d.KeyColumnQuals["display_name"].GetStringValue()

	// Empty check for nodeName
	if nodeName == "" {
		return nil, nil
	}

	client, err := Connect(ctx, d)
	if err != nil {
		logger.Error("getJenkinsNode", "connect_error", err)
		return nil, err
	}

	node, err := client.GetNode(ctx, nodeName)
	if err != nil {
		logger.Error("getJenkinsNode", "get_node_error", err)
		return nil, err
	}

	return node.Raw, nil
}
