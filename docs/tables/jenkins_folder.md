# Table: jenkins_folder

A container that stores job projects in it. Unlike view, which is just a filter, a folder creates a separate namespace, so you can have multiple things of the same name as long as they are in different folders.

## Examples

### Freestyle project jobs in queue of a folder

```sql
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
  jenkins_freestyle as fs
on
  fs.full_name = folder.full_name || '/' || (job ->> 'name')
where
  fs.in_queue and
  folder.full_name = 'corp-project';
```

### Number of Freestyle project jobs in queue in each folder

```sql
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
  jenkins_freestyle as fs
on
  fs.full_name = folder.full_name || '/' || (job ->> 'name')
where
  fs.in_queue
group by
  folder.full_name;
```

### Top bad health-scored jobs in a folder

```sql
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
  jenkins_freestyle as fs
on
  fs.full_name = folder.full_name || '/' || (job ->> 'name')
where
  folder.full_name = 'corp-project';
```

### Freestyle job's last successful build in a folder

```sql
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
  jenkins_freestyle as fs
on
  fs.full_name = folder.full_name || '/' || (job ->> 'name')
where
  folder.full_name = 'corp-project';
```

### Failed freestyle project in a folder

```sql
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
  jenkins_freestyle as fs
on
  fs.full_name = folder.full_name || '/' || (job ->> 'name')
where
  fs.last_build ->> 'Number' != '0' and
  fs.last_build ->> 'Number' = fs.last_unsuccessful_build ->> 'Number' and
  folder.full_name = 'corp-project'
order by
  fs.full_display_name;
```
