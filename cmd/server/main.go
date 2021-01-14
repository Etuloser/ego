package main

import (
	"net/http"
	"os"

	"ego/pkg/env"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

// Result Comment
type Result struct {
	Code    int         `json:"code" xml:"code"`
	Data    interface{} `json:"data" xml:"data"`
	Message string      `json:"message" xml:"message"`
	Success bool        `json:"success" xml:"success"`
}

func main() {
	// 设置环境变量
	os.Setenv("env", "development")

	// 从文件加载环境变量
	envPath := env.GetEnvPath()
	godotenv.Load(envPath)

	// 新建echo实例
	e := echo.New()

	// 路由表
	e.GET("/", endpoint)
	e.GET("/users/:id", getUser)
	e.GET("/test", test)

	e.Logger.Fatal(e.Start(":1323"))
}

func endpoint(c echo.Context) error {
	u := &Result{
		Code:    200,
		Data:    nil,
		Message: "endpoint is /",
		Success: true,
	}
	return c.JSON(http.StatusOK, u)
}

// e.GET("/users/:id", getUser)
func getUser(c echo.Context) error {
	// User ID from path `users/:id`
	id := c.Param("id")
	u := &Result{
		Code:    200,
		Data:    id,
		Message: "query success",
		Success: true,
	}
	return c.JSON(http.StatusOK, u)
}

func test(c echo.Context) error {
	env := os.Getenv("DEV_DATABASE_URI")
	u := &Result{
		Code:    200,
		Data:    env,
		Message: "query success",
		Success: true,
	}
	return c.JSON(http.StatusOK, u)
}
