package emailapp

import "github.com/gin-gonic/gin"

func Emailroute(c *gin.Context) {
	c.HTML(200, "email.html", gin.H{})
}
