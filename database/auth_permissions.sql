-- ============================================================
-- BI系统认证与权限模块数据库表结构
-- 版本: 2.1（简化版）
-- 创建时间: 2026-03-19
-- 更新: 移除密码过期策略字段
-- 描述: 简化版权限设计，适合数据查看为主的BI系统
-- 特性:
--   - 纯JWT无状态认证，无需会话表
--   - 菜单级权限控制，无需列级数据权限
--   - 5张核心表，结构清晰简单
--   - 移除密码过期策略，仅保留基本防暴力破解（登录失败次数限制）
-- ============================================================

-- 创建数据库（如果不存在）
CREATE DATABASE IF NOT EXISTS bi_system DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;

USE bi_system;

-- ============================================================
-- 表1: 用户表 (users)
-- 存储系统用户信息
-- ============================================================
DROP TABLE IF EXISTS `users`;
CREATE TABLE `users` (
  `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '用户ID',
  `username` VARCHAR(50) NOT NULL COMMENT '用户名',
  `password_hash` VARCHAR(255) NOT NULL COMMENT '密码哈希',
  `email` VARCHAR(100) NOT NULL COMMENT '邮箱',
  `phone` VARCHAR(20) DEFAULT NULL COMMENT '手机号',
  `avatar` VARCHAR(255) DEFAULT NULL COMMENT '头像URL',
  `real_name` VARCHAR(50) DEFAULT NULL COMMENT '真实姓名',
  `department` VARCHAR(100) DEFAULT NULL COMMENT '部门',
  `role_id` BIGINT UNSIGNED NOT NULL COMMENT '角色ID',
  `status` TINYINT NOT NULL DEFAULT 1 COMMENT '状态: 0-禁用, 1-正常',
  `is_admin` TINYINT NOT NULL DEFAULT 0 COMMENT '是否管理员: 0-否, 1-是',
  `last_login_at` DATETIME DEFAULT NULL COMMENT '最后登录时间',
  `last_login_ip` VARCHAR(45) DEFAULT NULL COMMENT '最后登录IP',
  `login_attempts` INT DEFAULT 0 COMMENT '连续登录失败次数',
  `lock_until` DATETIME DEFAULT NULL COMMENT '账户锁定截止时间',
  `created_by` BIGINT UNSIGNED DEFAULT NULL COMMENT '创建者ID',
  `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `deleted_at` DATETIME DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uk_username` (`username`),
  UNIQUE KEY `uk_email` (`email`),
  KEY `idx_role_id` (`role_id`),
  KEY `idx_status` (`status`),
  KEY `idx_department` (`department`),
  KEY `idx_created_at` (`created_at`),
  CONSTRAINT `fk_users_role` FOREIGN KEY (`role_id`) REFERENCES `roles` (`id`) ON DELETE RESTRICT ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='用户表';

-- ============================================================
-- 表2: 角色表 (roles)
-- 定义系统角色
-- ============================================================
DROP TABLE IF EXISTS `roles`;
CREATE TABLE `roles` (
  `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '角色ID',
  `name` VARCHAR(50) NOT NULL COMMENT '角色名称',
  `code` VARCHAR(50) NOT NULL COMMENT '角色代码(唯一标识)',
  `description` VARCHAR(255) DEFAULT NULL COMMENT '角色描述',
  `status` TINYINT NOT NULL DEFAULT 1 COMMENT '状态: 0-禁用, 1-正常',
  `is_system` TINYINT NOT NULL DEFAULT 0 COMMENT '是否系统角色: 0-否, 1-是',
  `sort_order` INT DEFAULT 0 COMMENT '排序号',
  `parent_id` BIGINT UNSIGNED DEFAULT NULL COMMENT '父级角色ID',
  `created_by` BIGINT UNSIGNED DEFAULT NULL COMMENT '创建者ID',
  `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `deleted_at` DATETIME DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uk_code` (`code`),
  KEY `idx_status` (`status`),
  KEY `idx_parent_id` (`parent_id`),
  KEY `idx_sort_order` (`sort_order`),
  CONSTRAINT `fk_roles_parent` FOREIGN KEY (`parent_id`) REFERENCES `roles` (`id`) ON DELETE SET NULL ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='角色表';

-- ============================================================
-- 表3: 权限表 (permissions)
-- 定义系统权限
-- ============================================================
DROP TABLE IF EXISTS `permissions`;
CREATE TABLE `permissions` (
  `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '权限ID',
  `name` VARCHAR(100) NOT NULL COMMENT '权限名称',
  `code` VARCHAR(100) NOT NULL COMMENT '权限代码(唯一标识)',
  `resource` VARCHAR(100) NOT NULL COMMENT '资源名称',
  `action` VARCHAR(50) NOT NULL COMMENT '操作类型: view, create, update, delete, manage',
  `type` VARCHAR(20) NOT NULL DEFAULT 'api' COMMENT '权限类型: menu, button, api, data',
  `description` VARCHAR(255) DEFAULT NULL COMMENT '权限描述',
  `parent_id` BIGINT UNSIGNED DEFAULT NULL COMMENT '父级权限ID',
  `path` VARCHAR(255) DEFAULT NULL COMMENT '菜单路径或API路径',
  `icon` VARCHAR(100) DEFAULT NULL COMMENT '菜单图标',
  `sort_order` INT DEFAULT 0 COMMENT '排序号',
  `status` TINYINT NOT NULL DEFAULT 1 COMMENT '状态: 0-禁用, 1-正常',
  `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `deleted_at` DATETIME DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uk_code` (`code`),
  KEY `idx_resource_action` (`resource`, `action`),
  KEY `idx_type` (`type`),
  KEY `idx_parent_id` (`parent_id`),
  KEY `idx_status` (`status`),
  KEY `idx_sort_order` (`sort_order`),
  CONSTRAINT `fk_permissions_parent` FOREIGN KEY (`parent_id`) REFERENCES `permissions` (`id`) ON DELETE SET NULL ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='权限表';

-- ============================================================
-- 表4: 角色权限关联表 (role_permissions)
-- 角色与权限的多对多关系
-- ============================================================
DROP TABLE IF EXISTS `role_permissions`;
CREATE TABLE `role_permissions` (
  `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `role_id` BIGINT UNSIGNED NOT NULL COMMENT '角色ID',
  `permission_id` BIGINT UNSIGNED NOT NULL COMMENT '权限ID',
  `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uk_role_permission` (`role_id`, `permission_id`),
  KEY `idx_permission_id` (`permission_id`),
  CONSTRAINT `fk_role_permissions_role` FOREIGN KEY (`role_id`) REFERENCES `roles` (`id`) ON DELETE CASCADE ON UPDATE CASCADE,
  CONSTRAINT `fk_role_permissions_permission` FOREIGN KEY (`permission_id`) REFERENCES `permissions` (`id`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='角色权限关联表';

-- ============================================================
-- 表5: 登录日志表 (login_logs)
-- 安全审计日志（纯JWT认证下用于登录审计）
-- 说明：本系统采用菜单级权限控制，不实现列级数据权限
-- ============================================================
DROP TABLE IF EXISTS `login_logs`;
CREATE TABLE `login_logs` (
  `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '日志ID',
  `user_id` BIGINT UNSIGNED DEFAULT NULL COMMENT '用户ID',
  `username` VARCHAR(50) NOT NULL COMMENT '用户名',
  `email` VARCHAR(100) DEFAULT NULL COMMENT '邮箱',
  `login_type` VARCHAR(20) NOT NULL DEFAULT 'password' COMMENT '登录类型: password, sms, oauth, ldap',
  `status` VARCHAR(20) NOT NULL COMMENT '登录状态: success, failed, locked, expired',
  `ip_address` VARCHAR(45) NOT NULL COMMENT 'IP地址',
  `user_agent` TEXT DEFAULT NULL COMMENT '用户代理字符串',
  `device_type` VARCHAR(50) DEFAULT NULL COMMENT '设备类型',
  `browser` VARCHAR(100) DEFAULT NULL COMMENT '浏览器',
  `os` VARCHAR(100) DEFAULT NULL COMMENT '操作系统',
  `location` VARCHAR(255) DEFAULT NULL COMMENT '登录地点',
  `error_message` VARCHAR(255) DEFAULT NULL COMMENT '错误信息',
  `login_duration` INT DEFAULT NULL COMMENT '登录耗时(毫秒)',
  `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '登录时间',
  PRIMARY KEY (`id`),
  KEY `idx_user_id` (`user_id`),
  KEY `idx_username` (`username`),
  KEY `idx_status` (`status`),
  KEY `idx_ip_address` (`ip_address`),
  KEY `idx_created_at` (`created_at`),
  KEY `idx_login_type` (`login_type`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='登录日志表';

-- ============================================================
-- 初始化数据
-- ============================================================

-- ----------------------------
-- 1. 初始化角色数据
-- ----------------------------
INSERT INTO `roles` (`id`, `name`, `code`, `description`, `status`, `is_system`, `sort_order`, `parent_id`, `created_by`, `created_at`) VALUES
(1, '超级管理员', 'super_admin', '系统超级管理员，拥有所有权限', 1, 1, 1, NULL, NULL, NOW()),
(2, '系统管理员', 'admin', '系统管理员，管理用户和角色', 1, 1, 2, NULL, NULL, NOW()),
(3, '数据分析师', 'analyst', '数据分析师，可以查询和分析数据', 1, 1, 3, NULL, NULL, NOW()),
(4, '普通用户', 'user', '普通用户，查看仪表盘和报表', 1, 1, 4, NULL, NULL, NOW()),
(5, '访客', 'guest', '访客，仅查看公开内容', 1, 1, 5, NULL, NULL, NOW());

-- ----------------------------
-- 2. 初始化权限数据
-- ----------------------------
INSERT INTO `permissions` (`id`, `name`, `code`, `resource`, `action`, `type`, `description`, `parent_id`, `path`, `icon`, `sort_order`, `status`, `created_at`) VALUES
-- 系统管理模块
(1, '系统管理', 'system', 'system', 'view', 'menu', '系统管理模块', NULL, '/system', 'Setting', 1, 1, NOW()),
(2, '用户管理', 'user', 'user', 'view', 'menu', '用户管理菜单', 1, '/system/user', 'User', 10, 1, NOW()),
(3, '用户列表', 'user:list', 'user', 'view', 'button', '查看用户列表', 2, NULL, NULL, 11, 1, NOW()),
(4, '创建用户', 'user:create', 'user', 'create', 'button', '创建新用户', 2, NULL, NULL, 12, 1, NOW()),
(5, '编辑用户', 'user:update', 'user', 'update', 'button', '编辑用户信息', 2, NULL, NULL, 13, 1, NOW()),
(6, '删除用户', 'user:delete', 'user', 'delete', 'button', '删除用户', 2, NULL, NULL, 14, 1, NOW()),
(7, '重置密码', 'user:reset_password', 'user', 'update', 'button', '重置用户密码', 2, NULL, NULL, 15, 1, NOW()),

-- 角色管理模块
(8, '角色管理', 'role', 'role', 'view', 'menu', '角色管理菜单', 1, '/system/role', 'Shield', 20, 1, NOW()),
(9, '角色列表', 'role:list', 'role', 'view', 'button', '查看角色列表', 8, NULL, NULL, 21, 1, NOW()),
(10, '创建角色', 'role:create', 'role', 'create', 'button', '创建新角色', 8, NULL, NULL, 22, 1, NOW()),
(11, '编辑角色', 'role:update', 'role', 'update', 'button', '编辑角色信息', 8, NULL, NULL, 23, 1, NOW()),
(12, '删除角色', 'role:delete', 'role', 'delete', 'button', '删除角色', 8, NULL, NULL, 24, 1, NOW()),
(13, '分配权限', 'role:assign_permission', 'role', 'update', 'button', '为角色分配权限', 8, NULL, NULL, 25, 1, NOW()),

-- 权限管理模块
(14, '权限管理', 'permission', 'permission', 'view', 'menu', '权限管理菜单', 1, '/system/permission', 'Key', 30, 1, NOW()),
(15, '权限列表', 'permission:list', 'permission', 'view', 'button', '查看权限列表', 14, NULL, NULL, 31, 1, NOW()),
(16, '创建权限', 'permission:create', 'permission', 'create', 'button', '创建新权限', 14, NULL, NULL, 32, 1, NOW()),
(17, '编辑权限', 'permission:update', 'permission', 'update', 'button', '编辑权限信息', 14, NULL, NULL, 33, 1, NOW()),
(18, '删除权限', 'permission:delete', 'permission', 'delete', 'button', '删除权限', 14, NULL, NULL, 34, 1, NOW()),

-- 数据源管理模块
(19, '数据源管理', 'datasource', 'datasource', 'view', 'menu', '数据源管理菜单', NULL, '/datasource', 'Database', 40, 1, NOW()),
(20, '数据源列表', 'datasource:list', 'datasource', 'view', 'button', '查看数据源列表', 19, NULL, NULL, 41, 1, NOW()),
(21, '创建数据源', 'datasource:create', 'datasource', 'create', 'button', '创建新数据源', 19, NULL, NULL, 42, 1, NOW()),
(22, '编辑数据源', 'datasource:update', 'datasource', 'update', 'button', '编辑数据源', 19, NULL, NULL, 43, 1, NOW()),
(23, '删除数据源', 'datasource:delete', 'datasource', 'delete', 'button', '删除数据源', 19, NULL, NULL, 44, 1, NOW()),
(24, '测试连接', 'datasource:test', 'datasource', 'manage', 'button', '测试数据源连接', 19, NULL, NULL, 45, 1, NOW()),

-- 数据查询模块
(25, '数据查询', 'query', 'query', 'view', 'menu', '数据查询菜单', NULL, '/query', 'Code', 50, 1, NOW()),
(26, '执行查询', 'query:execute', 'query', 'create', 'button', '执行SQL查询', 25, NULL, NULL, 51, 1, NOW()),
(27, '导出数据', 'query:export', 'query', 'manage', 'button', '导出查询结果', 25, NULL, NULL, 52, 1, NOW()),

-- 仪表盘管理模块
(28, '仪表盘管理', 'dashboard', 'dashboard', 'view', 'menu', '仪表盘管理菜单', NULL, '/dashboard', 'Dashboard', 60, 1, NOW()),
(29, '仪表盘列表', 'dashboard:list', 'dashboard', 'view', 'button', '查看仪表盘列表', 28, NULL, NULL, 61, 1, NOW()),
(30, '创建仪表盘', 'dashboard:create', 'dashboard', 'create', 'button', '创建新仪表盘', 28, NULL, NULL, 62, 1, NOW()),
(31, '编辑仪表盘', 'dashboard:update', 'dashboard', 'update', 'button', '编辑仪表盘', 28, NULL, NULL, 63, 1, NOW()),
(32, '删除仪表盘', 'dashboard:delete', 'dashboard', 'delete', 'button', '删除仪表盘', 28, NULL, NULL, 64, 1, NOW()),
(33, '分享仪表盘', 'dashboard:share', 'dashboard', 'manage', 'button', '分享仪表盘', 28, NULL, NULL, 65, 1, NOW()),

-- 图表管理模块
(34, '图表管理', 'chart', 'chart', 'view', 'menu', '图表管理菜单', NULL, '/chart', 'PieChart', 70, 1, NOW()),
(35, '图表列表', 'chart:list', 'chart', 'view', 'button', '查看图表列表', 34, NULL, NULL, 71, 1, NOW()),
(36, '创建图表', 'chart:create', 'chart', 'create', 'button', '创建新图表', 34, NULL, NULL, 72, 1, NOW()),
(37, '编辑图表', 'chart:update', 'chart', 'update', 'button', '编辑图表', 34, NULL, NULL, 73, 1, NOW()),
(38, '删除图表', 'chart:delete', 'chart', 'delete', 'button', '删除图表', 34, NULL, NULL, 74, 1, NOW()),

-- 系统设置模块
(39, '系统设置', 'settings', 'settings', 'view', 'menu', '系统设置菜单', 1, '/system/settings', 'Tools', 40, 1, NOW()),
(40, '操作日志', 'logs', 'logs', 'view', 'menu', '操作日志菜单', 1, '/system/logs', 'Document', 50, 1, NOW()),
(41, '查看日志', 'logs:list', 'logs', 'view', 'button', '查看操作日志', 40, NULL, NULL, 51, 1, NOW());

-- ----------------------------
-- 3. 初始化角色权限关联
-- ----------------------------
-- 超级管理员拥有所有权限
INSERT INTO `role_permissions` (`role_id`, `permission_id`, `created_at`)
SELECT 1, `id`, NOW() FROM `permissions` WHERE `status` = 1;

-- 系统管理员拥有除了权限管理之外的所有权限
INSERT INTO `role_permissions` (`role_id`, `permission_id`, `created_at`)
SELECT 2, `id`, NOW() FROM `permissions` WHERE `status` = 1 AND `id` NOT IN (15, 16, 17, 18);

-- 数据分析师拥有查询、仪表盘、图表相关权限
INSERT INTO `role_permissions` (`role_id`, `permission_id`, `created_at`)
SELECT 3, `id`, NOW() FROM `permissions`
WHERE `status` = 1 AND (
  `resource` IN ('query', 'dashboard', 'chart', 'datasource')
  OR `code` LIKE 'dashboard:%'
  OR `code` LIKE 'chart:%'
);

-- 普通用户拥有查看和创建仪表盘的权限
INSERT INTO `role_permissions` (`role_id`, `permission_id`, `created_at`)
SELECT 4, `id`, NOW() FROM `permissions`
WHERE `status` = 1 AND (
  `type` = 'menu' AND `resource` IN ('dashboard', 'chart')
  OR `code` IN ('dashboard:list', 'dashboard:view', 'chart:list', 'chart:view', 'query:execute', 'query:export')
);

-- 访客只有查看权限
INSERT INTO `role_permissions` (`role_id`, `permission_id`, `created_at`)
SELECT 5, `id`, NOW() FROM `permissions`
WHERE `status` = 1 AND (
  `type` = 'menu' AND `resource` IN ('dashboard', 'chart')
  OR `code` IN ('dashboard:list', 'chart:list')
);

-- ----------------------------
-- 4. 创建默认管理员用户
-- ----------------------------
-- 密码: admin123 (BCrypt加密)
INSERT INTO `users` (
  `id`, `username`, `password_hash`, `email`, `phone`, `real_name`,
  `department`, `role_id`, `status`, `is_admin`, `created_by`, `created_at`
) VALUES
(1, 'admin', '$2a$10$N.zmdr9k7uOCQb376NoUnuTJ8iAt6Z5EHsM8lE9lBOsl7iAt6Z5EH', 'admin@bi-system.com', '13800138000', '系统管理员', '技术部', 1, 1, 1, NULL, NOW());

-- 创建测试用户
INSERT INTO `users` (
  `id`, `username`, `password_hash`, `email`, `phone`, `real_name`,
  `department`, `role_id`, `status`, `is_admin`, `created_by`, `created_at`
) VALUES
(2, 'analyst', '$2a$10$N.zmdr9k7uOCQb376NoUnuTJ8iAt6Z5EHsM8lE9lBOsl7iAt6Z5EH', 'analyst@bi-system.com', '13800138001', '数据分析师', '数据分析部', 3, 1, 0, 1, NOW()),
(3, 'user', '$2a$10$N.zmdr9k7uOCQb376NoUnuTJ8iAt6Z5EHsM8lE9lBOsl7iAt6Z5EH', 'user@bi-system.com', '13800138002', '普通用户', '市场部', 4, 1, 0, 1, NOW());

-- ----------------------------
-- 5. 创建索引优化（定期清理登录日志）
-- ----------------------------
-- 建议定期执行: DELETE FROM login_logs WHERE created_at < DATE_SUB(NOW(), INTERVAL 90 DAY);

-- ----------------------------
-- 6. 视图定义（方便查询）
-- ----------------------------
-- 用户权限视图
CREATE OR REPLACE VIEW `v_user_permissions` AS
SELECT
  u.id AS user_id,
  u.username,
  r.id AS role_id,
  r.name AS role_name,
  r.code AS role_code,
  p.id AS permission_id,
  p.code AS permission_code,
  p.resource,
  p.action,
  p.type
FROM users u
INNER JOIN roles r ON u.role_id = r.id
INNER JOIN role_permissions rp ON r.id = rp.role_id
INNER JOIN permissions p ON rp.permission_id = p.id
WHERE u.status = 1 AND r.status = 1 AND p.status = 1;

-- ============================================================
-- 注释说明
-- ============================================================
-- 1. 所有表使用InnoDB引擎，支持事务和外键
-- 2. 字符集使用utf8mb4，支持emoji和特殊字符
-- 3. 所有时间字段使用DATETIME类型，存储精确到秒
-- 4. 软删除字段 deleted_at，支持数据恢复
-- 5. 密码使用BCrypt加密，默认密码 admin123
-- 6. 默认管理员账号: admin / admin123
-- 7. 建议定期清理90天前的登录日志: DELETE FROM login_logs WHERE created_at < DATE_SUB(NOW(), INTERVAL 90 DAY);

-- ============================================================
-- 使用说明
-- ============================================================
-- 1. 认证方式：纯JWT无状态认证，登录时签发JWT，后续请求验证JWT即可
-- 2. 权限控制：菜单级权限，通过 roles -> permissions 控制菜单/按钮可见性
-- 3. 默认账号：admin / admin123
-- 4. 定期维护：定期清理登录日志（建议90天）
-- ============================================================
