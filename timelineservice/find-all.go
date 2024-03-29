package timelineservice

import (
	"fmt"

	"github.com/qjouda/dignity-platform/backend/apperrors"
	"github.com/qjouda/dignity-platform/backend/datatype"
	"github.com/qjouda/dignity-platform/backend/helpers"
)

// FindAll finds all assets
func (db *Service) FindAll(
	sc datatype.ServiceContainer,
	user *datatype.User,
	timelineFilter datatype.TimelineFilter,
	offset int,
	limit int,
) ([]datatype.TimelineEntry, bool, error) {
	result := []datatype.TimelineEntry{}
	var clause string
	if timelineFilter.Type == datatype.HOMETIMELINE {
		clause = "WHERE b.status=1"
	} else if timelineFilter.Type == datatype.ASSETTIMELINE {
		clause = fmt.Sprintf(
			" WHERE b.status<>%d AND asset.id = %d",
			datatype.BlockRejected,
			timelineFilter.AssetID,
		)
	} else if timelineFilter.Type == datatype.USERTIMELINE {
		clause = fmt.Sprintf(
			" WHERE b.status<>%d AND (doer.id = %d OR oracle.id = %d)",
			datatype.BlockRejected,
			timelineFilter.UserID,
			timelineFilter.UserID,
		)
	}
	rows, err := db.Query(fmt.Sprintf(`
			SELECT
				b.id,
				b.userId,
				doer.username,
				doer.profileImageUrl,
				asset.id,
				asset.name,
				asset.symbol,
				oracle.id,
				oracle.username,
				b.text,
				b.status,
				b.ethereumTransactionAddress,
				b.videoID,
				b.favoritesCounter,
				b.createdAt,
				IF(favorites.blockId, TRUE, FALSE),
				IF(asset_favorites.assetId, TRUE, FALSE) as following
			FROM asset_block b
			LEFT JOIN asset asset ON b.assetId=asset.Id
			LEFT JOIN user doer ON doer.id=b.userId
			LEFT JOIN user oracle ON oracle.id=asset.creatorId
			LEFT JOIN asset_block_favorites favorites ON b.id=favorites.blockId AND favorites.userId=?
			LEFT JOIN asset_favorites ON asset.id=asset_favorites.assetId AND asset_favorites.userId=?
			%s
			ORDER BY b.createdAt DESC
			LIMIT ? OFFSET ?
			`, clause), user.ID, user.ID, limit+1, offset)
	if err != nil {
		return nil, false, err
	}
	defer rows.Close()
	for rows.Next() {
		var c datatype.TimelineEntry
		err := rows.Scan(
			&c.BlockID,
			&c.UserID,
			&c.UserName,
			&c.UserProfileImageURL,
			&c.AssetID,
			&c.AssetName,
			&c.AssetSymbol,
			&c.OracleID,
			&c.OracleName,
			&c.Text,
			&c.Status,
			&c.EthereumTransactionAddress,
			&c.YtVideoID,
			&c.FavoritesCount,
			&c.CreatedAt,
			&c.DidUserLike,
			&c.DidUserLikeTopic,
		)
		if timelineFilter.Type == datatype.HOMETIMELINE && c.DidUserLikeTopic == false {
			continue
		}
		c.CreatedAtHuman = helpers.DateToHuman(c.CreatedAt)
		if err != nil {
			apperrors.Critical("timelineservice:find-all:1", err)
			return nil, false, err
		}
		// TODO optimize fetching images, bring all images for all at once,
		// not query for each entry
		c.Images, err = sc.AssetService.GetAssetBlockImages(c.BlockID)
		if err != nil {
			apperrors.Critical("timelineservice:find-all:2", err)
			return nil, false, err
		}
		c.Reactions, err = db.FindClaimReactions(c.BlockID)
		if err != nil {
			apperrors.Critical("timelineservice:find-all:3", err)
			return nil, false, err
		}
		result = append(result, c)
	}
	if err := rows.Err(); err != nil {
		return nil, false, err
	}
	hasMore := len(result) == limit+1
	if hasMore {
		result = result[:len(result)-1]
	}
	return result, hasMore, nil
}
