package types

type ListOrder string

const (
	ListOrderName ListOrder = "name" // 按文件名称排序
	ListOrderTime ListOrder = "time" // 按文件修改时间排序
	ListOrderSize ListOrder = "size" // 按文件大小排序。
)

// 1 视频、2 音频、3 图片、4 文档、5 应用、6 其他、7 种子
type FileCategory int

const (
	FileCategoryVideo    FileCategory = 1
	FileCategoryAudio    FileCategory = 2
	FileCategoryImage    FileCategory = 3
	FileCategoryDocument FileCategory = 4
	FileCategoryApp      FileCategory = 5
	FileCategoryOther    FileCategory = 6
	FileCategoryTorrent  FileCategory = 7
)

type FileCreateRType int

const (
	FileCreateRTypePathRename          FileCreateRType = 1
	FileCreateRTypePathBlockListRename FileCreateRType = 2
	FileCreateRTypePathMove            FileCreateRType = 3
)

type CreateMode int

const (
	CreateModeUserUpload      CreateMode = 1 // 手动上传
	CreateModeUserMultiUpload CreateMode = 2 // 批量上传
	CreateModeAutoBackupFile  CreateMode = 3 // 文件自动备份
	CreateModeAutoBackupPhoto CreateMode = 4 // 相册自动备份
	CreateModeAutoBackupVideo CreateMode = 5 // 视频自动备份
)

type OndupMode string

const (
	OndupModeFail      OndupMode = "fail"      //直接返回失败
	OndupModeNewcopy   OndupMode = "newcopy"   // 重命名文件
	OndupModeOverwrite OndupMode = "overwrite" // 覆盖同名文件
	OndupModeSkip      OndupMode = "skip"      // 跳过
)
