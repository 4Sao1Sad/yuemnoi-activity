package main

import (
	"github.com/4Sao1Sad/yuemnoi-activity/db"
	"github.com/4Sao1Sad/yuemnoi-activity/internal/config"
	"github.com/4Sao1Sad/yuemnoi-activity/internal/model"
)

func main() {
	cfg := config.Load()
	// Initialize the DB connection
	db.InitDB(cfg)

	_ = db.DB.AutoMigrate(&model.ActivityLog{})

	err := db.ServerInit(cfg, db.DB)
	if err != nil {
		panic(err)
	}
}
