package polo

import "github.com/gin-gonic/gin"

const (
	polo = "polo"
)

func Polo(c *gin.Context) {
	c.String(200, polo)
}