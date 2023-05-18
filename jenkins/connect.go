package jenkins

import (
	"context"
	"os"

	"github.com/bndr/gojenkins"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func Connect(ctx context.Context, d *plugin.QueryData) (*gojenkins.Jenkins, error) {
	conn, err := connectCached(ctx, d, nil)
	if err != nil {
		return nil, err
	}
	return conn.(*gojenkins.Jenkins), nil
}

var connectCached = plugin.HydrateFunc(connectUncached).Memoize()

func connectUncached(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (any, error) {
	jenkinsConfig := GetConfig(d.Connection)

	var url, userId, apiToken string

	if jenkinsConfig.Domain != nil {
		url = *jenkinsConfig.Domain
	} else {
		url = os.Getenv("JENKINS_URL")
	}

	if jenkinsConfig.UserId != nil {
		userId = *jenkinsConfig.UserId
	} else {
		userId = os.Getenv("JENKINS_USER_ID")
	}

	if jenkinsConfig.ApiToken != nil {
		apiToken = *jenkinsConfig.ApiToken
	} else {
		apiToken = os.Getenv("JENKINS_API_TOKEN")
	}

	// TODO handle the bellow
	// if url != "" && userId != "" && apiToken != "" {
	// 	return nil, "Missing credentials"
	// }

	return gojenkins.CreateJenkins(nil, url, userId, apiToken).Init(ctx)
}
