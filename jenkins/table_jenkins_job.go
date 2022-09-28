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
		Description: "A job is a runnable entity on Jenkins.",
		List: &plugin.ListConfig{
			Hydrate: listJenkinsJobs,
			KeyColumns: plugin.KeyColumnSlice{
				{Name: "id", Require: plugin.Optional},
			},
		},

		Columns: []*plugin.Column{
			{Name: "actions", Type: proto.ColumnType_JSON, Transform: transform.FromField("Actions"), Description: ""},
			{Name: "buildable", Type: proto.ColumnType_BOOL, Transform: transform.FromField("Buildable"), Description: ""},
			{Name: "builds", Type: proto.ColumnType_JSON, Transform: transform.FromField("Builds"), Description: ""},
			{Name: "color", Type: proto.ColumnType_STRING, Transform: transform.FromField("Color"), Description: ""},
			{Name: "concurrent_build", Type: proto.ColumnType_BOOL, Transform: transform.FromField("ConcurrentBuild"), Description: ""},
			{Name: "description", Type: proto.ColumnType_STRING, Transform: transform.FromField("Description"), Description: ""},
			{Name: "display_name_or_null", Type: proto.ColumnType_BOOL, Transform: transform.FromField("DisplayNameOrNull"), Description: ""},
			{Name: "display_name", Type: proto.ColumnType_STRING, Transform: transform.FromField("DisplayName"), Description: ""},
			{Name: "downstream_projects", Type: proto.ColumnType_JSON, Transform: transform.FromField("DownstreamProjects"), Description: ""},
			{Name: "first_build", Type: proto.ColumnType_JSON, Transform: transform.FromField("FirstBuild"), Description: ""},
			{Name: "full_display_name", Type: proto.ColumnType_STRING, Transform: transform.FromField("FullDisplayName"), Description: ""},
			{Name: "full_name", Type: proto.ColumnType_STRING, Transform: transform.FromField("FullName"), Description: ""},
			{Name: "health_report", Type: proto.ColumnType_JSON, Transform: transform.FromField("HealthReport"), Description: ""},
			{Name: "id", Type: proto.ColumnType_STRING, Transform: transform.FromField("Name"), Description: "Unique key for job."},
			{Name: "in_queue", Type: proto.ColumnType_BOOL, Transform: transform.FromField("InQueue"), Description: ""},
			{Name: "jobs", Type: proto.ColumnType_JSON, Transform: transform.FromField("Jobs"), Description: ""},
			{Name: "keep_dependencies", Type: proto.ColumnType_BOOL, Transform: transform.FromField("KeepDependencies"), Description: ""},
			{Name: "last_build", Type: proto.ColumnType_JSON, Transform: transform.FromField("LastBuild"), Description: ""},
			{Name: "last_completed_build", Type: proto.ColumnType_JSON, Transform: transform.FromField("LastCompletedBuild"), Description: ""},
			{Name: "last_failed_build", Type: proto.ColumnType_JSON, Transform: transform.FromField("LastFailedBuild"), Description: ""},
			{Name: "last_stable_build", Type: proto.ColumnType_JSON, Transform: transform.FromField("LastStableBuild"), Description: ""},
			{Name: "last_successful_build", Type: proto.ColumnType_JSON, Transform: transform.FromField("LastSuccessfulBuild"), Description: ""},
			{Name: "last_unstable_build", Type: proto.ColumnType_JSON, Transform: transform.FromField("LastUnstableBuild"), Description: ""},
			{Name: "last_unsuccessful_build", Type: proto.ColumnType_JSON, Transform: transform.FromField("LastUnsuccessfulBuild"), Description: ""},
			{Name: "name", Type: proto.ColumnType_STRING, Transform: transform.FromField("Name"), Description: titleDescription},
			{Name: "next_build_number", Type: proto.ColumnType_INT, Transform: transform.FromField("NextBuildNumber"), Description: ""},
			{Name: "primary_view", Type: proto.ColumnType_BOOL, Transform: transform.FromField("PrimaryView"), Description: ""},
			{Name: "property", Type: proto.ColumnType_JSON, Transform: transform.FromField("Property"), Description: ""},
			{Name: "queue_item", Type: proto.ColumnType_BOOL, Transform: transform.FromField("QueueItem"), Description: ""},
			{Name: "scm", Type: proto.ColumnType_JSON, Transform: transform.FromField("Scm"), Description: ""},
			{Name: "upstream_projects", Type: proto.ColumnType_JSON, Transform: transform.FromField("UpstreamProjects"), Description: ""},
			{Name: "url", Type: proto.ColumnType_STRING, Transform: transform.FromField("URL"), Description: ""},
			{Name: "views", Type: proto.ColumnType_JSON, Transform: transform.FromField("Views"), Description: ""},
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
