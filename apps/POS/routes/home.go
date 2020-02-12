package qpos

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func IndexPage(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{ "title":"Q-POS"})
}
