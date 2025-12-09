package file_test

import (
	"testing"

	"github.com/lfhy/baidu-pan-client/file"
	"github.com/lfhy/baidu-pan-client/test"
)

func TestFileList(t *testing.T) {
	test.TestSetEnv(t)
	res, err := file.List(&file.ListReq{
		Dir: "/",
	})
	test.PrintRes(res, err)
}

func TestFileListAll(t *testing.T) {
	test.TestSetEnv(t)
	res, err := file.ListAll(&file.ListAllReq{
		Path: "/",
	})
	test.PrintRes(res, err)
}

func TestFilemeta(t *testing.T) {
	test.TestSetEnv(t)
	res, err := file.ListAll(&file.ListAllReq{
		Path: "/",
	})
	if err == nil {
		var req file.FilemetasReq
		for _, file := range res.List {
			req.FsIds = append(req.FsIds, file.FsId)
		}
		res, err := file.Filemetas(&req)
		test.PrintRes(res, err)
	}
}

func TestSearch(t *testing.T) {
	test.TestSetEnv(t)
	res, err := file.Search(&file.SearchReq{
		Dir: "/",
		Key: "test",
	})
	test.PrintRes(res, err)
}
