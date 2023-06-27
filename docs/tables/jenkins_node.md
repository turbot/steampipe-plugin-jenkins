# Table: jenkins_node

A machine which is part of the Jenkins environment and capable of executing Pipelines or jobs.

## Examples

### Total number of nodes

```sql
select
  count(1) as number_of_nodes
from
  jenkins_node;
```

### Number of idle nodes

```sql
select
  count(1) as number_of_nodes_in_idle
from
  jenkins_node
where
  idle;
```

### Get the offline nodes

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

### Nodes that allow manual launch

```sql
select
  display_name
from
  jenkins_node
where
  manual_launch_allowed;
```

### Nodes by the number of executors

```sql
select
  display_name,
  num_executors
from
  jenkins_node
order by
  num_executors desc;
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
  architecture desc;
```
