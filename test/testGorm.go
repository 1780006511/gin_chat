package main

import (
	"gin_chat/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	db, err := gorm.Open(mysql.Open("root:123456@tcp(127.0.0.1:3306)/gin_chat?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// 迁移 schema
	//fmt.Println(db.Migrator().HasTable(&models.UserBasic{}))

	//db.AutoMigrate(&models.UserBasic{})
	db.AutoMigrate(&models.Message{})

	//if err != nil {
	//	fmt.Println("err:", err)
	//	return
	//}

	// Create
	//user := &models.UserBasic{}
	//user.Name = "wanghao"
	//db.Create(user)

	// Read
	//a := db.Debug().First(&models.UserBasic{}, 1) // 根据整型主键查找
	//fmt.Println("a:", a)
	//db.First(&product, "code = ?", "D42") // 查找 code 字段值为 D42 的记录

	// Update - 将 product 的 price 更新为 200
	//fmt.Println("aaaaaaaaaaaaaaaa")
	//db.Model(user).Update("PassWord", "1234")
	// Update - 更新多个字段
	//db.Model(&product).Updates(Product{Price: 200, Code: "F42"}) // 仅更新非零值字段
	//db.Model(&product).Updates(map[string]interface{}{"Price": 200, "Code": "F42"})

	// Delete - 删除 product
	//db.Delete(&product, 1)
}
