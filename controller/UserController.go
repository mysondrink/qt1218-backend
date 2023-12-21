/*
 * @Author: mysondrink
 * @Date: 2023-12-20 09:00:30
 * @Last Modified by: mysondrink
 * @Last Modified time: 2023-12-20 09:53:25
 * @Description:  用户数据操作处理
 */
package controller

import (
	"fmt"
	"log"
	"net/http"
	"qt1218-backend/common"
	"qt1218-backend/dto"
	"qt1218-backend/model"
	"qt1218-backend/response"
	"qt1218-backend/util"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// 注册事务
func Register(ctx *gin.Context) {
	db := common.GetDB()
	// gin的绑定获取参数
	var requestUser = model.UserTable{}
	ctx.Bind(&requestUser)

	user_name := requestUser.User_name
	user_code := requestUser.User_code

	fmt.Println(user_name, user_code)

	// 数据验证
	// if len(user_code) < 6 {
	// 	response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "密码不能少于6位")
	// }

	// 如果名称没有传，给一个10位的随机字符串
	if len(user_name) == 0 {
		user_name = util.RandomString(10)
	}

	// 判断用户是否存在
	if isUserNameExist(db, user_name) {
		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "用户存在")
	}

	// 创建用户
	// 加密用户密码
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user_code), bcrypt.DefaultCost)
	if err != nil {
		response.Response(ctx, http.StatusUnprocessableEntity, 500, nil, "加密错误")
		return
	}
	newUser := model.UserTable{
		User_name: user_name,
		User_code: string(hashedPassword),
	}
	db.Create(&newUser)

	// 发放token到前端
	token, err := common.ReleaseToken(newUser)
	if err != nil {
		response.Response(ctx, http.StatusInternalServerError, 500, nil, "系统错误")
		log.Printf("token generate error : %v", err)
		return
	}

	// 返回结果
	response.Success(ctx, gin.H{"token": token}, "注册成功")
}

// 登录事务
func Login(ctx *gin.Context) {
	db := common.GetDB()
	var requestUser = model.UserTable{}
	ctx.Bind(&requestUser)

	// 获取参数
	user_name := requestUser.User_name
	user_code := requestUser.User_code

	// 判断用户是否存在
	var usertable model.UserTable
	db.Unscoped().Where("user_name = ?", user_name).First(&usertable)
	if usertable.ID == 0 {
		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "用户不存在")
		return
	}

	// 判断密码是否正确
	if err := bcrypt.CompareHashAndPassword([]byte(usertable.User_code), []byte(user_code)); err != nil {
		response.Response(ctx, http.StatusBadRequest, 400, nil, "密码错误")
		return
	}

	// 发放token
	token, err := common.ReleaseToken(usertable)
	if err != nil {
		response.Response(ctx, http.StatusInternalServerError, 500, nil, "系统错误")
		log.Printf("token generate error : %v", err)
		return
	}

	// 返回结果
	response.Success(ctx, gin.H{"token": token}, "登录成功")
}

// 登录获取用户信息
func Info(ctx *gin.Context) {
	user, _ := ctx.Get("user")
	response.Success(ctx, gin.H{"user": dto.ToUserDto(user.(model.UserTable))}, "ok")
}

// 验证电话是否存在
func isUserNameExist(db *gorm.DB, user_name string) bool {
	var UserTable model.UserTable
	db.Unscoped().Where("user_name = ?", user_name).First(&UserTable)

	return UserTable.ID != 0
}
