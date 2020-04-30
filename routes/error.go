package routes

import("github.com/gin-gonic/gin")

func statusError(c *gin.Context, errorMessage string){
	c.JSON(500, gin.H{"status": errorMessage})
	c.Abort()
}