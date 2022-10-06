package jenkins

import (
	"context"
	"strings"

	"github.com/turbot/steampipe-plugin-sdk/v4/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v4/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v4/plugin/transform"
)

//// TABLE DEFINITION

func tableJenkinsView() *plugin.Table {
	return &plugin.Table{
		Name:        "jenkins_view",
		Description: "A list view of jobs based on a filter or in a manual selection of jobs.",
		Get: &plugin.GetConfig{
			Hydrate:    getJenkinsView,
			KeyColumns: plugin.SingleColumn("name"),
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: isNotFoundError([]string{"No view found", "404"}),
			},
		},
		List: &plugin.ListConfig{
			Hydrate: listJenkinsViews,
		},

		Columns: []*plugin.Column{
			{Name: "description", Type: proto.ColumnType_STRING, Description: "A human-readable description of the view."},
			{Name: "jobs", Type: proto.ColumnType_JSON, Description: "List o jobs included in this view."},
			{Name: "name", Type: proto.ColumnType_STRING, Description: "Unique key for the view."},
			{Name: "title", Type: proto.ColumnType_STRING, Transform: transform.FromField("Name"), Description: titleDescription},
			{Name: "url", Type: proto.ColumnType_STRING, Transform: transform.FromField("URL"), Description: "Full URL to the view."},
		},
	}
}

//// LIST FUNCTION

func listJenkinsViews(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)
	client, err := Connect(ctx, d)
	if err != nil {
		logger.Error("jenkins_view.listJenkinsViews", "connect_error", err)
		return nil, err
	}

	views, err := client.GetAllViews(ctx)
	if err != nil {
		logger.Error("jenkins_view.listJenkinsViews", "list_views_error", err)
		if strings.Contains(err.Error(), "Not found") {
			return nil, nil
		}
		return nil, err
	}

	for _, view := range views {
		d.StreamListItem(ctx, view.Raw)

		// Context can be cancelled due to manual cancellation or the limit has been hit
		if d.QueryStatus.RowsRemaining(ctx) == 0 {
			return nil, nil
		}
	}

	return nil, err
}

//// HYDRATE FUNCTION

func getJenkinsView(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)
	logger.Trace("jenkins_view.getJenkinsView")
	viewName := d.KeyColumnQuals["name"].GetStringValue()

	// Empty check for viewName
	if viewName == "" {
		return nil, nil
	}

	client, err := Connect(ctx, d)
	if err != nil {
		logger.Error("jenkins_view.getJenkinsView", "connect_error", err)
		return nil, err
	}

	view, err := client.GetView(ctx, viewName)
	if err != nil {
		logger.Error("jenkins_view.getJenkinsView", "get_view_error", err)
		return nil, err
	}

	return view.Raw, nil
}
