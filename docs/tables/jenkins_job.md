# Table: jenkins_job

A user-configured description of work which Jenkins should perform, such as building a piece of software, etc.

## Examples

### Jobs in queue

```sql
select
  name,
  url
from
  jenkins_job
where
  in_queue;
```

### Top bad health-scored Jobs

```sql
select
  health_report -> 0 ->> 'score' as health_report_score,
  name,
  health_report -> 0 ->> 'description' as health_report_description
from
  jenkins_job
order by 
  health_report_score desc;
```

### Next build number of a job

```sql
select
  next_build_number
from
  jenkins_job
where
  name = 'my-job';
```

### Top 10 Jobs with most builds

```sql
select
  jsonb_array_length(builds) number_of_builds,
  name
from
  jenkins_job
order by
  number_of_builds desc
limit 10;
```

### Job's last successful build

```sql
select
  name,
  last_successful_build ->> 'URL' as last_successful_build
from
  jenkins_job
order by
  name;
```
