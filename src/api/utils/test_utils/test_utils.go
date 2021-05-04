package test_utils

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"net/http/httptest"
)

// GetMockedContext provides a testable *gin.Context
func GetMockedContext(request *http.Request, response *httptest.ResponseRecorder) *gin.Context {
	c, _ := gin.CreateTestContext(response)
	c.Request = request
	return c
}