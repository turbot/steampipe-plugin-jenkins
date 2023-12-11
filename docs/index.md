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
engines: ["steampipe", "sqlite", "postgres", "export"]
---

# Jenkins + Steampipe

[Jenkins](https://www.jenkins.io/) is the leading open source automation server that provides hundreds of plugins to support building, deploying and automating any project.

[Steampipe](https://steampipe.io) is an open source CLI to instantly query cloud APIs using SQL.

For example:

```sql
select
  color,
  name,
  in_queue,
  last_completed_build  ->> 'URL' as last_completed_build
from
  jenkins_freestyle;
```

```
+----------+----------------------+----------+---------------------------------------------------------+
| color    | name                 | in_queue | last_completed_build                                    |
+----------+----------------------+----------+---------------------------------------------------------+
| blue     | stage-deploy         | false    | https://ci-cd.mycorp.com/job/stage-deploy/350/          |
| red      | build-and-unit-test  | true     | https://ci-cd.mycorp.com/job/build-and-unit-test/245/   |
| blue     | production-deploy    | true     | https://ci-cd.mycorp.com/job/build-and-unit-test/241/   |
+----------+----------------------+----------+---------------------------------------------------------+
```

## Documentation

- **[Table definitions & examples →](/plugins/turbot/jenkins/tables)**

## Quick start

### Install

Download and install the latest Jenkins plugin:

```bash
steampipe plugin install jenkins
```

### Credentials

| Item        | Description                                                                                                                                                                                                                                                                                 |
|-------------|---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| Credentials | Jenkins requires `server_url`, `username` and [password](https://www.jenkins.io/doc/book/pipeline/jenkinsfile/#handling-credentials) for all requests.                                                                                                                                                                                 |
| Permissions | Usernames and passwords have the same permissions as the user who creates them, and if the user permissions change, the connection's permissions also change.                                                                                                                                               |
| Radius      | Each connection represents a single Jenkins Installation.                                                                                                                                                                                                                                   |
| Resolution  | 1. Credentials explicitly set in a steampipe config file (`~/.steampipe/config/jenkins.spc`)<br />2. Credentials specified in environment variables, e.g., `JENKINS_URL`, `JENKINS_USERNAME` and `JENKINS_PASSWORD`. |

### Configuration

Installing the latest jenkins plugin will create a config file (~/.steampipe/config/jenkins.spc) with a single connection named jenkins:

```hcl
connection "jenkins" {
  plugin = "jenkins"

  # The Jenkins server URL is required for all requests. Required.
  # It should be fully qualified (e.g. # https://...) and point to the root of the Jenkins server location.
  # Can also be set via the JENKINS_URL environment variable.
  # server_url = "https://ci-cd.internal.my-company.com"

  # The Jenkins username for authentication is required for requests. Required.
  # Can also be set via the JENKINS_USERNAME environment variable.
  # username = "admin"

  # Either the password or the API token is required for requests. Required. 
  # Can also be set via the JENKINS_PASSWORD environment variable.
  # password = "aqt*abc8vcf9abc.ABC"

  # Further information: https://www.jenkins.io/doc/book/using/using-credentials/   
}
```

Alternatively, you can also use the standard Jenkins environment variables to obtain credentials **only if other arguments (`server_url`, `username` and `password`) are not specified** in the connection:

```sh
export JENKINS_URL=https://ci-cd.internal.my-company.com
export JENKINS_USERNAME=admin
export JENKINS_PASSWORD=aqt*abc8vcf9abc.ABC
```


