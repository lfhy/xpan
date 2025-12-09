package types

import (
	"fmt"
	"io"
	"net/url"
	"reflect"
	"strings"
)

func GetReqParams(req any) (query string, body io.Reader) {
	if req == nil {
		return
	}
	// 如果req是指针则取req指向的对象
	for reflect.TypeOf(req).Kind() == reflect.Pointer {
		req = reflect.ValueOf(req).Elem().Interface()
	}
	// 判断是否req是结构体
	if reflect.TypeOf(req).Kind() != reflect.Struct {
		return
	}
	// 解析结构体内容
	params := make(url.Values)
	bodys := make(url.Values)
	for i := 0; i < reflect.TypeOf(req).NumField(); i++ {
		field := reflect.TypeOf(req).Field(i)
		value := getENV(field.Tag.Get("default"))
		// 判断tags
		query := field.Tag.Get("query")
		if query != "" {
			data := reflect.ValueOf(req).Field(i)
			if !data.IsZero() {
				if data.Kind() == reflect.Array || data.Kind() == reflect.Slice {
					var arrayData []string
					for i := 0; i < data.Len(); i++ {
						arrayData = append(arrayData, fmt.Sprint(data.Index(i).Interface()))
					}
					value = "[" + strings.Join(arrayData, ",") + "]"
				} else {
					value = fmt.Sprint(data.Interface())
				}
			}
			if value == "" {
				continue
			}
			params.Add(query, value)
		}
		body := field.Tag.Get("body")
		if body != "" {
			data := reflect.ValueOf(req).Field(i)
			if !data.IsZero() {
				if data.Kind() == reflect.Array || data.Kind() == reflect.Slice {
					var arrayData []string
					for i := 0; i < data.Len(); i++ {
						arrayData = append(arrayData, fmt.Sprint(data.Index(i).Interface()))
					}
					value = "[" + strings.Join(arrayData, ",") + "]"
				} else {
					value = fmt.Sprint(data.Interface())
				}
			}
			if value == "" {
				continue
			}
			bodys.Add(body, value)
		}
	}
	query = params.Encode()
	if len(bodys) > 0 {
		body = strings.NewReader(bodys.Encode())
	}
	return
}
