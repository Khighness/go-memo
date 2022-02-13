package model

// @Author KHighness
// @Update 2022-02-13

// 自动迁移
func migrate()  {
	DB.Set("gorm:table_options", "charset=utf8mb4")
	DB.AutoMigrate(&User{})
	DB.AutoMigrate(&Task{})
}