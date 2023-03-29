package utils

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"reflect"
	"strconv"
)

type QueryFunc = func(string) string

func queryStruct[Q any](qFunc QueryFunc, q *Q) error {
	t := reflect.TypeOf(*q)
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		tag, ok := field.Tag.Lookup("query")
		name := field.Name
		var value string
		if ok {
			value = qFunc(tag)
		} else {
			value = qFunc(name)
		}

		// 如果查询的参数为空 则直接忽略
		if value == "" {
			continue
		}

		val := reflect.ValueOf(q)
		typeName := field.Type.Name()
		switch typeName {
		case "int":
			parseInt, err := strconv.ParseInt(value, 10, 10)
			if err != nil {
				return err
			}
			val.Elem().FieldByName(name).SetInt(parseInt)
		case "string":
			val.Elem().FieldByName(name).SetString(value)
			// 有需要自行添加...
		default:
			return fmt.Errorf("%s field, of type %s, cannot be converted", name, typeName)
		}
	}
	return nil
}

// QueryStruct 将 url 上的参数 绑定的结构体内
func QueryStruct[Q any](c *gin.Context, q *Q) error {
	return queryStruct[Q](func(name string) string {
		return c.Query(name)
	}, q)
}

// URLSearchParams 搜索url参数 赋值给 q 参数
func URLSearchParams[Q any](q *Q, fun QueryFunc) error {
	return queryStruct[Q](fun, q)
}
