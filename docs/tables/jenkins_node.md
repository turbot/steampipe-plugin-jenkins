# Table: jenkins_node

A machine which is part of the Jenkins environment and capable of executing Pipelines or jobs.

## Examples

### Nodes offline

```sql
select
  display_name
from
  jenkins_node
where
  offline = true;
```

### Number of nodes in idle

```sql
select
  count(1)
from
  jenkins_node
where
  idle = true;
```
