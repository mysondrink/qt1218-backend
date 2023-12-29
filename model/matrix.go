/*
 * @Author: mysondrink
 * @Date: 2023-12-20 09:49:51
 * @Last Modified by: mysondrink
 * @Last Modified time: 2023-12-29 16:31:23
 * @Description:  试剂卡规格数据表
 */
package model

import (
	"time"

	"gorm.io/gorm"
)

type MatrixTable struct {
	gorm.Model
	Reagent_type        string    `gorm:"type:varchar(255);default:null"`
	Reagent_matrix      string    `gorm:"type:varchar(255);default:null"`
	Reagent_matrix_info string    `gorm:"type:varchar(255);default:null"`
	Status              int       `gorm:"default:0"`
	CreatedAt           time.Time `gorm:"not null"`
	UpdatedAt           time.Time `gorm:"not null"`
}

type TypeTable struct {
	Reagent_type string
}
