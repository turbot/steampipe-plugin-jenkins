---
title: "Steampipe Table: jenkins_plugin - Query Jenkins Plugins using SQL"
description: "Allows users to query Jenkins Plugins, providing insights into plugin details such as its name, version, enabled status, and more."
---

# Table: jenkins_plugin - Query Jenkins Plugins using SQL

Jenkins Plugins are integral parts of Jenkins that add additional feature sets to the Jenkins core functionality. They provide a wide range of capabilities, from building and testing code to managing deployments and automating tasks. Plugins can be installed, updated, and managed via the Jenkins administrative console.

## Table Usage Guide

The `jenkins_plugin` table provides insights into the plugins within Jenkins. As a Jenkins administrator or DevOps engineer, explore plugin-specific details through this table, including its name, version, enabled status, and more. Utilize it to manage and monitor the plugins, ensuring the optimal operation of your Jenkins environment.

## Examples

### Plugins with updates available
Discover the segments that have available updates for Jenkins plugins. This is useful for maintaining system efficiency and ensuring the use of the latest plugin features.

```sql
select
  short_name,
  version,
  long_name,
  url
from
  jenkins_plugin
where
  has_update
order by
  short_name;
```

### Inactive plugins
Determine the areas in which plugins are inactive in your Jenkins environment. This can help in identifying unused resources and optimizing system performance.

```sql
select
  short_name, 
  long_name,
  url
from
  jenkins_plugin
where
  not active
order by
  short_name;
```

### Plugins without a backup version
Identify Jenkins plugins that do not have a backup version. This is useful in understanding which plugins need attention for backup management, thereby reducing the risk of data loss.

```sql
select
  short_name, 
  long_name,
  url
from
  jenkins_plugin
where
  backup_version is null
order by
  short_name;
```

### Number of dependencies of each plugin
Explore the complexity of each plugin by determining the number of dependencies it has, which can help in understanding the intricacies and interconnectivity within your Jenkins environment.

```sql
select
  short_name, 
  long_name,
  jsonb_array_length(dependencies) number_of_dependencies
from
  jenkins_plugin
order by
  short_name;
```

### Plugins with no dependencies
This query is useful to identify plugins in your Jenkins environment that are not dependent on any others. This can help streamline your system by making it easier to manage or remove standalone plugins.

```sql
select
  short_name, 
  long_name
from
  jenkins_plugin
where
  dependencies = '[]'::jsonb
order by
  short_name;
```