/*
 * @Author: mysondrink
 * @Date: 2023-12-19 16:52:55
 * @Last Modified by: mysondrink
 * @Last Modified time: 2023-12-19 16:57:07
 * @Description:  存储用户数据模型
 */
package model

import (
	"time"

	"gorm.io/gorm"
)

type UserTable struct {
	gorm.Model
	Name      string    `gorm:"type:varchar(20);not null"`
	Telephone string    `gorm:"type:varchar(20);not null"`
	Password  string    `gorm:"size:255;not null"`
	CreatedAt time.Time `gorm:"not null"`
	UpdatedAt time.Time `gorm:"not null"`
}
