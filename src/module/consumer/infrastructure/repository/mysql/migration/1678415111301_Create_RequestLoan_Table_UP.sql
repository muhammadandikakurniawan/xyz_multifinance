use consumer;

CREATE TABLE `request_loan` (
    `id` bigint unsigned NOT NULL AUTO_INCREMENT,
    `consumer_id` varchar(500) NOT NULL,
    `contract_number` varchar(500) NOT NULL,
    `otr` double NOT NULL,
    `admin_fee` double NOT NULL,
    `is_approved` TINYINT DEFAULT NULL,
    `installment` double DEFAULT 0,
    `interest` double DEFAULT 0,
    `asset_name` varchar(500) NOT NULL,
    
    `created_at` bigint unsigned DEFAULT 0,
    `updated_at` bigint unsigned DEFAULT 0,
    `deleted_at` bigint unsigned DEFAULT 0,
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
