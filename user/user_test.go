package user_test

import (
	"testing"

	"github.com/lfhy/baidu-pan-client/test"
	"github.com/lfhy/baidu-pan-client/user"
)

func TestGetUserInfo(t *testing.T) {
	test.TestSetEnv(t)
	res, err := user.GetUserInfo(&user.UserInfoReq{})
	test.PrintRes(res, err)
}

func TestGetQuota(t *testing.T) {
	test.TestSetEnv(t)
	res, err := user.GetQuota(&user.QuotaReq{})
	test.PrintRes(res, err)
}
