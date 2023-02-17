package main

import (
	"embed"

	"github.com/framehack/reacthandler"

	"github.com/gin-gonic/gin"
)

//go:embed build
var app embed.FS

func main() {
	server := gin.Default()
	rh := reacthandler.NewHandler(app, "")
	server.GET("/*any", rh.GinHandler())
	server.Run()
}
