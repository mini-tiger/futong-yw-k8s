package cfg

import (
	"database/sql"
	"fmt"

	"ftk8s/base/enum"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

// Mysql config info
type MysqlConf struct {
	MysqlUsername     string
	MysqlPassword     string
	MysqlHost         string
	MysqlPort         int
	MysqlDatabase     string
	MysqlTimeout      string
	MysqlMaxIdleConns int
	MysqlMaxOpenConns int
}

// GetMysqlCliGormAndSql return mysql client of *gorm.DB and *sql.DB
func GetMysqlCliGormAndSql(mysqlConfObj *MysqlConf, appRunMode string) (*gorm.DB, *sql.DB) {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=UTC&timeout=%s",
		mysqlConfObj.MysqlUsername,
		mysqlConfObj.MysqlPassword,
		mysqlConfObj.MysqlHost,
		mysqlConfObj.MysqlPort,
		mysqlConfObj.MysqlDatabase,
		mysqlConfObj.MysqlTimeout,
	)

	// Get mysql client of *gorm.DB
	var gormLogLevel logger.Interface
	if appRunMode == enum.AppRunModeRelease {
		gormLogLevel = logger.Default.LogMode(logger.Warn)
	} else {
		gormLogLevel = logger.Default.LogMode(logger.Info)
	}
	gdb, err := gorm.Open(
		mysql.New(mysql.Config{DSN: dsn}),
		&gorm.Config{
			NamingStrategy: schema.NamingStrategy{SingularTable: true},
			Logger:         gormLogLevel,
			PrepareStmt:    true,
		},
	)
	if err != nil {
		Mlog.Panic("failed to get mysql client of *gorm.DB, error message: ", err.Error())
		return nil, nil
	}

	// Get mysql client of *sql.DB
	udb, err := gdb.DB()
	if err != nil {
		Mlog.Panic("failed to get mysql client of *sql.DB, error message: ", err.Error())
		return nil, nil
	}

	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	udb.SetMaxIdleConns(mysqlConfObj.MysqlMaxIdleConns)
	// SetMaxOpenConns sets the maximum number of open connections to the database.
	udb.SetMaxOpenConns(mysqlConfObj.MysqlMaxOpenConns)

	Mlog.Info("successfully to get mysql client of *gorm.DB and *sql.DB")
	return gdb, udb
}
