# Table: jenkins_node

A machine which is part of the Jenkins environment and capable of executing Pipelines or jobs.

## Examples

### Nodes offline

```sql
select
  display_name,
  offline_cause,
  offline_cause_reason
from
  jenkins_node
where
  offline;
```

### Number of nodes in idle

```sql
select
  count(1)
from
  jenkins_node
where
  idle;
```

### Nodes that allows manual launch

```sql
select
  display_name
from
  jenkins_node
where
  manual_launch_allowed;
```

### Nodes by number of executors

```sql
select
  display_name,
  num_executors
from
  jenkins_node
order by
  num_executors DESC;
```

### Number of nodes by OS and architecture type
```sql
select
  monitor_data ->> 'hudson.node_monitors.ArchitectureMonitor' as architecture,
  count(1) as nodes
from
  jenkins_node
group by
  architecture
order by
  architecture DESC;
```
