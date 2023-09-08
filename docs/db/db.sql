

select * from t_host_info where id=(select max(id) from t_host_info);

----------------------------------------
-- table

CREATE TABLE `t_host_info` (
    `id` int unsigned NOT NULL AUTO_INCREMENT,
    `itime` datetime(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3) COMMENT '入库时间',
    `ctime` datetime(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3) COMMENT '采集时间',
    `ip` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '0' COMMENT '主机地址',
    `cpu` float NOT NULL COMMENT 'CPU使用率',
    `random_num` int unsigned NOT NULL,
    `status` tinyint unsigned NOT NULL DEFAULT '0' COMMENT '数据状态',
    PRIMARY KEY (`id`),
    KEY `IDX_HOST_CPU_CTIME` (`ctime`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;



-- 阻塞
alter table t_host_info modify itime DATETIME(3) not null DEFAULT CURRENT_TIMESTAMP(3) COMMENT '入库时间';
alter table t_host_info modify ctime DATETIME(3) not null DEFAULT CURRENT_TIMESTAMP(3) COMMENT '采集时间';

-- 非阻塞
alter table t_host_info add status TINYINT unsigned not null default 0 COMMENT '数据状态';
