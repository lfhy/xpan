package client

import (
	"github.com/lfhy/xpan/auth"
	"github.com/lfhy/xpan/utils"
)

type Client struct {
}

func New(env ...*auth.AuthEnv) *Client {
	c := &Client{}
	if len(env) > 0 {
		c.SetAuth(env[0])
	}
	return c
}

func (c *Client) SetAuth(req *auth.AuthEnv) {
	auth.SetEnv(req)
}

func (c *Client) GetAuthCodeURL(req ...*auth.AuthCodeReq) string {
	return auth.GetAuthCodeURL(utils.GetOneOrDefault(req...))
}

func (c *Client) GetToken(code string, redirectUrl ...string) (*auth.GetTokenRes, error) {
	res, err := auth.GetToken(&auth.GetTokenReq{
		Code:        code,
		RedirectUri: utils.GetOneOrDefault(redirectUrl...),
	})
	if err != nil {
		return nil, err
	}
	c.SetAuth(&auth.AuthEnv{
		AccessToken:  res.AccessToken,
		RedirectUri:  utils.GetOneOrDefault(redirectUrl...),
		RefreshToken: res.RefreshToken,
	})
	return res, nil
}

func (c *Client) RefreshToken() (*auth.GetTokenRes, error) {
	res, err := auth.RefreshToken(&auth.RefreshTokenReq{})
	if err != nil {
		return nil, err
	}
	c.SetAuth(&auth.AuthEnv{
		AccessToken:  res.AccessToken,
		RefreshToken: res.RefreshToken,
	})
	return res, nil
}
