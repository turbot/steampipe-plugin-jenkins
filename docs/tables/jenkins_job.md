# Table: jenkins_job

A user-configured description of work which Jenkins should perform. This table is a generic representation of a Jenkins Job.

## Examples

### List Maven projects

```sql
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

```sql
select
  j ->> 'name' name,
  j ->> 'color' color,
  j ->> 'url' url
from
  jenkins_job m,
  jsonb_array_elements(properties -> 'jobs') as j
where
  m.class = 'org.jenkinsci.plugins.workflow.multibranch.WorkflowMultiBranchProject'
```

### Jobs in queue

```sql
select
  full_display_name,
  url
from
  jenkins_job
where
  (properties ->> 'inQueue')::boolean;
```

### Top bad health-scored jobs

```sql
select
  properties -> 'healthReport' -> 0 ->> 'score' as health_report_score,
  name,
  properties -> 'healthReport' -> 0 ->> 'description' as health_report_description
from
  jenkins_job
order by 
  health_report_score desc;
```

### Last successful build of a job

```sql
select
  full_display_name,
  properties -> 'lastSuccessfulBuild' ->> 'URL' as last_successful_build
from
  jenkins_job
order by
  full_display_name;
```

### Jobs that last build failed

```sql
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
