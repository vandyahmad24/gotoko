package config

import (
	"fmt"
	"vandyahmad/newgotoko/migration"
	"vandyahmad/newgotoko/seeder"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB
var e error

func InitDB() {
	fmt.Println("Trying to connect database :" + GetEnvVariable("MYSQL_DBNAME"))
	fmt.Println("Trying to connect MYSQL_HOST :" + GetEnvVariable("MYSQL_HOST"))
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		GetEnvVariable("MYSQL_USER"),
		GetEnvVariable("MYSQL_PASSWORD"),
		GetEnvVariable("MYSQL_HOST"),
		GetEnvVariable("MYSQL_PORT"),
		GetEnvVariable("MYSQL_DBNAME"))
	DB, e = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if e != nil {
		panic(e)
	}
	fmt.Println("Success connect database :" + GetEnvVariable("MYSQL_DBNAME"))

	InitMigrate()
	InitSeeder()
}

func InitMigrate() {
	fmt.Println("Jalankan migration")
	DB.AutoMigrate(
		&migration.Cashier{}, &migration.Category{}, &migration.Payment{},
	)
	fmt.Println("Selesai migration")

}

func InitSeeder() {
	fmt.Println("Jalankan Seeder")
	seeder := seeder.NewSeeder(DB)
	seeder.CashierSeeder()
	seeder.CategorySeeder()
	seeder.PaymentSeeder()
	fmt.Println("Selesai Seeder")
}
