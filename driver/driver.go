package driver

import (
	"fmt"
	"github.com/globalsign/mgo"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/kelseyhightower/envconfig"
	"log"
	"os"
)

const driver = "mysql"

type dbConfig struct {
	Db   string
	Host string
	Port int
	User string
	Pass string
}

var sqlDbc dbConfig
var mongoDbc dbConfig

func ConnectDB() *sqlx.DB {
	err := envconfig.Process("SQL", &sqlDbc)
	if err != nil {
		log.Fatal(err.Error())
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true", sqlDbc.User, sqlDbc.Pass, sqlDbc.Host, sqlDbc.Port, sqlDbc.Db)
	db, err := sqlx.Connect(driver, dsn)
	if err != nil {
		log.Println(dsn)
		log.Println(err.Error())
	}

	err = db.Ping()
	if err != nil {
		log.Println(err.Error())
	}

	return db
}

func ConnectMongo() *mgo.Session {
	err := envconfig.Process("MONGO", &mongoDbc)
	if err != nil {
		log.Println(err.Error())
	}

	var mongoUrl string
	if _, ok := os.LookupEnv("MONGO_USER"); ok {
		mongoUrl = fmt.Sprintf("mongodb://%s:%s@%s:%d/%s", mongoDbc.User, mongoDbc.Pass, mongoDbc.Host, mongoDbc.Port, mongoDbc.Db)
	} else {
		mongoUrl = fmt.Sprintf("mongodb://%s:%d/%s", mongoDbc.Host, mongoDbc.Port, mongoDbc.Db)
	}
	mongoSession, err := mgo.Dial(mongoUrl)
	if err != nil {
		//log.Println(mongoUrl)
		log.Println(err.Error())
	}

	return mongoSession
}
