package auth

import "github.com/lfhy/baidu-pan-client/types"

type AuthCodeReq struct {
	ResponseType string `query:"response_type" default:"code"`
	ClientId     string `query:"client_id" default:"$CLIENT_ID"`
	RedirectUri  string `query:"redirect_uri" default:"$REDIRECT_URI"`
	Scope        string `query:"scope" default:"basic,netdisk"`
	DeviceId     string `query:"device_id"`
}

func GetAuthCodeURL(req *AuthCodeReq) string {
	query, _ := types.GetReqParams(req)
	return BaseURL + AuthRoute + "?" + query
}
