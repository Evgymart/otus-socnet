CREATE TABLE `users` (
  `id` BIGINT NOT NULL AUTO_INCREMENT PRIMARY KEY,
  `first_name` VARCHAR(255) NOT NULL,
  `last_name` VARCHAR(255) NOT NULL,
  `birthdate` DATE NOT NULL,
  `gender` ENUM('none', 'male', 'female') NOT NULL,
  `email` VARCHAR(255) NOT NULL UNIQUE,
  `password` VARCHAR(225) NOT NULL,
  `biography` TEXT DEFAULT NULL,
  `city` VARCHAR(255) NOT NULL
);
