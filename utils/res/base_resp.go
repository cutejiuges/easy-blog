package res

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2024/4/18 上午8:49
 * @FilePath: base_resp
 * @Description: 返回结构体的通用部分
 */

type BaseResp struct {
	StatusCode    int32             `json:"status_code"`
	StatusMessage string            `json:"status_message"`
	Extra         map[string]string `json:"extra"`
}

func NewBaseResp() *BaseResp {
	return &BaseResp{
		StatusCode:    SUCCESS,
		StatusMessage: "",
	}
}
