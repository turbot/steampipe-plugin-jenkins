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

func tableJenkinsFolder() *plugin.Table {
	return &plugin.Table{
		Name:        "jenkins_folder",
		Description: "A user-configured description of work which Jenkins should perform, such as building a piece of software, etc.",
		Get: &plugin.GetConfig{
			Hydrate:    getJenkinsFolder,
			KeyColumns: plugin.SingleColumn("full_name"),
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: isNotFoundError([]string{"404"}),
			},
		},
		List: &plugin.ListConfig{
			Hydrate: listJenkinsFolders,
		},

		Columns: []*plugin.Column{
			{Name: "actions", Type: proto.ColumnType_JSON, Description: "Data about the folder trigger."},
			{Name: "description", Type: proto.ColumnType_STRING, Description: "An optional description that can be added to the folder."},
			{Name: "display_name", Type: proto.ColumnType_STRING, Description: "Human readable name of the folder."},
			{Name: "full_display_name", Type: proto.ColumnType_STRING, Description: "Human readable name of the folder, including parent folder."},
			{Name: "full_name", Type: proto.ColumnType_STRING, Description: "Unique key for the folder."},
			{Name: "name", Type: proto.ColumnType_STRING, Description: "Name of the folder, without its parent folder."},
			{Name: "primary_view", Type: proto.ColumnType_JSON, Description: "Main view of this folder."},
			{Name: "title", Type: proto.ColumnType_STRING, Transform: transform.FromField("DisplayName"), Description: titleDescription},
			{Name: "url", Type: proto.ColumnType_STRING, Transform: transform.FromField("URL"), Description: "Full URL to the folder."},
			{Name: "views", Type: proto.ColumnType_JSON, Description: "Views this folder is shows on."},
			{Name: "jobs", Type: proto.ColumnType_JSON, Description: "Child jobs."},
		},
	}
}

//// Recursively find sub folders

func handleFolders(folders []*gojenkins.Job, ctx context.Context, d *plugin.QueryData) {
	for _, folder := range folders {
		// Filter to Folder job type only
		if folder.Raw.Class != "com.cloudbees.hudson.plugins.folder.Folder" {
			continue
		}
		d.StreamListItem(ctx, folder.Raw)

		child_jobs, _ := folder.GetInnerJobs(ctx)
		handleFolders(child_jobs, ctx, d)
	}
}

//// LIST FUNCTION

func listJenkinsFolders(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)
	client, err := Connect(ctx, d)
	if err != nil {
		logger.Error("jenkins_folder.listJenkinsFolders", "connect_error", err)
		return nil, err
	}

	folders, err := client.GetAllJobs(ctx)
	if err != nil {
		logger.Error("jenkins_folder.listJenkinsFolders", "list_folders_error", err)
		if strings.Contains(err.Error(), "Not found") {
			return nil, nil
		}
		return nil, err
	}

	handleFolders(folders, ctx, d)

	// TODO not sure where to put this:
	// // Context can be cancelled due to manual cancellation or the limit has been hit
	// if d.QueryStatus.RowsRemaining(ctx) == 0 {
	// 	return nil, nil
	// }

	return nil, err
}

//// HYDRATE FUNCTION

func getJenkinsFolder(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)
	logger.Trace("jenkins_folder.getJenkinsFolder")
	folderFullName := d.KeyColumnQuals["full_name"].GetStringValue()

	// Empty check for folderFullName
	if folderFullName == "" {
		return nil, nil
	}

	client, err := Connect(ctx, d)
	if err != nil {
		logger.Error("jenkins_folder.getJenkinsFolder", "connect_error", err)
		return nil, err
	}

	folderFullNameList := strings.Split(folderFullName, "/")
	folderParentNames := folderFullNameList[0 : len(folderFullNameList)-1]
	folderName := folderFullNameList[len(folderFullNameList)-1]

	folder, err := client.GetJob(ctx, folderName, folderParentNames...)
	if err != nil {
		logger.Error("jenkins_folder.getJenkinsFolder", "get_folder_error", err)
		return nil, err
	}

	// Filter to Folder job type only
	if folder.Raw.Class != "com.cloudbees.hudson.plugins.folder.Folder" {
		return nil, nil
	}

	return folder.Raw, nil
}
