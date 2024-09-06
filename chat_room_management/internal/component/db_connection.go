package component

import (
	"fmt"
	"log"

	"ahyalfan.my.id/chat_rom_management/internal/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func GetConnection(cnf *config.Config) *gorm.DB {

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=%s",
		cnf.Databases.Host, cnf.Databases.Username, cnf.Databases.Password, cnf.Databases.Name, cnf.Databases.Port, "Asia/Jakarta")

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalln("failed to connect database: ", err.Error())
	}
	return db
}
