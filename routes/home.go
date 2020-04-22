package routes

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

func Home(c *gin.Context){
	c.HTML(http.StatusFound, "index.html", gin.H{})
}