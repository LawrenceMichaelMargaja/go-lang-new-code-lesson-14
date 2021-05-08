package app

import (
	"github.com/dembygenesis/go-rest-industry-standard/oauth-api/src/api/controllers/oauth"
	"github.com/dembygenesis/go-rest-industry-standard/src/api/controllers/polo"
)

func mapUrls() {
	router.GET("/marco", polo.Marco)

	router.POST("/oauth/access_token", oauth.CreateAccessToken)
	router.GET("/oauth/access_token/:token_id", oauth.GetAccessToken)
}