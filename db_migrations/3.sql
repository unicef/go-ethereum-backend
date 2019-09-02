-- +migrate Up

ALTER table asset ADD column ethereumAddress varchar(512) NOT NULL DEFAULT '' AFTER favoritesCount;
ALTER table asset ADD column ethereumTransactionAddress varchar(512) NOT NULL DEFAULT '' AFTER ethereumAddress;
ALTER table asset ADD column createdAt datetime NOT NULL DEFAULT NOW() AFTER ethereumTransactionAddress;

ALTER table asset_block ADD videoID varchar(256) NOT NULL DEFAULT '' AFTER status;
ALTER table asset_block ADD favoritesCounter int(10) unsigned NOT NULL DEFAULT '0' AFTER videoID;
ALTER table asset_block ADD ethereumTransactionAddress varchar(512) NOT NULL DEFAULT '' AFTER favoritesCounter;
