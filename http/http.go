package http

import (
	"bytes"
	"encoding/json"
	"io"
	"mime/multipart"
	"net/http"

	"github.com/lfhy/xpan/log"

	"github.com/lfhy/xpan/types"
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
