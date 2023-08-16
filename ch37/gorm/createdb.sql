/*
 * createdb.sql
 * Author : He Chen <chenhe@zuoyejia.com>
 * Date :   2019/4/26
 * Copyright(c) 2019
 */

 -- create database
 create database test default character set utf8mb4 default collate utf8mb4_unicode_ci;

 -- create user 'sse'@'localhost'
 -- drop user if exists 'sse'@'localhost'
 create user 'sse'@'localhost' identified with mysql_native_password by 'StrongeneDB123456!';
 grant select,insert,update,delete on test.* to 'sse'@'localhost';
 create user 'sse'@'%' identified with mysql_native_password by 'StrongeneDB123456!';
 create user 'root'@'%' identified with mysql_native_password by 'StrongeneDB123456!';
 grant select,insert,update,delete on test.* to 'sse'@'%';
 GRANT ALL PRIVILEGES ON *.* TO 'root'@'%'


 flush privileges;

 -- create replication user
 -- drop user if exists 'sseslave'@'%'
 create user 'sseslave'@'%' identified by 'Zuoyejia!1475369';
 grant replication slave on *.* to 'sseslave'@'%';
 flush privileges;

 -- select database
 use ssetoccontent_develop;

 -- create tables
-- 正式资源
--`deletedAt` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00',
`deletedAt` timestamp NULL DEFAULT NULL,
CREATE TABLE `product` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `code` varchar(32) DEFAULT NULL,
  `price` bigint(20) NOT NULL,
  `createdAt` timestamp NOT NULL DEFAULT current_timestamp(),
  `updatedAt` timestamp NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp(),
  `deletedAt` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00',
  PRIMARY KEY (`id`),
  KEY `code` (`code`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;


CREATE TABLE `product_xcu` (
   `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
   `code` varchar(32) DEFAULT NULL,
   `price` bigint(20) NOT NULL,
   `createdAt` timestamp NOT NULL DEFAULT current_timestamp(),
   `updatedAt` timestamp NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp(),
   `deletedAt` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00',
   `fileMd5` varchar(32) DEFAULT NULL,
   PRIMARY KEY (`id`),
   KEY `code` (`code`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;


CREATE TABLE `product_adnoa` (
   `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
   `code` varchar(32) DEFAULT NULL,
   `price` bigint(20) NOT NULL,
   `createdAt` timestamp NOT NULL DEFAULT current_timestamp(),
   `updatedAt` timestamp NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp(),
   `deletedAt` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00',
   `fileMd5` varchar(32) DEFAULT NULL,
   PRIMARY KEY (`id`),
   KEY `code` (`code`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
