package module

import (
	"github.com/GoAdminGroup/go-admin/modules/db"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var conn *gorm.DB

func InitDatabaseConn(dsn db.Connection)  {
	db, err := gorm.Open(postgres.New(postgres.Config{
		Conn:                 dsn.GetDB("default"),
		PreferSimpleProtocol: true,
	}), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		panic(err)
	}
	conn = db

	setup()
}

func setup()  {
	if err := conn.AutoMigrate(&BeeNode{}); err != nil {
		panic(err)
	}
}