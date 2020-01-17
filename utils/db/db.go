package db

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"goloveu/utils"
)

// GormModel represents a base model
type GormModel struct {
	Id int64 `gorm:"PRIMARY_KEY;AUTO_INCREMENT" json:"id" form:"id"`
}

var db *gorm.DB

// OpenMySQL performs to open mysql db instance
func OpenMySQL(url string, maxIdleConns, maxOpenConns int, enableLog bool, models ...interface{}) (err error) {
	return OpenDB("mysql", url, maxIdleConns, maxOpenConns, enableLog, models...)
}

// OpenDB performs to open db connection
func OpenDB(dialect string, url string, maxIdleConns, maxOpenConns int, enableLog bool, models ...interface{}) (err error) {
	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return "t_" + defaultTableName
	}

	if db, err = gorm.Open(dialect, url); err != nil {
		utils.LogErrorf("opens database failed: %s", err.Error())
		return
	}

	db.LogMode(enableLog)
	db.SingularTable(true) // 禁用表名负数
	db.DB().SetMaxIdleConns(maxIdleConns)
	db.DB().SetMaxOpenConns(maxOpenConns)

	if err = db.AutoMigrate(models...).Error; nil != err {
		utils.LogErrorf("auto migrate tables failed: %s", err.Error())
	}
	return
}

// DB perfomrs to get db
func DB() *gorm.DB {
	return db
}

// CloseDB performs to close db
func CloseDB() {
	if db == nil {
		return
	}
	if err := db.Close(); nil != err {
		utils.LogErrorf("Disconnect from database failed: %s", err.Error())
	}
}

// Tx perform to do transaction
func Tx(db *gorm.DB, txFunc func(tx *gorm.DB) error) (err error) {
	tx := db.Begin()
	if tx.Error != nil {
		return
	}

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
			panic(r) // re-throw panic after Rollback
		} else if err != nil {
			tx.Rollback()
		} else {
			err = tx.Commit().Error
		}
	}()

	err = txFunc(tx)
	return err
}