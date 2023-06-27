# Table: jenkins_build

Result of a single execution of a job.

## Examples

### Artifacts created by a build

```sql
select
  artifact ->> 'fileName' as file_name
from
  jenkins_build as build,
  jsonb_array_elements(artifacts) as artifact
where
  build.number = 128 and
  job_full_name = 'build-and-unit-test';
```

### Amount of failed builds by freestyle job

```sql
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

```sql
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

```sql
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
