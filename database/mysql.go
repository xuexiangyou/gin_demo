package database

import (
	"bytes"
	"log"
	"strconv"

	"github.com/jinzhu/gorm"
)

var Eloquent *gorm.DB

type Database interface {
	Open(dbType string, conn string) (db *gorm.DB, err error)
}

type Mysql struct {}

func (*Mysql) Open(dbType string, conn string) (db *gorm.DB, err error) {
	eloquent, err := gorm.Open(dbType, conn)
	return eloquent, err
}

func init() {
	dbType := "mysql"
	host := "mysql"
	port := 3306
	database := "galaxy"
	username := "root"
	password := "root"

	var err error

	var conn bytes.Buffer
	conn.WriteString(username)
	conn.WriteString(":")
	conn.WriteString(password)
	conn.WriteString("@tcp(")
	conn.WriteString(host)
	conn.WriteString(":")
	conn.WriteString(strconv.Itoa(port))
	conn.WriteString(")")
	conn.WriteString("/")
	conn.WriteString(database)
	conn.WriteString("?charset=utf8&parseTime=True&loc=Local&timeout=1000ms")

	var db Database
	db = new(Mysql)
	Eloquent, err = db.Open(dbType, conn.String())
	Eloquent.LogMode(true)

	if err != nil {
		log.Fatalln("%s connect error %v", dbType, err)
	} else {
		log.Println("%s connect success!", dbType)
	}
}
