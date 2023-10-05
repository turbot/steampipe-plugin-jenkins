## v0.1.1 [2023-10-05]

_Dependencies_

- Recompiled plugin with [steampipe-plugin-sdk v5.6.2](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v562-2023-10-03) which prevents nil pointer reference errors for implicit hydrate configs. ([#12](https://github.com/turbot/steampipe-plugin-jenkins/pull/12))

## v0.1.0 [2023-10-02]

_Dependencies_

- Upgraded to [steampipe-plugin-sdk v5.6.1](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v561-2023-09-29) with support for rate limiters. ([#9](https://github.com/turbot/steampipe-plugin-jenkins/pull/9))
- Recompiled plugin with Go version `1.21`. ([#9](https://github.com/turbot/steampipe-plugin-jenkins/pull/9))

## v0.0.1 [2023-06-27]

_What's new?_

- New tables added

  - [jenkins_build](https://hub.steampipe.io/plugins/turbot/jenkins/tables/jenkins_build)
  - [jenkins_folder](https://hub.steampipe.io/plugins/turbot/jenkins/tables/jenkins_folder)
  - [jenkins_freestyle_project](https://hub.steampipe.io/plugins/turbot/jenkins/tables/jenkins_freestyle_project)
  - [jenkins_node](https://hub.steampipe.io/plugins/turbot/jenkins/tables/jenkins_node)
  - [jenkins_pipeline](https://hub.steampipe.io/plugins/turbot/jenkins/tables/jenkins_pipeline)
  - [jenkins_plugin](https://hub.steampipe.io/plugins/turbot/jenkins/tables/jenkins_plugin)
