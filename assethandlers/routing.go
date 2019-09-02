package assethandlers

import (
	"github.com/qjouda/dignity-platform/backend/datatype"
	"github.com/qjouda/dignity-platform/backend/routermiddleware"
	"github.com/gin-gonic/gin"
)

//InjectHandlers injects asset handlers into the application
func InjectHandlers(sc datatype.ServiceContainer, rg *gin.RouterGroup) {
	authenticator := routermiddleware.SessionMustAuth()
	rg.GET("/v2/assets/:assetId", GetAsset(sc))
	rg.GET("/v2/assets", GetAssets(sc))
	rg.POST("/v2/assets", authenticator, CreateAsset(sc))
	rg.POST("/v2/assets/:assetId/toggle-favorite", authenticator, ToggleFavoriteAsset(sc))
	rg.POST("/v2/asset-blocks", authenticator, CreateAssetBlock(sc))
	rg.POST("/v2/asset-block/:blockId/verify", authenticator, VerifyAssetBlock(sc))
	rg.POST("/v2/asset-block/:blockId/toggle-favorite", authenticator, ToggleFavoriteBlock(sc))

	rg.POST("/transfer", authenticator, Transfer(sc))
}
