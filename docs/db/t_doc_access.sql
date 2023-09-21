-- MYSQL
CREATE DATABASE doc CHARACTER SET = 'utf8mb4';

-- doc access table
CREATE TABLE `t_doc_access` (
    `id` INT UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY,
    `create_time` DATETIME(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3) COMMENT '创建时间',
    `url` VARCHAR(200) NOT NULL COMMENT '文章url',
    `num` INT UNSIGNED NOT NULL default '0' COMMENT '访问次数',
    `status` TINYINT UNSIGNED NOT NULL default '0' COMMENT '状态'
);
