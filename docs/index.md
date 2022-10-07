---
organization: Turbot
category: ["saas"]
icon_url: "/images/plugins/turbot/jenkins.svg"
brand_color: "#41667E"
display_name: "Jenkins"
name: "jenkins"
description: "Steampipe plugin for querying resource jobs, builds, nodes, plugin and more from Jenkins."
og_description: "Query Jenkins with SQL! Open source CLI. No DB required."
og_image: "/images/plugins/turbot/jenkins-social-graphic.png"
---

# Jenkins + Steampipe

[Jenkins](https://www.jenkins.io/) is the leading open source automation server, Jenkins provides hundreds of plugins to support building, deploying and automating any project.

[Steampipe](https://steampipe.io) is an open source CLI to instantly query cloud APIs using SQL.

For example:

```sql
select
  color,
  name,
  in_queue,
  last_completed_build  ->> 'URL' as last_completed_build
from
  jenkins_job;
```

```
+----------+----------------------+----------+---------------------------------------------------------+
| color    | name                 | in_queue | last_completed_build                                    |
+----------+----------------------+----------+---------------------------------------------------------+
| blue     | stage-deploy         | false    | https://ci-cd.mycorp.com/job/stage-deploy/350/          |
| red      | build-and-unit-test  | true     | https://ci-cd.mycorp.com/job/build-and-unit-test/245/   |
| notbuilt | production-deploy    | true     |                                                         |
+----------+----------------------+----------+---------------------------------------------------------+
```

## Documentation

- **[Table definitions & examples →](/plugins/turbot/jenkins/tables)**

## Get started

### Install

Download and install the latest Jenkins plugin:

```bash
steampipe plugin install jenkins
```

### Credentials

| Item        | Description                                                                                                                                                                                                                                                                                 |
|-------------|---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| Credentials | Jenkins requires an [API token](https://www.jenkins.io/doc/book/using/using-credentials/) for all requests.                                                                                                                                                                                 |
| Permissions | API tokens have the same permissions as the user who creates them, and if the user permissions change, the API token permissions also change.                                                                                                                                               |
| Radius      | Each connection represents a single Jenkins Installation.                                                                                                                                                                                                                                   |
| Resolution  | 1. With configuration provided in connection in steampipe _**.spc**_ config file.<br />2. With jenkins environment variables.<br />3. An jenkins.yaml file in a .jenkins folder in the current user's home directory _**(~/.jenkins/jenkins.yaml or %userprofile\.jenkins\jenkins.yaml)**_. |

### Configuration

Installing the latest jenkins plugin will create a config file (~/.steampipe/config/jenkins.spc) with a single connection named jenkins:

```hcl
connection "jenkins" {
  plugin = "jenkins"

  # Get your API token from Jenkins https://www.jenkins.io/doc/book/using/using-credentials/

  # url = "http://<your_jenkins_domain>"
  # user_id = "admin"
  # api_token = "116af6f5cf749f31410983860c692850a2"
}
```

By default, all options are commented out in the default connection, thus Steampipe will resolve your credentials using the same order as mentioned in [Credentials](#credentials). This provides a quick way to get started with Steampipe, but you will probably want to customize your experience using configuration options for querying multiple organizations, configuring credentials from your jenkins configuration files, [environment variables](#credentials-from-environment-variables), etc.


## Get involved

- Open source: https://github.com/turbot/steampipe-plugin-jenkins
- Community: [Slack Channel](https://steampipe.io/community/join)

## Configuring Jenkins Credentials

### Credentials from Environment Variables

The Jenkins plugin will use the standard Jenkins environment variables to obtain credentials **only if other arguments (`domain`, `token`, `client_id`, `private_key`) are not specified** in the connection:

#### API Token

```sh
export JENKINS_URL=https://<your_jenkins_domain>
export JENKINS_USER_ID=admin
export JENKINS_API_TOKEN=116af6f5cf749f31410983860c692850a2
```
