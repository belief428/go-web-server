package orm

import (
	"github.com/belief428/go-web-server/orm/logic"
	"log"
	"os"
	"time"

	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"

	"gorm.io/gorm"
)

type Config struct {
	Debug       bool
	TablePrefix string
	Complex     bool // 是否复数

	MaxIdleConns int
	MaxOpenConns int
	MaxLifetime  int

	Option logic.IEngine
}

type DBMode string

const (
	DBModeForMysql  DBMode = "mysql"
	DBModeForSqlite DBMode = "sqlite"
)

var (
	orm *gorm.DB
)

func NewDB(config *Config) {
	option := &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   config.TablePrefix,
			SingularTable: !config.Complex,
		},
	}
	if config.Debug {
		option.Logger = logger.New(
			log.New(os.Stdout, "\r\n", log.LstdFlags),
			logger.Config{
				SlowThreshold:             time.Second,
				LogLevel:                  logger.Info,
				Colorful:                  false,
				IgnoreRecordNotFoundError: true,
			},
		)
	}
	db, err := gorm.Open(config.Option.DSN(), option)

	if err != nil {
		panic("Orm New Error：" + err.Error())
	}
	_db, _ := db.DB()
	_db.SetMaxIdleConns(config.MaxIdleConns)
	_db.SetMaxOpenConns(config.MaxOpenConns)
	_db.SetConnMaxLifetime(time.Duration(config.MaxLifetime) * time.Second)

	orm = db
}

func GetDB() *gorm.DB {
	if orm == nil {
		panic("Orm Dont Init")
	}
	return orm
}
