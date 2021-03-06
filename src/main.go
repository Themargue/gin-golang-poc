package main

import (
	"fmt"
	"io/ioutil"

	"github.com/gin-gonic/gin"
)

func HomePage(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Hello world",
	})
}

func PostHomePage(c *gin.Context) {
	body := c.Request.Body
	value, err := ioutil.ReadAll(body)
	if err != nil {
		fmt.Println(err.Error())
	}

	c.JSON(200, gin.H{
		"body": string(value),
	})
}

func QueryStrings(c *gin.Context) {
	name := c.Query("name")
	age := c.Query("age")

	c.JSON(200, gin.H{
		"name": name,
		"age": age,
	})
}

func PathParameters(c *gin.Context) {
	name := c.Param("name")
	age := c.Param("age")
	
	c.JSON(200, gin.H{
		"name": name,
		"age": age,
	})
}

func main() {
	fmt.Println("Hello World")

	router := gin.Default()
	router.GET("/", HomePage)
	router.POST("/", PostHomePage)
	router.GET("/query", QueryStrings) // /query?name=elliot&age=24
	router.GET("/path/:name/:age", PathParameters)
	router.Run()
}