package test

import (
	"reflect"
	"testing"

	"github.com/lfhy/flag"

	glog "log"

	"github.com/lfhy/xpan/log"
	"github.com/lfhy/xpan/types"
)

type TestLogger struct {
}

func (t *TestLogger) Printf(f string, v ...any) {
	glog.Printf(f, v...)
}

func (t *TestLogger) Println(v ...any) {
	glog.Println(v...)
}

func ReadConfig() {
	f := flag.NewFlagSet("test", flag.ContinueOnError)
	f.String("c", "../config.toml", "配置文件")
	f.StringConfigVar(&types.ClientId,
		"client_id",
		"auth", "client_id",
		"",
		"client_id",
	)
	f.StringConfigVar(&types.ClientSecret,
		"client_secret",
		"auth", "client_secret",
		"",
		"client_secret",
	)
	f.StringConfigVar(&types.RedirectUri,
		"redirect_uri",
		"auth", "redirect_uri",
		"",
		"redirect_uri",
	)
	f.StringConfigVar(&types.AccessToken,
		"access_token",
		"auth", "access_token",
		"",
		"access_token",
	)
	f.StringConfigVar(&types.RefreshToken,
		"refresh_token",
		"auth", "refresh_token",
		"",
		"refresh_token",
	)
	f.Parse(nil)
}

func TestSetEnv(t *testing.T) {
	ReadConfig()
	log.SetLogger(&TestLogger{})
}

// dereferenceValue 递归解引用值，处理指针、切片和结构体中的指针
func dereferenceValue(v reflect.Value) interface{} {
	switch v.Kind() {
	case reflect.Ptr:
		// 如果是nil指针，直接返回nil
		if v.IsNil() {
			return nil
		}
		// 解引用指针
		return dereferenceValue(v.Elem())

	case reflect.Struct:
		// 处理结构体，递归处理每个字段
		result := make(map[string]interface{})
		t := v.Type()
		for i := 0; i < v.NumField(); i++ {
			field := v.Field(i)
			fieldType := t.Field(i)
			// 获取字段名（考虑结构体标签）
			fieldName := fieldType.Name
			if tag := fieldType.Tag.Get("json"); tag != "" {
				fieldName = tag
			}
			result[fieldName] = dereferenceValue(field)
		}
		return result

	case reflect.Slice:
		// 处理切片
		if v.Len() == 0 {
			return v.Interface()
		}
		// 创建新的切片存储处理后的值
		result := make([]interface{}, v.Len())
		for i := 0; i < v.Len(); i++ {
			result[i] = dereferenceValue(v.Index(i))
		}
		return result

	case reflect.Array:
		// 处理数组
		if v.Len() == 0 {
			return v.Interface()
		}
		result := make([]interface{}, v.Len())
		for i := 0; i < v.Len(); i++ {
			result[i] = dereferenceValue(v.Index(i))
		}
		return result

	case reflect.Map:
		// 处理映射
		result := make(map[interface{}]interface{})
		for _, key := range v.MapKeys() {
			result[key.Interface()] = dereferenceValue(v.MapIndex(key))
		}
		return result

	default:
		// 其他类型直接返回值
		return v.Interface()
	}
}

// formatValue 格式化值用于打印
func formatValue(v interface{}) interface{} {
	if v == nil {
		return nil
	}

	rv := reflect.ValueOf(v)
	return dereferenceValue(rv)
}

func PrintRes(res any, err error) {
	if err != nil {
		log.Println("发生错误:", err)
	} else {
		formatted := formatValue(res)
		log.Printf("响应结果:%+v", formatted)
	}
}
