package middleware

import (
	"net/http"
	"strings"

	"github.com/aaguero96/technical_challenge_q2bank/utils"
	"github.com/gin-gonic/gin"
	"github.com/swaggo/swag/example/celler/httputil"
)

func Authorization(ctx *gin.Context) {
	var token string

	token, err := ctx.Cookie("token")
	if err != nil {
		bearerToken := ctx.Request.Header.Get("Authorization")

		if bearerToken == "" {
			httputil.NewError(ctx, http.StatusUnauthorized, err)
			ctx.Abort()
			return
		}

		token = strings.Split(bearerToken, "Bearer ")[1]
	}

	email, err := utils.DecriptJWT(token)
	if err != nil {
		httputil.NewError(ctx, http.StatusUnauthorized, err)
		ctx.Abort()
		return
	}

	ctx.Request.Header.Set("email", email)

	ctx.Next()
}
