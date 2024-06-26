package middleware

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/shawgichan/tourist/token"
	"github.com/shawgichan/tourist/utils"
)

var (
	authorizationHeaderKey  = "authorization"
	authorizationTypeBearer = "bearer"
	AuthorizationPayloadKey = "authorization_payload"
	TestUser                = "dreamer"
	AuthurizationHeader     = ""
	HeaderNotProvidedError  = "authorization header is not provided"
)

func AuthMiddleware(tokenMaker token.Maker) gin.HandlerFunc {
	return func(c *gin.Context) {
		AuthurizationHeader = c.GetHeader(authorizationHeaderKey)

		if len(AuthurizationHeader) == 0 {
			err := errors.New("authorization header is not provided")
			c.AbortWithStatusJSON(http.StatusInternalServerError, utils.ErrorResponse(err))
			return
		}
		fields := strings.Fields(AuthurizationHeader)
		if len(fields) < 2 {
			err := errors.New("invalid authurization header")
			c.AbortWithStatusJSON(http.StatusInternalServerError, utils.ErrorResponse(err))
			return
		}
		authorizationType := strings.ToLower(fields[0])
		if authorizationType != authorizationTypeBearer {
			err := fmt.Errorf("unsupported authorization type %s", authorizationType)
			c.AbortWithStatusJSON(http.StatusInternalServerError, utils.ErrorResponse(err))
			return
		}
		accessToken := fields[1]

		payload, err := tokenMaker.VerifyToken(accessToken)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, utils.ErrorResponse(err))
			return
		}
		c.Set(AuthorizationPayloadKey, payload)
		c.Next()
	}
}

func TimeoutMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		//!
		ctx, cancel := context.WithTimeout(c.Request.Context(), 5*time.Second)
		defer cancel()
		c.Request = c.Request.WithContext(ctx)
		c.Next()
		if ctx.Err() != nil || ctx.Err() == context.DeadlineExceeded {
			c.AbortWithStatusJSON(http.StatusRequestTimeout, utils.ErrorResponse(ctx.Err()))
		}
	}
}
