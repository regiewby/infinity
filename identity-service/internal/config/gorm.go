package config

import (
	"time"

	"github.com/infinity/identity-service/server/config"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func NewDatabase(appConfig *config.AppConfig, log *logrus.Logger) *gorm.DB {
	dsn := appConfig.Data.MySQL.Master.DSN
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.New(&logrusWriter{Logger: log}, logger.Config{
			SlowThreshold:             time.Second * 5,
			Colorful:                  false,
			IgnoreRecordNotFoundError: true,
			ParameterizedQueries:      true,
			LogLevel:                  logger.Info,
		}),
	})
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	connection, err := db.DB()
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	connection.SetMaxIdleConns(appConfig.Data.MySQL.Master.MaxIdle)
	connection.SetMaxOpenConns(appConfig.Data.MySQL.Master.MaxOpen)
	connection.SetConnMaxLifetime(time.Duration(appConfig.Data.MySQL.Master.ConMaxLifeTime))

	log.Info("database connected succesfully")
	return db
}

type logrusWriter struct {
	Logger *logrus.Logger
}

func (l *logrusWriter) Printf(message string, args ...interface{}) {
	l.Logger.Tracef(message, args...)
}
