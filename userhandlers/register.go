package userhandlers

import (
	"net/http"

	"github.com/qjouda/dignity-platform/backend/appstrings"
	"github.com/qjouda/dignity-platform/backend/auth"
	"github.com/qjouda/dignity-platform/backend/datatype"
	"github.com/gin-gonic/gin"
)

//Register register route
func Register(sc datatype.ServiceContainer) gin.HandlerFunc {
	return func(c *gin.Context) {
		body := struct {
			Email        string `json:"email"`
			Password     string `json:"password"`
			Username     string `json:"username"`
			AgreeToTerms bool   `json:"agreeToTerms"`
			IsFastSignup bool   `json:"isFastSignup"`
		}{}
		c.BindJSON(&body)
		if body.IsFastSignup {
			body.Password = appstrings.NewRandom().Generate(8)
			body.AgreeToTerms = true
		}
		ethereumAddress, err := sc.Ethereum.CreateNewAddress()
		if err != nil {
			c.String(http.StatusBadRequest, err.Error())
			return
		}
		user, err := sc.UserService.Register(
			body.Email,
			body.Password,
			body.Username,
			ethereumAddress,
			body.AgreeToTerms,
		)
		if err != nil {
			c.String(http.StatusBadRequest, err.Error())
			return
		}
		auth.Login(c, user)
		c.JSON(http.StatusOK, user)
	}
}
