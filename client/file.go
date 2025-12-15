package client

import (
	"errors"
	"path/filepath"

	"github.com/lfhy/xpan/file"
	"github.com/lfhy/xpan/types"
)

func (c *Client) ListObjects(dir string, opt ...*file.ListAllReq) (*file.ListRes, error) {
	var req *file.ListAllReq
	if len(opt) > 0 {
		req = opt[0]
	}
	if req == nil {
		req = &file.ListAllReq{}
	}
	req.Path = dir
	return file.ListAll(req)
}

func (c *Client) StatObject(path string) (*file.FilemetasItem, error) {
	fsid, err := c.GetPathFsid(path)
	if err != nil {
		return nil, err
	}
	return c.StatObjectUseFsId(fsid)
}

func (c *Client) StatObjectUseFsId(fsId uint64) (*file.FilemetasItem, error) {
	res, err := file.Filemetas(&file.FilemetasReq{
		FsIds: []uint64{fsId},
		Dlink: types.BoolIntTrue,
	})
	if err != nil {
		return nil, err
	}
	if len(res.List) == 0 {
		return nil, errors.New("file not found")
	}
	return res.List[0], nil
}

func (c *Client) StatObjects(fsIds ...uint64) (*file.FilemetasRes, error) {
	return file.Filemetas(&file.FilemetasReq{
		FsIds: fsIds,
		Dlink: types.BoolIntTrue,
	})
}

func (c *Client) StatObjectsPro(req *file.FilemetasReq) (*file.FilemetasRes, error) {
	return file.Filemetas(req)
}

func (c *Client) CopyObject(srcFilePath string, destDir string, newName ...string) (*file.FilemanagerItem, error) {
	if len(newName) == 0 {
		newName = []string{filepath.Base(srcFilePath)}
	}
	res, err := file.Filemanager(&file.FilemanagerReq[file.FileCopyAndMoveItem]{
		Filelist: []file.FileCopyAndMoveItem{{
			Path:    srcFilePath,
			Dest:    destDir,
			Newname: newName[0],
		}},
		Opera: file.Copy,
		Ondup: types.OndupModeFail,
	})
	if err != nil {
		return nil, err
	}
	if len(res.List) == 0 {
		return nil, errors.New("copy object failed")
	}
	return res.List[0], nil
}

func (c *Client) MoveObject(srcFilePath string, destDir string, newName ...string) (*file.FilemanagerItem, error) {
	if len(newName) == 0 {
		newName = []string{filepath.Base(srcFilePath)}
	}
	res, err := file.Filemanager(&file.FilemanagerReq[file.FileCopyAndMoveItem]{
		Filelist: []file.FileCopyAndMoveItem{{
			Path:    srcFilePath,
			Dest:    destDir,
			Newname: newName[0],
		}},
		Opera: file.Move,
		Ondup: types.OndupModeFail,
	})
	if err != nil {
		return nil, err
	}
	if len(res.List) == 0 {
		return nil, errors.New("move object failed")
	}
	return res.List[0], nil
}

func (c *Client) RenameObject(srcFilePath string, newName string) (*file.FilemanagerItem, error) {
	res, err := file.Filemanager(&file.FilemanagerReq[file.FileRenameItem]{
		Filelist: []file.FileRenameItem{{
			Path:    srcFilePath,
			Newname: newName,
		}},
		Opera: file.Rename,
		Ondup: types.OndupModeFail,
	})
	if err != nil {
		return nil, err
	}
	if len(res.List) == 0 {
		return nil, errors.New("rename object failed")
	}
	return res.List[0], nil
}
