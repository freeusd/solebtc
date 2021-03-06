package middlewares

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/solefaucet/sole-server/errors"
	"github.com/solefaucet/sole-server/models"
)

type authRequiredDependencyGetAuthToken func(authTokenString string) (models.AuthToken, error)

// AuthRequired checks if user is authorized
func AuthRequired(
	getAuthToken authRequiredDependencyGetAuthToken,
	authTokenLifetime time.Duration,
) gin.HandlerFunc {
	return func(c *gin.Context) {
		authTokenHeader := c.Request.Header.Get("Auth-Token")
		if authTokenHeader == "" {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		authToken, err := getAuthToken(authTokenHeader)
		if err != nil && err != errors.ErrNotFound {
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}

		if authToken.CreatedAt.Add(authTokenLifetime).Before(time.Now()) {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		c.Set("auth_token", authToken)
		c.Next()
	}
}
