use consumer;

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
