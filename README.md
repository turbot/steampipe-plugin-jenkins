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
