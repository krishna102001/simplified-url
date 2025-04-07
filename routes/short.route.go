package routes

import (
	"github.com/gin-gonic/gin"
)

func GetSimplifyUrl(c *gin.Context) {
	c.JSON(200, gin.H{"message": "hello to simplified url"})
}
