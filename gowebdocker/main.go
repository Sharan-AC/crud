package main
import (
	"github.com/gin-gonic/gin"
)
func Getme(c *gin.Context){
	c.JSON(200, gin.H{
		"message":"hi am sharan ",
})
}
func Getdetails(c *gin.Context){
	c.JSON(200,gin.H{
		"message":"hi am sharan living in tirunelveli and have completed my bachelors degree waiting for my graduation",
	})
}
func main() {
	r := gin.Default()
	r.GET("/about",Getme)
	r.GET("/details",Getdetails)
	r.Run(":8080")
}