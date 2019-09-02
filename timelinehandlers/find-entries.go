package timelinehandlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/qjouda/dignity-platform/backend/auth"
	"github.com/qjouda/dignity-platform/backend/datatype"
	"github.com/qjouda/dignity-platform/backend/dbservice"
)

// FindEntries get time line entries handler
func FindEntries(sc datatype.ServiceContainer, timelineType datatype.TimelineType) gin.HandlerFunc {
	return func(c *gin.Context) {
		user := auth.MustGetUser(c)
		page, err := strconv.Atoi(c.Query("page"))
		if err != nil {
			page = 1
		}
		limit := 25
		offset := page*limit - limit
		timelineFilter := datatype.TimelineFilter{}

		timelineFilter.Type = timelineType
		assetID, _ := dbservice.StringToID(c.Params.ByName("assetID"))
		timelineFilter.AssetID = assetID
		userID, _ := dbservice.StringToID(c.Params.ByName("userID"))
		timelineFilter.UserID = userID
		entries, hasMore, err := sc.TimelineService.FindAll(
			sc,
			user,
			timelineFilter,
			offset,
			limit,
		)
		if err != nil {
			c.JSON(http.StatusInternalServerError, err)
			return
		}
		// assetsReactions, err := sc.TimelineService.FindAssetsReactions(entries)
		// for i, entry := range entries {
		// 	for _, postEmoji := range assetsReactions {
		// 		// if entry.BlockID == postEmoji. {
		// 		// 	entries[i].Reactions = append(entries[i].Reactions, postEmoji)
		// 		// }
		// 	}
		// }
		c.JSON(http.StatusOK, struct {
			HasMore bool
			Page    int
			Entries []datatype.TimelineEntry
		}{
			HasMore: hasMore,
			Page:    page,
			Entries: entries,
		})
	}
}
