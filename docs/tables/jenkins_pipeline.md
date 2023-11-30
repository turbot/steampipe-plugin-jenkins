---
title: "Steampipe Table: jenkins_pipeline - Query Jenkins Pipelines using SQL"
description: "Allows users to query Jenkins Pipelines, specifically providing insights into pipeline configurations and their status."
---

# Table: jenkins_pipeline - Query Jenkins Pipelines using SQL

Jenkins is a self-contained, open-source automation server which can be used to automate all sorts of tasks related to building, testing, and delivering or deploying software. Jenkins Pipelines are a suite of plugins supporting implementation and integration of continuous delivery pipelines into Jenkins. It provides an extensible set of tools for modeling simple-to-complex delivery pipelines as code.

## Table Usage Guide

The `jenkins_pipeline` table provides insights into Jenkins Pipelines within the Jenkins automation server. As a DevOps engineer, explore pipeline-specific details through this table, including configurations, status, and associated metadata. Utilize it to uncover information about pipelines, such as their current status, the steps involved in the pipeline, and the verification of pipeline configurations.

## Examples

### Pipelines in queue
Discover the segments that are currently in queue within the Jenkins pipeline, providing you with a quick overview and access to their URLs. This is useful to prioritize tasks and manage workflow efficiently.

```sql
select
  full_display_name,
  url
from
  jenkins_pipeline
where
  in_queue;
```

### Top bad health-scored pipelines
Uncover the details of your Jenkins pipelines with the lowest health scores to understand potential areas of improvement. This query is particularly useful in identifying pipelines that might require immediate attention due to their poor health scores.

```sql
select
  health_report -> 0 ->> 'score' as health_report_score,
  full_display_name,
  health_report -> 0 ->> 'description' as health_report_description
from
  jenkins_pipeline
order by 
  health_report_score desc;
```

### Health color of a pipeline
Analyze the health status of a specific pipeline in a Jenkins project. This query is particularly useful for understanding the operational state of a pipeline, which can guide troubleshooting efforts or inform operational decisions.

```sql
select
  full_display_name as pipeline,
  color as health_color
from
  jenkins_pipeline
where
  full_name = 'corp-project/build-and-test-pipeline';
```

### How long a pipeline usually takes to run (in seconds)?
Analyze the average time it takes for a successful pipeline to run in your project, aiding in performance optimization and resource planning. This can help identify any potential bottlenecks and improve overall efficiency.

```sql
select
  ROUND(avg(b.duration)/1000) as average_duration
from
  jenkins_pipeline as p
join
  jenkins_build as b
on
  b.job_full_name = p.full_name
where
  b.result = 'SUCCESS' and
  p.full_name = 'corp-project/build-and-test-pipeline'
group by
  b.result;
```

### Pipeline's last successful build
Explore which Jenkins pipelines had successful builds last, providing a quick overview of successful deployments. This can help in assessing the stability and reliability of different pipelines.

```sql
select
  full_display_name,
  last_successful_build ->> 'URL' as last_successful_build
from
  jenkins_pipeline
order by
  full_display_name;
```

### Pipelines where the last build failed
This query helps identify pipelines where the most recent build was unsuccessful, providing insights into potential issues and facilitating quick troubleshooting. It's useful for maintaining the health and efficiency of your Jenkins pipelines.

```sql
select
  full_display_name as pipeline,
  color,
  health_report -> 0 ->> 'score' as health_report_score,
  health_report -> 0 ->> 'description' as health_report_description,
  last_unsuccessful_build ->> 'URL' as last_unsuccessful_build
from
  jenkins_pipeline
where
  last_build ->> 'Number' != '0' and
  last_build ->> 'Number' = last_unsuccessful_build ->> 'Number'
order by
  full_display_name;
```