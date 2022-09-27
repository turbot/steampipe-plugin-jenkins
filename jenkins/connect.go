package jenkins

import (
	"context"
	"os"

	"github.com/bndr/gojenkins"
	"github.com/turbot/steampipe-plugin-sdk/v4/plugin"
)

func Connect(ctx context.Context, d *plugin.QueryData) (*gojenkins.Jenkins, error) {
	// have we already created and cached the session?
	sessionCacheKey := "JenkinsSession"
	if cachedData, ok := d.ConnectionManager.Cache.Get(sessionCacheKey); ok {
		return cachedData.(*gojenkins.Jenkins), nil
	}

	jenkinsConfig := GetConfig(d.Connection)

	var domain, user, password string

	if jenkinsConfig.Domain != nil {
		domain = *jenkinsConfig.Domain
	} else {
		domain = os.Getenv("JENKINS_CLIENT_DOMAIN")
	}

	if jenkinsConfig.User != nil {
		user = *jenkinsConfig.User
	} else {
		user = os.Getenv("JENKINS_CLIENT_CLIENTID")
	}

	if jenkinsConfig.Password != nil {
		password = *jenkinsConfig.Password
	} else {
		password = os.Getenv("JENKINS_CLIENT_PASSWORD")
	}

	// TODO handle the bellow
	// if domain != "" && user != "" && password != "" {
	// 	return nil, "Missing credentials"
	// }

	client, err := gojenkins.CreateJenkins(nil, domain, user, password).Init(ctx)
	if err != nil {
		return nil, err
	}
	// Save session into cache
	d.ConnectionManager.Cache.Set(sessionCacheKey, client)
	return client, err

}
