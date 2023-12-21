/*
 * @Author: mysondrink
 * @Date: 2023-12-20 10:34:14
 * @Last Modified by: mysondrink
 * @Last Modified time: 2023-12-20 10:34:37
 * @Description:  试剂数据传输对象定义
 */
package dto

import (
	"qt1218-backend/model"
)

type ReagentDto struct {
	Reagent_type        string `json:"reagent_type"`
	Patient_id          int    `json:"patient_id"`
	Reagent_photo       string `json:"reagent_photo"`
	Reagent_time        string `json:"reagent_time"`
	Reagent_id          int    `json:"reagent_id"`
	Reagent_code        string `json:"reagent_code"`
	Doctor              string `json:"doctor"`
	Depart              string `json:"depart"`
	Reagent_matrix      string `json:"reagent_matrix"`
	Reagent_time_detail string `json:"reagent_time_detail"`
	Reagent_matrix_info string `json:"reagent_matrix_info"`
	Patient_name        string `json:"patient_name"`
	Patient_age         string `json:"patient_age"`
	Patient_gender      string `json:"patient_gender"`
}

func ToReagentDto(reagent []model.ReagentCopy1) []ReagentDto {
	var result []ReagentDto
	for _, item := range reagent {
		template := ReagentDto{
			Reagent_type:        item.Reagent_type,
			Patient_id:          item.Patient_id,
			Reagent_photo:       item.Reagent_photo,
			Reagent_time:        item.Reagent_time,
			Reagent_id:          item.Reagent_id,
			Reagent_code:        item.Reagent_code,
			Doctor:              item.Doctor,
			Depart:              item.Depart,
			Reagent_matrix:      item.Reagent_matrix,
			Reagent_time_detail: item.Reagent_time_detail,
			Reagent_matrix_info: item.Reagent_matrix_info,
			Patient_name:        item.Patient_name,
			Patient_age:         item.Patient_age,
			Patient_gender:      item.Patient_gender,
		}
		result = append(result, template)
	}
	return result
}
