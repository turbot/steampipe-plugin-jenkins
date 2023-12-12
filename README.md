![image](https://hub.steampipe.io/images/plugins/turbot/jenkins-social-graphic.png)

# Jenkins Plugin for Steampipe

Use SQL to query jobs, builds, nodes, plugin and more from Jenkins.

- **[Get started →](https://hub.steampipe.io/plugins/turbot/jenkins)**
- Documentation: [Table definitions & examples](https://hub.steampipe.io/plugins/turbot/jenkins/tables)
- Community: [Join #steampipe on Slack →](https://turbot.com/community/join)
- Get involved: [Issues](https://github.com/turbot/steampipe-plugin-jenkins/issues)

## Quick start

### Install

Download and install the latest Steampipe plugin:

```bash
steampipe plugin install jenkins
```

Configure your [credentials](https://hub.steampipe.io/plugins/turbot/jenkins#credentials) and [config file](https://hub.steampipe.io/plugins/turbot/jenkins#configuration).

Configure your account details in `~/.steampipe/config/jenkins.spc`:

```hcl
connection "jenkins" {
  plugin = "jenkins"

  # Your Jenkins instance URL
  server_url = "https://ci-cd.internal.my-company.com"

  # Authentication information
  username = "admin"
  password = "aqt*abc8vcf9abc.ABC"
}
```

Or through environment variables:

```sh
export JENKINS_URL=https://ci-cd.internal.my-company.com
export JENKINS_USERNAME=admin
export JENKINS_PASSWORD=aqt*abc8vcf9abc.ABC
```

Run steampipe:

```shell
steampipe query
```

List FreeStyle jobs on your Jenkins instance:

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

## Engines

This plugin is available for the following engines:

| Engine        | Description
|---------------|------------------------------------------
| [Steampipe](https://steampipe.io/docs) | The Steampipe CLI exposes APIs and services as a high-performance relational database, giving you the ability to write SQL-based queries to explore dynamic data. Mods extend Steampipe's capabilities with dashboards, reports, and controls built with simple HCL. The Steampipe CLI is a turnkey solution that includes its own Postgres database, plugin management, and mod support.
| [Postgres FDW](https://steampipe.io/docs/steampipe_postgres/index) | Steampipe Postgres FDWs are native Postgres Foreign Data Wrappers that translate APIs to foreign tables. Unlike Steampipe CLI, which ships with its own Postgres server instance, the Steampipe Postgres FDWs can be installed in any supported Postgres database version.
| [SQLite Extension](https://steampipe.io/docs//steampipe_sqlite/index) | Steampipe SQLite Extensions provide SQLite virtual tables that translate your queries into API calls, transparently fetching information from your API or service as you request it.
| [Export](https://steampipe.io/docs/steampipe_export/index) | Steampipe Plugin Exporters provide a flexible mechanism for exporting information from cloud services and APIs. Each exporter is a stand-alone binary that allows you to extract data using Steampipe plugins without a database.
| [Turbot Pipes](https://turbot.com/pipes/docs) | Turbot Pipes is the only intelligence, automation & security platform built specifically for DevOps. Pipes provide hosted Steampipe database instances, shared dashboards, snapshots, and more.

## Developing

Prerequisites:

- [Steampipe](https://steampipe.io/downloads)
- [Golang](https://golang.org/doc/install)

Clone:

```sh
git clone https://github.com/turbot/steampipe-plugin-jenkins.git
cd steampipe-plugin-jenkins
```

Build, which automatically installs the new version to your `~/.steampipe/plugins` directory:

```
make
```

Configure the plugin:

```
cp config/* ~/.steampipe/config
vi ~/.steampipe/config/jenkins.spc
```

Try it!

```
steampipe query
> .inspect jenkins
```

Further reading:

- [Writing plugins](https://steampipe.io/docs/develop/writing-plugins)
- [Writing your first table](https://steampipe.io/docs/develop/writing-your-first-table)

## Open Source & Contributing

This repository is published under the [Apache 2.0](https://www.apache.org/licenses/LICENSE-2.0) (source code) and [CC BY-NC-ND](https://creativecommons.org/licenses/by-nc-nd/2.0/) (docs) licenses. Please see our [code of conduct](https://github.com/turbot/.github/blob/main/CODE_OF_CONDUCT.md). We look forward to collaborating with you!

[Steampipe](https://steampipe.io) is a product produced from this open source software, exclusively by [Turbot HQ, Inc](https://turbot.com). It is distributed under our commercial terms. Others are allowed to make their own distribution of the software, but cannot use any of the Turbot trademarks, cloud services, etc. You can learn more in our [Open Source FAQ](https://turbot.com/open-source).

## Get Involved

**[Join #steampipe on Slack →](https://turbot.com/community/join)**

Want to help but don't know where to start? Pick up one of the `help wanted` issues:

- [Steampipe](https://github.com/turbot/steampipe/labels/help%20wanted)
- [Jenkins Plugin](https://github.com/turbot/steampipe-plugin-jenkins/labels/help%20wanted)
