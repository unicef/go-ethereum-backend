package timelineservice

import (
	"github.com/qjouda/dignity-platform/backend/apperrors"
	"github.com/qjouda/dignity-platform/backend/datatype"
)

// UserReaction ...
type UserReaction struct {
	UserID    datatype.ID
	BlockID   datatype.ID
	EmojiID   datatype.ID
	EmojiName string
	EmojiLogo string
}

func remove(slice []UserReaction, s int) []UserReaction {
	return append(slice[:s], slice[s+1:]...)
}

// FindClaimReactions finds claim reactions
func (db *Service) FindClaimReactions(
	claimID datatype.ID,
) ([]datatype.PostEmoji, error) {

	result := []UserReaction{}
	stmt := `SELECT
		user_asset_block_reaction.userId,
		user_asset_block_reaction.assetBlockId,
		reactions.id,
		reactions.name,
		reactions.logo
	FROM user_asset_block_reaction
	LEFT JOIN reactions ON reactions.id=user_asset_block_reaction.reactionId
	WHERE user_asset_block_reaction.assetBlockId = ?`
	rows, err := db.Query(stmt, claimID)
	if err != nil {
		apperrors.Critical("timelineservice:find-assets-reactions:1", err)
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var c UserReaction
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
		result = append(result, c)
	}
	if err := rows.Err(); err != nil {
		apperrors.Critical("timelineservice:find-assets-reactions:3", err)
		return nil, err
	}
	// @TODO optimize this somehow, its a dump way to do it
	reactions := []datatype.PostEmoji{}
	for j := 0 + 1; j < len(result); j++ {
		userRaction := result[j]
		found := false
		for _, r := range reactions {
			if r.EmojiID == userRaction.EmojiID {
				found = true
				break
			}
		}
		if found {
			continue
		}
		var c datatype.PostEmoji
		users := []datatype.ID{userRaction.UserID}
		for i := j + 1; i < len(result); i++ {
			if uint64(result[i].EmojiID) == uint64(userRaction.EmojiID) {
				users = append(users, result[i].UserID)
				result = append(result[:i], result[i+1:]...)
			}
		}
		c.EmojiID = userRaction.EmojiID
		c.EmojiLogo = userRaction.EmojiLogo
		c.EmojiName = userRaction.EmojiName
		c.Users = users
		reactions = append(reactions, c)
	}
	return reactions, nil
}
