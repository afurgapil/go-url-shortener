CREATE TABLE `urls` (
  `id` int NOT NULL AUTO_INCREMENT,
  `long_url` varchar(256) NOT NULL,
  `short_url` varchar(256) NOT NULL,
  `usage_count` int NOT NULL DEFAULT '0',
  `pass_key` varchar(32) NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `id_UNIQUE` (`id`),
  UNIQUE KEY `pass_key_UNIQUE` (`pass_key`)
) ENGINE=InnoDB AUTO_INCREMENT=11 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
