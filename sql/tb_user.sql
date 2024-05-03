-- easy_blog_db.tb_user definition

CREATE TABLE `tb_user` (
                           `id` int unsigned NOT NULL COMMENT '用户id',
                           `user_name` varchar(20) NOT NULL COMMENT '用户名',
                           `password` varchar(100) DEFAULT NULL COMMENT '用户密码',
                           `avatar` varchar(255) DEFAULT NULL COMMENT '用户头像',
                           `email` varchar(350) DEFAULT NULL COMMENT '用户邮箱',
                           `tel` varchar(20) DEFAULT NULL COMMENT '用户手机号',
                           `foreign_id` varchar(100) NOT NULL COMMENT '其他平台的唯一id',
                           `ip` varchar(20) DEFAULT NULL COMMENT 'ip地址',
                           `role` tinyint NOT NULL COMMENT '用户角色',
                           `sign_origin` tinyint DEFAULT NULL COMMENT '注册来源',
                           `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
                           `update_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '更新时间',
                           `deleted_at` datetime DEFAULT NULL COMMENT '删除时间',
                           `status` tinyint DEFAULT NULL COMMENT '用户状态',
                           PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='easy_blog的用户信息表';