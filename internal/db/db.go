package db

import (
	"fmt"
	"log"

	"github.com/nvlhnn/url-shortener/internal/config"
	gormMySql "gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/golang-migrate/migrate/v4"
	goMigrateSql "github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

type Database struct {
	config     *config.Database
	Conn       *gorm.DB
}

func NewDatabase(config *config.Database) *Database {
	return &Database{
		config: config,
	}
}

func (db *Database) Connect() *gorm.DB {

	// connect to the database using gorm mysql database driver
	dsn := db.ConnString()
	conn, err :=  gorm.Open(gormMySql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to create a connection to database")
	}
	db.Conn = conn

	// tes ping the connection
	if err := db.Conn.Exec("SELECT 1").Error; err != nil {
		panic("Failed to ping database")
	}

	
	// run the migration
	sqlconn, _ := db.Conn.DB()
	driver, _ := goMigrateSql.WithInstance(sqlconn, &goMigrateSql.Config{})
	m, err := migrate.NewWithDatabaseInstance(
		"file://./internal/db/migration",
		"mysql", 
		driver,
	)

	if err != nil {
		log.Println("Failed to create a migration")
		panic(err)
	}
	
	m.Up()

	return db.Conn
	
}

func (db *Database) Close() {
	// close the database connection
	sqlDB, err := db.Conn.DB()
	if err != nil {
		panic("Failed to close the database connection")
	}

	sqlDB.Close()
}

func (db *Database) ConnString() string {
	
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		db.config.User,
		db.config.Password,
		db.config.Host,
		db.config.Port,
		db.config.Name,
	)
}
