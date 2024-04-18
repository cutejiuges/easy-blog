package errno

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2024/4/18 上午9:04
 * @FilePath: errno
 * @Description: 存放错误码和错误描述信息
 */

const (
	SUCCESS      = 0
	ParamIllegal = -10001
	SystemError  = -10002
)

var ErrDesc = map[int32]string{
	int32(SUCCESS):      "success",
	int32(ParamIllegal): "请求参数非法",
	int32(SystemError):  "系统内部异常",
}
