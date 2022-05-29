-- テーブル定義
CREATE TABLE IF NOT EXISTS `users` (
  `id`         varchar(40)  NOT NULL DEFAULT (UUID()) PRIMARY KEY COMMENT 'id',
  `name`       varchar(255) NOT NULL                              COMMENT 'name',

  `created_at` datetime     DEFAULT CURRENT_TIMESTAMP                             COMMENT 'create datetime',
  `updated_at` datetime     DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT 'last update datetime',
  `deleted_at` datetime     DEFAULT NULL                                          COMMENT 'delete datetime'
) ENGINE=InnoDB COMMENT='users';

insert into users (name) values
('test user');