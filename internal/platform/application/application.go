package application

import (
	"context"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type App struct {
	DB struct {
		Mongo *mongo.Database
	}
}

// NewApp sets up a new application instance
// This method is a constructor function that creates a new instance of the App struct.
func NewApp() *App {
	app := &App{}
	if err := app.RegisterMongo(); err != nil {
		log.Fatal("Connection Failed!")
	}
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
