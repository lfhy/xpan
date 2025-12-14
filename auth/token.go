package auth

import (
	"github.com/lfhy/baidu-pan-client/http"
	"github.com/lfhy/baidu-pan-client/types"
)

type GetTokenReq struct {
	ClientId     string `query:"client_id" default:"$CLIENT_ID"`
	ClientSecret string `query:"client_secret" default:"$CLIENT_SECRET"`
	Code         string `query:"code"`
	GrantType    string `query:"grant_type" default:"authorization_code"`
	RedirectUri  string `query:"redirect_uri" default:"$REDIRECT_URI"`
}

type GetTokenRes struct {
	AccessToken   string `json:"access_token"`
	ExpiresIn     int64  `json:"expires_in"`
	RefreshToken  string `json:"refresh_token"`
	Scope         string `json:"scope"`
	SessionKey    string `json:"session_key"`
	SessionSecret string `json:"session_secret"`
}

func GetToken(req *GetTokenReq) (*GetTokenRes, error) {
	api := http.Request[*GetTokenReq, *GetTokenRes]{
		BaseURL:    types.AuthBaseURL,
		Route:      types.TokenRoute,
		HTTPMethod: http.GET,
		Request:    req,
	}
	return api.Do()
}

type RefreshTokenReq struct {
	ClientId     string `query:"client_id" default:"$CLIENT_ID"`
	ClientSecret string `query:"client_secret" default:"$CLIENT_SECRET"`
	GrantType    string `query:"grant_type" default:"refresh_token"`
	RefreshToken string `query:"refresh_token" default:"$REFRESH_TOKEN"`
}

func RefreshToken(req *RefreshTokenReq) (*GetTokenRes, error) {
	api := http.Request[*RefreshTokenReq, *GetTokenRes]{
		BaseURL:    types.AuthBaseURL,
		Route:      types.TokenRoute,
		HTTPMethod: http.GET,
		Request:    req,
	}
	return api.Do()
}
