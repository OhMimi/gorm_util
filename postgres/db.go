package postgres

import (
	DBConf "bobby/package/source/db_util/config"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

const (
	defaultMaxIdleConn = 5
	defaultMaxOpenConn = 10
	defaultMaxLifeTime = 120
)

// getDsn
func getDsn(c DBConf.Config) string {
	// Example: "host=localhost user=gorm password=gorm dbname=gorm port=9920 sslmode=disable TimeZone=Asia/Taipei"
	connectStr := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Taipei", c.Host, c.Account, c.Password, c.DBName, c.Port)
	return connectStr
}

// OpenDBConnection 開啟DB連線
func OpenDBConnection(c DBConf.Config) *gorm.DB {
	dbLogger := logger.New(
		log.New(os.Stdout, "[DB]", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Second,  // 慢 SQL 阈值
			LogLevel:      logger.Error, // Log level
			Colorful:      true,         // 禁用彩色打印
		},
	)

	dsn := getDsn(c)
	conn, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: dbLogger,
	})
	if err != nil {
		msg := "postgres,func: OpenDBConnection, DB 連線 " + dsn
		log.Fatal(msg + err.Error())
		return nil
	}

	// 設定MaxIdleConn
	maxIdleConn := defaultMaxIdleConn
	if c.MaxIdleConn != 0 {
		maxIdleConn = c.MaxIdleConn
	}
	// 設定MaxOpenConn
	maxOpenConn := defaultMaxOpenConn
	if c.MaxOpenConn != 0 {
		maxOpenConn = c.MaxOpenConn
	}
	// 設定MaxLifeTime
	maxLifeTime := defaultMaxLifeTime
	if c.MaxLifeTime != 0 {
		maxLifeTime = c.MaxLifeTime
	}

	dbInstance, err := conn.DB()
	if err != nil {
		log.Fatal("get dbInstance failed:", err)
		return nil
	}

	dbInstance.SetMaxIdleConns(maxIdleConn)
	dbInstance.SetMaxOpenConns(maxOpenConn)
	dbInstance.SetConnMaxLifetime(time.Duration(maxLifeTime) * time.Second)

	return conn
}
