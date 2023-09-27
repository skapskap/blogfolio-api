package main

import (
	"database/sql"
	"flag"
	"fmt"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/skapskap/blogfolio-api/internal/data"
	"github.com/skapskap/blogfolio-api/internal/database"
	"log"
	"net/http"
	"os"
	"time"
)

const version = "1.0.0"

type config struct {
	port int
	env  string
}

type application struct {
	config config
	logger *log.Logger
	db     *sql.DB
	models data.Models
}

func main() {
	var cfg config

	// Verifique se o arquivo .env existe
	if _, err := os.Stat("./app.env"); err == nil {
		// Se o arquivo .env existe, carregue as variáveis de ambiente a partir dele
		err := godotenv.Load("./app.env")
		if err != nil {
			log.Fatal("Error loading .env file:", err)
		}
	}

	flag.IntVar(&cfg.port, "port", 4869, "Porta da API")
	flag.StringVar(&cfg.env, "env", "development", "Environment (development|staging|production)")
	flag.Parse()

	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	db, err := database.SetupDatabase()
	if err != nil {
		logger.Fatal(err)
	}
	defer db.Close()

	logger.Printf("database connection pool established")

	app := &application{
		config: cfg,
		logger: logger,
		db:     db,
		models: data.NewModels(db),
	}

	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.port),
		Handler:      app.routes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	logger.Printf("starting %s server on %s", cfg.env, srv.Addr)
	err = srv.ListenAndServe()
	logger.Fatal(err)
}
