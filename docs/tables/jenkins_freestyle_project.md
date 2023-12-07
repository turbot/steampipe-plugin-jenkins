---
title: "Steampipe Table: jenkins_freestyle_project - Query Jenkins Freestyle Projects using SQL"
description: "Allows users to query Jenkins Freestyle Projects, specifically the project name, description, URL, and build details, providing insights into project configurations and build statuses."
---

# Table: jenkins_freestyle_project - Query Jenkins Freestyle Projects using SQL

Jenkins is an open-source automation server that enables developers to build, test, and deploy their software. It provides numerous plugins to support building, deploying, and automating any project. A Freestyle Project in Jenkins is a type of project that provides maximum flexibility and simplicity for the users. It is the simplest and the most flexible way to design a build, especially when the build process involves several different steps.

## Table Usage Guide

The `jenkins_freestyle_project` table provides insights into Freestyle Projects within Jenkins. As a DevOps engineer, explore project-specific details through this table, including project name, description, URL, and build details. Utilize it to uncover information about projects, such as their configurations, the status of the builds, and the details of the last build.

## Examples

### Freestyle project jobs in queue
Explore which freestyle project jobs are currently in queue to manage your workload and prioritize tasks effectively. This helps in optimizing your project pipeline and ensuring smooth operations.

```sql+postgres
select
  full_display_name,
  url
from
  jenkins_freestyle_project
where
  in_queue;
```

```sql+sqlite
select
  full_display_name,
  url
from
  jenkins_freestyle_project
where
  in_queue = 1;
```

### Top bad health-scored freestyle project jobs
Determine the areas in which freestyle project jobs have poor health scores. This helps prioritize and address issues to improve overall project performance.

```sql+postgres
select
  health_report -> 0 ->> 'score' as health_report_score,
  name,
  health_report -> 0 ->> 'description' as health_report_description
from
  jenkins_freestyle_project
order by 
  health_report_score desc;
```

```sql+sqlite
select
  json_extract(json_extract(health_report, '$[0]'), '$.score') as health_report_score,
  name,
  json_extract(json_extract(health_report, '$[0]'), '$.description') as health_report_description
from
  jenkins_freestyle_project
order by 
  health_report_score desc;
```

### Health color of a freestyle project and its downstream projects
This query helps to assess the health status of a specific project and its related downstream projects in a Jenkins environment. It is particularly useful for monitoring project health in real time, enabling proactive issue detection and resolution.

```sql+postgres
select
  job.name as job_name,
  job.color as job_health_color
from
  jenkins_freestyle_project as job
where
  job.full_name = 'corp-project/build-and-test'
union
select
  ds_job ->> 'name' as job_name,
  ds_job ->> 'color' as job_health_color
from
  jenkins_freestyle_project as job,
  jsonb_array_elements(downstream_projects) as ds_job
where
  job.full_name = 'corp-project/build-and-test';
```

```sql+sqlite
select
  job.name as job_name,
  job.color as job_health_color
from
  jenkins_freestyle_project as job
where
  job.full_name = 'corp-project/build-and-test'
union
select
  json_extract(ds_job.value, '$.name') as job_name,
  json_extract(ds_job.value, '$.color') as job_health_color
from
  jenkins_freestyle_project as job,
  json_each(downstream_projects) as ds_job
where
  job.full_name = 'corp-project/build-and-test';
```

### Top 10 freestyle projects with most builds
Analyze your Jenkins freestyle projects to identify the top ten with the most builds. This can help prioritize maintenance efforts and understand where your resources are most heavily utilized.

```sql+postgres
select
  jsonb_array_length(builds) number_of_builds,
  full_display_name
from
  jenkins_freestyle_project
order by
  number_of_builds desc
limit 10;
```

```sql+sqlite
select
  json_array_length(builds) as number_of_builds,
  full_display_name
from
  jenkins_freestyle_project
order by
  number_of_builds desc
limit 10;
```

### Freestyle project's last successful build
This query helps you identify the last successful build of each freestyle project in your Jenkins environment, which can assist in tracking project progress and ensuring builds are completing successfully. It is particularly useful for maintaining build quality and identifying issues early by pinpointing the projects where the most recent build was successful.

```sql+postgres
select
  full_display_name,
  last_successful_build ->> 'URL' as last_successful_build
from
  jenkins_freestyle_project
order by
  full_display_name;
```

```sql+sqlite
select
  full_display_name,
  json_extract(last_successful_build, '$.URL') as last_successful_build
from
  jenkins_freestyle_project
order by
  full_display_name;
```

### Freestyle projects where the last build failed
Identify freestyle projects in Jenkins where the most recent build was unsuccessful. This can help in quickly pinpointing problematic projects, allowing for timely troubleshooting and resolution.

```sql+postgres
select
  full_display_name as freestyle,
  color,
  health_report -> 0 ->> 'score' as health_report_score,
  health_report -> 0 ->> 'description' as health_report_description,
  last_unsuccessful_build ->> 'URL' as last_unsuccessful_build
from
  jenkins_freestyle_project
where
  last_build ->> 'Number' != '0' and
  last_build ->> 'Number' = last_unsuccessful_build ->> 'Number'
order by
  full_display_name;
```

```sql+sqlite
select
  full_display_name as freestyle,
  color,
  json_extract(health_report, '$[0].score') as health_report_score,
  json_extract(health_report, '$[0].description') as health_report_description,
  json_extract(last_unsuccessful_build, '$.URL') as last_unsuccessful_build
from
  jenkins_freestyle_project
where
  json_extract(last_build, '$.Number') != '0' and
  json_extract(last_build, '$.Number') = json_extract(last_unsuccessful_build, '$.Number')
order by
  full_display_name;
```