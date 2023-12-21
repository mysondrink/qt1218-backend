/*
 * @Author: mysondrink
 * @Date: 2023-12-20 10:33:51
 * @Last Modified by: mysondrink
 * @Last Modified time: 2023-12-20 11:40:07
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

func GetReagentInfo(ctx *gin.Context) {
	db := common.GetDB()
	// 获取参数
	var requestReagent = model.ReagentCopy1{}
	ctx.Bind(&requestReagent)

	// 获取id
	var reagenttables []model.ReagentCopy1
	if result := db.Find(&reagenttables); result.Error != nil {
		response.Response(ctx, http.StatusBadRequest, 400, nil, "查询失败")
		return
	}

	// 返回信息
	response.Success(ctx, gin.H{"data": dto.ToReagentDto(reagenttables)}, "查询成功")
}

func SearchReagentInfo(ctx *gin.Context) {
	db := common.GetDB()
	// 获取参数
	var requestReagent = model.ReagentCopy1{}
	ctx.Bind(&requestReagent)

	reagent_time := requestReagent.Reagent_time

	// 获取id
	var reagenttables []model.ReagentCopy1
	if result := db.Unscoped().Where("reagent_time = ?", reagent_time).Find(&reagenttables); result.Error != nil {
		response.Response(ctx, http.StatusBadRequest, 400, nil, "查询失败")
		return
	}

	// 返回信息
	response.Success(ctx, gin.H{"data": dto.ToReagentDto(reagenttables)}, "查询成功")
}
