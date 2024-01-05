/*
 * @Author: mysondrink
 * @Date: 2024-01-02 16:55:03
 * @Last Modified by: mysondrink
 * @Last Modified time: 2024-01-03 11:04:29
 * @Description:  获取服务器图片
 */
package controller

import (
	"encoding/base64"
	"fmt"
	"net/http"
	"os"
	"qt1218-backend/common"
	"qt1218-backend/dto"
	"qt1218-backend/model"
	"qt1218-backend/response"

	"github.com/gin-gonic/gin"
)

func GetImage(ctx *gin.Context) {
	// 获取试剂id
	reagent_id := ctx.Query("Reagent_id")
	// 结果获取
	var reagenttables []model.ReagentCopy1
	// 查询数据库
	db := common.GetDB()
	if result := db.Unscoped().Where("reagent_id = ?", reagent_id).First(&reagenttables); result.Error != nil {
		response.Response(ctx, http.StatusBadRequest, 400, nil, "查询失败")
		return
	}
	if len(reagenttables) == 0 {
		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "数据不存在")
		return
	}
	photo_name := reagenttables[0].Reagent_photo
	dir_name := reagenttables[0].Reagent_time[:10]
	file_path := ".jpeg"
	file_path = fmt.Sprintf("/home/orangpei/Desktop/qt0922/img/%s/%s.jpeg", dir_name, photo_name)
	fmt.Println(file_path)
	file_name := "C:\\Users\\YGS_Tu\\Desktop\\1109\\test.jpeg"
	// 图片读取到变量中
	// file, _ := os.ReadFile(file_path)
	file, _ := os.ReadFile(file_name)
	// 发送给前端
	// ctx.Writer.WriteString(string(file))
	base64img := base64.StdEncoding.EncodeToString(file)
	response.Success(ctx, gin.H{"data": dto.ToReagentDto(reagenttables), "img": base64img}, "查询成功")
}
