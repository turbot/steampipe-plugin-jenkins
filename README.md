![image](https://hub.steampipe.io/images/plugins/turbot/jenkins-social-graphic.png)

# Jenkins Plugin for Steampipe

Use SQL to query jobs, builds, nodes, plugin and more from Jenkins.

- **[Get started →](https://hub.steampipe.io/plugins/turbot/jenkins)**
- Documentation: [Table definitions & examples](https://hub.steampipe.io/plugins/turbot/jenkins/tables)
- Community: [Slack Channel](https://steampipe.io/community/join)
- Get involved: [Issues](https://github.com/turbot/steampipe-plugin-jenkins/issues)

## Quick start

Install the plugin with [Steampipe](https://steampipe.io):

```shell
steampipe plugin install jenkins
```

Run a query:

```sql
select
  color,
  name,
  in_queue,
  last_completed_build  ->> 'URL' as last_completed_build
from
  jenkins_job;
```

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

## Contributing

Please see the [contribution guidelines](https://github.com/turbot/steampipe/blob/main/CONTRIBUTING.md) and our [code of conduct](https://github.com/turbot/steampipe/blob/main/CODE_OF_CONDUCT.md). All contributions are subject to the [Apache 2.0 open source license](https://github.com/turbot/steampipe-plugin-jenkins/blob/main/LICENSE).

`help wanted` issues:

- [Steampipe](https://github.com/turbot/steampipe/labels/help%20wanted)
- [Jenkins Plugin](https://github.com/turbot/steampipe-plugin-jenkins/labels/help%20wanted)
