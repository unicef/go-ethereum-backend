package userhandlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/qjouda/dignity-platform/backend/auth"
	"github.com/qjouda/dignity-platform/backend/datatype"
)

//Login login route
func Login(sc datatype.ServiceContainer) gin.HandlerFunc {
	return func(c *gin.Context) {
		body := struct {
			Name string `json:"name"`
		}{}
		c.BindJSON(&body)
		// This is just for demonstration purposes
		email := body.Name + "@dignity.network"
		password := "PASS_PLACEOLDER"
		user, err := sc.UserService.FindByEmail(email)
		if user == nil {
			// HACK --> quickly signup a user with only a name
			ethereumAddress, err := sc.Ethereum.CreateNewAddress()
			if err != nil {
				c.String(http.StatusBadRequest, err.Error())
				return
			}
			user, err = sc.UserService.Register(
				email,
				password,
				body.Name,
				ethereumAddress,
				true,
			)
			if err != nil {
				c.String(http.StatusBadRequest, "Invalid login")
				return
			}
		} else if err != nil || !user.IsPassword(password) {
			c.String(http.StatusBadRequest, "Invalid login")
			return
		}
		auth.Login(c, user)
		c.JSON(http.StatusOK, user)
	}
}
