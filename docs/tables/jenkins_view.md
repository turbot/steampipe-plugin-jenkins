# Table: jenkins_view

A list view of jobs based on a filter or in a manual selection of jobs.

## Examples

### Amount of jobs in views

```sql
select
  name as view,
  jsonb_array_length(jobs) as jobs
from
  jenkins_view
order by
  jobs;
```

### Job's health colors and names of a view

```sql
select
  job ->> 'color' as job_color,
  job ->> 'name' as job_name
from
  jenkins_view as view,
  jsonb_array_elements(jobs) as job
where
  view.name = 'dev-phase';
```

### Amount of jobs in queue per view

```sql
select
  view.name as view,
  count(1) as jobs_in_queue
from
  jenkins_view as view,
  jsonb_array_elements(jobs) as view_job
left join
  jenkins_job as job
on
  job.name = view_job ->> 'name'
where
  job.in_queue
group by
  view.name;
```

### Job's last successful build

```sql
select
  job.name as job,
  job.last_successful_build ->> 'URL' as last_successful_build
from
  jenkins_view as view,
  jsonb_array_elements(jobs) as view_job
left join
  jenkins_job as job
on
  job.name = view_job ->> 'name'
where
  view.name = 'dev-phase';
```
