/*
 * @Author: mysondrink
 * @Date: 2023-12-20 10:33:51
 * @Last Modified by: mysondrink
 * @Last Modified time: 2023-12-29 16:41:08
 * @Description:  试剂数据操作处理
 */
package controller

import (
	"net/http"
	"qt1218-backend/common"
	"qt1218-backend/dto"
	"qt1218-backend/model"
	"qt1218-backend/response"

	"github.com/gin-gonic/gin"
)

// 获取试剂类型，同时测试网络联通性
func GetReagentInfo(ctx *gin.Context) {
	db := common.GetDB()

	// 获取试剂类型总数
	var typeTable []model.TypeTable
	if result := db.Model(&model.MatrixTable{}).Select("DISTINCT reagent_type").Where("status = ?", 0).Find(&typeTable); result.Error != nil {
		response.Response(ctx, http.StatusBadRequest, 400, nil, "查询失败")
		return
	}
	// 返回信息
	response.Success(ctx, gin.H{"data": dto.ToTypeDto(typeTable)}, "查询成功")
}

func SearchReagentInfo(ctx *gin.Context) {
	var total int64
	var mode int
	db := common.GetDB()
	// 获取参数
	var requestReagent = model.ReagentSearch{}
	ctx.Bind(&requestReagent)

	reagent_time := requestReagent.Reagent_time
	reagent_type := requestReagent.Reagent_type
	pagesize := requestReagent.Page_size
	page := requestReagent.Page

	// 结果获取
	var reagenttables []model.ReagentCopy1

	mode = 1
	offset := pagesize * (page - 1)

	if reagent_type != "" {
		mode = 2
	}
	switch mode {
	case 1:
		// 获取总页数
		if err := db.Unscoped().Where("reagent_time = ?", reagent_time).Find(&reagenttables).Count(&total).Error; err != nil {
			response.Response(ctx, http.StatusBadRequest, 400, nil, "查询失败")
			return
		}
		// 查询
		if result := db.Unscoped().Offset(offset).Limit(pagesize).Where("reagent_time = ?", reagent_time).Find(&reagenttables); result.Error != nil {
			response.Response(ctx, http.StatusBadRequest, 400, nil, "查询失败")
			return
		}
	case 2:
		// 获取总页数
		if err := db.Unscoped().Where("reagent_time = ? AND reagent_type = ?", reagent_time, reagent_type).Find(&reagenttables).Count(&total).Error; err != nil {
			response.Response(ctx, http.StatusBadRequest, 400, nil, "查询失败")
			return
		}

		// 查询
		if result := db.Unscoped().Offset(offset).Limit(pagesize).Where("reagent_time = ? AND reagent_type = ?", reagent_time, reagent_type).Find(&reagenttables); result.Error != nil {
			response.Response(ctx, http.StatusBadRequest, 400, nil, "查询失败")
			return
		}
	}

	// 返回信息
	response.Success(ctx, gin.H{"data": dto.ToReagentDto(reagenttables), "total": total}, "查询成功")
}
