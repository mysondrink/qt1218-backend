/*
 * @Author: mysondrink
 * @Date: 2023-12-19 16:56:44
 * @Last Modified by: mysondrink
 * @Last Modified time: 2023-12-19 17:27:35
 * @Description:  存储试剂数据模型
 */
package model

import "gorm.io/gorm"

type ReagentCopy1 struct {
	gorm.Model
	Reagent_type        string `gorm:"type:varchar(255);not null"`
	Patient_id          int    `gorm:"not null"`
	Reagent_photo       string `gorm:"type:varchar(255);default:null"`
	Reagent_time        string `gorm:"type:datetime;not null"`
	Reagent_id          int    `gorm:"not null;primary_key;auto_increment"`
	Reagent_code        string `gorm:"type:varchar(255);default:null"`
	Doctor              string `gorm:"type:varchar(255);default:null"`
	Depart              string `gorm:"type:varchar(255);default:null"`
	Reagent_matrix      string `gorm:"type:varchar(255);default:null"`
	Reagent_time_detail string `gorm:"type:varchar(255);default:null"`
	Reagent_matrix_info string `gorm:"type:varchar(255);default:null"`
	Patient_name        string `gorm:"type:varchar(255);default:null"`
	Patient_age         string `gorm:"type:varchar(255);default:null"`
	Patient_gender      string `gorm:"type:varchar(255);default:null"`
	// CreatedAt time.Time `gorm:"not null"`
	// UpdatedAt time.Time `gorm:"not null"`
}
