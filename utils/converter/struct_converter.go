package converter

import "reflect"

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2024/5/4 下午6:28
 * @FilePath: struct_converter
 * @Description: 完成2个结构体之间的互转
 */

// StructConverter 进行结构体同名字段的converter，从source结构体转换到binding结构体
func StructConverter(binding any, source any) {
	bindingVal := reflect.ValueOf(binding).Elem()
	sourceVal := reflect.ValueOf(source).Elem()
	sourceValType := sourceVal.Type()
	for i, n := 0, sourceVal.NumField(); i < n; i++ {
		//在目标结构体中查询同名字段，有则修改
		name := sourceValType.Field(i).Name
		if ok := bindingVal.FieldByName(name).IsValid(); ok {
			bindingVal.FieldByName(name).Set(reflect.ValueOf(sourceVal.Field(i).Interface()))
		}
	}
}
