/*
 * @Author: mysondrink
 * @Date: 2023-12-20 09:40:12
 * @Last Modified by: mysondrink
 * @Last Modified time: 2023-12-20 09:53:03
 * @Description:  用户数据传输对象定义
 */
package dto

import (
	"qt1218-backend/model"
	"time"
)

type UserDto struct {
	User_name string    `json:"user_name"`
	User_code string    `json:"user_code"`
	Status    int       `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// 反序列化
func ToUserDto(user model.UserTable) UserDto {
	return UserDto{
		User_name: user.User_name,
		User_code: user.User_code,
	}
}
