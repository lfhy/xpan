package file

import "github.com/lfhy/xpan/types"

func Mkdir(path string) (*CreateRes, error) {
	return Create(&CreateReq{
		Path:  path,
		IsDir: types.BoolIntTrue,
	})
}
