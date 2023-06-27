# Table: jenkins_pipeline

Orchestrates long-running activities that can span multiple build agents. Suitable for building pipelines (formerly known as workflows) and/or organizing complex activities that do not easily fit in free-style job type.

## Examples

### Pipelines in queue

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
