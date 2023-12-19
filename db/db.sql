ConCREATE DATABASE IF NOT EXISTS `homeconomy`;
USE `homeconomy`;

CREATE TABLE IF NOT EXISTS `users`
(
    `id`   BINARY(16)  NOT NULL,
    `name` varchar(255) NOT NULL,
    PRIMARY KEY (`id`)
    ) ENGINE = InnoDB
    DEFAULT CHARSET = latin1;

CREATE TABLE IF NOT EXISTS `accounts`
(
    `id`      BINARY(16)  NOT NULL,
    `name`    varchar(255) NOT NULL,
    `balance` double       NOT NULL,
    `user_id` BINARY(16)  NOT NULL,
    PRIMARY KEY (`id`),
    FOREIGN KEY (`user_id`) REFERENCES `users` (`id`)
    ) ENGINE = InnoDB
    DEFAULT CHARSET = latin1;

CREATE TABLE IF NOT EXISTS `companies`
(
    `id`   BINARY(16)  NOT NULL,
    `name` varchar(255) NOT NULL,
    PRIMARY KEY (`id`)
    ) ENGINE = InnoDB
    DEFAULT CHARSET = latin1;

CREATE TABLE IF NOT EXISTS `transactions`
(
    `id`          BINARY(16)  NOT NULL,
    `account_id`  BINARY(16)  NOT NULL,
    `amount`      double       NOT NULL,
    `date`        datetime     NOT NULL,
    `description` varchar(255) NOT NULL,
    `company_id`  BINARY(16)  NOT NULL,
    PRIMARY KEY (`id`),
    FOREIGN KEY (`company_id`) REFERENCES `companies` (`id`),
    FOREIGN KEY (`account_id`) REFERENCES `accounts` (`id`)
    ) ENGINE = InnoDB
    DEFAULT CHARSET = latin1;
