package routes

import("github.com/gin-gonic/gin")

func statusError(c *gin.Context, errorMessage string, status int){
	c.JSON(status, gin.H{"error": errorMessage})
	c.Abort()
}