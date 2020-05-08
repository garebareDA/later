package routes

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

//Home ホームの表示
func Home(c *gin.Context){
	c.HTML(http.StatusFound, "index.html", gin.H{})
}