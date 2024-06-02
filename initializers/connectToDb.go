package initializers

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDb() {
	var err error
	dsn := "user='koyeb-adm' password=QoRkYEXByW96 host=ep-black-silence-a13pxavi.ap-southeast-1.pg.koyeb.app dbname='koyebdb'"
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		PrepareStmt: false,
	})
	if err != nil {
		panic("failed to connect database")
	}
}
