package file

import (
	"fmt"
	"io"
	"net/http"

	"github.com/lfhy/baidu-pan-client/types"

	phttp "github.com/lfhy/baidu-pan-client/http"
)

type DownloadRange struct {
	Start types.SizeB
	End   types.SizeB
}

// 获取下载链接
func DownloadUrl(dlink string) string {
	return types.PCSBaseURL + dlink + "&access_token=" + types.AccessToken
}

// 执行下载
func Download(dlink string, rangeBytes ...DownloadRange) (io.ReadCloser, error) {
	req, err := http.NewRequest("GET", DownloadUrl(dlink), nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("User-Agent", "pan.baidu.com")
	if len(rangeBytes) > 0 {
		req.Header.Set("Range", "bytes="+fmt.Sprintf("%d-%d", rangeBytes[0].Start, rangeBytes[0].End))
	}
	resp, err := phttp.GetClient().Do(req)
	if err != nil {
		return nil, err
	}
	return resp.Body, nil
}
