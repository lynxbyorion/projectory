package main

import (
	"github.com/gin-gonic/gin"

	"github.com/lynxbyorion/projectory/app"
)

func main() {
	var app = app.App{}
	app.InitDB()
	defer app.Database.Close()

	r := gin.Default()

	v1 := r.Group("api/v1")
	{
		v1.GET("/projects", app.GetProjects)
	}

	r.Run(":1333")
}
