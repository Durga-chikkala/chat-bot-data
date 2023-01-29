package main

import (
	"github.com/Dataservicee/handlers/query"

	handleruser "github.com/Dataservicee/handlers/user"
	serviceuser "github.com/Dataservicee/services/user"
	storequery "github.com/Dataservicee/store/query"
	"github.com/Dataservicee/store/user"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"

	"github.com/Dataservicee/models"
	"github.com/gin-gonic/gin"
)

func main() {
	app := gin.Default()
	PORT := ":3000"

	dsn := "host=localhost user=postgres password=password dbname=queries port=5432"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Printf("DB connection failed with Err:%v", err)
		return
	}

	// Migration
	err = db.AutoMigrate(models.QueryInfo{})
	if err != nil {
		log.Printf("Error during Migrations, err:%v", err)
		return
	}

	err = db.AutoMigrate(models.UserInfo{})
	if err != nil {
		log.Printf("Error during Migrations, err:%v", err)
		return
	}

	//Questions Injection
	queryStore := storequery.New(db)
	queryHandler := query.New(queryStore)

	//User Injection
	userStore := user.New(db)
	userService := serviceuser.New(userStore)
	userHandler := handleruser.New(userService)

	//Questions Endpoints
	app.POST("/chatbot", queryHandler.Create)
	app.GET("/chatbot", queryHandler.Get)
	app.GET("/chatbot/:question", queryHandler.GetByQuestion)
	app.GET("/chatbot/frequentQuestions", queryHandler.GetFrequentQuestions)
	app.PATCH("/chatbot/:question", queryHandler.PatchByQuestion)

	// User Details
	app.POST("/user/signup", userHandler.Create)
	app.GET("/user/login", userHandler.Get)
	app.GET("user/:id", userHandler.GetByID)
	app.PATCH("user/:id", userHandler.PatchByID)

	log.Printf("The server is running at port:%v", PORT)
	err = app.Run(PORT)
	if err != nil {
		log.Printf("Port is already in use")
		return
	}
}
