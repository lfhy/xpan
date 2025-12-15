package file

import (
	"github.com/lfhy/xpan/http"
	"github.com/lfhy/xpan/types"
)

// NOTE:分页???
type SearchReq struct {
	Key       string             `query:"key"`                   // 搜索关键字，最大30字符（UTF8格式）
	Dir       string             `query:"dir" default:"/"`       // 搜索目录，默认根目录
	Category  types.FileCategory `query:"category"`              //	文件类型，1 视频、2 音频、3 图片、4 文档、5 应用、6 其他、7 种子
	Recursion types.BoolInt      `query:"recursion" default:"0"` // 递归获取全部子目录下文件列表，默认为不需要
	Web       types.BoolInt      `query:"web"`                   // 返回dir_empty属性和缩略图数据；不传该参数，则不返回缩略图地址
	DeviceId  string             `query:"device_id"`             // 设备ID，硬件设备必传
}

type SearchRes struct {
	HasMore types.BoolInt `json:"has_more"` // 是否还有下一页，0表示无，1表示有
	List    []*ListItem   `json:"list"`     // 文件列表
}

func Search(req *SearchReq) (*SearchRes, error) {
	api := &http.Request[*SearchReq, *SearchRes]{
		AccessToken: types.AccessToken,
		BaseURL:     types.PanBaseURL,
		HTTPMethod:  http.GET,
		Method:      "search",
		Request:     req,
		Route:       types.FileRoute,
	}
	return api.Do()
}
