package jenkins

import (
	"context"
	"os"

	"github.com/bndr/gojenkins"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func Connect(ctx context.Context, d *plugin.QueryData) (*gojenkins.Jenkins, error) {
	// have we already created and cached the session?
	sessionCacheKey := "JenkinsSession"
	if cachedData, ok := d.ConnectionManager.Cache.Get(sessionCacheKey); ok {
		return cachedData.(*gojenkins.Jenkins), nil
	}

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

	client, err := gojenkins.CreateJenkins(nil, url, userId, apiToken).Init(ctx)
	if err != nil {
		return nil, err
	}
	// Save session into cache
	d.ConnectionManager.Cache.Set(sessionCacheKey, client)
	return client, err

}
