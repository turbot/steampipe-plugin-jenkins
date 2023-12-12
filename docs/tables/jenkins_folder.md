---
title: "Steampipe Table: jenkins_folder - Query Jenkins Folders using SQL"
description: "Allows users to query Jenkins Folders, specifically providing insights into the organization and structure of jobs within Jenkins."
---

# Table: jenkins_folder - Query Jenkins Folders using SQL

Jenkins Folders are a resource within the Jenkins Continuous Integration and Continuous Delivery (CI/CD) server. They allow for the organization and grouping of jobs, creating a hierarchical structure within Jenkins. This is particularly useful in large projects, where managing a large number of jobs can become complex.

## Table Usage Guide

The `jenkins_folder` table provides insights into the organization and structuring of jobs within Jenkins. As a DevOps engineer, explore folder-specific details through this table, including the name, URL, and associated jobs. Utilize it to uncover information about the structure and organization of jobs, helping to manage and navigate large Jenkins projects.

## Examples

### Freestyle project jobs in queue of a folder
Discover the segments that are part of a specific project folder and are queued for execution, providing a way to monitor and manage your pipeline efficiently. This could be particularly useful in large projects where multiple jobs are in queue and prioritization is needed.

```sql+postgres
select
  folder.full_name,
  fs.full_display_name,
  fs.url
from
  jenkins_folder as folder
join
  jsonb_array_elements(jobs) as job
on
  true
join
  jenkins_freestyle_project as fs
on
  fs.full_name = folder.full_name || '/' || (job ->> 'name')
where
  fs.in_queue and
  folder.full_name = 'corp-project';
```

```sql+sqlite
select
  folder.full_name,
  fs.full_display_name,
  fs.url
from
  jenkins_folder as folder,
  json_each(jobs) as job
join
  jenkins_freestyle_project as fs
on
  fs.full_name = folder.full_name || '/' || json_extract(job.value, '$.name')
where
  fs.in_queue and
  folder.full_name = 'corp-project';
```

### Number of Freestyle project jobs in queue in each folder
Explore the distribution of pending tasks across different project categories in a Jenkins environment. This can help optimize task scheduling and resource allocation by identifying areas with high demand.

```sql+postgres
select
  folder.full_name as folder,
  count(1) as jobs_in_queue
from
  jenkins_folder as folder
join
  jsonb_array_elements(jobs) as job
on
  true
join
  jenkins_freestyle_project as fs
on
  fs.full_name = folder.full_name || '/' || (job ->> 'name')
where
  fs.in_queue
group by
  folder.full_name;
```

```sql+sqlite
select
  folder.full_name as folder,
  count(1) as jobs_in_queue
from
  jenkins_folder as folder,
  json_each(jobs) as job
join
  jenkins_freestyle_project as fs
on
  fs.full_name = folder.full_name || '/' || json_extract(job.value, '$.name')
where
  fs.in_queue
group by
  folder.full_name;
```

### Top bad health-scored jobs in a folder
Discover the segments that have the worst health scores within a specific project folder. This is useful to identify areas in need of immediate attention or improvement.

```sql+postgres
select
  fs.health_report -> 0 ->> 'score' as health_report_score,
  fs.full_display_name,
  fs.health_report -> 0 ->> 'description' as health_report_description
from
  jenkins_folder as folder
join
  jsonb_array_elements(jobs) as job
on
  true
join
  jenkins_freestyle_project as fs
on
  fs.full_name = folder.full_name || '/' || (job ->> 'name')
where
  folder.full_name = 'corp-project';
```

```sql+sqlite
select
  json_extract(fs.health_report, '$[0].score') as health_report_score,
  fs.full_display_name,
  json_extract(fs.health_report, '$[0].description') as health_report_description
from
  jenkins_folder as folder,
  json_each(folder.jobs) as job
join
  jenkins_freestyle_project as fs
on
  fs.full_name = folder.full_name || '/' || json_extract(job.value, '$.name')
where
  folder.full_name = 'corp-project';
```

### Freestyle job's last successful build in a folder
Explore the last successful build of a freestyle job within a specific project folder. This can help in understanding the project's build history and identifying any potential issues or areas for improvement.

```sql+postgres
select
  fs.full_display_name,
  fs.last_successful_build ->> 'URL' as last_successful_build
from
  jenkins_folder as folder
join
  jsonb_array_elements(jobs) as job
on
  true
join
  jenkins_freestyle_project as fs
on
  fs.full_name = folder.full_name || '/' || (job ->> 'name')
where
  folder.full_name = 'corp-project';
```

```sql+sqlite
select
  fs.full_display_name,
  json_extract(fs.last_successful_build, '$.URL') as last_successful_build
from
  jenkins_folder as folder,
  json_each(folder.jobs) as job
join
  jenkins_freestyle_project as fs
on
  fs.full_name = folder.full_name || '/' || json_extract(job.value, '$.name')
where
  folder.full_name = 'corp-project';
```

### Failed freestyle project in a folder
Determine the health status and details of unsuccessful freestyle projects within a specific Jenkins folder. This aids in identifying problematic areas and taking corrective measures promptly.

```sql+postgres
select
  fs.full_display_name as job,
  fs.color,
  fs.health_report -> 0 ->> 'score' as health_report_score,
  fs.health_report -> 0 ->> 'description' as health_report_description,
  fs.last_unsuccessful_build ->> 'URL' as last_unsuccessful_build
from
  jenkins_folder as folder
join
  jsonb_array_elements(jobs) as job
on
  true
join
  jenkins_freestyle_project as fs
on
  fs.full_name = folder.full_name || '/' || (job ->> 'name')
where
  fs.last_build ->> 'Number' != '0' and
  fs.last_build ->> 'Number' = fs.last_unsuccessful_build ->> 'Number' and
  folder.full_name = 'corp-project'
order by
  fs.full_display_name;
```

```sql+sqlite
Error: The corresponding SQLite query is unavailable. 
```