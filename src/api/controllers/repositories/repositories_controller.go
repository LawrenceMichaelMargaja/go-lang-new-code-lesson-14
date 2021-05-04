package repositories

import (
	"github.com/dembygenesis/go-rest-industry-standard/src/api/domain/repositories"
	"github.com/dembygenesis/go-rest-industry-standard/src/api/services"
	"github.com/dembygenesis/go-rest-industry-standard/src/api/utils/errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateRepo(c *gin.Context) {
	var request repositories.CreateRepoRequest
	// Checks if json parameters received can bind with the request object
	if err := c.ShouldBindJSON(&request); err != nil {
		apiErr := errors.NewBadRequestError("invalid json body")
		c.JSON(apiErr.Status(), apiErr)
		return
	}

	result, err := services.RepositoryService.CreateRepo(request)
	if err != nil {
		c.JSON(err.Status(), err)
		return
	}
	c.JSON(http.StatusCreated, result)
}