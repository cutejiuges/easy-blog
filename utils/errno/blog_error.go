package errno

import (
	"errors"
	"fmt"
	"runtime"
)

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2024/4/18 上午9:25
 * @FilePath: blog_error
 * @Description: 自定义一些常见错误
 */

type BlogError struct {
	Code int32  `json:"code"`
	Msg  string `json:"msg"`
	fn   string
	line int
}

func (err *BlogError) Error() string {
	return fmt.Sprintf("%s:%d, code = %d, msg = %s", err.fn, err.line, err.Code, err.Msg)
}

func NewBlogError(code int32) *BlogError {
	_, fn, line, _ := runtime.Caller(1)
	return &BlogError{
		Code: code,
		Msg:  ErrDesc[code],
		fn:   fn,
		line: line,
	}
}

func NewBlogErrorWithMsg(code int32, msg string) *BlogError {
	_, fn, line, _ := runtime.Caller(1)
	return &BlogError{
		Code: code,
		Msg:  msg,
		fn:   fn,
		line: line,
	}
}

func IsSameError(err error, code int32) bool {
	if err == nil && code == SUCCESS {
		return true
	}
	var t *BlogError
	switch {
	case errors.As(err, &t):
		if t == nil {
			return code == SUCCESS
		}
		return t.Code == code
	}
	return false
}
