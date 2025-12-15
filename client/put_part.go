package client

import (
	"io"

	"github.com/lfhy/baidu-pan-client/file"
	"github.com/lfhy/baidu-pan-client/types"
)

func (c *Client) InitMultiPartUpload(path string, size types.SizeB, blockList ...string) (*file.PreCreateRes, error) {
	return file.PreCreate(&file.PreCreateReq{
		Path:      path,
		Size:      size,
		BlockList: blockList,
	})
}

func (c *Client) UploadPart(path string, uploadId string, partSeq int, data io.Reader) (*file.UploadChunkRes, error) {
	res, err := c.GetUploadHost(path, uploadId)
	if err != nil {
		return nil, err
	}
	return file.UploadChunk(&file.UploadChunkReq{
		Path:     path,
		UploadId: uploadId,
		PartSeq:  partSeq,
		File:     data,
	}, res.Host)
}

func (c *Client) ComplateMultiPartUpload(path string, size types.SizeB, uploadId string, blockMd5List []string) (*file.CreateRes, error) {
	return file.Create(&file.CreateReq{
		Path:      path,
		Size:      size,
		UploadId:  uploadId,
		BlockList: blockMd5List,
	})
}
