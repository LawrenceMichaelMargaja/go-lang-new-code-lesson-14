package app

import (
	"github.com/dembygenesis/go-rest-industry-standard/src/api/controllers/polo"
	"github.com/dembygenesis/go-rest-industry-standard/src/api/controllers/repositories"
)

func mapUrls() {
	router.GET("/marco", polo.Marco)
	// Old
	// router.POST("/repositories", repositories.CreateRepo)

	// New - supports multi-create
	router.POST("/repositories", repositories.CreateRepos)
}