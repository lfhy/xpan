# xpan - Baidu Netdisk Go SDK

A Go SDK for Baidu Netdisk Open Platform API, providing convenient wrapper functions for rapid integration of Baidu Netdisk functionalities.

## Features

- User authentication (OAuth2)
- Retrieve user info and disk quota
- File management (upload, download, move, copy, delete, rename, etc.)
- File listing
- Support for large file chunked upload

## Installation

```bash
go get github.com/lfhy/xpan
```

## Quick Start

### Initialize Client

```go
import "github.com/lfhy/xpan/client"

// Create client instance
c := client.New()

// Set authentication info
c.SetAuth(&auth.AuthEnv{
    ClientID:     "your-client-id",
    ClientSecret: "your-client-secret",
    AccessToken:  "your-access-token",
    RefreshToken: "your-refresh-token",
})
```

### Get User Info

```go
import "github.com/lfhy/xpan/user"

userInfo, err := user.GetUserInfo(&user.UserInfoReq{})
if err != nil {
    log.Fatal(err)
}
fmt.Printf("Username: %s\n", userInfo.NetDiskName)
```

### File Operations Example

```go
import "github.com/lfhy/xpan/file"

// Create directory
err := file.Mkdir(&file.MkdirReq{
    Path: "/my-folder",
})

// List files
files, err := file.List(&file.ListReq{
    Path: "/",
})
```

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.