package jenkins

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
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
			{Name: "FullName", Type: proto.ColumnType_STRING, Hydrate: getlistJenkinsUsers, Description: "User's full name."},
			{Name: "AbsoluteURL", Type: proto.ColumnType_STRING, Hydrate: getlistJenkinsUsers, Description: "User's absolute URL."},
		},
	}
}

//// LIST FUNCTION

func getlistJenkinsUsers(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	client, err := Connect(ctx, d)
	if err != nil {
		return nil, err
	}
	users, err := client.GetAllUsers(ctx)
	if err != nil {
		return nil, err
	}
	for _, user := range users.Raw.Users {
		d.StreamListItem(ctx, map[string]interface{}{
			"FullName":    user.User.FullName,
			"AbsoluteURL": user.User.AbsoluteURL,
		})
	}
	return nil, nil
}
