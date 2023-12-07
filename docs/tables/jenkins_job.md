---
title: "Steampipe Table: jenkins_job - Query Jenkins Jobs using SQL"
description: "Allows users to query Jenkins Jobs, specifically providing details about each job such as its name, description, URL, and build status."
---

# Table: jenkins_job - Query Jenkins Jobs using SQL

Jenkins is an open-source automation server that enables developers to build, test, and deploy their software. It facilitates continuous integration and continuous delivery (CI/CD) by automating parts of the development process, with a focus on testing and deployment. A Jenkins Job is a runnable task that is controlled and monitored by Jenkins.

## Table Usage Guide

The `jenkins_job` table provides insights into Jobs within Jenkins. As a DevOps engineer, explore job-specific details through this table, including job name, description, URL, and build status. Utilize it to monitor the status of various tasks, identify any jobs that may have failed, and verify the details of each job.

## Examples

### List Maven projects
Explore which Jenkins jobs are configured as Maven projects to manage and understand your build process better. This can help streamline your development workflow and troubleshoot potential issues.

```sql+postgres
select
  full_display_name,
  url,
  properties
from
  jenkins_job
where
  class = 'hudson.maven.MavenModuleSet';
```

```sql+sqlite
select
  full_display_name,
  url,
  properties
from
  jenkins_job
where
  class = 'hudson.maven.MavenModuleSet';
```

### List child jobs of a Multibranch Pipeline
Explore the different tasks under a multi-branch pipeline to understand its structure and workflow. This query is useful in identifying the various jobs within a pipeline, their status, and their respective URLs for easy access and management.

```sql+postgres
select
  j ->> 'name' name,
  j ->> 'color' color,
  j ->> 'url' url
from
  jenkins_job m,
  jsonb_array_elements(properties -> 'jobs') as j
where
  m.class = 'org.jenkinsci.plugins.workflow.multibranch.WorkflowMultiBranchProject';
```

```sql+sqlite
select
  json_extract(j.value, '$.name') name,
  json_extract(j.value, '$.color') color,
  json_extract(j.value, '$.url') url
from
  jenkins_job m,
  json_each(m.properties, '$.jobs') as j
where
  m.class = 'org.jenkinsci.plugins.workflow.multibranch.WorkflowMultiBranchProject';
```

### Jobs in queue
Discover the segments that are currently in the Jenkins job queue, allowing you to prioritize tasks and manage workflow effectively.

```sql+postgres
select
  full_display_name,
  url
from
  jenkins_job
where
  (properties ->> 'inQueue')::boolean;
```

```sql+sqlite
select
  full_display_name,
  url
from
  jenkins_job
where
  json_extract(properties, '$.inQueue') = 'true';
```

### Top bad health-scored jobs
Identify instances where jobs have a poor health score in order to prioritize and address them. This helps in maintaining the overall health and efficiency of the system.

```sql+postgres
select
  properties -> 'healthReport' -> 0 ->> 'score' as health_report_score,
  name,
  properties -> 'healthReport' -> 0 ->> 'description' as health_report_description
from
  jenkins_job
order by 
  health_report_score desc;
```

```sql+sqlite
select
  json_extract(json_extract(properties, '$.healthReport[0]'), '$.score') as health_report_score,
  name,
  json_extract(json_extract(properties, '$.healthReport[0]'), '$.description') as health_report_description
from
  jenkins_job
order by 
  health_report_score desc;
```

### Last successful build of a job
Uncover the details of the most recent successful build for a job in Jenkins. This can be useful to identify potential issues and improve future builds.

```sql+postgres
select
  full_display_name,
  properties -> 'lastSuccessfulBuild' ->> 'URL' as last_successful_build
from
  jenkins_job
order by
  full_display_name;
```

```sql+sqlite
select
  full_display_name,
  json_extract(json_extract(properties, '$.lastSuccessfulBuild'), '$.URL') as last_successful_build
from
  jenkins_job
order by
  full_display_name;
```

### Jobs that last build failed
Identify instances where the most recent job execution was unsuccessful, enabling you to analyze and rectify the issues causing the failure.

```sql+postgres
select
  full_display_name as job,
  properties ->> 'color' as color,
  properties -> 'healthReport' -> 0 ->> 'score' as health_report_score,
  properties -> 'healthReport' -> 0 ->> 'description' as health_report_description,
  properties -> 'lastUnsuccessfulBuild' ->> 'URL' as last_unsuccessful_build
from
  jenkins_job
where
  properties -> 'last_build' ->> 'Number' != '0' and
  properties -> 'last_build' ->> 'Number' = properties -> 'lastUnsuccessfulBuild' ->> 'Number'
order by
  full_display_name;
```

```sql+sqlite
select
  full_display_name as job,
  json_extract(properties, '$.color') as color,
  json_extract(json_extract(properties, '$.healthReport'), '$[0].score') as health_report_score,
  json_extract(json_extract(properties, '$.healthReport'), '$[0].description') as health_report_description,
  json_extract(json_extract(properties, '$.lastUnsuccessfulBuild'), '$.URL') as last_unsuccessful_build
from
  jenkins_job
where
  json_extract(json_extract(properties, '$.last_build'), '$.Number') != '0' and
  json_extract(json_extract(properties, '$.last_build'), '$.Number') = json_extract(json_extract(properties, '$.lastUnsuccessfulBuild'), '$.Number')
order by
  full_display_name;
```