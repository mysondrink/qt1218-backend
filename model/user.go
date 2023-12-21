/*
 * @Author: mysondrink
 * @Date: 2023-12-19 16:52:55
 * @Last Modified by: mysondrink
 * @Last Modified time: 2023-12-20 10:57:29
 * @Description:  存储用户数据模型
 */
package model

import (
	"time"

	"gorm.io/gorm"
)

type UserTable struct {
	gorm.Model
	User_name string    `gorm:"type:varchar(20);not null"`
	User_code string    `gorm:"type:varchar(20);not null"`
	Status    int       `gorm:"default:0"`
	CreatedAt time.Time `gorm:"not null"`
	UpdatedAt time.Time `gorm:"not null"`
}
