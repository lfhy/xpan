package user_test

import (
	"testing"

	"github.com/lfhy/xpan/test"
	"github.com/lfhy/xpan/user"
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
