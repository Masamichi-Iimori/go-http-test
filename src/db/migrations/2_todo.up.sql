CREATE TABLE `todo` (
  `id` bigint(2) NOT NULL AUTO_INCREMENT,
  `task` varchar(20) NOT NULL,
  `limit_date` datetime NOT NULL,
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL,
  `deleted_at` datetime,
  PRIMARY KEY (`id`)
) DEFAULT CHARSET=utf8mb4;