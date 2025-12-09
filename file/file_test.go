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

func TestCopyFile(t *testing.T) {
	test.TestSetEnv(t)
	res, err := file.Filemanager(&file.FilemanagerReq[file.FileCopyAndMoveItem]{
		Opera: file.Copy,
		// cp /test.txt /test/
		Filelist: []file.FileCopyAndMoveItem{
			{
				Path:    "/test.txt",
				Dest:    "/test",
				Newname: "test.txt",
			},
		},
	})
	test.PrintRes(res, err)
}

func TestMoveFile(t *testing.T) {
	test.TestSetEnv(t)
	res, err := file.Filemanager(&file.FilemanagerReq[file.FileCopyAndMoveItem]{
		Opera: file.Move,
		// mv /test.txt /test/
		Filelist: []file.FileCopyAndMoveItem{
			{
				Path:    "/test.txt",
				Dest:    "/test",
				Newname: "test.txt",
			},
		},
	})
	test.PrintRes(res, err)
}

func TestRenameFile(t *testing.T) {
	test.TestSetEnv(t)
	res, err := file.Filemanager(&file.FilemanagerReq[file.FileRenameItem]{
		Opera: file.Rename,
		// mv /test.txt /test2.txt
		Filelist: []file.FileRenameItem{
			{
				Path:    "/test.txt",
				Newname: "test2.txt",
			},
		},
	})
	test.PrintRes(res, err)
}

func TestDeleteFile(t *testing.T) {
	test.TestSetEnv(t)
	res, err := file.Filemanager(&file.FilemanagerReq[file.FileDeleteItem]{
		Opera: file.Delete,
		// rm /test.txt
		Filelist: []file.FileDeleteItem{
			"/test.txt",
		},
	})
	test.PrintRes(res, err)
}
