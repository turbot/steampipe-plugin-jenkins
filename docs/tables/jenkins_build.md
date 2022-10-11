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
  job_name = 'build-and-unit-test';
```

### Amount of failed builds of a job

```sql
select
  count(1) as failed_builds
from
  jenkins_build
where
  job_name = 'build-and-unit-test' and
  result = 'FAILURE'
group by
  result;
```

### Average execution time duration of successful builds of a job (in seconds)
```sql
select
  ROUND(avg(duration)/1000) as average_duration
from
  jenkins_build
where
  job_name = 'build-and-unit-test' and
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
  job_name = 'build-and-unit-test' and
  duration > estimated_duration
order by
  timestamp desc;
```
