package file

import (
	"github.com/lfhy/xpan/http"
	"github.com/lfhy/xpan/types"
)

type FilemetasReq struct {
	FsIds     []uint64      `query:"fsids"`      // 文件id数组，数组中元素是uint64类型，数组大小上限是：100
	Dlink     types.BoolInt `query:"dlink"`      // 是否需要下载地址 获取到dlink后，参考下载文档进行下载操作
	Path      string        `query:"path"`       // 查询共享目录或专属空间内文件时需要。含义并非待查询文件的路径，而是查询特定目录下文件的开关参数。共享目录格式： /uk-fsid 其中uk为共享目录创建者id， fsid对应共享目录的fsid 专属空间格式：/_pcs_.appdata/xpan/。此参数生效时，不会返回正常目录下文件的信息。
	Thumb     types.BoolInt `query:"thumb"`      // 是否需要缩略图地址
	Extra     types.BoolInt `query:"extra"`      //	图片是否需要拍摄时间、原图分辨率等其他信息
	Needmedia types.BoolInt `query:"needmedia"`  // 视频是否需要展示时长信息，返回 duration 信息时间单位为秒 （s），转换为向上取整。
	Detail    types.BoolInt `query:"detail"`     // 视频是否需要展示长，宽等信息。返回信息在media_info字段内
	DeviceId  string        `query:"device_id"`  // 设备ID，硬件设备必传
	FromApaas types.BoolInt `query:"from_apaas"` // 为下载地址(dlink)附加极速流量权益。用户通过此dlink产生下载行为时，消耗等同于文件大小的极速流量权益。此权益为付费权益，否则此参数无效。
}

type FilemetasItem struct {
	FsId        uint64             `json:"fs_id"`        // 文件在云端的唯一标识ID
	Category    types.FileCategory `json:"category"`     //	文件类型，1 视频、2 音频、3 图片、4 文档、5 应用、6 其他、7 种子
	Dlink       string             `json:"dlink"`        // 	文件下载地址，参考下载文档进行下载操作。注意unicode解码处理。
	Filename    string             `json:"filename"`     // 文件名
	IsDir       types.BoolInt      `json:"isdir"`        //	是否为目录，0 文件、1 目录
	ServerMtime types.Time         `json:"server_mtime"` //	文件在服务器修改时间
	ServerCtime types.Time         `json:"server_ctime"` //	文件在服务器创建时间
	LocalMtime  *types.Time        `json:"local_mtime"`  // 文件在客户端修改时间
	LocalCtime  *types.Time        `json:"local_ctime"`  // 文件在客户端创建时间
	Size        types.SizeB        `json:"size"`         //	文件大小，单位B
	PhotoFileInfo
	VideoFileInfo
}

type PhotoFileInfo struct {
	Thumbs      *Thumbs    `json:"thumbs"`      // 图片缩略图 包含4个尺寸
	Height      *int       `json:"height"`      // 图片高度
	Width       *int       `json:"width"`       // 图片宽度
	DateTaken   types.Time `json:"date_taken"`  // 图片拍摄时间
	Orientation *string    `json:"orientation"` // 图片旋转方向信息
}

type VideoFileInfo struct {
	MediaInfo *MediaInfo `json:"media_info"` // 视频信息
}

type MediaInfo struct {
	Channels   int     `json:"channels"`
	Duration   int     `json:"duration"`
	DurationMs int     `json:"duration_ms"`
	ExtraInfo  string  `json:"extra_info"`
	FileSize   string  `json:"file_size"`
	FrameRate  float64 `json:"frame_rate"`
	Height     int     `json:"height"`
	MetaInfo   string  `json:"meta_info"`
	Resolution string  `json:"resolution"`
	Rotate     int     `json:"rotate"`
	SampleRate int     `json:"sample_rate"`
	UseSegment int     `json:"use_segment"`
	Width      int     `json:"width"`
}

type FilemetasRes struct {
	Names any              `json:"names"` //如果查询共享目录，该字段为共享目录文件上传者的uk和账户名称
	List  []*FilemetasItem `json:"list"`
}

func Filemetas(req *FilemetasReq) (*FilemetasRes, error) {
	api := &http.Request[*FilemetasReq, *FilemetasRes]{
		AccessToken: types.AccessToken,
		BaseURL:     types.PanBaseURL,
		HTTPMethod:  http.GET,
		Method:      "filemetas",
		Request:     req,
		Route:       types.MultimediaRoute,
	}
	return api.Do()
}
