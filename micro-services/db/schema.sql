CREATE DATABASE IF NOT EXISTS counter_db;
USE counter_db;
CREATE TABLE IF NOT EXISTS counter (
    id INT AUTO_INCREMENT PRIMARY KEY,
    count INT NOT NULL DEFAULT 0
);
INSERT INTO counter (count)
VALUES (0);