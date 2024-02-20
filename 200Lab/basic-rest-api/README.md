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

CREATE TABLE `todo_items`(
  `id` INT NOT NULL PRIMARY KEY AUTO_INCREMENT,
  `title` VARCHAR(255),
  `image` JSON,
  `description` TEXT,
  `status` ENUM('Doing', 'Done', 'Deleted') DEFAULT 'Doing',
  `created_at` DATETIME DEFAULT NOW(),
  `updated_at` DATETIME DEFAULT NOW() ON UPDATE CURRENT_TIMESTAMP
);


CREATE INDEX idx_status ON `todo_items` (`status`);

INSERT INTO `todo_items` (`title`, `description`, `status`)
VALUES
  ('Task 1', 'Description for Task 1', 'Doing'),
  ('Task 2', 'Description for Task 2', 'Doing'),
  ('Task 3', 'Description for Task 3', 'Done'),
  ('Task 4', 'Description for Task 4', 'Done'),
  ('Task 5', 'Description for Task 5', 'Deleted'),
  ('Task 6', 'Description for Task 6', 'Doing'),
  ('Task 7', 'Description for Task 7', 'Done'),
  ('Task 8', 'Description for Task 8', 'Doing'),
  ('Task 9', 'Description for Task 9', 'Doing'),
  ('Task 10', 'Description for Task 10', 'Done');

----------------------------------------------

CREATE TABLE `todo_user_like_items`(
  `user_id` INT NOT NULL AUTO_INCREMENT,
  `item_id` INT NOT NULL,
  `created_at` DATETIME DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`user_id`, `item_id`)
);

CREATE INDEX idx_item_id ON `todo_user_like_items` (`item_id`);


```
