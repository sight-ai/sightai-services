CREATE TABLE `account`
(
    `id`             BIGINT UNSIGNED auto_increment,
    `address`        VARCHAR(255) NOT NULL,
    `role`           VARCHAR(255) NOT NULL,
    `nonce`          BIGINT UNSIGNED DEFAULT 0,
    `hold`           DECIMAL(36, 18) NOT NULL DEFAULT '0.000000000000000000',
    `available`      DECIMAL(36, 18) NOT NULL DEFAULT '0.000000000000000000',
    `created_at`     DATETIME NOT NULL,
    `updated_at`     DATETIME NOT NULL,
    `deleted_at`     DATETIME DEFAULT NULL,
    PRIMARY KEY (`id`),
    UNIQUE KEY `idx_address` (`address`)
)
    engine = innodb
    charset = utf8mb4
    COLLATE utf8mb4_unicode_ci;

CREATE TABLE `gateway`
(
    `id`             BIGINT UNSIGNED auto_increment,
    `address`        VARCHAR(255) NOT NULL,
    `endpoint`       VARCHAR(255) NOT NULL,
    `name`           VARCHAR(255) NOT NULL,
    `created_at`     DATETIME NOT NULL,
    `updated_at`     DATETIME NOT NULL,
    `deleted_at`     DATETIME DEFAULT NULL,
    PRIMARY KEY (`id`),
    UNIQUE KEY `idx_address` (`address`),
    UNIQUE KEY `idx_name` (`name`)
)
    engine = innodb
    charset = utf8mb4
    COLLATE utf8mb4_unicode_ci;

CREATE TABLE `transaction`
(
    `id`         BIGINT UNSIGNED auto_increment,
    `account_id` BIGINT UNSIGNED NOT NULL,
    `available`  DECIMAL(36, 18) NOT NULL DEFAULT '0.000000000000000000',
    `hold`       DECIMAL(36, 18) NOT NULL DEFAULT '0.000000000000000000',
    `type`       VARCHAR(255) NOT NULL,
    `notes`      VARCHAR(255) DEFAULT NULL,
    `created_at` DATETIME NOT NULL,
    `updated_at` DATETIME NOT NULL,
    `deleted_at` DATETIME DEFAULT NULL,
    PRIMARY KEY (`id`),
    KEY `idx_account_id` (`account_id`)
)
    engine = innodb
    charset = utf8mb4
    COLLATE utf8mb4_unicode_ci;

CREATE TABLE `receipt`
(
    `id`                BIGINT UNSIGNED auto_increment,
    `user_id`           VARCHAR(255) NOT NULL,
    `gateway_id`        VARCHAR(255) NOT NULL,
    `finished_at`       DATETIME DEFAULT NULL,
    `cost`              DECIMAL(36, 18) NOT NULL DEFAULT '0.000000000000000000',
    `proof`             VARCHAR(255) NOT NULL,
    `txn_id`            VARCHAR(255) NOT NULL,
    `status`            VARCHAR(255) NOT NULL,
    `created_at`        DATETIME NOT NULL,
    `updated_at`        DATETIME NOT NULL,
    `deleted_at`        DATETIME DEFAULT NULL,
    PRIMARY KEY (`id`),
    KEY `idx_txn_id` (`txn_id`)
)
    engine = innodb
    charset = utf8mb4
    COLLATE utf8mb4_unicode_ci;

CREATE TABLE `deposit_event`
(
    `id`                BIGINT UNSIGNED auto_increment,
    `txn_hash`          VARCHAR(255) NOT NULL,
    `block_timestamp`   DATETIME NOT NULL,
    `from_addr`         VARCHAR(255) NOT NULL,
    `to_addr`           VARCHAR(255) NOT NULL,
    `amount`            VARCHAR(255) NOT NULL,
    `created_at`        DATETIME NOT NULL,
    `updated_at`        DATETIME NOT NULL,
    `deleted_at`        DATETIME DEFAULT NULL,
    PRIMARY KEY (`id`),
    UNIQUE KEY `idx_txn_hash` (`txn_hash`)
)
    engine = innodb
    charset = utf8mb4
    COLLATE utf8mb4_unicode_ci;

CREATE TABLE `withdraw_event`
(
    `id`                BIGINT UNSIGNED auto_increment,
    `txn_hash`          VARCHAR(255) NOT NULL,
    `block_timestamp`   DATETIME NOT NULL,
    `to_addr`           VARCHAR(255) NOT NULL,
    `amount`            VARCHAR(255) NOT NULL,
    `nonce`             BIGINT UNSIGNED DEFAULT 0,
    `created_at`        DATETIME NOT NULL,
    `updated_at`        DATETIME NOT NULL,
    `deleted_at`        DATETIME DEFAULT NULL,
    PRIMARY KEY (`id`),
    UNIQUE KEY `idx_txn_hash` (`txn_hash`),
    UNIQUE KEY `idx_to_addr_nonce` (`to_addr`, `nonce`)
)
    engine = innodb
    charset = utf8mb4
    COLLATE utf8mb4_unicode_ci;

CREATE TABLE `allowance`
(
    `id`                BIGINT UNSIGNED auto_increment,
    `from_account_id`   BIGINT UNSIGNED NOT NULL,
    `to_account_id`     BIGINT UNSIGNED NOT NULL,
    `allowance`         DECIMAL(36, 18) NOT NULL DEFAULT '0.000000000000000000',
    `version`           BIGINT UNSIGNED NOT NULL DEFAULT 0,
    `created_at`        DATETIME NOT NULL,
    `updated_at`        DATETIME NOT NULL,
    `deleted_at`        DATETIME DEFAULT NULL,
    PRIMARY KEY (`id`),
    UNIQUE KEY `idx_from_to` (`from_account_id`, `to_account_id`)
)
    engine = innodb
    charset = utf8mb4
    COLLATE utf8mb4_unicode_ci;

