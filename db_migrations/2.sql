-- +migrate Up

ALTER table user ADD column profileImageUrl varchar(256) NOT NULL DEFAULT '' AFTER username;


DROP TABLE IF EXISTS `asset_block_favorites`;
CREATE TABLE `asset_block_favorites` (
  `userId` int(10) unsigned NOT NULL,
  `blockId` int(10) unsigned NOT NULL,
  `createdAt` datetime NOT NULL,
  UNIQUE KEY `pk_userid_assetid` (`blockId`,`userId`),
  KEY `userid_idx` (`userId`),
  CONSTRAINT `asset_block_favorites_fk_blockid` FOREIGN KEY (`blockId`) REFERENCES `asset_block` (`id`) ON DELETE CASCADE ON UPDATE CASCADE,
  CONSTRAINT `asset_block_favorites_fk_userid` FOREIGN KEY (`userId`) REFERENCES `user` (`id`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
