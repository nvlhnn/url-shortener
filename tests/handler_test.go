package tests

import (
	"fmt"
	"log"
	"os"
	"testing"
	"time"

	"github.com/joho/godotenv"
	"github.com/nvlhnn/url-shortener/internal/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB  *gorm.DB

func TestMain(m *testing.M) {
	err := godotenv.Load(os.ExpandEnv("../.env"))
	if err != nil {
		log.Fatalf("Error getting env %v\n", err)
	}

	Database()


	os.Exit(m.Run())
}


func Database(){

	DBURL := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", os.Getenv("TestDbUser"), os.Getenv("TestDbPassword"), os.Getenv("TestDbHost"), os.Getenv("TestDbPort"), os.Getenv("TestDbName"))
	conn, err :=  gorm.Open(mysql.Open(DBURL), &gorm.Config{})
	if err != nil {
		fmt.Printf("Cannot connect to mysql database\n")
		log.Fatal("This is the error:", err)
	} else {
		fmt.Printf("We are connected to the mysql database\n")
	}

	DB = conn
	refreshTable()
}


func refreshTable() error {

	// drop table if exist
	err := DB.Migrator().DropTable(&model.URL{})
	if err != nil {
		fmt.Printf("Cannot drop table: %v\n", err)
		return err
	}
	err = DB.AutoMigrate(&model.URL{})
	if err != nil {
		fmt.Printf("Cannot migrate table: %v\n", err)
		return err
	}

	log.Printf("Successfully refreshed tables")
	return nil
}


func CreateTestData() (model.URL, error) {

	refreshTable()

	// create sample url model
	url := model.URL{
		ShortenedURL: "test",
		OriginalURL: "https://www.google.com",
		ExpiredAt: time.Now().Add(time.Hour * 24),

	}
	
	err := DB.Create(&url).Error
	if err != nil {
		return model.URL{}, err
	}

	log.Printf("Successfully create test data")
	return url, nil
}