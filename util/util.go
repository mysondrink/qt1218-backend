/*
 * @Author: mysondrink
 * @Date: 2023-12-20 10:37:08
 * @Last Modified by:   mysondrink
 * @Last Modified time: 2023-12-20 10:37:08
 * @Description:  随机数生成工具
 */
package util

import "math/rand"

func RandomString(n int) string {
	var letters = []byte("qwertyuiopasdfghjklzxcvbnmQWERTYUIOPASDFGHJKLZXCVBNM")
	result := make([]byte, n)

	for i := range result {
		result[i] = letters[rand.Intn(len(letters))]
	}

	return string(result)
}
