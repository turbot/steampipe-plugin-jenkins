package jenkins

import (
	"context"
	"strings"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableJenkinsNode() *plugin.Table {
	return &plugin.Table{
		Name:        "jenkins_node",
		Description: "A machine which is part of the Jenkins environment and capable of executing Pipelines or jobs.",
		Get: &plugin.GetConfig{
			Hydrate:    getJenkinsNode,
			KeyColumns: plugin.SingleColumn("display_name"),
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: isNotFoundError([]string{"No node found", "404"}),
			},
		},
		List: &plugin.ListConfig{
			Hydrate: listJenkinsNodes,
		},

		Columns: []*plugin.Column{
			{Name: "display_name", Type: proto.ColumnType_STRING, Description: "Unique key for the node."},
			{Name: "executors", Type: proto.ColumnType_JSON, Description: "List of executors, which are slots for execution of tasks in a node."},
			{Name: "icon_class_name", Type: proto.ColumnType_STRING, Description: "An HTML/CSS class indicating the status of the node, such as online/offline."},
			{Name: "icon", Type: proto.ColumnType_STRING, Description: "Image indicating the status of the node, such as online/offline."},
			{Name: "idle", Type: proto.ColumnType_BOOL, Description: "Boolean to indicate whether the node is not currently running any build."},
			{Name: "jnlp_agent", Type: proto.ColumnType_BOOL, Description: "Boolean to indicate whether the node uses a Java Network Launch Protocol agent to connect to master node."},
			{Name: "manual_launch_allowed", Type: proto.ColumnType_BOOL, Description: "Boolean to indicate whether a manual launch can be performed on the node."},
			{Name: "monitor_data", Type: proto.ColumnType_JSON, Description: "Node OS data used for motoring purpose such as, disk space, memory, etc."},
			{Name: "num_executors", Type: proto.ColumnType_INT, Description: "Number of executors in the node. The higher the more parallel task can be performed."},
			{Name: "offline_cause_reason", Type: proto.ColumnType_STRING, Description: "Detailed cause of why node is offline"},
			{Name: "offline_cause", Type: proto.ColumnType_JSON, Description: "Cause of why node is offline"},
			{Name: "offline", Type: proto.ColumnType_BOOL, Description: "Boolean to indicate whether the node is offline"},
			{Name: "temporarily_offline", Type: proto.ColumnType_BOOL, Description: "Boolean to indicate whether the node is marked to be offline."},
			{Name: "title", Type: proto.ColumnType_STRING, Transform: transform.FromField("DisplayName"), Description: titleDescription},
		},
	}
}

//// LIST FUNCTION

func listJenkinsNodes(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)
	client, err := Connect(ctx, d)
	if err != nil {
		logger.Error("jenkins_node.listJenkinsNodes", "connect_error", err)
		return nil, err
	}

	nodes, err := client.GetAllNodes(ctx)
	if err != nil {
		logger.Error("jenkins_node.listJenkinsNodes", "query_error", err)
		if strings.Contains(err.Error(), "Not found") {
			return nil, nil
		}
		return nil, err
	}

	for _, node := range nodes {
		d.StreamListItem(ctx, node.Raw)

		// Context can be cancelled due to manual cancellation or the limit has been hit
		if d.RowsRemaining(ctx) == 0 {
			return nil, nil
		}
	}

	return nil, err
}

//// HYDRATE FUNCTION

func getJenkinsNode(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)
	logger.Debug("jenkins_node.getJenkinsNode")
	nodeName := d.EqualsQualString("display_name")

	// Empty check for nodeName
	if nodeName == "" {
		return nil, nil
	}

	client, err := Connect(ctx, d)
	if err != nil {
		logger.Error("jenkins_node.getJenkinsNode", "connect_error", err)
		return nil, err
	}

	node, err := client.GetNode(ctx, nodeName)
	if err != nil {
		logger.Error("jenkins_node.getJenkinsNode", "query_error", err)
		return nil, err
	}

	return node.Raw, nil
}
