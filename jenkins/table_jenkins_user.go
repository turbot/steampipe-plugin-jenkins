package jenkins

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"

	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

//// TABLE DEFINITION

func tableJenkinsUser() *plugin.Table {
	return &plugin.Table{
		Name:        "jenkins_user",
		Description: "An extension to Jenkins functionality provided separately from Jenkins Core.",
		List: &plugin.ListConfig{
			Hydrate: getlistJenkinsUsers,
		},

		Columns: []*plugin.Column{
			{Name: "FullName", Type: proto.ColumnType_STRING, Hydrate: getlistJenkinsUsers, Description: "String to indicate whether the plugin is active."},
			{Name: "AbsoluteURL", Type: proto.ColumnType_STRING, Hydrate: getlistJenkinsUsers, Description: "String to indicate whether the plugin is active."},
		},
	}
}

//// LIST FUNCTION

// this function takes a lot of time to get the
func getlistJenkinsUsers(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {

	//	logger := plugin.Logger(ctx)

	client, err := Connect(ctx, d)
	if err != nil {
		return nil, err
	}
	users, err := client.GetAllUsers(ctx)
	for _, user := range users.Raw.Users {
		d.StreamListItem(ctx, user.User)

	}
	return nil, nil

}
