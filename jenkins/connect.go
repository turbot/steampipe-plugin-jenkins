package jenkins

import (
	"context"
	"errors"
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

	var server_url, username, password string

	if jenkinsConfig.ServerURL != nil {
		server_url = *jenkinsConfig.ServerURL
	} else {
		server_url = os.Getenv("JENKINS_URL")
	}

	if jenkinsConfig.Username != nil {
		username = *jenkinsConfig.Username
	} else {
		username = os.Getenv("JENKINS_USERNAME")
	}

	if jenkinsConfig.Password != nil {
		password = *jenkinsConfig.Password
	} else {
		password = os.Getenv("JENKINS_PASSWORD")
	}

	// Error if the minimum config is not set
	if server_url == "" || username == "" || password == "" {
		return nil, errors.New("'server_url', 'username' and 'password' must be set in the connection configuration. Edit your connection configuration file and then restart Steampipe.")
	}	

	return gojenkins.CreateJenkins(nil, server_url, username, password).Init(ctx)
}
