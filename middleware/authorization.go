package middleware

import (
	"errors"
	"net/http"
	"strings"

	"github.com/aaguero96/technical_challenge_q2bank/utils"
	"github.com/gin-gonic/gin"
	"github.com/swaggo/swag/example/celler/httputil"
)

func Authorization(ctx *gin.Context) {
	var token string
	var err error

	bearerToken := ctx.Request.Header.Get("Authorization")
	if bearerToken == "" {
		token, err = ctx.Cookie("token")
		if err != nil {
			httputil.NewError(ctx, http.StatusUnauthorized, errors.New("token is not in cookies and it is not in authorizarion bearer"))
			ctx.Abort()
			return
		}
	} else {
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
