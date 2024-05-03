package entities

import (
	"database/sql/driver"
	"strings"
)

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2024/4/23 下午11:55
 * @FilePath: model
 * @Description: 数据定义的通用结构和类型
 */

type RoleType int

const (
	PermissionAdmin   RoleType = 1 //管理员
	PermissionUser    RoleType = 2 //普通用户
	PermissionVisitor RoleType = 3 //游客
	PermissionDisable RoleType = 4 //封禁用户
)

func ParseRole(role RoleType) string {
	switch role {
	case PermissionAdmin:
		return "PermissionAdmin"
	case PermissionUser:
		return "PermissionUser"
	case PermissionVisitor:
		return "PermissionVisitor"
	case PermissionDisable:
		return "PermissionDisable"
	default:
		return "UnKnown"
	}
}

type SignOriginType int

const (
	SignQQ     SignOriginType = 1 //QQ登陆
	SignEmail  SignOriginType = 2 //邮箱登陆
	SignGithub SignOriginType = 3 //github登陆
)

func ParseSignOrigin(signOrigin SignOriginType) string {
	switch signOrigin {
	case SignQQ:
		return "SignQQ"
	case SignEmail:
		return "SignEmail"
	case SignGithub:
		return "SignGithub"
	default:
		return "UnKnown"
	}
}

type Array []string

func (array *Array) Scan(value interface{}) error {
	v, _ := value.([]byte)
	if string(v) == "" {
		*array = []string{}
		return nil
	}
	*array = strings.Split(string(v), "\n")
	return nil
}

func (array Array) Value() (driver.Value, error) {
	//将数字转换为值
	return strings.Join(array, "\n"), nil
}
