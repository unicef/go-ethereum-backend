-- +migrate Up

DROP TABLE IF EXISTS `not_message`;
CREATE TABLE `not_message` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `senderId` int(10) unsigned NOT NULL,
  `receiverId` int(10) unsigned NOT NULL,
  `topicId` varchar(100) NOT NULL,
  `type` int(10) unsigned NOT NULL,
  `isRead` tinyint(1) NOT NULL DEFAULT '0',
  `createdAt` datetime NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=48 DEFAULT CHARSET=utf8mb4;
