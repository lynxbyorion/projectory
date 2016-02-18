package app

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/boltdb/bolt"
	"github.com/gin-gonic/gin"
)

// ProjectDescription содержит информацию по конкретному проекту.
type ProjectDescription struct {
	ID int64 `json:"id"`
	// Name Имя проекта.
	Name string `json:"name"`
	// URLToRepo содержит url до репозитория данного проекта.
	URLToRepo string `json:"url_to_repo"`
	// LeadDeveloper пользователь являющийся ведущим разработчиком данного
	// проекта.
	LeadDeveloper string `json:"lead_developer"`
}

// App структура приложения
type App struct {
	Database *bolt.DB
}

// InitDB функция инициализации базы данных boltdb.
func (app *App) InitDB() error {
	var err error
	app.Database, err = bolt.Open("projectory.db", 0644, nil)
	if err != nil {
		log.Fatal(err)
	}

	app.Database.Update(func(tx *bolt.Tx) error {
		projectsSotre, err := tx.CreateBucketIfNotExists([]byte("Projects"))
		if err != nil {
			return fmt.Errorf("Create bucket: %s", err)
		}

		type Projects []ProjectDescription
		var projects = Projects{
			ProjectDescription{ID: 1, Name: "src77ya6vp", URLToRepo: "http://develop.res/rtimints/src77ay6vp", LeadDeveloper: "IODor"},
			ProjectDescription{ID: 2, Name: "src7mcf3", URLToRepo: "http://develop.res/rtimints/src7mcf3", LeadDeveloper: "AStankevich"},
		}

		for _, p := range projects {
			enc, err := json.Marshal(p)
			if err != nil {
				return err
			}
			projectsSotre.Put([]byte(p.Name), enc)
		}

		return nil
	})

	return nil
}

// GetProjects отдает json с списокм проектов.
func (app *App) GetProjects(c *gin.Context) {
	var projects []ProjectDescription
	var p *ProjectDescription
	app.Database.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("Projects"))
		b.ForEach(func(k, v []byte) error {
			err := json.Unmarshal(b.Get(k), &p)
			if err != nil {
				log.Fatal(err)
			}
			projects = append(projects, *p)
			return nil
		})
		return nil
	})

	c.JSON(200, projects)
}
