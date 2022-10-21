package jenkins

import (
	"context"
	"strings"

	"github.com/bndr/gojenkins"
	"github.com/turbot/steampipe-plugin-sdk/v4/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v4/plugin/transform"

	"github.com/turbot/steampipe-plugin-sdk/v4/plugin"
)

//// TABLE DEFINITION

func tableJenkinsPipeline() *plugin.Table {
	return &plugin.Table{
		Name:        "jenkins_pipeline",
		Description: "Orchestrates long-running activities that can span multiple build agents. Suitable for building pipelines (formerly known as workflows) and/or organizing complex activities that do not easily fit in free-style job type.",
		Get: &plugin.GetConfig{
			Hydrate:    getJenkinsPipeline,
			KeyColumns: plugin.SingleColumn("full_name"),
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: isNotFoundError([]string{"404"}),
			},
		},
		List: &plugin.ListConfig{
			ParentHydrate: listJenkinsFolders,
			Hydrate:       listJenkinsPipelines,
		},

		Columns: []*plugin.Column{
			{Name: "actions", Type: proto.ColumnType_JSON, Description: "Data about the pipeline trigger."},
			{Name: "buildable", Type: proto.ColumnType_BOOL, Description: "Boolean to indicate whether the pipeline is able to run a build."},
			{Name: "builds", Type: proto.ColumnType_JSON, Description: "List of builds of the pipeline."},
			{Name: "color", Type: proto.ColumnType_STRING, Description: "Color indicating the health of the pipeline based on the result of recent builds."},
			{Name: "concurrent_build", Type: proto.ColumnType_BOOL, Description: "Boolean to indicate whether the pipeline is able to run builds in parallel."},
			{Name: "description", Type: proto.ColumnType_STRING, Description: "An optional description that can be added to the pipeline."},
			{Name: "display_name", Type: proto.ColumnType_STRING, Description: "Human readable name of the pipeline."},
			{Name: "first_build", Type: proto.ColumnType_JSON, Description: "First build of the pipeline."},
			{Name: "full_display_name", Type: proto.ColumnType_STRING, Description: "Human readable name of the pipeline, including folder."},
			{Name: "full_name", Type: proto.ColumnType_STRING, Description: "Unique key for the pipeline."},
			{Name: "health_report", Type: proto.ColumnType_JSON, Description: "Health data about recent builds."},
			{Name: "in_queue", Type: proto.ColumnType_BOOL, Description: "Boolean to indicate whether the pipeline is waiting on queue to run a build."},
			{Name: "keep_dependencies", Type: proto.ColumnType_BOOL, Description: "Boolean to indicate whether the pipeline will keep dependencies used on last build."},
			{Name: "last_build", Type: proto.ColumnType_JSON, Description: "Last build of the pipeline."},
			{Name: "last_completed_build", Type: proto.ColumnType_JSON, Description: "Last completed build of the pipeline."},
			{Name: "last_failed_build", Type: proto.ColumnType_JSON, Description: "Last failed build of the pipeline."},
			{Name: "last_stable_build", Type: proto.ColumnType_JSON, Description: "Last stable build of the pipeline."},
			{Name: "last_successful_build", Type: proto.ColumnType_JSON, Description: "Last successful build of the pipeline."},
			{Name: "last_unstable_build", Type: proto.ColumnType_JSON, Description: "Last unstable build of the pipeline."},
			{Name: "last_unsuccessful_build", Type: proto.ColumnType_JSON, Description: "Last unsuccessful build of the pipeline."},
			{Name: "name", Type: proto.ColumnType_STRING, Description: "Name of the pipeline."},
			{Name: "next_build_number", Type: proto.ColumnType_INT, Description: "Number that will be assigned to build on next"},
			{Name: "property", Type: proto.ColumnType_JSON, Description: "Properties of the pipeline."},
			{Name: "title", Type: proto.ColumnType_STRING, Transform: transform.FromField("DisplayName"), Description: titleDescription},
			{Name: "url", Type: proto.ColumnType_STRING, Transform: transform.FromField("URL"), Description: "Full URL to the pipeline."},
		},
	}
}

//// LIST FUNCTION

func listJenkinsPipelines(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)

	folder := h.Item.(*gojenkins.Job)

	pipelines, err := folder.GetInnerJobs(ctx)
	if err != nil {
		logger.Error("jenkins_pipeline.listJenkinsPipelines", "list_pipelines_error", err)
		if strings.Contains(err.Error(), "Not found") {
			return nil, nil
		}
		return nil, err
	}

	for _, pipeline := range pipelines {
		// Filter to Pipeline job type only
		if pipeline.Raw.Class != "org.jenkinsci.plugins.workflow.job.WorkflowJob" {
			continue
		}
		d.StreamListItem(ctx, pipeline.Raw)

		// Context can be cancelled due to manual cancellation or the limit has been hit
		if d.QueryStatus.RowsRemaining(ctx) == 0 {
			return nil, nil
		}
	}

	return nil, err
}

//// HYDRATE FUNCTION

func getJenkinsPipeline(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)
	logger.Trace("jenkins_pipeline.getJenkinsPipeline")
	pipelineFullName := d.KeyColumnQuals["full_name"].GetStringValue()

	// Empty check for pipelineFullName
	if pipelineFullName == "" {
		return nil, nil
	}

	client, err := Connect(ctx, d)
	if err != nil {
		logger.Error("jenkins_pipeline.getJenkinsPipeline", "connect_error", err)
		return nil, err
	}

	pipelineFullNameList := strings.Split(pipelineFullName, "/")
	pipelineParentNames := pipelineFullNameList[0 : len(pipelineFullNameList)-1]
	pipelineName := pipelineFullNameList[len(pipelineFullNameList)-1]

	pipeline, err := client.GetJob(ctx, pipelineName, pipelineParentNames...)
	if err != nil {
		logger.Error("jenkins_pipeline.getJenkinsPipeline", "get_pipeline_error", err)
		return nil, err
	}

	// Filter to Pipeline job type only
	if pipeline.Raw.Class != "org.jenkinsci.plugins.workflow.job.WorkflowJob" {
		return nil, nil
	}

	return pipeline.Raw, nil
}
