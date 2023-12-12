---
title: "Steampipe Table: jenkins_build - Query Jenkins Builds using SQL"
description: "Allows users to query Jenkins Builds, specifically the build details, providing insights into build status, duration, and other related details."
---

# Table: jenkins_build - Query Jenkins Builds using SQL

Jenkins is an open-source automation server that enables developers to build, test, and deploy their software. It provides numerous plugins to support building, deploying, and automating any project. A Jenkins Build is a single execution of a Jenkins Job, which includes steps and post-build actions, and contains all the results of the steps.

## Table Usage Guide

The `jenkins_build` table provides insights into Jenkins Builds within the Jenkins automation server. As a DevOps engineer, explore build-specific details through this table, including build status, duration, and associated metadata. Utilize it to uncover information about builds, such as those with failed tests, the duration of each build, and the verification of build results.

## Examples

### Artifacts created by a build
Analyze the artifacts produced by a specific build process to understand what files were generated. This can help in assessing the output of a build process and identifying any unexpected or missing files.

```sql+postgres
select
  artifact ->> 'fileName' as file_name
from
  jenkins_build as build,
  jsonb_array_elements(artifacts) as artifact
where
  build.number = 128 and
  job_full_name = 'build-and-unit-test';
```

```sql+sqlite
select
  json_extract(artifact.value, '$.fileName') as file_name
from
  jenkins_build as build,
  json_each(build.artifacts) as artifact
where
  build.number = 128 and
  job_full_name = 'build-and-unit-test';
```

### Amount of failed builds by freestyle job
Determine the areas in which your freestyle projects are experiencing the most build failures. This can help you identify problematic projects and prioritize them for troubleshooting and optimization.

```sql+postgres
select
  j.full_name as job,
  count(1) as failed_builds
from
  jenkins_freestyle_project as j
join
  jenkins_build b
on
  b.job_full_name = j.full_name
where
  b.result = 'FAILURE'
group by
  j.full_name
order by
  failed_builds desc;
```

```sql+sqlite
select
  j.full_name as job,
  count(1) as failed_builds
from
  jenkins_freestyle_project as j
join
  jenkins_build b
on
  b.job_full_name = j.full_name
where
  b.result = 'FAILURE'
group by
  j.full_name
order by
  failed_builds desc;
```

### Average execution time duration of successful builds of a job (in seconds)
Determine the average duration of successful builds for a specific job to gain insights into performance efficiency and identify potential areas for process optimization.

```sql+postgres
select
  ROUND(avg(duration)/1000) as average_duration
from
  jenkins_build
where
  job_full_name = 'corp-project/build-and-test' and
  result = 'SUCCESS'
group by
  result;
```

```sql+sqlite
select
  ROUND(avg(duration)/1000) as average_duration
from
  jenkins_build
where
  job_full_name = 'corp-project/build-and-test' and
  result = 'SUCCESS'
group by
  result;
```

### Builds that took longer than estimated to execute (in seconds)
Determine the instances where certain build processes took longer than anticipated in a specific production project. This could be useful in identifying inefficiencies and areas for improvement in the production process.

```sql+postgres
select
  full_display_name,
  result,
  (duration - estimated_duration) / 1000 as difference,
  url
from
  jenkins_build
where
  job_full_name = 'corp-project/production/deploy-to-prod' and
  duration > estimated_duration
order by
  timestamp desc;
```

```sql+sqlite
select
  full_display_name,
  result,
  (duration - estimated_duration) / 1000 as difference,
  url
from
  jenkins_build
where
  job_full_name = 'corp-project/production/deploy-to-prod' and
  duration > estimated_duration
order by
  timestamp desc;
```