package db

import (
	"log"
	"parking-app-go/model"

	_ "github.com/jmoiron/sqlx"
	"golang.org/x/exp/errors/fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type mySQLTable struct {
	DB *gorm.DB
}

func (c mySQLTable) Find(result, query interface{}, args ...interface{}) {
	if query == nil {
		c.DB.Find(result)
		return
	}
	c.DB.Where(query, args...).Find(result)
	log.Println()
	return
}

func (c mySQLTable) FindOne(id, result interface{}) {
	c.DB.Find(result, id)
	return
}

func (c mySQLTable) Insert(doc interface{}) {
	c.DB.Create(doc)
	return
}

func (c mySQLTable) Join(t1, query string, results interface{}, args ...interface{}) {
	c.DB.Table(t1).Joins(query, args...).Scan(results)
}

func (c mySQLTable) Update(doc interface{}, query interface{}, args ...interface{}) {
	c.DB.Where(query, args).UpdateColumns(doc)
	return
}

type mySQLClient struct {
	db *gorm.DB
}

func (m mySQLClient) TableInstance(md interface{}) Table {
	m.db.AutoMigrate(md)
	return &mySQLTable{
		DB: m.db,
	}
}

func newMySQLClient(c model.Config, db string) *mySQLClient {
	connURI := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True", c.DBUser, c.DBPass, c.DBHost, c.DBPort, db)
	dbInstance, err := gorm.Open(mysql.Open(connURI), &gorm.Config{})
	if err != nil {
		// TODO - Retry to connect again in exponential time intervals
		log.Fatalln("Unable to connect to mysql DB...", err.Error())
	}
	return &mySQLClient{
		db: dbInstance,
	}
}
