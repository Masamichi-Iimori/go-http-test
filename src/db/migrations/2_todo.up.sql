CREATE TABLE `todos` (
  `id` bigint(2) NOT NULL AUTO_INCREMENT,
  `task` varchar(20) NOT NULL,
  `status` boolean NOT NULL,
  `limitDate` datetime NOT NULL,
  `createdAt` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updatedAt` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `deletedAt` datetime,
  PRIMARY KEY (`id`)
) DEFAULT CHARSET=utf8mb4;