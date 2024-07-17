package configs

import (
	"golang.org/x/exp/slog"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDatabaseConnect(config *Config) *gorm.DB {
	dsn := "host=" + config.Pg.Host + " user=" + config.Pg.User + " password=" + config.Pg.Password + " dbname=" + config.Pg.Dbname + " port=" + config.Pg.Port + " sslmode=disable TimeZone=Asia/Shanghai"
	slog.Info("PG : DB CONNECTED")
	slog.Info(config.Pg.Host)
	slog.Info(config.Pg.Port)
	slog.Info(config.Pg.User)

	pg := postgres.Open(dsn)
	db, err := gorm.Open(pg, &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db.AutoMigrate()

	return db
}
