---
title: "Steampipe Table: jenkins_node - Query Jenkins Nodes using SQL"
description: "Allows users to query Jenkins Nodes, specifically detailed information about each node including the name, description, number of executors, labels, and status."
---

# Table: jenkins_node - Query Jenkins Nodes using SQL

Jenkins Nodes are the worker machines that are part of the Jenkins distributed build system. They are responsible for executing the build jobs dispatched by the master. Each Jenkins Node can have different operating systems and architecture, allowing for diverse build environments.

## Table Usage Guide

The `jenkins_node` table provides insights into Jenkins Nodes within Jenkins distributed build system. As a DevOps engineer, explore node-specific details through this table, including name, description, number of executors, labels, and status. Utilize it to uncover information about nodes, such as those with high executor counts, the labels associated with each node, and the status of each node.

## Examples

### Total number of nodes
Explore the total count of nodes in your Jenkins environment to understand the scale of your build and test infrastructure. This information can be useful for capacity planning and resource allocation.

```sql+postgres
select
  count(1) as number_of_nodes
from
  jenkins_node;
```

```sql+sqlite
select
  count(1) as number_of_nodes
from
  jenkins_node;
```

### Number of idle nodes
Explore how many nodes are currently idle in the Jenkins system. This can help in assessing system resource utilization and planning capacity.

```sql+postgres
select
  count(1) as number_of_nodes_in_idle
from
  jenkins_node
where
  idle;
```

```sql+sqlite
select
  count(1) as number_of_nodes_in_idle
from
  jenkins_node
where
  idle = 1;
```

### Get the offline nodes
Explore which Jenkins nodes are offline and understand the underlying reasons for their status. This can help in identifying issues and implementing appropriate solutions to restore these nodes.

```sql+postgres
select
  display_name,
  offline_cause,
  offline_cause_reason
from
  jenkins_node
where
  offline;
```

```sql+sqlite
select
  display_name,
  offline_cause,
  offline_cause_reason
from
  jenkins_node
where
  offline = 1;
```

### Nodes that allow manual launch
Discover the segments where manual launch is permitted, offering you more control and flexibility in your operations. This can be useful in situations where automated launches may not be ideal or in testing environments.

```sql+postgres
select
  display_name
from
  jenkins_node
where
  manual_launch_allowed;
```

```sql+sqlite
select
  display_name
from
  jenkins_node
where
  manual_launch_allowed = 1;
```

### Nodes by the number of executors
Analyze the settings to understand the distribution of executors across different nodes in a Jenkins environment. This can help in balancing workload and optimizing resource utilization.

```sql+postgres
select
  display_name,
  num_executors
from
  jenkins_node
order by
  num_executors desc;
```

```sql+sqlite
select
  display_name,
  num_executors
from
  jenkins_node
order by
  num_executors desc;
```

### Number of nodes by OS and architecture type
Explore the distribution of nodes by operating system and architecture type, allowing you to understand your system's structure and diversity. This can be particularly useful for planning updates or assessing compatibility requirements.

```sql+postgres
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

```sql+sqlite
select
  json_extract(monitor_data, '$.hudson.node_monitors.ArchitectureMonitor') as architecture,
  count(1) as nodes
from
  jenkins_node
group by
  architecture
order by
  architecture desc;
```