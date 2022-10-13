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
		Get: &plugin.GetConfig{
			Hydrate:    getJenkinsJob,
			KeyColumns: plugin.SingleColumn("name"),
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: isNotFoundError([]string{"404"}),
			},
		},
		List: &plugin.ListConfig{
			Hydrate: listJenkinsJobs,
		},

		Columns: []*plugin.Column{
			{Name: "actions", Type: proto.ColumnType_JSON, Description: "Data about the job trigger."},
			{Name: "buildable", Type: proto.ColumnType_BOOL, Description: "Boolean to indicate whether the job is able to run a build."},
			{Name: "builds", Type: proto.ColumnType_JSON, Description: "List of builds of the job."},
			{Name: "color", Type: proto.ColumnType_STRING, Description: "Color indicating the health of the job based on the result of recent builds."},
			{Name: "concurrent_build", Type: proto.ColumnType_BOOL, Description: "Boolean to indicate whether the job is able to run builds in parallel."},
			{Name: "description", Type: proto.ColumnType_STRING, Description: "An optional description that can be added to the job."},
			{Name: "display_name", Type: proto.ColumnType_STRING, Description: "Human readable name of the job."},
			{Name: "downstream_projects", Type: proto.ColumnType_JSON, Description: "Jobs called after build execution."},
			{Name: "first_build", Type: proto.ColumnType_JSON, Description: "First build of the job."},
			{Name: "full_display_name", Type: proto.ColumnType_STRING, Description: "Human readable name of the job, including folder."},
			{Name: "full_name", Type: proto.ColumnType_STRING, Description: "Job name including folder."},
			{Name: "health_report", Type: proto.ColumnType_JSON, Description: "Health data about recent builds."},
			{Name: "in_queue", Type: proto.ColumnType_BOOL, Description: "Boolean to indicate whether the job is waiting on queue to run a build."},
			{Name: "jobs", Type: proto.ColumnType_JSON, Description: "Child jobs."},
			{Name: "keep_dependencies", Type: proto.ColumnType_BOOL, Description: "Boolean to indicate whether the job will keep dependencies used on last build."},
			{Name: "last_build", Type: proto.ColumnType_JSON, Description: "Last build of the job."},
			{Name: "last_completed_build", Type: proto.ColumnType_JSON, Description: "Last completed build of the job."},
			{Name: "last_failed_build", Type: proto.ColumnType_JSON, Description: "Last failed build of the job."},
			{Name: "last_stable_build", Type: proto.ColumnType_JSON, Description: "Last stable build of the job."},
			{Name: "last_successful_build", Type: proto.ColumnType_JSON, Description: "Last successful build of the job."},
			{Name: "last_unstable_build", Type: proto.ColumnType_JSON, Description: "Last unstable build of the job."},
			{Name: "last_unsuccessful_build", Type: proto.ColumnType_JSON, Description: "Last unsuccessful build of the job."},
			{Name: "name", Type: proto.ColumnType_STRING, Description: "Unique key for the job."},
			{Name: "next_build_number", Type: proto.ColumnType_INT, Description: "Number that will be assigned to build on next"},
			{Name: "primary_view", Type: proto.ColumnType_JSON, Description: "Main view of this job."},
			{Name: "property", Type: proto.ColumnType_JSON, Description: "Properties of the job."},
			{Name: "scm", Type: proto.ColumnType_JSON, Description: "Source code management set on this job."},
			{Name: "title", Type: proto.ColumnType_STRING, Transform: transform.FromField("DisplayName"), Description: titleDescription},
			{Name: "upstream_projects", Type: proto.ColumnType_JSON, Description: "Jobs that calls this job after their build finishes."},
			{Name: "url", Type: proto.ColumnType_STRING, Transform: transform.FromField("URL"), Description: "Full URL to the job."},
			{Name: "views", Type: proto.ColumnType_JSON, Description: "Views this job is shows on."},
		},
	}
}

//// LIST FUNCTION

func listJenkinsJobs(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)
	client, err := Connect(ctx, d)
	if err != nil {
		logger.Error("jenkins_job.listJenkinsJobs", "connect_error", err)
		return nil, err
	}

	jobs, err := client.GetAllJobs(ctx)
	if err != nil {
		logger.Error("jenkins_job.listJenkinsJobs", "list_jobs_error", err)
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

//// HYDRATE FUNCTION

func getJenkinsJob(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)
	logger.Trace("jenkins_job.getJenkinsJob")
	jobName := d.KeyColumnQuals["name"].GetStringValue()

	// Empty check for jobName
	if jobName == "" {
		return nil, nil
	}

	client, err := Connect(ctx, d)
	if err != nil {
		logger.Error("jenkins_job.getJenkinsJob", "connect_error", err)
		return nil, err
	}

	job, err := client.GetJob(ctx, jobName)
	if err != nil {
		logger.Error("jenkins_job.getJenkinsJob", "get_job_error", err)
		return nil, err
	}

	return job.Raw, nil
}
