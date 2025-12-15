package client

import (
	"github.com/lfhy/baidu-pan-client/user"
	"github.com/lfhy/baidu-pan-client/utils"
)

func (c *Client) Quota(req ...*user.QuotaReq) (*user.QuotaRes, error) {
	return user.GetQuota(utils.GetOneOrDefault(req...))
}

func (c *Client) UserInfo(v2 ...bool) (*user.UserInfoRes, error) {
	var v string
	if v2 != nil && v2[0] {
		v = "v2"
	}
	return user.GetUserInfo(&user.UserInfoReq{
		VipVersion: v,
	})
}
