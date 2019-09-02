-- +migrate Up

DROP TABLE IF EXISTS `reactions`;
CREATE TABLE `reactions` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(32) NOT NULL,
  `logo` varchar(512) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=48 DEFAULT CHARSET=utf8mb4;

DROP TABLE IF EXISTS `user_asset_block_reaction`;
CREATE TABLE `user_asset_block_reaction` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `userId` int(10) unsigned NOT NULL,
  `assetBlockId` int(10) unsigned NOT NULL,
  `reactionId` varchar(100) NOT NULL,
  `createdAt` datetime NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=48 DEFAULT CHARSET=utf8mb4;

INSERT INTO reactions (name, logo) VALUES ('fb-heart', 'fb-heart.gif');
INSERT INTO reactions (name, logo) VALUES ('all_the_things', 'all_the_things.jpg');
INSERT INTO reactions (name, logo) VALUES ('success-kid', 'success-kid.png');

-- INSERT INTO reactions (name, logo) VALUES ('party_parrot', 'party_parrot.gif');
-- INSERT INTO reactions (name, logo) VALUES ('nyancat_big', 'nyancat_big.gif');
-- INSERT INTO reactions (name, logo) VALUES ('troll', 'troll.png');
-- INSERT INTO reactions (name, logo) VALUES ('aaw_yeah', 'aaw_yeah.gif');
-- INSERT INTO reactions (name, logo) VALUES ('arya', 'arya.png');
-- INSERT INTO reactions (name, logo) VALUES ('aw_yeah', 'aw_yeah.gif');
-- INSERT INTO reactions (name, logo) VALUES ('bananadance', 'bananadance.gif');
-- INSERT INTO reactions (name, logo) VALUES ('blues', 'blues.png');
-- INSERT INTO reactions (name, logo) VALUES ('carlton', 'carlton.gif');
-- INSERT INTO reactions (name, logo) VALUES ('charmander_dancing', 'charmander_dancing.gif');
-- INSERT INTO reactions (name, logo) VALUES ('deal_with_it_parrot', 'deal_with_it_parrot.gif');
-- INSERT INTO reactions (name, logo) VALUES ('doge', 'doge.png');
-- INSERT INTO reactions (name, logo) VALUES ('facepalm', 'facepalm.png');
-- INSERT INTO reactions (name, logo) VALUES ('fast_parrot', 'fast_parrot.gif');
-- INSERT INTO reactions (name, logo) VALUES ('fidget_spinner', 'fidget_spinner.gif');
-- INSERT INTO reactions (name, logo) VALUES ('homer-disappear', 'homer-disappear.gif');
-- INSERT INTO reactions (name, logo) VALUES ('lightsaber', 'lightsaber.png');
-- INSERT INTO reactions (name, logo) VALUES ('mario_luigi_dance', 'mario_luigi_dance.gif');
-- INSERT INTO reactions (name, logo) VALUES ('nightking', 'nightking.png');
-- INSERT INTO reactions (name, logo) VALUES ('pikachu', 'pikachu.png');
-- INSERT INTO reactions (name, logo) VALUES ('slam', 'slam.gif');
