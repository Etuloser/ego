package main

import (
	"ego/pkg/gdbc"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/v1/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.GET("/v1/db/init", func(c *gin.Context) {
		gdbc.InitDb()
		c.JSON(200, gin.H{
			"data":    "",
			"code":    200,
			"info":    "success",
			"success": "true",
		})
	})
	r.GET("/v1/product/create", func(c *gin.Context) {
		gdbc.ProductCreate()
		c.JSON(200, gin.H{
			"data":    "",
			"code":    200,
			"info":    "success",
			"success": "true",
		})
	})
	r.GET("/v1/product/get", func(c *gin.Context) {
		result := gdbc.ProductGet()
		c.JSON(200, gin.H{
			"data":    result,
			"code":    200,
			"info":    "success",
			"success": "true",
		})
	})
	r.Run() // 监听并在 0.0.0.0:8080 上启动服务
}
