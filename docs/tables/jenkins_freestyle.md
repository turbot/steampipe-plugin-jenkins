# Table: jenkins_freestyle

A user-configured description of work which Jenkins should perform, such as building a piece of software, etc.

## Examples

### Freestyle project jobs in queue

```sql
select
  full_display_name,
  url
from
  jenkins_freestyle
where
  in_queue;
```

### Top bad health-scored freestyle project jobs

```sql
select
  health_report -> 0 ->> 'score' as health_report_score,
  name,
  health_report -> 0 ->> 'description' as health_report_description
from
  jenkins_freestyle
order by 
  health_report_score desc;
```

### Health color of a freestyle project and its downstream projects

```sql
select
  job.name as job_name,
  job.color as job_health_color
from
  jenkins_freestyle as job
where
  job.full_name = 'corp-project/build-and-test'
union
select
  ds_job ->> 'name' as job_name,
  ds_job ->> 'color' as job_health_color
from
  jenkins_freestyle as job,
  jsonb_array_elements(downstream_projects) as ds_job
where
  job.full_name = 'corp-project/build-and-test';
```

### Top 10 freestyle project with most builds

```sql
select
  jsonb_array_length(builds) number_of_builds,
  full_display_name
from
  jenkins_freestyle
order by
  number_of_builds desc
limit 10;
```

### Freestyle project's last successful build

```sql
select
  full_display_name,
  last_successful_build ->> 'URL' as last_successful_build
from
  jenkins_freestyle
order by
  full_display_name;
```

### Freestyle projects that last build failed

```sql
select
  full_display_name as freestyle,
  color,
  health_report -> 0 ->> 'score' as health_report_score,
  health_report -> 0 ->> 'description' as health_report_description,
  last_unsuccessful_build ->> 'URL' as last_unsuccessful_build
from
  jenkins_freestyle
where
  last_build ->> 'Number' != '0' and
  last_build ->> 'Number' = last_unsuccessful_build ->> 'Number'
order by
  full_display_name;
```
