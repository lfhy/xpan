package file

import (
	"github.com/lfhy/baidu-pan-client/http"
	"github.com/lfhy/baidu-pan-client/types"
)

type ListReq struct {
	Dir       string          `query:"dir" default:"/"`      //	需要list的目录，以/开头的绝对路径, 默认为/
	Order     types.ListOrder `query:"name" default:"name"`  // 	排序字段：默认为name； time按修改时间排序；name表示按文件名称排序 size表示按文件大小排序。
	Desc      types.BoolInt   `query:"desc"`                 // 默认为升序 设置为true实现降序 （注：排序的对象是当前目录下所有文件，不是当前分页下的文件）
	Start     int             `query:"start" default:"0"`    //	起始位置，从0开始
	Limit     int             `query:"limit" default:"1000"` // 查询数目，默认为1000，建议最大不超过1000
	Web       types.BoolInt   `query:"web"`                  // 返回dir_empty属性和缩略图数据；不传该参数，则不返回缩略图地址
	Folder    types.BoolInt   `query:"folder"`               //	是否只返回文件夹且属性只返回path字段
	Showempty types.BoolInt   `query:"showempty"`            // 是否返回dir_empty属性，0 不返回，1 返回
}

type Thumbs struct {
	Icon string `json:"icon"`
	Url1 string `json:"url1"`
	Url2 string `json:"url2"`
	Url3 string `json:"url3"`
}

func List(req *ListReq) ([]*ListItem, error) {
	api := &http.Request[*ListReq, *ListRes]{
		AccessToken: types.AccessToken,
		BaseURL:     types.PanBaseURL,
		HTTPMethod:  http.GET,
		Method:      "list",
		Request:     req,
		Route:       types.FileRoute,
	}
	res, err := api.Do()
	if err != nil {
		return nil, err
	}
	return res.List, nil
}
