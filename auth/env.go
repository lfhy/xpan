package auth

import "github.com/lfhy/baidu-pan-client/types"

type AuthEnv struct {
	ClientId     string
	ClientSecret string
	RedirectUri  string
	AccessToken  string
	RefreshToken string
}

func SetEnv(env *AuthEnv) {
	if env.ClientId != "" {
		types.ClientId = env.ClientId
	}
	if env.ClientSecret != "" {
		types.ClientSecret = env.ClientSecret
	}
	if env.RedirectUri != "" {
		types.RedirectUri = env.RedirectUri
	}
	if env.AccessToken != "" {
		types.AccessToken = env.AccessToken
	}
	if env.RefreshToken != "" {
		types.RefreshToken = env.RefreshToken
	}
}
