CREATE TABLE IF NOT EXISTS `books` (
  `id` char(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '对象Id',
  `create_at` bigint NOT NULL COMMENT '创建时间(13位时间戳)',
  `create_by` varchar(255) COLLATE utf8mb4_general_ci NOT NULL COMMENT '创建人',
  `update_at` bigint NOT NULL COMMENT '更新时间',
  `update_by` varchar(255) COLLATE utf8mb4_general_ci NOT NULL COMMENT '更新人',
  `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '书名',
  `author` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '作者',
  PRIMARY KEY (`id`),
  KEY `idx_name` (`name`) USING BTREE COMMENT '用于书名搜索',
  KEY `idx_author` (`author`) USING BTREE COMMENT '用于作者搜索'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

CREATE TABLE IF NOT EXISTS `resource` (
    `id` char(64) NOT NULL COMMENT '对象Id',
    `resource_type` tinyint(2) NOT NULL,
    `vendor` tinyint(1) NOT NULL,
    `region` varchar(64) NOT NULL,
    `zone` varchar(64) NOT NULL,
    `create_at` bigint NOT NULL COMMENT '创建时间(13位时间戳)',
    `expire_at` bigint NOT NULL COMMENT '创建时间(13位时间戳)',
    `category`  varchar(64) NOT NULL,
    `type` varchar(120) NOT NULL,
    `name` varchar(255) NOT NULL,
    `description` varchar(512) DEFAULT NULL,
    `status` varchar(255) NOT NULL,
    `update_at` bigint(13) DEFAULT NULL COMMENT '更新时间',
    `sync_at` bigint(13) DEFAULT NULL COMMENT '更新时间',
    `sync_account` varchar(255) DEFAULT NULL,
    `public_ip` varchar(64) DEFAULT NULL,
    `private_ip` varchar(64) DEFAULT NULL,
    `pay_type` varchar(255) DEFAULT NULL,
    `describe_hash` varchar(255) NOT NULL,
    `resource_hash` varchar(255) NOT NULL,
    `secret_id` varchar(64) NOT NULL,
    `domain` varchar(255) NOT NULL,
    `namespace` varchar(255) NOT NULL,
    `env` varchar(255) NOT NULL,
    `usage_mode` tinyint(2) NOT NULL,
    PRIMARY KEY (`id`) USING BTREE,
    KEY `idx_name` (`name`) USING BTREE,
    KEY `idx_status` (`status`) USING BTREE,
    KEY `idx_private_ip` (`private_ip`) USING BTREE,
    KEY `idx_public_ip` (`public_ip`) USING BTREE,
    KEY `idx_domain` (`domain`) USING HASH,
    KEY `idx_namespace` (`namespace`) USING HASH,
    KEY `idx_env` (`env`) USING HASH
    ) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

CREATE TABLE IF NOT EXISTS `resource_tag` (
    `id` char(64) NOT NULL,
    `t_key` varchar(255) NOT NULL,
    `t_value` varchar(255) NOT NULL,
    `description` varchar(255) NOT NULL,
    `resource_id` varchar(255) CHARACTER SET latin1 NOT NULL,
    `weight` int(11) NOT NULL,
    `type` tinyint(4) NOT NULL,
    `create_at` bigint NOT NULL,
    PRIMARY KEY (`id`),
    UNIQUE KEY `idx_id` (`t_key`, `t_value`, `resource_id`),
    KEY `idx_t_key` (`t_key`) USING HASH,
    KEY `idx_t_value` (`t_value`) USING HASH,
    KEY `idx_resource_id` (`resource_id`) USING HASH
    ) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

CREATE TABLE IF NOT EXISTS `resource_host` (
    `resource_id` varchar(64) CHARACTER SET latin1 NOT NULL COMMENT '关联的资源ID',
    `cpu` tinyint(4) NOT NULL COMMENT 'cpu核算',
    `memory` int(13) NOT NULL COMMENT '内存大小',
    `gpu_amount` tinyint(4) DEFAULT NULL COMMENT 'gpu核数',
    `gpu_spec` varchar(255) CHARACTER SET latin1 DEFAULT NULL COMMENT 'gpu规格',
    `os_type` varchar(255) CHARACTER SET latin1 DEFAULT NULL COMMENT '操作系统类型',
    `os_name` varchar(255) CHARACTER SET utf8mb4 DEFAULT NULL,
    `serial_number` varchar(120) CHARACTER SET latin1 DEFAULT NULL COMMENT '系统序列号',
    `image_id` char(64) CHARACTER SET latin1 DEFAULT NULL,
    `internet_max_bandwidth_out` int(10) DEFAULT NULL,
    `internet_max_bandwidth_in` int(10) DEFAULT NULL,
    `key_pair_name` varchar(255) DEFAULT NULL,
    `security_groups` varchar(255) DEFAULT NULL,
    PRIMARY KEY (`resource_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

CREATE TABLE IF NOT EXISTS `secret` (
    `id` varchar(64) NOT NULL COMMENT '凭证Id',
    `create_at` bigint(13) NOT NULL COMMENT '创建时间',
    `description` varchar(255) NOT NULL COMMENT '凭证描述',
    `vendor` tinyint(1) NOT NULL COMMENT '资源提供商',
    `address` varchar(255)  NOT NULL COMMENT '体验提供方访问地址',
    `allow_regions` text  NOT NULL COMMENT '允许同步的Region列表',
    `crendential_type` tinyint(1) NOT NULL COMMENT '凭证类型',
    `api_key` varchar(255) NOT NULL COMMENT '凭证key',
    `api_secret` text  NOT NULL COMMENT '凭证secret',
    `request_rate` int(11) NOT NULL COMMENT '请求速率',
    PRIMARY KEY (`id`),
    UNIQUE KEY `idx_key` (`api_key`) USING BTREE
    ) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='资源提供商同步凭证管理';

CREATE TABLE IF NOT EXISTS `task` (
    `id` varchar(64) NOT NULL COMMENT '任务Id',
    `region` varchar(64) NOT NULL COMMENT '资源所属Region',
    `resource_type` tinyint(1) NOT NULL COMMENT '资源类型',
    `secret_id` varchar(64) NOT NULL COMMENT '用于操作资源的凭证Id',
    `secret_desc` text NOT NULL COMMENT '凭证描述',
    `timeout` int(11) NOT NULL COMMENT '任务超时时间',
    `status` tinyint(1) NOT NULL COMMENT '任务当前状态',
    `message` text NOT NULL COMMENT '任务失败相关信息',
    `start_at` bigint(20) NOT NULL COMMENT '任务开始时间',
    `end_at` bigint(20) NOT NULL COMMENT '任务结束时间',
    `total_succeed` int(11) NOT NULL COMMENT '总共操作成功的资源数量',
    `total_failed` int(11) NOT NULL COMMENT '总共操作失败的资源数量',
    PRIMARY KEY (`id`)
    ) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='资源操作任务管理';