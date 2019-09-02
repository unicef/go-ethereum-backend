package timelineservice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/qjouda/dignity-platform/backend/apperrors"
	"github.com/qjouda/dignity-platform/backend/datatype"
)

// FindAssetsReactions finds assets reactions
func (db *Service) FindAssetsReactions(
	entries []datatype.TimelineEntry,
) (map[datatype.ID]datatype.PostEmoji, error) {
	type UserPostReaction struct {
		UserID    datatype.ID
		BlockID   datatype.ID
		EmojiID   datatype.ID
		EmojiName string
		EmojiLogo string
	}
	result := []UserPostReaction{}
	idsList := ""
	for idx, entry := range entries {
		if idx == len(entries)-1 {
			idsList += strconv.FormatUint(uint64(entry.BlockID), 10)
		} else {
			idsList += strconv.FormatUint(uint64(entry.BlockID), 10) + ", "
		}
	}
	stmt := fmt.Sprintf(`
	SELECT
		user_asset_block_reaction.userId,
		user_asset_block_reaction.assetBlockId,
		reactions.id,
		reactions.name,
		reactions.logo
	FROM user_asset_block_reaction
	LEFT JOIN reactions ON reactions.id=user_asset_block_reaction.reactionId
	WHERE user_asset_block_reaction.assetBlockId IN (%s)
	`, idsList)
	rows, err := db.Query(stmt)
	if err != nil {
		apperrors.Critical("timelineservice:find-assets-reactions:1", err)
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var c UserPostReaction
		err := rows.Scan(
			&c.UserID,
			&c.BlockID,
			&c.EmojiID,
			&c.EmojiName,
			&c.EmojiLogo,
		)
		if err != nil {
			apperrors.Critical("timelineservice:find-assets-reactions:2", err)
			return nil, err
		}
	}
	if err := rows.Err(); err != nil {
		apperrors.Critical("timelineservice:find-assets-reactions:3", err)
		return nil, err
	}
	// reactionMap := map[datatype.ID]datatype.PostEmoji{} // BlockID to reactions
	// for _, v := range result {
	// 	if val, ok := reactionMap[v.BlockID]; ok {
	// 		users := reactionMap[v.BlockID].Users
	// 		users = append(users, v.UserID)
	// 		reactionMap[v.BlockID].Users = users
	// 	} else {
	// 		e := datatype.PostEmoji{
	// 			EmojiID:   v.EmojiID,
	// 			EmojiName: v.EmojiName,
	// 			EmojiLogo: v.EmojiLogo,
	// 			Users:     []datatype.ID{v.UserID},
	// 		}
	// 		reactionMap[v.BlockID] = e
	// 	}
	// }
	log.Println(result)
	return nil, nil
}
