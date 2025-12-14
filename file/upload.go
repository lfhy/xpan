package file

import (
	"io"

	"github.com/lfhy/baidu-pan-client/http"
	"github.com/lfhy/baidu-pan-client/types"
)

type UploadReq struct {
	Path  string          `query:"path"`  // 上传的文件绝对路径
	Ondup types.OndupMode `query:"ondup"` // 遇到重复文件的处理策略
	File  io.Reader       `file:"file"`   // 上传的文件内容
}

type UploadRes struct {
	Path  string      `json:"path"`  // 文件的绝对路径
	Size  types.SizeB `json:"size"`  // 文件大小，单位B
	Ctime types.Time  `json:"ctime"` // 文件创建时间
	Mtime types.Time  `json:"mtime"` // 文件修改时间
	Md5   string      `json:"md5"`   // 文件的MD5，只有提交文件时才返回，提交目录时没有该值
	FsId  uint64      `json:"fs_id"` // 文件在云端的唯一标识ID
}

// 单步上传
// 上传文件大小上限为2GB
// 此接口可能有一定限制，推荐使用主流接口分片上传
func Upload(req *UploadReq, uploadHost string) (*UploadRes, error) {
	api := &http.Request[*UploadReq, *UploadRes]{
		AccessToken: types.AccessToken,
		BaseURL:     uploadHost,
		HTTPMethod:  http.POST,
		Method:      "upload",
		Request:     req,
		Route:       types.UploadRoute,
	}
	return api.Do()
}
