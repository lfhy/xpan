# xpan - 百度网盘 Go SDK

[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

百度网盘开放平台 Go SDK，提供对百度网盘 API 的封装，方便开发者快速集成百度网盘功能。

## 功能特性

- 用户认证（OAuth2）
- 获取用户信息和网盘配额
- 文件管理（上传、下载、移动、复制、删除、重命名等）
- 文件列表查询
- 支持大文件分片上传

## 安装

```bash
go get github.com/lfhy/xpan
```

## 快速开始

### 初始化客户端

```go
import "github.com/lfhy/xpan/client"

// 创建客户端实例
c := client.New()

// 设置认证信息
c.SetAuth(&auth.AuthEnv{
    ClientID:     "your-client-id",
    ClientSecret: "your-client-secret",
    AccessToken:  "your-access-token",
    RefreshToken: "your-refresh-token",
})
```

### 获取用户信息

```go
import "github.com/lfhy/xpan/user"

userInfo, err := user.GetUserInfo(&user.UserInfoReq{})
if err != nil {
    log.Fatal(err)
}
fmt.Printf("用户名: %s\n", userInfo.NetDiskName)
```

### 文件操作示例

```go
import "github.com/lfhy/xpan/file"

// 创建目录
err := file.Mkdir(&file.MkdirReq{
    Path: "/my-folder",
})

// 列出文件
files, err := file.List(&file.ListReq{
    Path: "/",
})
```

## 许可证

本项目采用 MIT 许可证，详情请见 [LICENSE](LICENSE) 文件。