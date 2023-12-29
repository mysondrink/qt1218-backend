/*
 * @Author: mysondrink
 * @Date: 2023-12-29 14:26:31
 * @Last Modified by: mysondrink
 * @Last Modified time: 2023-12-29 16:36:51
 * @Description:  试剂类型传输对象定义
 */
package dto

import (
	"qt1218-backend/model"
	"time"
)

type MatrixDto struct {
	Reagent_type        string    `json:"reagent_type"`
	Reagent_matrix      string    `json:"reagent_matrix"`
	Reagent_matrix_info string    `json:"reagent_matrix_info"`
	Status              int       `json:"status"`
	CreatedAt           time.Time `json:"created_at"`
	UpdatedAt           time.Time `json:"updated_at"`
}

type TypeDto struct {
	Reagent_type string `json:"reagent_type"`
}

func ToMatrixDto(matrix []model.MatrixTable) []MatrixDto {
	var result []MatrixDto
	for _, item := range matrix {
		template := MatrixDto{
			Reagent_type:        item.Reagent_type,
			Reagent_matrix:      item.Reagent_matrix,
			Reagent_matrix_info: item.Reagent_matrix_info,
			Status:              item.Status,
			CreatedAt:           item.CreatedAt,
			UpdatedAt:           item.UpdatedAt,
		}
		result = append(result, template)
	}
	return result
}

func ToTypeDto(data []model.TypeTable) []TypeDto {
	var result []TypeDto
	for _, item := range data {
		template := TypeDto{
			Reagent_type: item.Reagent_type,
		}
		result = append(result, template)
	}
	return result
}
