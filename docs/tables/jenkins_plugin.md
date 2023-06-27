# Table: jenkins_plugin

An extension to Jenkins functionality provided separately from Jenkins Core.

## Examples

### Plugins with available update

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
