-- テーブル定義
CREATE TABLE IF NOT EXISTS `users` (
  `id`         int unsigned NOT NULL AUTO_INCREMENT PRIMARY KEY COMMENT 'id',
  `name`       varchar(255) NOT NULL                            COMMENT 'name',

  `created_at` datetime     DEFAULT NOW()                       COMMENT 'create datetime',
  `updated_at` datetime     DEFAULT NULL                        COMMENT 'last update datetime',
  `deleted_at` datetime     DEFAULT NULL                        COMMENT 'delete datetime'
) ENGINE=InnoDB COMMENT='users';

insert into users (name) values
('test user');