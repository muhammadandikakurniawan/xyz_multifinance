CREATE DATABASE consumer

CREATE TABLE `consumer` (
    `id` varchar(500) NOT NULL,

    `nik` varchar(20) NOT NULL,
    `full_name` varchar(100) NOT NULL,
    `legal_name` varchar(100) NOT NULL,
    `place_of_birth` varchar(100) NOT NULL,
    `date_of_birth` bigint unsigned NOT NULL,
    `ktp_image_url` varchar(1000) NOT NULL,
    `selfie_image_url` varchar(1000) NOT NULL,
    `salary` double NOT NULL,
    
    `created_at` bigint unsigned DEFAULT 0,
    `updated_at` bigint unsigned DEFAULT 0,
    `deleted_at` bigint unsigned DEFAULT 0,
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE `tenor_limit` (
    `consumer_id` varchar(500) NOT NULL,
    `month` int NOT NULL,
    `limit_value` double NOT NULL,
    
    `created_at` bigint unsigned DEFAULT 0,
    `updated_at` bigint unsigned DEFAULT 0,
    `deleted_at` bigint unsigned DEFAULT 0,
    UNIQUE KEY `unique_limit` (`consumer_id`,`month`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

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