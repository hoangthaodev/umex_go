-- +goose Up
-- +goose StatementBegin
CREATE table IF NOT EXISTS `tb_users` (
    `usr_id` INT(11) unsigned not null AUTO_INCREMENT,
    `usr_email` VARCHAR(30) NOT NULL default '',
    `usr_phone` VARCHAR(15) NOT NULL default '',
    `usr_username` VARCHAR(30) NOT NULL default '',
    `usr_password` VARCHAR(32) NOT NULL default '',
    `usr_created_at` int(11) NOT NULL DEFAULT '0',
    `usr_updated_at` int(11) NOT NULL DEFAULT '0',
    `usr_create_ip_at` VARCHAR(12) not NULL DEFAULT '',
    `usr_last_login_at` int(11) NOT NULL DEFAULT '0',
    `usr_last_login_ip_at` VARCHAR(12) NOT NULL DEFAULT '',
    `usr_login_times` int(11) NOT NULL DEFAULT '0',
    `usr_status` tinyint(1) NOT NULL DEFAULT '0' COMMENT 'Status 1: enabled, 0: disabled, -1: deleted',
    PRIMARY KEY (`usr_id`),
    KEY `idx_email` (`usr_email`),
    KEY `idx_phone` (`usr_phone`),
    KEY `idx_username` (`usr_username`)
    ) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS `tb_users`;
-- +goose StatementEnd
