package assethandlers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/qjouda/dignity-platform/backend/apperrors"
	"github.com/qjouda/dignity-platform/backend/appstrings"
	"github.com/qjouda/dignity-platform/backend/auth"
	"github.com/qjouda/dignity-platform/backend/datatype"
	"github.com/qjouda/dignity-platform/backend/filestorage"
	"github.com/qjouda/dignity-platform/backend/img"
)

// CreateAssetBlock creat a new asset
func CreateAssetBlock(sc datatype.ServiceContainer) gin.HandlerFunc {
	return func(c *gin.Context) {
		user := auth.MustGetUser(c)
		body := struct {
			AssetID   datatype.ID `json:"assetId"`
			BlockText string      `json:"blockText"`
			Images    []struct {
				Contents string `json:"contents"`
				FileName string `json:"filename"`
			}
		}{}
		c.BindJSON(&body)
		if len(body.Images) > 3 {
			c.String(http.StatusBadRequest, "Block images are limited to 3 images")
			return
		}
		var logoPaths []string
		{ // upload photos
			for _, image := range body.Images {
				var logoPath string
				if image.Contents != "" {
					imgData, contentType, ext, err := img.FromBase64(image.Contents)
					logoPath = fmt.Sprintf(
						"posts-images/%d-%s.%s",
						user.ID,
						appstrings.NewRandom().Generate(20),
						ext,
					)
					err = sc.FileStorage.Put(
						logoPath,
						contentType,
						imgData,
						filestorage.AclPublicRead,
					)
					if err != nil {
						apperrors.Critical("assethandlers:CreateAssetBlock:1", err)
						c.String(http.StatusBadRequest, err.Error())
						return
					}
					logoPath = "https://dignity-unicef.s3.amazonaws.com/" + logoPath
				}
				logoPaths = append(logoPaths, logoPath)
			}
		} // end of uploading images
		block, err := sc.AssetService.CreateAssetBlock(
			user.ID,
			body.AssetID,
			body.BlockText,
			logoPaths,
		)
		if err != nil {
			c.String(http.StatusBadRequest, err.Error())
			return
		}
		asset, _ := sc.AssetService.FindByID(body.AssetID)
		sc.UserService.Notify(user.ID, asset.CreatorID, body.AssetID, datatype.NEWPOST)
		c.JSON(http.StatusOK, block)
	}
}
