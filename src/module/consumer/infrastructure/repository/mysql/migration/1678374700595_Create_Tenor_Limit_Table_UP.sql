use consumer;

CREATE TABLE `tenor_limit` (
    `consumer_id` varchar(500) NOT NULL,
    `month` int NOT NULL,
    `limit_value` double NOT NULL,
    
    `created_at` bigint unsigned DEFAULT 0,
    `updated_at` bigint unsigned DEFAULT 0,
    `deleted_at` bigint unsigned DEFAULT 0,
    UNIQUE KEY `unique_limit` (`consumer_id`,`month`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
