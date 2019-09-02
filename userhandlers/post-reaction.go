package userhandlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/qjouda/dignity-platform/backend/auth"
	"github.com/qjouda/dignity-platform/backend/datatype"
)

// PostReaction submit user reaction to a post
func PostReaction(sc datatype.ServiceContainer) gin.HandlerFunc {
	return func(c *gin.Context) {
		user := auth.MustGetUser(c)
		body := struct {
			BlockID datatype.ID `json:"blockId"`
			EmojiID datatype.ID `json:"emojiId"`
		}{}
		c.BindJSON(&body)
		err := sc.UserService.AddUserPostReaction(
			user,
			body.BlockID,
			body.EmojiID,
		)
		if err != nil {
			c.String(http.StatusBadRequest, err.Error())
			return
		}
		c.JSON(http.StatusOK, nil)
	}
}
