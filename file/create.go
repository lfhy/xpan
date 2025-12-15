package file

import (
	"github.com/lfhy/xpan/http"
	"github.com/lfhy/xpan/types"
)

type CreateReq struct {
	Path       string                `body:"path"`        // 上传后使用的文件绝对路径，需要urlencode，需要与预上传precreate接口中的path保持一致
	Size       types.SizeB           `body:"size"`        // 文件或目录的大小，必须要和文件真实大小保持一致，需要与预上传precreate接口中的size保持一致
	IsDir      types.BoolInt         `body:"isdir"`       // 是否目录，0 文件、1 目录，需要与预上传precreate接口中的isdir保持一致
	BlockList  []string              `body:"block_list"`  // 文件各分片md5数组的json串 需要与预上传precreate接口中的block_list保持一致，同时对应分片上传superfile2接口返回的md5，且要按照序号顺序排列，组成md5数组的json串。
	UploadId   string                `body:"uploadid"`    // 预上传precreate接口下发的uploadid
	RType      types.FileCreateRType `body:"rtype"`       // 文件命名策略
	LocalCtime types.Time            `body:"local_ctime"` // 客户端创建时间(精确到秒)，默认为当前时间戳
	LocalMtime types.Time            `body:"local_mtime"` // 客户端修改时间(精确到秒)，默认为当前时间戳
	ZipQuality int                   `body:"zip_quality"` // 图片压缩程度，有效值50、70、100（带此参数时，zip_sign 参数需要一并带上）
	ZipSign    string                `body:"zip_sign"`    // 未压缩原始图片文件真实md5（带此参数时，zip_quality 参数需要一并带上）
	IsRevision types.BoolInt         `body:"is_revision"` // 是否需要多版本支持
	Mode       types.CreateMode      `body:"mode"`        // 上传方式
	ExifInfo   *ExifInfo             `body:"exif_info"`   // orientation、width、height、recovery为必传字段，其他字段如果没有可以不传
}

type ExifInfo struct {
	Height            int    `json:"height"`
	DateTimeOriginal  string `json:"date_time_original"`
	Model             string `json:"model"`
	Width             int    `json:"width"`
	DateTimeDigitized string `json:"date_time_digitized"`
	DateTime          string `json:"date_time"`
	Orientation       int    `json:"orientation"`
	Recovery          int    `json:"recovery"`
}

type CreateRes struct {
	FsId           uint64             `json:"fs_id"`           // 文件在云端的唯一标识ID
	Md5            string             `json:"md5"`             // 文件的MD5，只有提交文件时才返回，提交目录时没有该值
	ServerFilename string             `json:"server_filename"` // 文件名
	Category       types.FileCategory `json:"category"`        // 分类类型
	Path           string             `json:"path"`            // 上传后使用的文件绝对路径
	Size           types.SizeB        `json:"size"`            // 文件大小，单位B
	Ctime          types.Time         `json:"ctime"`           // 文件创建时间
	Mtime          types.Time         `json:"mtime"`           // 文件修改时间
	IsDir          types.BoolInt      `json:"isdir"`           // 是否目录
}

// 创建文件 合块
func Create(req *CreateReq) (*CreateRes, error) {
	api := &http.Request[*CreateReq, *CreateRes]{
		AccessToken: types.AccessToken,
		BaseURL:     types.PanBaseURL,
		HTTPMethod:  http.POST,
		Method:      "create",
		Request:     req,
		Route:       types.FileRoute,
	}
	return api.Do()
}
