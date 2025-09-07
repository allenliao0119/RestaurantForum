package database

import (
	"fmt"
	"log"

	"github.com/allenliao0119/RestaurantForum/internal/config"

	_ "github.com/lib/pq"
	"xorm.io/xorm"
)

var Engine *xorm.Engine

func InitDB(cfg config.Database) {
	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		cfg.Username,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.Name,
	)

	engine, err := xorm.NewEngine("postgres", connStr)
    if err != nil {
        log.Fatalf("Failed to connect to DB: %v", err)
    }

    if err := engine.Ping(); err != nil {
        log.Fatalf("DB not reachable: %v", err)
    }

    log.Println("Database connected successfully 🚀")

	Engine = engine
}
