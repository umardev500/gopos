package database

import (
	"context"
	stdLog "log"
	"os"
	"sync"
	"time"

	"github.com/rs/zerolog/log"
	"gitub.com/umardev500/gopos/pkg/constant"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type GormInstance struct {
	DB *gorm.DB
}

var (
	gormInstance *GormInstance
	gormOnce     sync.Once
)

func GetGormInstance() *GormInstance {
	gormOnce.Do(func() {
		host := os.Getenv("DB_HOST")
		port := os.Getenv("DB_PORT")
		user := os.Getenv("DB_USER")
		pass := os.Getenv("DB_PASS")
		name := os.Getenv("DB_NAME")
		dsn := "host=" + host + " port=" + port + " user=" + user + " password=" + pass + " dbname=" + name + " sslmode=disable"
		newLogger := logger.New(
			stdLog.New(os.Stdout, "\r\n", stdLog.LstdFlags), // io writer
			logger.Config{
				SlowThreshold:             time.Second, // Slow SQL threshold
				LogLevel:                  logger.Info, // Log level
				IgnoreRecordNotFoundError: true,        // Ignore ErrRecordNotFound error for logger
				ParameterizedQueries:      true,        // Don't include params in the SQL log
				Colorful:                  true,        // Disable color
			},
		)

		db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
			Logger: newLogger,
		})
		if err != nil {
			log.Fatal().Err(err).Msg("Failed to connect to database")
		}

		gormInstance = &GormInstance{
			DB: db,
		}
	})
	return gormInstance
}

func (g *GormInstance) GetConn(ctx context.Context) *gorm.DB {
	if tx, ok := ctx.Value(constant.TransactionContextKey).(*gorm.DB); ok {
		return g.switchScope(tx, ctx)
	}
	return g.switchScope(g.DB, ctx)
}

func (g *GormInstance) switchScope(db *gorm.DB, ctx context.Context) *gorm.DB {
	if unscoped, ok := ctx.Value(constant.ScopeContextKey).(bool); ok && unscoped {
		return db.Unscoped().WithContext(ctx)
	}
	return db.WithContext(ctx)
}

func (g *GormInstance) WithTransaction(ctx context.Context, fn func(context.Context) error) error {
	return g.DB.Transaction(func(tx *gorm.DB) error {
		ctx = context.WithValue(ctx, constant.TransactionContextKey, tx)
		return fn(ctx)
	})
}
