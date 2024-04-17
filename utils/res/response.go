package res

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2024/4/14 下午11:57
 * @FilePath: response
 * @Description: 封装response结构
 */

type Response struct {
	Code int    `json:"code"`
	Data any    `json:"data"`
	Msg  string `json:"msg"`
}

const (
	SUCCESS = 0
	ERROR   = -1
)

// Result 通用结果返回函数
func Result(c *gin.Context, code int, data any, msg string) {
	c.JSON(http.StatusOK, Response{
		Code: code,
		Data: data,
		Msg:  msg,
	})
}

// Ok 返回成功，但不携带任何信息
func Ok(c *gin.Context) {
	Result(c, SUCCESS, map[string]interface{}{}, "操作成功!")
}

// OkWithData 返回成功，携带数据
func OkWithData(c *gin.Context, data any) {
	Result(c, SUCCESS, data, "操作成功!")
}

// OkWithMsg 返回成功，携带消息
func OkWithMsg(c *gin.Context, msg string) {
	Result(c, SUCCESS, map[string]interface{}{}, msg)
}
