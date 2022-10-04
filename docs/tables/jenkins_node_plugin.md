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
  has_update = true
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
  active = false
order by
  short_name;
```
