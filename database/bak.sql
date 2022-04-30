/*
 Navicat Premium Data Transfer

 Source Server         : ginblog
 Source Server Type    : MySQL
 Source Server Version : 80027
 Source Host           : localhost:3306
 Source Schema         : ginblog

 Target Server Type    : MySQL
 Target Server Version : 80027
 File Encoding         : 65001

 Date: 30/04/2022 19:31:29
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for user
-- ----------------------------
DROP TABLE IF EXISTS `user`;
CREATE TABLE `user`  (
                         `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
                         `created_at` datetime(3) NULL DEFAULT NULL,
                         `updated_at` datetime(3) NULL DEFAULT NULL,
                         `deleted_at` datetime(3) NULL DEFAULT NULL,
                         `username` varchar(20) CHARACTER SET utf8 COLLATE utf8_bin NOT NULL,
                         `password` varchar(500) CHARACTER SET utf8 COLLATE utf8_bin NOT NULL,
                         `role` bigint NULL DEFAULT 2,
                         PRIMARY KEY (`id`) USING BTREE,
                         INDEX `idx_user_deleted_at`(`deleted_at` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 8 CHARACTER SET = utf8 COLLATE = utf8_bin ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of user
-- ----------------------------
INSERT INTO `user` VALUES (1, '2021-02-09 17:05:14.764', '2021-02-09 17:05:14.764', NULL, 'admin', '2a10$YGL5a9z7ykG6BWOo.XhJU.h8r98BD5IvAmLISBB9rFIefbDzrv58O', 1);

-- ----------------------------
-- Table structure for profile
-- ----------------------------
DROP TABLE IF EXISTS `profile`;
CREATE TABLE `profile`  (
                            `id` bigint NOT NULL AUTO_INCREMENT,
                            `name` varchar(20) CHARACTER SET utf8 COLLATE utf8_bin NULL DEFAULT NULL,
                            `desc` varchar(200) CHARACTER SET utf8 COLLATE utf8_bin NULL DEFAULT NULL,
                            `qqchat` varchar(200) CHARACTER SET utf8 COLLATE utf8_bin NULL DEFAULT NULL,
                            `wechat` varchar(100) CHARACTER SET utf8 COLLATE utf8_bin NULL DEFAULT NULL,
                            `weibo` varchar(200) CHARACTER SET utf8 COLLATE utf8_bin NULL DEFAULT NULL,
                            `bili` varchar(200) CHARACTER SET utf8 COLLATE utf8_bin NULL DEFAULT NULL,
                            `email` varchar(200) CHARACTER SET utf8 COLLATE utf8_bin NULL DEFAULT NULL,
                            `img` varchar(200) CHARACTER SET utf8 COLLATE utf8_bin NULL DEFAULT NULL,
                            `avatar` varchar(200) CHARACTER SET utf8 COLLATE utf8_bin NULL DEFAULT NULL,
                            `icp_record` varchar(200) CHARACTER SET utf8 COLLATE utf8_bin NULL DEFAULT NULL,
                            PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 2 CHARACTER SET = utf8 COLLATE = utf8_bin ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for comment
-- ----------------------------
DROP TABLE IF EXISTS `comment`;
CREATE TABLE `comment`  (
                            `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
                            `created_at` datetime(3) NULL DEFAULT NULL,
                            `updated_at` datetime(3) NULL DEFAULT NULL,
                            `deleted_at` datetime(3) NULL DEFAULT NULL,
                            `user_id` bigint UNSIGNED NULL DEFAULT NULL,
                            `article_id` bigint UNSIGNED NULL DEFAULT NULL,
                            `title` longtext CHARACTER SET utf8 COLLATE utf8_bin NULL,
                            `username` longtext CHARACTER SET utf8 COLLATE utf8_bin NULL,
                            `content` varchar(500) CHARACTER SET utf8 COLLATE utf8_bin NOT NULL,
                            `status` tinyint NULL DEFAULT 2,
                            PRIMARY KEY (`id`) USING BTREE,
                            INDEX `idx_comment_deleted_at`(`deleted_at` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 3 CHARACTER SET = utf8 COLLATE = utf8_bin ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for category
-- ----------------------------
DROP TABLE IF EXISTS `category`;
CREATE TABLE `category`  (
                             `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
                             `name` varchar(20) CHARACTER SET utf8 COLLATE utf8_bin NOT NULL,
                             PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 7 CHARACTER SET = utf8 COLLATE = utf8_bin ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for article
-- ----------------------------
DROP TABLE IF EXISTS `article`;
CREATE TABLE `article`  (
                            `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
                            `created_at` datetime(3) NULL DEFAULT NULL,
                            `updated_at` datetime(3) NULL DEFAULT NULL,
                            `deleted_at` datetime(3) NULL DEFAULT NULL,
                            `title` varchar(100) CHARACTER SET utf8 COLLATE utf8_bin NOT NULL,
                            `cid` bigint UNSIGNED NOT NULL,
                            `desc` varchar(200) CHARACTER SET utf8 COLLATE utf8_bin NULL DEFAULT NULL,
                            `content` longtext CHARACTER SET utf8 COLLATE utf8_bin NULL,
                            `img` varchar(100) CHARACTER SET utf8 COLLATE utf8_bin NULL DEFAULT NULL,
                            `comment_count` bigint NOT NULL DEFAULT 0,
                            `read_count` bigint NOT NULL DEFAULT 0,
                            PRIMARY KEY (`id`) USING BTREE,
                            INDEX `idx_article_deleted_at`(`deleted_at` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 2 CHARACTER SET = utf8 COLLATE = utf8_bin ROW_FORMAT = Dynamic;

SET FOREIGN_KEY_CHECKS = 1;
