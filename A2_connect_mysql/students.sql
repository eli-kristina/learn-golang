CREATE DATABASE learn_golang;
USE learn_golang;

CREATE TABLE `students` (
  `id` int(10) UNSIGNED NOT NULL,
  `name` varchar(100) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `age` int(10) UNSIGNED NOT NULL DEFAULT '0',
  `grade` int(10) UNSIGNED NOT NULL DEFAULT '0'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

INSERT INTO `students` (`id`, `name`, `age`, `grade`) VALUES
(1, 'Anna', 17, 1),
(2, 'Billy', 18, 2),
(3, 'Charlie', 16, 2),
(4, 'Donna', 15, 1),
(5, 'Ellie', 18, 2);

ALTER TABLE `students` ADD PRIMARY KEY(`id`);