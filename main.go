package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	setupRoutes(r)
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

func setupRoutes(r *gin.Engine) {
	r.GET("/city/:name", Dummy)
}

//Dummy function
func Dummy(c *gin.Context) {
	name, ok := c.Params.Get("name")
	if ok == false {
		res := gin.H{
			"error": "name is missing",
		}
		c.JSON(http.StatusOK, res)
		return
	}
	/*
		city := ""
	*/
	res := gin.H{
		"name": name,
		"city": "hajipur",
	}
	c.JSON(http.StatusOK, res)
}
