package userservice

import (
	"time"

	"github.com/qjouda/dignity-platform/backend/apperrors"
	"github.com/qjouda/dignity-platform/backend/datatype"
)

// AddUserPostReaction adds a user reaction to a post
func (db *Service) AddUserPostReaction(
	user *datatype.User,
	blockID datatype.ID,
	reactionID datatype.ID,
) error {
	res, err := db.Exec(`
		INSERT INTO user_asset_block_reaction
		SET 
			userId = ?, 
			assetBlockId = ?, 
			reactionId = ?, 
			createdAt = ?
		`,
		user.ID, blockID, reactionID, time.Now(),
	)
	if err != nil {
		apperrors.Critical("user:userservice:add-user-post-reaction:1", err)
		return err
	}
	_, err = res.LastInsertId()
	if err != nil {
		apperrors.Critical("user:userservice:add-user-post-reaction:2", err)
		return err
	}
	return nil
}
