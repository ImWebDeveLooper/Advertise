package application

import (
	"agahi/internal/entity/users"
	"agahi/internal/platform/repository"
	"context"
	"errors"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type App struct {
	DB struct {
		Mongo *mongo.Database
	}
	Repository struct {
		UserRepo users.Repository
	}
	Router *gin.Engine
}

// NewApp sets up a new application instance
// This method is a constructor function that creates a new instance of the App struct.
func NewApp() *App {
	app := &App{}
	if err := app.RegisterMongo(); err != nil {
		log.Fatal("Connection Failed!")
	}
	if err := app.RegisterUserRepo(); err != nil {
		log.Fatal(err)
	}
	if err := app.RegisterRouter(); err != nil {
		log.Fatal(err)
	}
	app.RegisterRoutes()
	return app
}

// RegisterMongo Establish a Connection to a MongoDB Database
func (a *App) RegisterMongo() error {
	// Make a Connection for database
	log.Debug("Database is Connecting...")
	dbURI := "mongodb://localhost:27017"
	dbName := "agahi"
	clientOptions := options.Client().ApplyURI(dbURI)

	// Connect to Database
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Error(err)
		return err
	}

	// Access the database
	a.DB.Mongo = client.Database(dbName)

	log.Debug("Database Connected Successfully.")
	return nil
}

// RegisterUserRepo Registers Users Repository to Contact with Database.
// to Create or Make Changes or Delete Users
func (a *App) RegisterUserRepo() error {
	if a.DB.Mongo == nil {
		return errors.New("can't Connect to Database")
	}
	a.Repository.UserRepo = repository.UserRepo{DB: a.DB.Mongo}
	return nil
}

// RegisterRouter Make a Router with Gin Web Framework
func (a *App) RegisterRouter() error {
	gin.SetMode(gin.ReleaseMode)
	a.Router = gin.Default()
	return nil
}

// RunRouter Run and Serve the Router that Created with Gin
func (a *App) RunRouter() {
	router := a.Router
	log.Println("Router is Running...")
	log.Println("Server Started On Port 3000")
	if err := router.Run(":3000"); err != nil {
		log.Fatal(err)
	}
}
