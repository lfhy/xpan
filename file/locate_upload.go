package file

import (
	"github.com/lfhy/baidu-pan-client/http"
	"github.com/lfhy/baidu-pan-client/types"
)

type LocateUploadReq struct {
	AppId         int    `query:"appid" default:"250528"`       // 应用ID，本接口固定为250528
	UploadVersion string `query:"upload_version" default:"2.0"` // 版本号，本接口固定为2.0
	Path          string `query:"path"`                         // 上传后使用的文件绝对路径，需要urlencode
	UploadId      string `query:"uploadid"`                     // 上传ID
}

type Server struct {
	Server string `json:"server"`
}

// 上传文件数据时，需要先通过此接口获取上传域名。可使用返回结果servers字段中的 https 协议的任意一个域名
type LocateUploadRes struct {
	BakServer   []Server   `json:"bak_server"`
	BakServers  []Server   `json:"bak_servers"`
	ClientIp    string     `json:"client_ip"`
	ExpireSec   int        `json:"expire"`
	Host        string     `json:"host"`
	Newno       string     `json:"newno"`
	QuicServer  []Server   `json:"quic_server"`
	QuicServers []Server   `json:"quic_servers"`
	Server      []Server   `json:"server"`
	ServerTime  types.Time `json:"server_time"`
	Servers     []Server   `json:"servers"`
	Sl          int        `json:"sl"`
}

func LocateUpload(req *LocateUploadReq) (*LocateUploadRes, error) {
	api := &http.Request[*LocateUploadReq, *LocateUploadRes]{
		AccessToken: types.AccessToken,
		BaseURL:     types.UploadBaseURL,
		HTTPMethod:  http.GET,
		Method:      "locateupload",
		Request:     req,
		Route:       types.UploadRoute,
	}
	return api.Do()
}
