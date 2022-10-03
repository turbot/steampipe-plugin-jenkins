package jenkins

import (
	"context"
	"strings"

	"github.com/turbot/steampipe-plugin-sdk/v4/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v4/plugin/transform"

	"github.com/turbot/steampipe-plugin-sdk/v4/plugin"
)

//// TABLE DEFINITION

func tableJenkinsJob() *plugin.Table {
	return &plugin.Table{
		Name:        "jenkins_job",
		Description: "A user-configured description of work which Jenkins should perform, such as building a piece of software, etc.",
		List: &plugin.ListConfig{
			Hydrate: listJenkinsJobs,
		},

		Columns: []*plugin.Column{
			{Name: "actions", Type: proto.ColumnType_JSON, Description: ""},
			{Name: "buildable", Type: proto.ColumnType_BOOL, Description: ""},
			{Name: "builds", Type: proto.ColumnType_JSON, Description: ""},
			{Name: "color", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "concurrent_build", Type: proto.ColumnType_BOOL, Description: ""},
			{Name: "description", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "display_name_or_null", Type: proto.ColumnType_BOOL, Description: ""},
			{Name: "display_name", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "downstream_projects", Type: proto.ColumnType_JSON, Description: ""},
			{Name: "first_build", Type: proto.ColumnType_JSON, Description: ""},
			{Name: "full_display_name", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "full_name", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "health_report", Type: proto.ColumnType_JSON, Description: ""},
			{Name: "in_queue", Type: proto.ColumnType_BOOL, Description: ""},
			{Name: "jobs", Type: proto.ColumnType_JSON, Description: ""},
			{Name: "keep_dependencies", Type: proto.ColumnType_BOOL, Description: ""},
			{Name: "last_build", Type: proto.ColumnType_JSON, Description: ""},
			{Name: "last_completed_build", Type: proto.ColumnType_JSON, Description: ""},
			{Name: "last_failed_build", Type: proto.ColumnType_JSON, Description: ""},
			{Name: "last_stable_build", Type: proto.ColumnType_JSON, Description: ""},
			{Name: "last_successful_build", Type: proto.ColumnType_JSON, Description: ""},
			{Name: "last_unstable_build", Type: proto.ColumnType_JSON, Description: ""},
			{Name: "last_unsuccessful_build", Type: proto.ColumnType_JSON, Description: ""},
			{Name: "name", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "next_build_number", Type: proto.ColumnType_INT, Description: ""},
			{Name: "primary_view", Type: proto.ColumnType_BOOL, Description: ""},
			{Name: "property", Type: proto.ColumnType_JSON, Description: ""},
			{Name: "queue_item", Type: proto.ColumnType_BOOL, Description: ""},
			{Name: "scm", Type: proto.ColumnType_JSON, Description: ""},
			{Name: "title", Type: proto.ColumnType_STRING, Transform: transform.FromField("DisplayName"), Description: titleDescription},
			{Name: "upstream_projects", Type: proto.ColumnType_JSON, Description: ""},
			{Name: "url", Type: proto.ColumnType_STRING, Transform: transform.FromField("URL"), Description: ""},
			{Name: "views", Type: proto.ColumnType_JSON, Description: ""},
		},
	}
}

//// LIST FUNCTION

func listJenkinsJobs(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)
	client, err := Connect(ctx, d)
	if err != nil {
		logger.Error("listJenkinsJobs", "connect_error", err)
		return nil, err
	}

	jobs, err := client.GetAllJobs(ctx)
	if err != nil {
		logger.Error("listJenkinsJobs", "list_jobs_error", err)
		if strings.Contains(err.Error(), "Not found") {
			return nil, nil
		}
		return nil, err
	}

	for _, job := range jobs {
		d.StreamListItem(ctx, job.Raw)

		// Context can be cancelled due to manual cancellation or the limit has been hit
		if d.QueryStatus.RowsRemaining(ctx) == 0 {
			return nil, nil
		}
	}

	return nil, err
}
