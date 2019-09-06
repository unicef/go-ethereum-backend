package userhandlers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/qjouda/dignity-platform/backend/apperrors"
	"github.com/qjouda/dignity-platform/backend/appstrings"
	"github.com/qjouda/dignity-platform/backend/auth"
	"github.com/qjouda/dignity-platform/backend/datatype"
	"github.com/qjouda/dignity-platform/backend/filestorage"
	"github.com/qjouda/dignity-platform/backend/img"
)

// UploadProfileImage uploads new profile image
func UploadProfileImage(sc datatype.ServiceContainer) gin.HandlerFunc {
	return func(c *gin.Context) {
		user := auth.MustGetUser(c)
		body := struct {
			Image string `json:"image"`
		}{}
		c.BindJSON(&body)
		if len(body.Image) == 0 {
			c.String(http.StatusBadRequest, "Missing image")
			return
		}
		var logoPath string
		if body.Image != "" {
			imgData, contentType, ext, err := img.FromBase64(body.Image)
			logoPath = fmt.Sprintf(
				"profile-pictures/%d-%s.%s",
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
				apperrors.Critical("assethandlers:upload-profile-image:1", err)
				c.String(http.StatusBadRequest, err.Error())
				return
			}
			logoPath = "https://dignity-unicef.s3.amazonaws.com/" + logoPath
		}
		log.Println("---------------------------", logoPath)
		err := sc.UserService.UpdateProfileImage(user, logoPath)
		if err != nil {
			c.String(http.StatusBadRequest, err.Error())
			return
		}
		c.JSON(http.StatusOK, nil)
	}
}
