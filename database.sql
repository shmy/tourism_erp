DROP DATABASE IF EXISTS `jx_erp`;
CREATE DATABASE `jx_erp` DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_croatian_ci;
USE `jx_erp`;

CREATE TABLE `organization` (
  `id` INT UNSIGNED AUTO_INCREMENT,
  `name` VARCHAR(32) NOT NULL COMMENT '组织名称',
  `describe` VARCHAR(128) NOT NULL COMMENT '组织简介',
  `pid` INT UNSIGNED COMMENT '父ID',
  `created_at` DATETIME NOT NULL,
  `updated_at` DATETIME DEFAULT NULL,
  `deleted_at` DATETIME DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=INNOID COMMENT='组织表';

CREATE TABLE `department` (
  `id` INT UNSIGNED AUTO_INCREMENT,
  `name` VARCHAR(32) NOT NULL COMMENT '部门名称',
  `describe` VARCHAR(128) NOT NULL COMMENT '部门简介',
  `pid` INT UNSIGNED COMMENT '父ID',
  `organization_id` INT UNSIGNED COMMENT '父ID',
  KEY `department_organization_id` (`organization_id`),
  CONSTRAINT `department_organization_id` FOREIGN KEY (`organization_id`) REFERENCES `organization` (`id`) ON DELETE CASCADE,
  `created_at` DATETIME NOT NULL,
  `updated_at` DATETIME DEFAULT NULL,
  `deleted_at` DATETIME DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=INNOID COMMENT='部门表';

CREATE TABLE `people` (
  `id` INT UNSIGNED AUTO_INCREMENT,
  `username` VARCHAR(32) NOT NULL COMMENT '用户名',
  `password` VARCHAR(32) NOT NULL COMMENT '姓名',
  `realname` VARCHAR(32) NOT NULL COMMENT '真名',
  `avatar` VARCHAR(32) NOT NULL COMMENT '头像',
  `phone` VARCHAR(32) NOT NULL COMMENT '手机号码',
  `socket_id` varchar(255) DEFAULT NULL COMMENT 'socketId', # socket_id 预留 用于实时通知
  `locked` TINYINT DEFAULT 0 COMMENT '用户是否被锁定: 0 否 1 是',
  `lock_why` varchar(255) DEFAULT NULL COMMENT '被锁定的原因',
  `created_at` DATETIME NOT NULL,
  `updated_at` DATETIME DEFAULT NULL,
  `deleted_at` DATETIME DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=INNOID COMMENT='人员表';

CREATE TABLE `merge_organization_department_people` (
  `id` INT UNSIGNED AUTO_INCREMENT,
  `organization_id` INT UNSIGNED COMMENT '组织ID',
  `department_id` INT UNSIGNED COMMENT '部门ID',
  `people_id` INT UNSIGNED COMMENT '人员ID',
  `created_at` DATETIME NOT NULL,
  `updated_at` DATETIME DEFAULT NULL,
  `deleted_at` DATETIME DEFAULT NULL,
  # 关联外键
  KEY `merge_organization_department_people_organization_id_foreign` (`organization_id`),
  KEY `merge_organization_department_people_department_id_foreign` (`department_id`),
  KEY `merge_organization_department_people_people_id_foreign` (`people_id`),
  CONSTRAINT `merge_organization_department_people_organization_id_foreign` FOREIGN KEY (`organization_id`) REFERENCES `organization` (`id`) ON DELETE CASCADE,
  CONSTRAINT `merge_organization_department_people_department_id_foreign` FOREIGN KEY (`department_id`) REFERENCES `department` (`id`) ON DELETE CASCADE,
  CONSTRAINT `merge_organization_department_people_people_id_foreign` FOREIGN KEY (`people_id`) REFERENCES `people` (`id`) ON DELETE CASCADE,
  PRIMARY KEY (`id`)
) ENGINE=INNOID COMMENT='组织部门人员中间表';

CREATE TABLE `role` (
  `id` INT UNSIGNED AUTO_INCREMENT,
  `name` VARCHAR(32) NOT NULL COMMENT '名称',
  `describe` VARCHAR(128) NOT NULL COMMENT '描述',
  `created_at` DATETIME NOT NULL,
  `updated_at` DATETIME DEFAULT NULL,
  `deleted_at` DATETIME DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=INNOID COMMENT='角色表';

CREATE TABLE `permissions` (
  `id` INT UNSIGNED AUTO_INCREMENT,
  `name` VARCHAR(32) NOT NULL COMMENT '名称',
  `describe` VARCHAR(128) NOT NULL COMMENT '描述',
  `code` VARCHAR(32) NOT NULL COMMENT '对应的接口匹配代码',
  `created_at` DATETIME NOT NULL,
  `updated_at` DATETIME DEFAULT NULL,
  `deleted_at` DATETIME DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=INNOID COMMENT='权限表';

CREATE TABLE `merge_role_permissions` (
  `id` INT UNSIGNED AUTO_INCREMENT,
  `role_id` INT UNSIGNED COMMENT '角色ID',
  `permissions_id` INT UNSIGNED COMMENT '权限ID',
  `created_at` DATETIME NOT NULL,
  `updated_at` DATETIME DEFAULT NULL,
  `deleted_at` DATETIME DEFAULT NULL,
  # 关联外键
  KEY `merge_role_permissions_role_id_foreign` (`role_id`),
  KEY `merge_role_permissions_permissions_id_foreign` (`permissions_id`),
  CONSTRAINT `merge_role_permissions_role_id_foreign` FOREIGN KEY (`role_id`) REFERENCES `role` (`id`) ON DELETE CASCADE,
  CONSTRAINT `merge_role_permissions_permissions_id_foreign` FOREIGN KEY (`permissions_id`) REFERENCES `permissions` (`id`) ON DELETE CASCADE,
  PRIMARY KEY (`id`)
) ENGINE=INNOID COMMENT='角色权限中间表';

CREATE TABLE `merge_role_people` (
  `id` INT UNSIGNED AUTO_INCREMENT,
  `role_id` INT UNSIGNED COMMENT '角色ID',
  `people_id` INT UNSIGNED COMMENT '人员ID',
  `created_at` DATETIME NOT NULL,
  `updated_at` DATETIME DEFAULT NULL,
  `deleted_at` DATETIME DEFAULT NULL,
  # 关联外键
  KEY `merge_role_people_role_id_foreign` (`role_id`),
  KEY `merge_role_people_people_id_foreign` (`people_id`),
  CONSTRAINT `merge_role_people_role_id_foreign` FOREIGN KEY (`role_id`) REFERENCES `role` (`id`) ON DELETE CASCADE,
  CONSTRAINT `merge_role_people_people_id_foreign` FOREIGN KEY (`people_id`) REFERENCES `people` (`id`) ON DELETE CASCADE,
  PRIMARY KEY (`id`)
) ENGINE=INNOID COMMENT='角色人员中间表';

CREATE TABLE `token` (
  `id` INT UNSIGNED AUTO_INCREMENT,
  `type` TINYINT NOT NULL COMMENT '令牌类型(1 后台系统, 2, app客户端)',
  `token` TEXT COMMENT '对应token',
  `people_id` INT UNSIGNED COMMENT '人员ID',
  `created_at` DATETIME NOT NULL,
  `updated_at` DATETIME DEFAULT NULL,
  `deleted_at` DATETIME DEFAULT NULL,
  # 关联外键
  KEY `token_people_id_foreign` (`people_id`),
  CONSTRAINT `token_people_id_foreign` FOREIGN KEY (`people_id`) REFERENCES `people` (`id`) ON DELETE CASCADE,
  PRIMARY KEY (`id`)
) ENGINE=INNOID COMMENT='令牌表';

# 插入组织机构数据
INSERT INTO `organization` VALUES
(1, '建华旅游', '四川建华旅游有限公司', NULL, NOW(), NOW(), NULL),
(2, '建华旅游绵阳分店', '四川建华旅游有限公司绵阳分店', 1, NOW(), NOW(), NULL),
(3, '建华旅游德阳分店', '四川建华旅游有限公司德阳分店', 1, NOW(), NOW(), NULL),
(4, '建华旅游广元分店', '四川建华旅游有限公司广元分店', 1, NOW(), NOW(), NULL);

# 插入部门数据
INSERT INTO `department` VALUES
(1, '管理层', '负责领导', NULL, 2, NOW(), NOW(), NULL),
(2, '销售部', '负责销售', NULL, 2, NOW(), NOW(), NULL),
(3, '司机部', '负责接送', NULL, 2, NOW(), NOW(), NULL),
(4, '导游部', '负责旅游', NULL, 2, NOW(), NOW(), NULL);

# 插入人员数据
INSERT INTO `people` VALUES
(1, 'admin', '123456', '管理员', NULL, '13888888888', NULL, 0, NULL, NOW(), NOW(), NULL),
(2, 'guanliyuan', '123456', '管理员一号', NULL, '13888888888', NULL, 0, NULL, NOW(), NOW(), NULL),
(3, 'wenyuan1', '123456', '文员一号', NULL, '13888888888', NULL, 0, NULL, NOW(), NOW(), NULL);

# 插入角色数据
INSERT INTO `role` VALUES
(1, '超级管理员', '系统最高权限者', NOW(), NOW(), NULL),
(2, '管理员', '普通管理员', NOW(), NOW(), NULL),
(3, '文员', '负责整理和录入资料', NOW(), NOW(), NULL);

# 插入权限数据
INSERT INTO `permissions` VALUES
(1, '新增人员', '是否可以新增人员', 'PEOPLE_CREATE', NOW(), NOW(), NULL),
(2, '更新人员', '是否可以更新人员', 'PEOPLE_UPDATE', NOW(), NOW(), NULL),
(3, '删除/锁定人员', '是否可以删除/锁定人员', 'PEOPLE_DELETE', NOW(), NOW(), NULL);

# 插入角色权限中间表数据
INSERT INTO `merge_role_permissions` (`role_id`, `permissions_id`, `created_at`, `updated_at`) VALUES
(1, 1, NOW(), NOW()),
(1, 2, NOW(), NOW()),
(1, 3, NOW(), NOW()),
(2, 1, NOW(), NOW()),
(2, 3, NOW(), NOW()),
(3, 1, NOW(), NOW());


# 插入角色人员中间表数据
INSERT INTO `merge_role_people` (`role_id`, `people_id`, `created_at`, `updated_at`) VALUES
(1, 1, NOW(), NOW()),
(2, 2, NOW(), NOW()),
(3, 3, NOW(), NOW());