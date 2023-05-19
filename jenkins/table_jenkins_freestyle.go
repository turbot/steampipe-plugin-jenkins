package jenkins

import (
	"context"
	"strings"

	"github.com/bndr/gojenkins"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"

	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

//// TABLE DEFINITION

func tableJenkinsFreestyle() *plugin.Table {
	return &plugin.Table{
		Name:        "jenkins_freestyle",
		Description: "A user-configured description of work which Jenkins should perform, such as building a piece of software, etc.",
		Get: &plugin.GetConfig{
			Hydrate:    getJenkinsFreestyle,
			KeyColumns: plugin.SingleColumn("full_name"),
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: isNotFoundError([]string{"404"}),
			},
		},
		List: &plugin.ListConfig{
			ParentHydrate: listJenkinsFolders,
			Hydrate:       listJenkinsFreestyles,
		},

		Columns: []*plugin.Column{
			{Name: "actions", Type: proto.ColumnType_JSON, Description: "Data about the freestyle trigger."},
			{Name: "buildable", Type: proto.ColumnType_BOOL, Description: "Boolean to indicate whether the freestyle is able to run a build."},
			{Name: "builds", Type: proto.ColumnType_JSON, Description: "List of builds of the freestyle."},
			{Name: "color", Type: proto.ColumnType_STRING, Description: "Color indicating the health of the freestyle based on the result of recent builds."},
			{Name: "concurrent_build", Type: proto.ColumnType_BOOL, Description: "Boolean to indicate whether the freestyle is able to run builds in parallel."},
			{Name: "description", Type: proto.ColumnType_STRING, Description: "An optional description that can be added to the freestyle."},
			{Name: "display_name", Type: proto.ColumnType_STRING, Description: "Human readable name of the freestyle."},
			{Name: "downstream_projects", Type: proto.ColumnType_JSON, Description: "Jobs called after build execution."},
			{Name: "first_build", Type: proto.ColumnType_JSON, Description: "First build of the freestyle."},
			{Name: "full_display_name", Type: proto.ColumnType_STRING, Description: "Human readable name of the freestyle, including folder."},
			{Name: "full_name", Type: proto.ColumnType_STRING, Description: "Unique key for the freestyle."},
			{Name: "health_report", Type: proto.ColumnType_JSON, Description: "Health data about recent builds."},
			{Name: "in_queue", Type: proto.ColumnType_BOOL, Description: "Boolean to indicate whether the freestyle is waiting on queue to run a build."},
			{Name: "keep_dependencies", Type: proto.ColumnType_BOOL, Description: "Boolean to indicate whether the freestyle will keep dependencies used on last build."},
			{Name: "last_build", Type: proto.ColumnType_JSON, Description: "Last build of the freestyle."},
			{Name: "last_completed_build", Type: proto.ColumnType_JSON, Description: "Last completed build of the freestyle."},
			{Name: "last_failed_build", Type: proto.ColumnType_JSON, Description: "Last failed build of the freestyle."},
			{Name: "last_stable_build", Type: proto.ColumnType_JSON, Description: "Last stable build of the freestyle."},
			{Name: "last_successful_build", Type: proto.ColumnType_JSON, Description: "Last successful build of the freestyle."},
			{Name: "last_unstable_build", Type: proto.ColumnType_JSON, Description: "Last unstable build of the freestyle."},
			{Name: "last_unsuccessful_build", Type: proto.ColumnType_JSON, Description: "Last unsuccessful build of the freestyle."},
			{Name: "name", Type: proto.ColumnType_STRING, Description: "Name of the freestyle."},
			{Name: "next_build_number", Type: proto.ColumnType_INT, Description: "Number that will be assigned to build on next"},
			{Name: "property", Type: proto.ColumnType_JSON, Description: "Properties of the freestyle."},
			{Name: "scm", Type: proto.ColumnType_JSON, Description: "Source code management set on this freestyle."},
			{Name: "title", Type: proto.ColumnType_STRING, Transform: transform.FromField("DisplayName"), Description: titleDescription},
			{Name: "upstream_projects", Type: proto.ColumnType_JSON, Description: "Jobs that calls this freestyle after their build finishes."},
			{Name: "url", Type: proto.ColumnType_STRING, Transform: transform.FromField("URL"), Description: "Full URL to the freestyle."},
		},
	}
}

//// LIST FUNCTION

func listJenkinsFreestyles(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)

	folder := h.Item.(*gojenkins.Job)

	freestyles, err := folder.GetInnerJobs(ctx)
	if err != nil {
		logger.Error("jenkins_freestyle.listJenkinsFreestyles", "query_error", err)
		if strings.Contains(err.Error(), "Not found") {
			return nil, nil
		}
		return nil, err
	}

	for _, freestyle := range freestyles {
		// Filter to Freestyle job type only
		if freestyle.Raw.Class != "hudson.model.FreeStyleProject" {
			continue
		}
		d.StreamListItem(ctx, freestyle.Raw)

		// Context can be cancelled due to manual cancellation or the limit has been hit
		if d.RowsRemaining(ctx) == 0 {
			return nil, nil
		}
	}

	return nil, err
}

//// HYDRATE FUNCTION

func getJenkinsFreestyle(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)
	logger.Trace("jenkins_freestyle.getJenkinsFreestyle")
	freestyleFullName := d.EqualsQualString("full_name")

	// Empty check for freestyleFullName
	if freestyleFullName == "" {
		return nil, nil
	}

	client, err := Connect(ctx, d)
	if err != nil {
		logger.Error("jenkins_freestyle.getJenkinsFreestyle", "connect_error", err)
		return nil, err
	}

	freestyleFullNameList := strings.Split(freestyleFullName, "/")
	freestyleParentNames := freestyleFullNameList[0 : len(freestyleFullNameList)-1]
	freestyleName := freestyleFullNameList[len(freestyleFullNameList)-1]

	freestyle, err := client.GetJob(ctx, freestyleName, freestyleParentNames...)
	if err != nil {
		logger.Error("jenkins_freestyle.getJenkinsFreestyle", "query_error", err)
		return nil, err
	}

	// Filter to Freestyle job type only
	if freestyle.Raw.Class != "hudson.model.FreeStyleProject" {
		return nil, nil
	}

	return freestyle.Raw, nil
}
