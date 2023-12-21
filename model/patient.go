/*
 * @Author: mysondrink
 * @Date: 2023-12-19 17:26:55
 * @Last Modified by: mysondrink
 * @Last Modified time: 2023-12-20 11:22:45
 * @Description:  存储病人数据信息
 */
package model

import "gorm.io/gorm"

type PatientCopy1 struct {
	gorm.Model
	Name       string `gorm:"type:varchar(255);default:null"`
	Age        string `gorm:"type:varchar(255);default:null"`
	Gender     string `gorm:"type:varchar(255);default:null"`
	Patient_id int    `gorm:"primaryKey"`
	Id         int    `gorm:"primaryKey;autoIncrement:false"`
}
