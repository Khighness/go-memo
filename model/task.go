package model

import "github.com/jinzhu/gorm"

// @Author KHighness
// @Update 2022-02-13

// 任务模型
type Task struct {
	gorm.Model
	User      User
	Uid       uint    `gorm:"not null"`
	Title     string  `gorm:"index;not null"`
	Status    int8    `gorm:"default:'0'"`
	Content   string  `gorm:"type:longtext"`
	StartTime int64   `gorm:"not null"`
	EndTime   int64   `gorm:"default:'0'"`
}

