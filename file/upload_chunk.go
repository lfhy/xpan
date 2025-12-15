package file

import (
	"io"
	"strings"

	"github.com/lfhy/baidu-pan-client/http"
	"github.com/lfhy/baidu-pan-client/types"
)

type UploadChunkReq struct {
	Type     string    `query:"type" default:"tmpfile"` //	固定值 tmpfile
	Path     string    `query:"path"`                   // 上传后使用的文件绝对路径，需要与上一个阶段预上传precreate接口中的path保持一致
	UploadId string    `query:"uploadid"`               // 上一个阶段预上传precreate接口下发的uploadid
	PartSeq  int       `query:"partseq"`                // 文件分片的位置序号，从0开始，参考上一个阶段预上传precreate接口返回的block_list
	File     io.Reader `file:"file"`                    // 上传的文件内容
}

type UploadChunkRes struct {
	Md5 string `json:"md5"` // 云端的md5 不是文件本身的
}

// 上传主机名通过locate_upload接口获取
func UploadChunk(req *UploadChunkReq, uploadHost string) (*UploadChunkRes, error) {
	if !strings.HasPrefix(uploadHost, "http") {
		uploadHost = "https://" + uploadHost
	}
	api := &http.Request[*UploadChunkReq, *UploadChunkRes]{
		AccessToken: types.AccessToken,
		BaseURL:     uploadHost,
		HTTPMethod:  http.POST,
		Method:      "upload",
		Request:     req,
		Route:       types.SuperFileRoute,
	}
	return api.Do()
}
