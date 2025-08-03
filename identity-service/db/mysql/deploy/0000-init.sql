BEGIN;


CREATE TABLE IF NOT EXISTS `users`
(
    `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
    `safe_id` VARCHAR(36) NOT NULL DEFAULT '' COMMENT 'user identifier key',
    `phone_number` VARCHAR(50) NOT NULL DEFAULT '' COMMENT 'user phone number',
    `email` VARCHAR(128) NOT NULL DEFAULT '' COMMENT 'user email',
    `pin` VARCHAR(128) NOT NULL DEFAULT '' COMMENT 'user pin',
    `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,

    PRIMARY KEY (`id`),
    UNIQUE KEY `uk_user_identity` (`safe_id`, `phone_number`, `email`)
) DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_unicode_ci;


COMMIT;