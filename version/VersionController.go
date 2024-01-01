package version

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetVersion(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, "v1.0")
}
