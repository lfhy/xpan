package client

import (
	"io"

	"github.com/lfhy/baidu-pan-client/file"
	"github.com/lfhy/baidu-pan-client/types"
	"github.com/lfhy/baidu-pan-client/utils"
)

func (c *Client) Mkdir(path string) (*file.CreateRes, error) {
	return file.Mkdir(path)
}

func (c *Client) PutObject(path string, data io.Reader, ondup ...types.OndupMode) (*file.UploadRes, error) {
	res, err := c.GetUploadHost(path)
	if err != nil {
		return nil, err
	}
	return file.Upload(&file.UploadReq{
		File:  data,
		Ondup: utils.GetOneOrDefault(ondup...),
		Path:  path,
	}, res.Host)
}

func (c *Client) GetUploadHost(path string, uploadId ...string) (*file.LocateUploadRes, error) {
	return file.LocateUpload(&file.LocateUploadReq{
		Path:     path,
		UploadId: utils.GetOneOrDefault(uploadId...),
	})
}
