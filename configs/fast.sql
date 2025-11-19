-- ---------------------------------------------------------
-- fast.sql
-- Initialization script based on your current live database
-- ---------------------------------------------------------

-- Drop and recreate the database
DROP DATABASE IF EXISTS `fastgo`;
CREATE DATABASE `fastgo` DEFAULT CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci;
USE `fastgo`;

-- =========================================================
-- Table: user
-- =========================================================
DROP TABLE IF EXISTS `user`;

CREATE TABLE `user` (
                        `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
                        `userID` varchar(36) NOT NULL DEFAULT '' COMMENT 'User unique ID',
                        `username` varchar(255) NOT NULL DEFAULT '' COMMENT 'Username (unique)',
                        `password` varchar(255) NOT NULL DEFAULT '' COMMENT 'User password (encrypted)',
                        `nickname` varchar(30) NOT NULL DEFAULT '' COMMENT 'User nickname',
                        `email` varchar(256) NOT NULL DEFAULT '' COMMENT 'User email address',
                        `phone` varchar(16) NOT NULL DEFAULT '' COMMENT 'User phone number',
                        `createdAt` datetime NOT NULL DEFAULT current_timestamp() COMMENT 'User creation time',
                        `updatedAt` datetime NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp() COMMENT 'User last update time',
                        PRIMARY KEY (`id`),
                        UNIQUE KEY `user.userID` (`userID`),
                        UNIQUE KEY `user.username` (`username`),
                        UNIQUE KEY `user.phone` (`phone`)
) ENGINE=MyISAM AUTO_INCREMENT=99 DEFAULT CHARSET=utf8mb3 COLLATE=utf8mb3_general_ci COMMENT='User table';

-- Insert the initial root user
INSERT INTO `user` (`id`, `userID`, `username`, `password`, `nickname`, `email`, `phone`, `createdAt`, `updatedAt`) VALUES
    (96, 'user-000000', 'root', '$2a$10$ctsFXEUAMd7rXXpmccNlO.ZRiYGYz0eOfj8EicPGWqiz64YBBgR1y', 'colin404', 'colin404@foxmail.com', '18110000000', '2024-12-12 03:55:25', '2024-12-12 03:55:25');

-- =========================================================
-- Table: post
-- =========================================================
DROP TABLE IF EXISTS `post`;

CREATE TABLE `post` (
                        `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
                        `userID` varchar(36) NOT NULL DEFAULT '' COMMENT 'User unique ID',
                        `postID` varchar(35) NOT NULL DEFAULT '' COMMENT 'Post unique ID',
                        `title` varchar(256) NOT NULL DEFAULT '' COMMENT 'Post title',
                        `content` longtext NOT NULL COMMENT 'Post content',
                        `createdAt` datetime NOT NULL DEFAULT current_timestamp() COMMENT 'Post creation time',
                        `updatedAt` datetime NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp() COMMENT 'Post last update time',
                        PRIMARY KEY (`id`),
                        UNIQUE KEY `post.postID` (`postID`),
                        KEY `idx_post_userID` (`userID`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb3 COLLATE=utf8mb3_general_ci COMMENT='Post table';

-- No initial data for the post table
