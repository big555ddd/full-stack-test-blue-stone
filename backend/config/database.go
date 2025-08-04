package config

import (
	"app/internal/logger"
	"context"
	"database/sql"
	"errors"
	"fmt"
	"sync"

	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/mysqldialect"
	"github.com/uptrace/bun/extra/bundebug"
)

func Database() {
	// Connect to database
	Register(
		&db,
		&DBOption{
			Host:      confString("DB_HOST", "127.0.0.1"),
			Port:      confInt64("DB_PORT", 3306),
			Database:  confString("DB_DATABASE", "testdb"),
			Username:  confString("DB_USER", "root"),
			Password:  confString("DB_PASSWORD", "secret"),
			Charset:   confString("DB_CHARSET", "utf8mb4"),
			ParseTime: confString("DB_PARSETIME", "true"),
			Loc:       confString("DB_LOC", "Local"),
		},
	)
	logger.Info("database connected success")
}

type DBOption struct {
	DSN       string
	Host      string
	Port      int64
	Database  string
	Username  string
	Password  string
	Charset   string
	ParseTime string
	Loc       string
}

func Register(conn **bun.DB, conf *DBOption) {
	if conf.DSN == "" {
		conf.DSN = generateDSN(conf)
	}

	sqldb, err := sql.Open("mysql", conf.DSN)
	if err != nil {
		logger.Errf("Failed to open database: %v", err)
	}

	*conn = bun.NewDB(sqldb, mysqldialect.New())

	if viper.GetBool("DEBUG") {
		(*conn).AddQueryHook(bundebug.NewQueryHook(bundebug.WithVerbose(true)))
	}

	err = (*conn).Ping()
	if err != nil {
		logger.Errf("Database connection failed: %v", err)
	}
}

func generateDSN(conf *DBOption) string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=%s&loc=%s",
		conf.Username, conf.Password, conf.Host, conf.Port, conf.Database, conf.Charset, conf.ParseTime, conf.Loc)
}

var (
	db     *bun.DB
	dbMap  = make(map[string]*bun.DB)
	dbLock sync.RWMutex
)

func GetDB() *bun.DB {
	return db
}

func DB(name ...string) *bun.DB {
	dbLock.RLock()
	defer dbLock.RUnlock()
	if dbMap == nil {
		panic("database not initialized")
	}
	if len(name) == 0 {
		return dbMap["default"]
	}

	database, ok := dbMap[name[0]]
	if !ok {
		panic("database not initialized")
	}
	return database
}

func Open(ctx context.Context) error {
	if db == nil {
		return errors.New("database not initialized")
	}
	return db.Ping()
}

// Close function to close database connection
func Close(ctx context.Context) error {
	if db == nil {
		return nil
	}
	return db.Close()
}
