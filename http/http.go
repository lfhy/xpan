package http

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"mime/multipart"
	"net/http"
	"sync"

	"github.com/lfhy/xpan/log"

	"github.com/lfhy/xpan/types"
	"golang.org/x/time/rate"
)

type HTTPMethod string

const (
	GET    HTTPMethod = "GET"
	POST   HTTPMethod = "POST"
	PUT    HTTPMethod = "PUT"
	DELETE HTTPMethod = "DELETE"
)

type Client interface {
	Do(req *http.Request) (*http.Response, error)
}

var client Client

// 用于存储每个接口的限流器
var limiters = sync.Map{}

// 用于存储正在进行的GET请求
var pendingRequests = sync.Map{}

// PendingRequest 用于保存正在进行的请求和等待结果的调用方
type PendingRequest struct {
	mutex    sync.Mutex
	wg       sync.WaitGroup
	data     []byte
	err      error
	done     bool
	Response any // 用于缓存成功响应的结果
}

// 获取指定接口的限流器
func getLimiter(route, method string) *rate.Limiter {
	key := route + "|" + method
	limiter, ok := limiters.Load(key)
	if !ok {
		// 每分钟10个请求
		limiter, _ = limiters.LoadOrStore(key, rate.NewLimiter(rate.Limit(float64(10)/60), 10))
		return limiter.(*rate.Limiter)
	}
	return limiter.(*rate.Limiter)
}

func GetClient() Client {
	if client == nil {
		client = http.DefaultClient
	}
	return client
}

func SetClient(c Client) {
	client = c
}

type Request[Req any, Res any] struct {
	BaseURL     string
	HTTPMethod  HTTPMethod
	Route       string
	Method      string
	AccessToken string
	Request     Req
	Response    Res
}

func (api *Request[Req, Res]) Do() (Res, error) {
	// 对于GET请求，尝试合并相同请求
	if api.HTTPMethod == GET {
		return api.doWithMerging()
	}
	
	// 应用限流控制
	limiter := getLimiter(api.Route, api.Method)
	// 等待直到有可用的令牌
	ctx := context.Background()
	if err := limiter.Wait(ctx); err != nil {
		var res Res
		return res, err
	}

	reqURl := api.BaseURL + api.Route
	var isFirstQuery = true
	if api.Method != "" {
		if isFirstQuery {
			isFirstQuery = false
			reqURl += "?"
		} else {
			reqURl += "&"
		}
		reqURl += "method=" + api.Method
	}

	if api.AccessToken != "" {
		if isFirstQuery {
			isFirstQuery = false
			reqURl += "?"
		} else {
			reqURl += "&"
		}
		reqURl += "access_token=" + api.AccessToken
	}
	query, body, file := types.GetReqParams(api.Request)
	if query != "" {
		if isFirstQuery {
			isFirstQuery = false
			reqURl += "?"
		} else {
			reqURl += "&"
		}
		reqURl += query
	}

	contentType := "application/x-www-form-urlencoded"
	if file != nil {
		bodyBuf := &bytes.Buffer{}
		bodyWriter := multipart.NewWriter(bodyBuf)
		fileWriter, err := bodyWriter.CreateFormFile("file", "file")
		if err != nil {
			var res Res
			return res, err
		}
		io.Copy(fileWriter, file)
		bodyWriter.Close()
		contentType = bodyWriter.FormDataContentType()
		body = bodyBuf
	}
	log.Println("API Request:", reqURl)
	req, err := http.NewRequest(string(api.HTTPMethod), reqURl, body)
	if err != nil {
		return api.Response, err
	}
	req.Header.Add("Content-Type", contentType)
	req.Header.Add("User-Agent", "pan.baidu.com")
	resp, err := GetClient().Do(req)
	if err != nil {
		return api.Response, err
	}
	defer resp.Body.Close()
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return api.Response, err
	}
	log.Println("API Response:", string(data))
	var errCode types.Error
	json.Unmarshal(data, &errCode)
	if errCode.IsError() {
		return api.Response, errCode
	}
	var res Res
	err = json.Unmarshal(data, &res)
	if err != nil {
		return api.Response, err
	}
	api.Response = res
	return res, nil
}

// doWithMerging 处理GET请求的合并逻辑
func (api *Request[Req, Res]) doWithMerging() (Res, error) {
	// 构建完整URL用于标识相同请求
	reqURl := api.BaseURL + api.Route
	var isFirstQuery = true
	if api.Method != "" {
		if isFirstQuery {
			isFirstQuery = false
			reqURl += "?"
		} else {
			reqURl += "&"
		}
		reqURl += "method=" + api.Method
	}

	if api.AccessToken != "" {
		if isFirstQuery {
			isFirstQuery = false
			reqURl += "?"
		} else {
			reqURl += "&"
		}
		reqURl += "access_token=" + api.AccessToken
	}
	query, _, _ := types.GetReqParams(api.Request)
	if query != "" {
		if isFirstQuery {
			isFirstQuery = false
			reqURl += "?"
		} else {
			reqURl += "&"
		}
		reqURl += query
	}

	// 检查是否已有相同请求正在进行
	if pending, loaded := pendingRequests.LoadOrStore(reqURl, &PendingRequest{}); loaded {
		// 存在相同请求，等待其完成
		pr := pending.(*PendingRequest)
		pr.wg.Add(1)
		pr.wg.Wait()

		// 返回已缓存的结果
		var res Res
		if pr.err != nil {
			return res, pr.err
		}
		
		var errCode types.Error
		json.Unmarshal(pr.data, &errCode)
		if errCode.IsError() {
			return res, errCode
		}
		
		err := json.Unmarshal(pr.data, &res)
		if err != nil {
			return res, err
		}
		return res, nil
	}

	// 当前请求是第一个，执行实际请求
	pr := &PendingRequest{}
	pr.wg.Add(1)
	pendingRequests.Store(reqURl, pr)
	
	// 执行完成后清理
	defer func() {
		pr.mutex.Lock()
		pr.done = true
		pr.wg.Done()
		pendingRequests.Delete(reqURl)
		pr.mutex.Unlock()
	}()

	// 应用限流控制
	limiter := getLimiter(api.Route, string(GET))
	ctx := context.Background()
	if err := limiter.Wait(ctx); err != nil {
		pr.err = err
		var res Res
		return res, err
	}

	// 发起实际请求
	contentType := "application/x-www-form-urlencoded"
	log.Println("API Request:", reqURl)
	req, err := http.NewRequest(string(api.HTTPMethod), reqURl, nil)
	if err != nil {
		pr.err = err
		var res Res
		return res, err
	}
	req.Header.Add("Content-Type", contentType)
	req.Header.Add("User-Agent", "pan.baidu.com")
	resp, err := GetClient().Do(req)
	if err != nil {
		pr.err = err
		var res Res
		return res, err
	}
	defer resp.Body.Close()
	
	pr.data, err = io.ReadAll(resp.Body)
	if err != nil {
		pr.err = err
		var res Res
		return res, err
	}
	
	log.Println("API Response:", string(pr.data))
	
	var errCode types.Error
	json.Unmarshal(pr.data, &errCode)
	if errCode.IsError() {
		pr.err = errCode
		var res Res
		return res, errCode
	}
	
	var res Res
	err = json.Unmarshal(pr.data, &res)
	if err != nil {
		pr.err = err
		return res, err
	}
	
	pr.Response = res
	return res, nil
}