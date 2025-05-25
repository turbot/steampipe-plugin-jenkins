---
title: "Steampipe Table: jenkins_user - Query Jenkins Users using SQL"
description: "Allows users to query Jenkins Users, providing insights into users details such as its full_name, absolute_url and more."
---

# Table: jenkins_user – Query Jenkins Users Using SQL

Jenkins users are entities configured within a Jenkins instance, typically representing developers, admins, or service accounts interacting with the Jenkins system. This table provides insights into the users configured in Jenkins, including their full names and profile URLs.

## Table Usage Guide
The jenkins_user table allows Jenkins administrators, DevOps engineers, and auditors to query details about users in a Jenkins instance. This includes identifying active users, reviewing user profiles, and auditing user access or existence.

Each row in this table represents a single Jenkins user account retrieved from the Jenkins user database or API.

## Examples

### List all users with full profile URLs
This query retrieves the full_name and absolute_url of all users configured in your Jenkins environment. It's useful for getting a complete overview of every user account, including their access URLs, which can help with user audits, role reviews, or general system documentation.

```sql+postgres
select
  full_name,
  absolute_url
from
  jenkins_user;
```

```sql+sqlite
select
  full_name,
  absolute_url
from
  jenkins_user;
```

### List specfic users with full profile URLs
This query filters the list of Jenkins users to show only those whose full_name contains the term "admin" (case-insensitive). It's particularly helpful for identifying administrative or privileged accounts, ensuring proper access control, and monitoring for potential misconfigured or unauthorized users.

```sql+postgres
select
  full_name,
  absolute_url
from
  jenkins_user
where
  full_name ilike '%admin%';
```

```sql+sqlite
select
  full_name,
  absolute_url
from
  jenkins_user
where
  lower(full_name) like '%admin%';
```