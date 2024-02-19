# Basic REST API

## Setup

```bash
go mod init basic-rest-api
docker run --name mysql-200lab-basic-restapi -e MYSQL_ROOT_PASSWORD=my-secret-pw -d -p 3307:3306 mysql:8.0.31

```

## Database

```sql
CREATE DATABASE `todo-list`
    DEFAULT CHARACTER SET = 'utf8mb4';

----------------------------------------------

```
