package auth_test

import (
	"fmt"
	"testing"

	"github.com/lfhy/xpan/auth"
	"github.com/lfhy/xpan/test"
)

func TestGetAuthReq(t *testing.T) {
	test.TestSetEnv(t)
	authUrl := auth.GetAuthCodeURL(&auth.AuthCodeReq{})
	fmt.Printf("authUrl: %v\n", authUrl)
}

func TestGetToken(t *testing.T) {
	test.TestSetEnv(t)
	// https://openapi.baidu.com/oauth/2.0/token?grant_type=authorization_code&code=d5a53cd0ca7799d033399487b23ec992&client_id=EVaI5x0U6lEmP125G0Su55ROEXZtItdD&client_secret=VPgfmrt8UBM5kgkeUemwRVmr5AjhFuEV&redirect_uri=oob
	res, err := auth.GetToken(&auth.GetTokenReq{Code: "5355bc4353b7563885135b51a123d7bf"})
	test.PrintRes(res, err)
}
