package user

import (
	"github.com/lfhy/xpan/http"
	"github.com/lfhy/xpan/types"
)

type QuotaReq struct {
	Checkfree   types.BoolInt `query:"checkfree"`   // 是否检查免费信息，0为不查，1为查，默认为0
	Checkexpire types.BoolInt `query:"checkexpire"` // 是否检查过期信息，0为不查，1为查，默认为0
}

type QuotaRes struct {
	Total  types.SizeB `json:"total"`  // 总空间大小，单位B
	Expire bool        `json:"expire"` // 7天内是否有容量到期
	Used   types.SizeB `json:"used"`   // 已使用大小，单位B
	Free   types.SizeB `json:"free"`   // 免费容量，单位B
}

func GetQuota(req *QuotaReq) (*QuotaRes, error) {
	api := http.Request[*QuotaReq, *QuotaRes]{
		BaseURL:     types.PanBaseURL,
		Route:       types.QuotaRoute,
		AccessToken: types.AccessToken,
		HTTPMethod:  http.GET,
		Request:     req,
	}
	return api.Do()
}
