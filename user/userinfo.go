package user

import (
	"github.com/lfhy/baidu-pan-client/http"
	"github.com/lfhy/baidu-pan-client/types"
)

type UserInfoReq struct {
	VipVersion string `query:"vip_version"` // 可以选择v2
}

type UserInfoRes struct {
	BaiduName   string        `json:"baidu_name"`   // 百度账号
	NetDiskName string        `json:"netdisk_name"` // 网盘账号
	AvatarUrl   string        `json:"avatar_url"`   // 头像地址
	VipType     types.VipType `json:"vip_type"`     // 会员类型。旧含义0普通用户、1普通会员、2超级会员。指定vip_version字段为v2时，0普通用户、1会员VIP用户、2超级会员SVIP用户
	Uk          int           `json:"uk"`           // 用户ID
}

func GetUserInfo(req *UserInfoReq) (*UserInfoRes, error) {
	api := &http.Request[*UserInfoReq, *UserInfoRes]{
		AccessToken: types.AccessToken,
		BaseURL:     types.PanBaseURL,
		HTTPMethod:  http.GET,
		Method:      "uinfo",
		Request:     req,
		Route:       types.NasRoute,
	}
	return api.Do()
}
