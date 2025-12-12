package file

import (
	"github.com/lfhy/baidu-pan-client/http"
	"github.com/lfhy/baidu-pan-client/types"
)

type PreCreateReq struct {
	Path      string        `body:"path"`                 // 上传后使用的文件绝对路径
	Size      types.SizeB   `body:"size"`                 // 文件和目录两种情况：上传文件时，表示文件的大小，单位B；上传目录时，表示目录的大小，目录的话大小默认为0
	IsDir     types.BoolInt `body:"isdir"`                // 是否为目录，0 文件，1 目录
	BlockList []string      `body:"block_list"`           // 文件各分片MD5数组的json串。block_list的含义如下，如果上传的文件小于4MB，其md5值（32位小写）即为block_list字符串数组的唯一元素；如果上传的文件大于4MB，需要将上传的文件按照4MB大小在本地切分成分片，不足4MB的分片自动成为最后一个分片，所有分片的md5值（32位小写）组成的字符串数组即为block_list。
	AutoInit  int           `body:"autoinit" default:"1"` // 固定值1

	// 1 表示当path冲突时，进行重命名
	// 2 表示当path冲突且block_list不同时，进行重命名
	// 3 当云端存在同名文件时，对该文件进行覆盖
	RType types.FileCreateRType `body:"rtype"` // 文件命名策略。

	UploadId   string `body:"uploadid"`    // 上传ID
	ContentMd5 string `body:"content-md5"` // 文件MD5，32位小写
	SliceMd5   string `body:"slice-md5"`   // 文件校验段的MD5，32位小写，校验段对应文件前256KB
	LocalCtime string `body:"local_ctime"` // 客户端创建时间， 默认为当前时间戳
	LocalMtime string `body:"local_mtime"` // 客户端修改时间，默认为当前时间戳
}

type PreCreateRes struct {
	Path       string `json:"path"`        // 文件的绝对路径
	UploadId   string `json:"uploadid"`    // 上传唯一ID标识此上传任务
	ReturnType int    `json:"return_type"` // 返回类型，系统内部状态字段
	BlockList  []int  `json:"block_list"`  // 需要上传的分片序号列表，索引从0开始
}

func PreCreate(req *PreCreateReq) (*PreCreateRes, error) {
	api := &http.Request[*PreCreateReq, *PreCreateRes]{
		AccessToken: types.AccessToken,
		BaseURL:     types.PanBaseURL,
		HTTPMethod:  http.POST,
		Method:      "precreate",
		Request:     req,
		Route:       types.SuperFileRoute,
	}
	return api.Do()
}
