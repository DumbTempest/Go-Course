package app

import (
	"database/sql"
	"fmt"
	"go/goProject/internal/api"
	"go/goProject/internal/migrations"
	"go/goProject/internal/store"
	"log"
	"net/http"
	"os"
)

type Application struct {
	Logger         *log.Logger
	WorkoutHandler *api.WorkoutHandler
	DB             *sql.DB
}

func NewApplication() (*Application, error) {
	pgDB, err := store.Open()
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	err = store.MigrateFS(pgDB, migrations.FS, ".")
	if err != nil {
		return nil, fmt.Errorf("failed to migrate database: %w", err)
	}
	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	WorkoutHandler := api.NewWorkoutHandler()

	app := &Application{
		Logger:         logger,
		WorkoutHandler: WorkoutHandler,
		DB:             pgDB,
	}
	return app, nil
}

func (app *Application) HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "status: ok")
	app.Logger.Println("Health check endpoint hit")
}
