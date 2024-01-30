package config

import (
	"fmt"
	_ "github.com/lib/pq"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB

type DBConfig struct {
	Host     string
	Port     int
	User     string
	DBName   string
	Password string
}

func init() {
	//db
	dsn := "sqlserver://sa:root@localhost:1433?database=social-media"
	db, err := gorm.Open(sqlserver.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("database connected")
	DB = db
}

func BuildConfig() *DBConfig {
	return &DBConfig{
		Host:     "localhost",
		Port:     5432,
		User:     "postgres",
		DBName:   "social-media",
		Password: "root",
	}
}

func DbURL(dbConfig *DBConfig) string {
	s := fmt.Sprintf(
		"host=%s port=%d user=%s dbname=%s password=%s sslmode=disable",
		dbConfig.Host, dbConfig.Port, dbConfig.User, dbConfig.DBName, dbConfig.Password,
	)
	return s
}

func GetDb() *gorm.DB {
	return DB
}
