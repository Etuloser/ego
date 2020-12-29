package gdbc

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// Product struct
type Product struct {
	gorm.Model
	Code  string
	Price uint
}

// GetDb Comment
func GetDb() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	return db
}

// InitDb Comment
func InitDb() {
	db := GetDb()
	// 迁移 schema
	db.AutoMigrate(&Product{})
}

func ProductCreate() {
	db := GetDb()
	db.Create(&Product{Code: "D42", Price: 100})
}

func ProductGet() Product {
	db := GetDb()
	var product Product
	db.First(&product, "code = ?", "D42") // 查找 code 字段值为 D42 的记录
	return product
}
