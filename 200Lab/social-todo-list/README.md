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

###### users

```sql
CREATE TABLE `users` (
    `id` INT NOT NULL PRIMARY KEY AUTO_INCREMENT,
    `created_at` DATETIME DEFAULT CURRENT_TIMESTAMP,
    `updated_at` DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    `email` VARCHAR(100),
    `password` VARCHAR(100),
    `salt` VARCHAR(100),
    `first_name` VARCHAR(100),
    `last_name` VARCHAR(100),
    `phone` VARCHAR(20),
    `role` ENUM("user", "admin", "shipper", "mod") DEFAULT "user",
    `status` INT DEFAULT 1
);

CREATE UNIQUE INDEX email_unique_index ON `users` (`email`);

INSERT INTO `users` (`email`, `password`, `salt`, `first_name`, `last_name`, `phone`, `role`, `status`)
VALUES
    ('user1@example.com', 'password1', 'salt1', 'John', 'Doe', '1234567890', 'user', 1),
    ('user2@example.com', 'password2', 'salt2', 'Jane', 'Smith', '9876543210', 'admin', 1),
    ('user3@example.com', 'password3', 'salt3', 'Bob', 'Johnson', '5555555555', 'shipper', 1),
    ('user4@example.com', 'password4', 'salt4', 'Alice', 'Williams', '3333333333', 'mod', 1),
    ('user5@example.com', 'password5', 'salt5', 'Charlie', 'Brown', '7777777777', 'user', 0),
    ('user6@example.com', 'password6', 'salt6', 'Eva', 'Miller', '1111111111', 'admin', 0),
    ('user7@example.com', 'password7', 'salt7', 'David', 'Davis', '9999999999', 'shipper', 1),
    ('user8@example.com', 'password8', 'salt8', 'Grace', 'Taylor', '4444444444', 'mod', 0),
    ('user9@example.com', 'password9', 'salt9', 'Frank', 'Anderson', '6666666666', 'user', 1),
    ('user10@example.com', 'password10', 'salt10', 'Sophie', 'White', '8888888888', 'admin', 1);


```

###### todo_items

```sql
CREATE TABLE `todo_items` (
    `id` INT NOT NULL PRIMARY KEY AUTO_INCREMENT,
    `user_id` INT NOT NULL,
    `created_at` DATETIME DEFAULT CURRENT_TIMESTAMP,
    `updated_at` DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    `title` VARCHAR(150),
    `description` TEXT,
    `image` JSON,
    `status` ENUM('Doing', 'Done', 'Deleted') DEFAULT 'Doing'
);

CREATE INDEX idx_status ON `todo_items` (`status`);

INSERT INTO `todo_items` (`user_id`, `title`, `description`, `status`)
VALUES
  (1, 'Task 1', 'Description for Task 1', 'Doing'),
  (1, 'Task 2', 'Description for Task 2', 'Doing'),
  (2, 'Task 3', 'Description for Task 3', 'Doing'),
  (2, 'Task 4', 'Description for Task 4', 'Done'),
  (3, 'Task 5', 'Description for Task 5', 'Done'),
  (3, 'Task 6', 'Description for Task 6', 'Deleted'),
  (4, 'Task 7', 'Description for Task 7', 'Doing'),
  (4, 'Task 8', 'Description for Task 8', 'Done'),
  (5, 'Task 9', 'Description for Task 9', 'Doing'),
  (5, 'Task 10', 'Description for Task 10', 'Done');



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
go get github.com/golang-jwt/jwt
go get -u github.com/golang-jwt/jwt/v5

```

## Links

[GORM](https://gorm.io/docs/)
