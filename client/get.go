package client

import (
	"errors"
	"io"
	"path/filepath"

	"github.com/lfhy/xpan/file"
	"github.com/lfhy/xpan/types"
)

func (c *Client) GetPathFsid(path string) (uint64, error) {
	dir := filepath.Dir(path)
	for file := range c.ListObjectsCursor(dir) {
		if file.Path == path {
			return file.FsId, nil
		}
	}
	return 0, errors.New("file not found")
}

func (c *Client) GetObject(path string) (io.ReadCloser, error) {
	res, err := c.StatObject(path)
	if err != nil {
		return nil, err
	}
	if res.Size > types.SizeMB(50).ToB() {
		// 分块下载
		return &FileReader{
			client: c,
			size:   res.Size,
			meta:   res,
		}, nil
	}
	return file.Download(res.Dlink)
}

type FileReader struct {
	client  *Client
	size    types.SizeB
	nowRead types.SizeB
	meta    *file.FilemetasItem
	buffer  []byte
	bufPos  int
	bufLen  int
}

func (d *FileReader) Read(p []byte) (n int, err error) {
	// 如果缓冲区为空或已读完，则加载更多数据
	if d.bufPos >= d.bufLen {
		// 计算下次需要读取的数据范围
		remaining := d.size - d.nowRead
		if remaining <= 0 {
			return 0, io.EOF
		}

		// 每次最多下载50MB
		chunkSize := types.SizeMB(50).ToB()
		if remaining < chunkSize {
			chunkSize = remaining
		}

		// 下载数据块
		reader, err := file.Download(d.meta.Dlink, file.DownloadRange{
			Start: d.nowRead,
			End:   d.nowRead + chunkSize - 1,
		})
		if err != nil {
			return 0, err
		}
		defer reader.Close()

		// 读取数据到缓冲区
		d.buffer = make([]byte, chunkSize)
		d.bufLen, err = io.ReadFull(reader, d.buffer)
		if err != nil && err != io.EOF && err != io.ErrUnexpectedEOF {
			return 0, err
		}

		d.bufPos = 0
		d.nowRead += types.SizeB(chunkSize)
	}

	// 从缓冲区复制数据到目标切片
	bytesToCopy := d.bufLen - d.bufPos
	if bytesToCopy > len(p) {
		bytesToCopy = len(p)
	}

	copy(p[:bytesToCopy], d.buffer[d.bufPos:d.bufPos+bytesToCopy])
	d.bufPos += bytesToCopy

	return bytesToCopy, nil
}

func (d *FileReader) Close() error {
	return nil
}
