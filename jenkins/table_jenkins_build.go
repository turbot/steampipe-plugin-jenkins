package jenkins

import (
	"context"
	"strconv"
	"strings"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableJenkinsBuild() *plugin.Table {
	return &plugin.Table{
		Name:        "jenkins_build",
		Description: "Result of a single execution of a job",
		Get: &plugin.GetConfig{
			Hydrate:    getJenkinsBuild,
			KeyColumns: plugin.AllColumns([]string{"job_full_name", "number"}),
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: isNotFoundError([]string{"No build found", "404"}),
			},
		},
		List: &plugin.ListConfig{
			Hydrate:    listJenkinsBuilds,
			KeyColumns: plugin.SingleColumn("job_full_name"),
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: isNotFoundError([]string{"No build found", "404"}),
			},
		},

		Columns: []*plugin.Column{
			{Name: "actions", Type: proto.ColumnType_JSON, Hydrate: getJenkinsBuild, Description: "Data about the build trigger."},
			{Name: "artifacts", Type: proto.ColumnType_JSON, Hydrate: getJenkinsBuild, Description: "Files created as a result of the build execution."},
			{Name: "building", Type: proto.ColumnType_BOOL, Hydrate: getJenkinsBuild, Description: "Boolean to indicate whether the build is executing."},
			{Name: "built_on", Type: proto.ColumnType_STRING, Hydrate: getJenkinsBuild, Description: "Node where the build was executed."},
			{Name: "change_set", Type: proto.ColumnType_JSON, Hydrate: getJenkinsBuild, Description: "SCM changes between builds."},
			{Name: "culprits", Type: proto.ColumnType_JSON, Hydrate: getJenkinsBuild, Description: "People involved to the build."},
			{Name: "description", Type: proto.ColumnType_STRING, Hydrate: getJenkinsBuild, Description: "An optional description that can be added to the build."},
			{Name: "display_name", Type: proto.ColumnType_STRING, Hydrate: getJenkinsBuild, Description: "The name of the build, defaults to the build number."},
			{Name: "estimated_duration", Type: proto.ColumnType_DOUBLE, Hydrate: getJenkinsBuild, Description: "The expected amount of building time."},
			{Name: "executor", Type: proto.ColumnType_JSON, Hydrate: getJenkinsBuild, Description: "The executor where the build ran."},
			{Name: "finger_print", Type: proto.ColumnType_JSON, Hydrate: getJenkinsBuild, Description: "MD5 checksum fingerprint of the artifact file."},
			{Name: "full_display_name", Type: proto.ColumnType_STRING, Hydrate: getJenkinsBuild, Description: "Stands for the job name plus the display name."},
			{Name: "id", Type: proto.ColumnType_STRING, Hydrate: getJenkinsBuild, Transform: transform.FromField("ID"), Description: "Same as the build number, but as string."},
			{Name: "job_full_name", Type: proto.ColumnType_STRING, Description: "Full name of the job which defines the build. This column is required on any query because a build cannot exist without a job"},
			{Name: "keep_log", Type: proto.ColumnType_BOOL, Hydrate: getJenkinsBuild, Description: "Boolean to indicate whether the build kept the log."},
			{Name: "maven_artifacts", Type: proto.ColumnType_JSON, Hydrate: getJenkinsBuild, Description: "Maven artifacts generated during the build execution, if any."},
			{Name: "maven_version_used", Type: proto.ColumnType_STRING, Hydrate: getJenkinsBuild, Description: "Version of Maven used to execute the build."},
			{Name: "number", Type: proto.ColumnType_INT, Description: "Unique key for the build."},
			{Name: "queue_id", Type: proto.ColumnType_INT, Hydrate: getJenkinsBuild, Transform: transform.FromField("QueueID"), Description: "The queue ID assigned to the build. Each queue ID is unique."},
			{Name: "result", Type: proto.ColumnType_STRING, Hydrate: getJenkinsBuild, Description: "Result of the build execution."},
			{Name: "timestamp", Type: proto.ColumnType_INT, Hydrate: getJenkinsBuild, Description: "Time when the build started."},
			{Name: "title", Type: proto.ColumnType_STRING, Transform: transform.FromField("FullDisplayName"), Description: titleDescription},
			{Name: "url", Type: proto.ColumnType_STRING, Transform: transform.FromField("URL"), Description: "Full URL to the build."},
			{Name: "duration", Type: proto.ColumnType_DOUBLE, Hydrate: getJenkinsBuild, Description: "Actual amount of time took for the build execution."},
		},
	}
}

//// LIST FUNCTION

func listJenkinsBuilds(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)
	jobFullName := d.EqualsQualString("job_full_name")

	// Empty check for jobFullName
	if jobFullName == "" {
		return nil, nil
	}

	client, err := Connect(ctx, d)
	if err != nil {
		logger.Error("jenkins_build.listJenkinsBuilds", "connect_error", err)
		return nil, err
	}

	jobFullNameList := strings.Split(jobFullName, "/")
	jobParentNames := jobFullNameList[0 : len(jobFullNameList)-1]
	jobName := jobFullNameList[len(jobFullNameList)-1]

	job, err := client.GetJob(ctx, jobName, jobParentNames...)
	if err != nil {
		logger.Error("jenkins_build.listJenkinsBuilds", "get_job", "query_error", err)
		return nil, err
	}

	builds, err := job.GetAllBuildIds(ctx)
	if err != nil {
		logger.Error("jenkins_build.listJenkinsBuilds", "list_builds", "query_error", err)
		if strings.Contains(err.Error(), "Not found") {
			return nil, nil
		}
		return nil, err
	}

	for _, build := range builds {
		buildMap := map[string]interface{}{
			"Number":      build.Number,
			"URL":         build.URL,
			"JobFullName": jobFullName,
		}
		d.StreamListItem(ctx, buildMap)

		// Context can be cancelled due to manual cancellation or the limit has been hit
		if d.RowsRemaining(ctx) == 0 {
			return nil, nil
		}
	}

	return nil, err
}

//// HYDRATE FUNCTION

func getJenkinsBuild(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)
	logger.Trace("jenkins_build.getJenkinsBuild")
	jobFullName := d.EqualsQualString("job_full_name")

	// Empty check for jobFullName
	if jobFullName == "" {
		return nil, nil
	}

	buildNumberQual := d.EqualsQualString("number")
	buildNumber, err := strconv.ParseInt(buildNumberQual, 10, 64)
	if err != nil {
		buildNumber = 0
	}

	// Empty check for buildNumber
	if buildNumber == 0 {
		buildNumber = h.Item.(map[string]interface{})["Number"].(int64)
		if buildNumber == 0 {
			return nil, nil
		}
	}

	client, err := Connect(ctx, d)
	if err != nil {
		logger.Error("jenkins_build.getJenkinsBuild", "connect_error", err)
		return nil, err
	}

	jobFullNameList := strings.Split(jobFullName, "/")
	jobParentNames := jobFullNameList[0 : len(jobFullNameList)-1]
	jobName := jobFullNameList[len(jobFullNameList)-1]

	job, err := client.GetJob(ctx, jobName, jobParentNames...)
	if err != nil {
		logger.Error("jenkins_build.listJenkinsBuilds", "get_job", "query_error", err)
		return nil, err
	}

	build, err := job.GetBuild(ctx, buildNumber)
	if err != nil {
		logger.Error("jenkins_build.getJenkinsBuild", "get_build", "query_error", err)
		return nil, err
	}

	buildMap := map[string]interface{}{
		"JobName":           build.Job.Raw.Name,
		"Actions":           build.Raw.Actions,
		"Artifacts":         build.Raw.Artifacts,
		"Building":          build.Raw.Building,
		"BuiltOn":           build.Raw.BuiltOn,
		"ChangeSet":         build.Raw.ChangeSet,
		"ChangeSets":        build.Raw.ChangeSets,
		"Culprits":          build.Raw.Culprits,
		"Description":       build.Raw.Description,
		"Duration":          build.Raw.Duration,
		"EstimatedDuration": build.Raw.EstimatedDuration,
		"Executor":          build.Raw.Executor,
		"DisplayName":       build.Raw.DisplayName,
		"FullDisplayName":   build.Raw.FullDisplayName,
		"ID":                build.Raw.ID,
		"KeepLog":           build.Raw.KeepLog,
		"Number":            build.Raw.Number,
		"QueueID":           build.Raw.QueueID,
		"Result":            build.Raw.Result,
		"Timestamp":         build.Raw.Timestamp,
		"URL":               build.Raw.URL,
		"MavenArtifacts":    build.Raw.MavenArtifacts,
		"MavenVersionUsed":  build.Raw.MavenVersionUsed,
		"FingerPrint":       build.Raw.FingerPrint,
		"Runs":              build.Raw.Runs,
	}

	return buildMap, nil
}
