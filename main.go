package main

import "github.com/gin-gonic/gin"

// ProjectDescription содержит информацию по конкретному проекту.
type ProjectDescription struct {
	ID            int64  `json:"id"`
	Name          string `json:"name"`
	URLToRepo     string `json:"url_to_repo"`
	LeadDeveloper string `json:"lead_developer"`
}

func main() {
	r := gin.Default()

	v1 := r.Group("api/v1")
	{
		v1.GET("/projects", GetProjects)
	}

	r.Run(":1333")
}

// GetProjects отдает json с списокм проектов.
func GetProjects(c *gin.Context) {
	type Projects []ProjectDescription

	var projects = Projects{
		ProjectDescription{ID: 1, Name: "src77ya6vp", URLToRepo: "http://develop.res/rtimints/src77ay6vp", LeadDeveloper: "IODor"},
		ProjectDescription{ID: 2, Name: "src7mcf3", URLToRepo: "http://develop.res/rtimints/src7mcf3", LeadDeveloper: "AStankevich"},
	}

	c.JSON(200, projects)
}
