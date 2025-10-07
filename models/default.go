package models

import (
	"log"
	"os"

	"github.com/beego/beego/v2/adapter/orm"
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
)

type ContactModel struct {
	Id      uint64 `orm:"auto"`        // this automatically creates an integer primary key
	Name    string `orm:"size(100)"`   // 100 characters max
	Email   string `orm:"size(255)"`   // 255 characters max
	Message string `form:"type(text)"` // any size string
}

var O orm.Ormer

func InitDB() {
	driverName := os.Getenv("DATABASE_DRIVER")
	dataSource := os.Getenv("DATABASE_URL")
	driver := orm.DRSqlite
	if driverName == "postgres" {
		driver = orm.DRPostgres
	} else {
		driverName = "sqlite3"
		dataSource = "./stopwatch.db"
	}
	orm.RegisterDriver(driverName, driver)
	orm.RegisterDataBase("default", driverName, dataSource)
	// orm.RegisterDriver("sqlite3", orm.DRSqlite)
	// orm.RegisterDataBase("default", "sqlite3", "./goBee.db")
	// this function can take a list, e.g. orm.RegisterModel(new(M1), new(M2), ...)
	orm.RegisterModel(new(ContactModel), new(User), new(Visits))
	O = orm.NewOrm()

	err := orm.RunSyncdb("default", false, true)
	if err != nil {
		log.Fatalf("Failed to sync database: %v", err)
	}
}

// user stuff
type User struct {
	Id       uint64 `orm:"auto"` // this automatically creates an integer primary key
	Name     string `orm:"size(100)"`
	Email    string `orm:"size(255);unique"`
	Password string `orm:"size(255)"`
}

type Visits struct {
	Visits uint64 `orm:"size(500)"`
}
