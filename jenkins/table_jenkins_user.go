package jenkins

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableJenkinsUser() *plugin.Table {
	return &plugin.Table{
		Name:        "jenkins_user",
		Description: "An extension to Jenkins functionality provided separately from Jenkins Core.",
		List: &plugin.ListConfig{
			Hydrate: listJenkinsUsers,
		},

		Columns: []*plugin.Column{
			{Name: "full_name", Type: proto.ColumnType_STRING, Hydrate: listJenkinsUsers, Description: "User's full name."},
			{Name: "absolute_url", Type: proto.ColumnType_STRING, Hydrate: listJenkinsUsers, Transform: transform.FromField("AbsoluteURL"), Description: "User's absolute URL."},
		},
	}
}

//// LIST FUNCTION

func listJenkinsUsers(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
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
