-- +migrate Up

SET FOREIGN_KEY_CHECKS=0;

DROP TABLE IF EXISTS `country`;
CREATE TABLE `country` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(128) DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `name_UNIQUE` (`name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;


DROP TABLE IF EXISTS `crypto_deposit`;
CREATE TABLE `crypto_deposit` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `assetId` int(10) unsigned NOT NULL,
  `txId` varchar(512) NOT NULL,
  `userId` int(10) unsigned NOT NULL,
  `createdAt` datetime NOT NULL,
  `amount` decimal(30,18) unsigned NOT NULL,
  `address` varchar(256) NOT NULL,
  `confirmations` int(10) unsigned NOT NULL DEFAULT '0',
  `isConfirmed` tinyint(3) unsigned NOT NULL,
  `isProcessed` tinyint(3) unsigned NOT NULL,
  PRIMARY KEY (`id`),
  KEY `asset_fkk` (`assetId`),
  KEY `user_fkk` (`userId`),
  CONSTRAINT `asset_fkk` FOREIGN KEY (`assetId`) REFERENCES `asset` (`id`) ON DELETE NO ACTION ON UPDATE CASCADE,
  CONSTRAINT `user_fkk` FOREIGN KEY (`userId`) REFERENCES `user` (`id`) ON DELETE NO ACTION ON UPDATE CASCADE
) ENGINE=InnoDB AUTO_INCREMENT=48 DEFAULT CHARSET=utf8;


DROP TABLE IF EXISTS `asset_user_address`;
CREATE TABLE `asset_user_address` (
  `userId` int(10) unsigned NOT NULL,
  `assetId` int(10) unsigned NOT NULL,
  `address` varchar(512) NOT NULL,
  UNIQUE KEY `pk` (`userId`,`assetId`),
  KEY `address` (`address`),
  KEY `asset_fkk_idx` (`assetId`),
  CONSTRAINT `asset_fkk_idx` FOREIGN KEY (`assetId`) REFERENCES `asset` (`id`) ON DELETE CASCADE ON UPDATE CASCADE,
  CONSTRAINT `user_fk` FOREIGN KEY (`userId`) REFERENCES `user` (`id`) ON DELETE NO ACTION ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8;


DROP TABLE IF EXISTS `gorp_migrations`;
CREATE TABLE `gorp_migrations` (
  `id` varchar(255) NOT NULL,
  `applied_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

DROP TABLE IF EXISTS `asset`;
CREATE TABLE `asset` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(255) NOT NULL,
  `symbol` varchar(45) NOT NULL,
  `creatorId` int(10) unsigned NOT NULL,
  `supply` int(10) unsigned NOT NULL DEFAULT '0',
  `decimals` int(10) unsigned NOT NULL DEFAULT '0',
  `description` text NOT NULL,
  `minersCounter` int(10) unsigned NOT NULL DEFAULT '0',
  `favoritesCount` int(10) unsigned NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`),
  UNIQUE KEY `id_UNIQUE` (`id`),
  UNIQUE KEY `name_UNIQUE` (`name`),
  UNIQUE KEY `symbol_UNIQUE` (`symbol`)
) ENGINE=InnoDB AUTO_INCREMENT=250 DEFAULT CHARSET=utf8;

DROP TABLE IF EXISTS `user`;
CREATE TABLE `user` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `email` varchar(255) NOT NULL,
  `username` varchar(64) NOT NULL,
  `password` varchar(255) NOT NULL,
  `salt` varchar(32) NOT NULL,
  `ethereumAddress` varchar(512) NOT NULL,
  `createdAt` datetime NOT NULL,
  `updatedAt` datetime NOT NULL,
  `agreeToTerms` tinyint(1) unsigned NOT NULL DEFAULT '1',
  `isDeleted` tinyint(1) NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`),
  UNIQUE KEY `id_UNIQUE` (`id`),
  UNIQUE KEY `email_UNIQUE` (`email`),
  UNIQUE KEY `username_UNIQUE` (`username`)
) ENGINE=InnoDB AUTO_INCREMENT=10001 DEFAULT CHARSET=utf8;


DROP TABLE IF EXISTS `user_balance`;
CREATE TABLE `user_balance` (
  `userId` int(10) unsigned NOT NULL,
  `assetId` int(10) unsigned NOT NULL,
  `balance` decimal(30,8) NOT NULL DEFAULT '0.00000000',
  `reserved` decimal(30,8) NOT NULL DEFAULT '0.00000000',
  PRIMARY KEY (`userId`,`assetId`),
  KEY `assetId` (`assetId`),
  CONSTRAINT `assetID_FKK` FOREIGN KEY (`assetId`) REFERENCES `asset` (`id`) ON DELETE NO ACTION ON UPDATE NO ACTION,
  CONSTRAINT `userID_FKK` FOREIGN KEY (`userId`) REFERENCES `user` (`id`) ON DELETE NO ACTION ON UPDATE NO ACTION
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

DROP TABLE IF EXISTS `user_change_email_confirm`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `user_change_email_confirm` (
  `userId` int(10) unsigned NOT NULL,
  `email` varchar(255) NOT NULL,
  `token` varchar(255) NOT NULL,
  `createdAt` datetime NOT NULL,
  PRIMARY KEY (`userId`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;


DROP TABLE IF EXISTS `user_password_reset`;
CREATE TABLE `user_password_reset` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `userId` int(10) unsigned NOT NULL,
  `token` varchar(255) DEFAULT NULL,
  `createdAt` datetime DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `userId` (`userId`),
  CONSTRAINT `userId` FOREIGN KEY (`userId`) REFERENCES `user` (`id`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8;


CREATE TABLE `asset_favorites` (
  `userId` int(10) unsigned NOT NULL,
  `assetId` int(10) unsigned NOT NULL,
  `addedAt` datetime NOT NULL,
  UNIQUE KEY `pk_userid_assetid` (`assetId`,`userId`),
  KEY `userid_idx` (`userId`),
  CONSTRAINT `asset_favorites_fk_assetid` FOREIGN KEY (`assetId`) REFERENCES `asset` (`id`) ON DELETE CASCADE ON UPDATE CASCADE,
  CONSTRAINT `asset_favorites_fk_userid` FOREIGN KEY (`userId`) REFERENCES `user` (`id`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;


DROP TABLE IF EXISTS `asset_block`;
CREATE TABLE `asset_block` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `assetId` int(10) unsigned NOT NULL,
  `userId` int(10) unsigned NOT NULL,
  `text` varchar(10000),
  `status` tinyint(2) unsigned NOT NULL,
  `createdAt` datetime NOT NULL,
  PRIMARY KEY (`id`),
  KEY `asset_fkk` (`assetId`),
  KEY `user_fkk` (`userId`)
) ENGINE=InnoDB AUTO_INCREMENT=48 DEFAULT CHARSET=utf8;

DROP TABLE IF EXISTS `asset_block_image`;
CREATE TABLE `asset_block_image` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `blockId` int(10) unsigned NOT NULL,
  `filepath` VARCHAR(512) NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=48 DEFAULT CHARSET=utf8;

SET FOREIGN_KEY_CHECKS=1;
