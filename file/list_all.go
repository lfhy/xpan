package file

import (
	"github.com/lfhy/baidu-pan-client/http"
	"github.com/lfhy/baidu-pan-client/types"
)

type ListAllReq struct {
	Path      string          `query:"path" default:"/"`      // 需要list的目录，以/开头的绝对路径, 默认为/
	Recursion types.BoolInt   `query:"recursion" default:"0"` // 递归获取全部子目录下文件列表，默认为不需要
	Order     types.ListOrder `query:"name" default:"name"`   // 排序字段：默认为name； time按修改时间排序；name表示按文件名称排序 size表示按文件大小排序。
	Desc      types.BoolInt   `query:"desc"`                  // 默认为升序 设置为true实现降序 （注：排序的对象是当前目录下所有文件，不是当前分页下的文件）
	Start     int             `query:"start" default:"0"`     // 起始位置，从0开始
	Limit     int             `query:"limit" default:"1000"`  // 查询数目，默认为1000，建议最大不超过1000
	Ctime     types.Time      `query:"ctime"`                 // 文件上传时间，设置此参数，表示只返回上传时间大于ctime的文件
	Mtime     types.Time      `query:"mtime"`                 // 文件修改时间，设置此参数，表示只返回修改时间大于mtime的文件
	Web       types.BoolInt   `query:"web"`                   // 返回dir_empty属性和缩略图数据；不传该参数，则不返回缩略图地址
	DeviceId  string          `query:"device_id"`             // 设备ID，硬件设备必传
}

type ListItem struct {
	FsId        uint64             `json:"fs_id"`           // 文件在云端的唯一标识ID
	Path        string             `json:"path"`            // 文件的绝对路径
	Name        string             `json:"server_filename"` //	文件名称
	Size        types.SizeB        `json:"size"`            //	文件大小，单位B
	ServerMtime types.Time         `json:"server_mtime"`    //	文件在服务器修改时间
	ServerCtime types.Time         `json:"server_ctime"`    //	文件在服务器创建时间
	LocalMtime  types.Time         `json:"local_mtime"`     //	文件在客户端修改时间
	LocalCtime  types.Time         `json:"local_ctime"`     //	文件在客户端创建时间
	IsDir       types.BoolInt      `json:"isdir"`           //	是否为目录，0 文件、1 目录
	Category    types.FileCategory `json:"category"`        //	文件类型，1 视频、2 音频、3 图片、4 文档、5 应用、6 其他、7 种子
	ServerMd5   string             `json:"md5"`             //	云端哈希（非文件真实MD5），只有是文件类型时，该字段才存在
	DirEmpty    types.BoolInt      `json:"dir_empty"`       //	该目录是否存在子目录，只有请求参数web=1且该条目为目录时，该字段才存在， 0为存在， 1为不存在
	Thumbs      *Thumbs            `json:"thumbs"`          //	只有请求参数web=1且该条目分类为图片时，该字段才存在，包含三个尺寸的缩略图URL；不传web参数，则不返回缩略图地址
}

type ListRes struct {
	HasMore types.BoolInt `json:"has_more"` //	是否还有下一页，0表示无，1表示有
	Cursor  int           `json:"cursor"`   // 当还有下一页时，为下一次查询的起点
	List    []*ListItem   `json:"list"`     // 文件列表
}

func ListAll(req *ListAllReq) (*ListRes, error) {
	api := &http.Request[*ListAllReq, *ListRes]{
		AccessToken: types.AccessToken,
		BaseURL:     types.PanBaseURL,
		HTTPMethod:  http.GET,
		Method:      "listall",
		Request:     req,
		Route:       types.MultimediaRoute,
	}
	return api.Do()
}
