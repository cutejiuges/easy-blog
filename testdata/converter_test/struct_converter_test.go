package converter_test

import (
	"cutejiuges/easy-blog/utils/converter"
	"log"
	"testing"
)

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2024/5/4 下午6:55
 * @FilePath: struct_converter_test
 * @Description:
 */
type personA struct {
	Age   int
	Name  string
	Score float64
}

type personB struct {
	Age  int
	Name string
	Addr string
}

func TestStructConverter(t *testing.T) {
	pa := personA{
		Age:   11,
		Name:  "hello",
		Score: 33.3,
	}
	pb := personB{
		Age:  22,
		Name: "world",
		Addr: "浙江杭州",
	}
	converter.StructConverter(&pb, &pa)
	log.Println(pb)
}
