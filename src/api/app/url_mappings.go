package app

import (
	"github.com/dembygenesis/go-rest-industry-standard/src/api/controllers/polo"
	"github.com/dembygenesis/go-rest-industry-standard/src/api/controllers/repositories"
)

func mapUrls() {
	router.GET("/marco", polo.Marco)
	router.POST("/repositories", repositories.CreateRepo)
}