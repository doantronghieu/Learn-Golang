# Description

## Database

### MySQL

#### [Docker](https://hub.docker.com/_/mysql)

```console
docker run --name mysql-200lab -e MYSQL_ROOT_PASSWORD=my-secret-pw -d -p 3306:3306 mysql:8.0.31
docker ps
```

#### Connection

- Connection name: MySQL-200Lab
- Server Address: 127.0.0.1
- Port: 3306
- Username: root
- Password: my-secret-pw
- SSL: Disabled

#### Databases

##### social-todo-list

```sql
CREATE DATABASE `social-todo-list`
    DEFAULT CHARACTER SET = 'utf8mb4';
```

###### todo_items

```sql
CREATE TABLE `todo_items` (
    `id` INT NOT NULL PRIMARY KEY AUTO_INCREMENT,
    `created_at` DATETIME DEFAULT CURRENT_TIMESTAMP,
    `updated_at` DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    `title` VARCHAR(150),
    `description` TEXT,
    `image` JSON,
    `status` ENUM('Doing', 'Done', 'Deleted') DEFAULT 'Doing'
);
```

```sql
CREATE INDEX idx_status ON `todo_items` (`status`);
```

```sql
INSERT INTO `todo_items` (`title`, `description`, `status`)
VALUES
    ('Task 1 title', 'Task 1 description'),
    ('Task 2 title', 'Task 2 description', 'Doing'),
    ('Task 3 title', 'Task 3 description', 'Done'),
    ('Task 4 title', 'Task 4 description', 'Deleted'),
    ('Task 5 title', 'Task 5 description', 'Doing'),
    ('Task 6 title', 'Task 6 description', 'Doing'),
    ('Task 7 title', 'Task 7 description', 'Done'),
    ('Task 8 title', 'Task 8 description', 'Doing'),
    ('Task 9 title', 'Task 9 description', 'Done'),
    ('Task 10 title', 'Task 10 description', 'Deleted'),
    ('Task 11 title', 'Task 11 description', 'Doing');
```

###### user_like_items

```sql
CREATE TABLE `user_like_items` (
    `user_id` INT NOT NULL,
    `item_id` INT NOT NULL,
    `created_at` DATETIME DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (`user_id`, `item_id`)
);

```

```sql
CREATE INDEX idx_item_id ON `user_like_items` (`item_id`);

```

## Golang

```bash
go mod init social-todo-list
go get -u github.com/gin-gonic/gin
go get -u gorm.io/gorm
go get -u gorm.io/driver/mysql
go get github.com/joho/godotenv

```

## Links

[GORM](https://gorm.io/docs/)
