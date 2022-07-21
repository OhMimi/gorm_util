package db_util

import (
	DBConf "bobby/package/source/db_util/config"
	Postgres "bobby/package/source/db_util/postgres"
	"gorm.io/gorm"
)

var dbSet *databaseSet

// databaseSet DB資料庫
type databaseSet struct {
	master *gorm.DB
	slave  *gorm.DB
}

func GetMasterDB() *gorm.DB {
	return dbSet.master
}

func GetSlaveDB() *gorm.DB {
	return dbSet.slave
}

// DBSetting DB 設定
type DBSetting struct {
	MasterConf DBConf.Config
	SlaveConf  DBConf.Config
}

func Init(setting DBSetting) {
	masterDB := Postgres.OpenDBConnection(setting.MasterConf)
	slaveDB := Postgres.OpenDBConnection(setting.SlaveConf)

	dbSet = &databaseSet{masterDB, slaveDB}
}
