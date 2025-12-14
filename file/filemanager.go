package file

import (
	"github.com/lfhy/baidu-pan-client/http"
	"github.com/lfhy/baidu-pan-client/types"
)

type FileOp = string

const (
	Move   FileOp = "move"
	Copy   FileOp = "copy"
	Rename FileOp = "rename"
	Delete FileOp = "delete"
)

type AsyncMode int

const (
	Sync  AsyncMode = 0 // 同步
	Auto  AsyncMode = 1 // 自适应
	Async AsyncMode = 2 // 异步
)

type FileCopyAndMoveItem struct {
	Path    string `json:"path"`    // 操作文件路径
	Dest    string `json:"dest"`    // 目标的目录
	Newname string `json:"newname"` // 操作文件的新名称
}

type FileRenameItem struct {
	Path    string `json:"path"`    // 操作文件路径
	Newname string `json:"newname"` // 操作文件的新名称
}

type FileDeleteItem string

type FilemanagerReq[FileOpItem FileCopyAndMoveItem | FileRenameItem | FileDeleteItem] struct {
	Opera    FileOp          `query:"opera"`   // move/copy/rename/delete
	Async    AsyncMode       `query:"async"`   // 异步模式
	Filelist []FileOpItem    `body:"filelist"` // 操作文件列表
	Ondup    types.OndupMode `body:"ondup"`    // 遇到重复文件的处理策略
}

type FilemanagerItem struct {
	Path  string `json:"path"`  // 文件路径
	Errno int    `json:"errno"` // 是否发生错误
}

type FilemanagerRes struct {
	TaskId uint64             `json:"taskid"` // 异步任务ID
	List   []*FilemanagerItem `json:"list"`   // 同步任务响应的操作结果
}

// 复制 删除 重命名 移动
func Filemanager[FileOpItem FileCopyAndMoveItem | FileRenameItem | FileDeleteItem](req *FilemanagerReq[FileOpItem]) (*FilemanagerRes, error) {
	api := &http.Request[*FilemanagerReq[FileOpItem], *FilemanagerRes]{
		AccessToken: types.AccessToken,
		BaseURL:     types.PanBaseURL,
		HTTPMethod:  http.POST,
		Method:      "filemanager",
		Request:     req,
		Route:       types.FileRoute,
	}
	return api.Do()
}
