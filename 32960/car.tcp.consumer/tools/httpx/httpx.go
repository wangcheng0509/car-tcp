package httpx

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"

	"github.com/pkg/errors"
)

// RequestOption 自定义处理请求
type RequestOption func(*http.Request) (*http.Request, error)

// NewDefaultHTTPClient 创建默认的HTTP客户端
func NewDefaultHTTPClient() *http.Transport {
	tr := &http.Transport{
		Proxy: http.ProxyFromEnvironment,
		DialContext: (&net.Dialer{
			Timeout:   30 * time.Second,
			KeepAlive: 30 * time.Second,
			DualStack: true,
		}).DialContext,
		IdleConnTimeout:       90 * time.Second,
		TLSHandshakeTimeout:   10 * time.Second,
		ExpectContinueTimeout: 1 * time.Second,
		DisableKeepAlives:     true,
		MaxIdleConnsPerHost:   -1,
	}

	return tr
}

// RequestURL  get request url
// 拼接host和router，确保host和router中间只有一个“/”
func RequestURL(base, router string) string {
	var buf bytes.Buffer
	if l := len(base); l > 0 {
		if base[l-1] == '/' {
			base = base[:l-1]
		}
		buf.WriteString(base)

		if rl := len(router); rl > 0 {
			if router[0] != '/' {
				buf.WriteByte('/')
			}
		}
	}
	buf.WriteString(router)
	return buf.String()
}

// Request do request
func Request(ctx context.Context, urlStr, method string, body io.Reader, options ...RequestOption) (Responser, error) {
	if ctx == nil {
		ctx = context.Background()
	}
	req, err := http.NewRequest(method, urlStr, body)
	if err != nil {
		return nil, err
	}

	if len(options) > 0 {
		req, err = options[0](req)
		if err != nil {
			return nil, err
		}
	}

	var resper Responser
	err = request(ctx, req, func(resp *http.Response, err error) error {
		if err != nil {
			return err
		}

		resper = newResponse(resp)
		return nil
	})
	if err != nil {
		return nil, err
	}

	if resper.Response().StatusCode != http.StatusOK {
		respBody, _ := resper.String()
		b, _ := io.ReadAll(body)
		log.Printf("req: %s\n resp: status:%v header:%v body:%s\n", b,
			resper.Response().Status, resper.Response().Header, respBody)

		return nil, fmt.Errorf("resp: status:%v\n header:%v body:%s",
			resper.Response().Status, resper.Response().Header, respBody)
	}

	return resper, nil
}

// RequestJSONWithToken 使用令牌发送json请求
func RequestJSONWithToken(ctx context.Context, urlStr, method string, tokenFunc func() (string, error), body, result interface{}) error {
	w := new(bytes.Buffer)
	if err := json.NewEncoder(w).Encode(body); err != nil {
		return err
	}

	resp, err := Request(ctx, urlStr, method, w, func(req *http.Request) (*http.Request, error) {
		token, err := tokenFunc()
		if err != nil {
			return req, err
		}

		req.Header.Set(TokenKey, bearerToken(token))
		req.Header.Set("Content-Type", "application/json")
		return req, nil
	})
	if err != nil {
		return err
	}

	return parseResponseJSON(resp, result)
}

func bearerToken(token string) string {
	if len(token) < len(Bearer) {
		return Bearer + token
	}

	if token[:len(Bearer)] == Bearer {
		return token
	}
	return Bearer + token
}

// HTTP请求处理
func request(ctx context.Context, req *http.Request, f func(*http.Response, error) error) error {
	tr := NewDefaultHTTPClient()

	cli := &http.Client{Transport: tr}
	c := make(chan error, 1)
	go func() { c <- f(cli.Do(req)) }()
	select {
	case <-ctx.Done():
		tr.CancelRequest(req)
		<-c
		return ctx.Err()
	case err := <-c:
		return err
	}
}

// Responser HTTP response interface
type Responser interface {
	String() (string, error)
	Bytes() ([]byte, error)
	JSON(v interface{}) error
	Response() *http.Response
	Close()
}

func newResponse(resp *http.Response) *response {
	return &response{resp}
}

type response struct {
	resp *http.Response
}

// ErrorResult 错误结果
type ErrorResult struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (r *ErrorResult) Error() string {
	return r.Message
}

// 解析响应JSON
func parseResponseJSON(resp Responser, result interface{}) error {
	if resp.Response().StatusCode != 200 {
		buf, err := resp.Bytes()
		if err != nil {
			return err
		}

		errResult := &ErrorResult{}
		err = json.Unmarshal(buf, errResult)
		if err == nil &&
			(errResult.Code != 0 || errResult.Message != "") {
			return errResult
		}

		return fmt.Errorf("%s", buf)
	} else if result == nil {
		resp.Close()
		return nil
	}
	defer resp.Close()

	return resp.JSON(result)
}

// ParseResponseJSON ..
func ParseResponseJSON(resp Responser, result interface{}) error {
	return parseResponseJSON(resp, result)
}

func (r *response) Response() *http.Response {
	return r.resp
}

func (r *response) String() (string, error) {
	b, err := r.Bytes()
	if err != nil {
		return "", err
	}
	return string(b), nil
}

func (r *response) Bytes() ([]byte, error) {
	defer r.resp.Body.Close()

	buf, err := ioutil.ReadAll(r.resp.Body)
	if err != nil {
		return nil, err
	}
	return buf, nil
}

func (r *response) JSON(v interface{}) error {
	defer r.resp.Body.Close()

	return json.NewDecoder(r.resp.Body).Decode(v)
}

func (r *response) Close() {
	if !r.resp.Close {
		r.resp.Body.Close()
	}
}

// Get get request
func Get(ctx context.Context, urlStr string, param url.Values, options ...RequestOption) (Responser, error) {
	if param != nil {
		c := '?'
		if strings.IndexByte(urlStr, '?') != -1 {
			c = '&'
		}
		urlStr = fmt.Sprintf("%s%c%s", urlStr, c, param.Encode())
	}

	return Request(ctx, urlStr, http.MethodGet, nil, options...)
}

// TokenKey token key
const (
	TokenKey = "Authorization"
	Bearer   = "Bearer "
)

// GetWithToken 携带令牌发送GET请求
func GetWithToken(ctx context.Context, urlStr string, tokenFunc func() (string, error), param url.Values, result interface{}) error {
	resp, err := Get(ctx, urlStr, param, func(req *http.Request) (*http.Request, error) {
		token, err := tokenFunc()
		if err != nil {
			return req, err
		}

		req.Header.Set(TokenKey, bearerToken(token))
		return req, nil
	})
	if err != nil {
		return err
	}

	return parseResponseJSON(resp, result)
}

// PostJSON 发送post请求
func PostJSON(ctx context.Context, urlStr string, body interface{}, result interface{}) error {
	b, err := json.Marshal(body)
	if err != nil {
		return errors.WithStack(err)
	}

	res, err := Request(ctx, urlStr, http.MethodPost, bytes.NewBuffer(b))
	if err != nil {
		return errors.WithStack(err)
	}

	return parseResponseJSON(res, result)
}

// PostJSONWithToken 使用令牌发送post请求
func PostJSONWithToken(ctx context.Context, urlStr string, tokenFunc func() (string, error), body, result interface{}) error {
	return RequestJSONWithToken(ctx, urlStr, http.MethodPost, tokenFunc, body, result)
}

// PutJSON 发送put请求
func PutJSON(ctx context.Context, urlStr string, body io.Reader, result interface{}) error {
	res, err := Request(ctx, urlStr, http.MethodPost, body)
	if err != nil {
		return err
	}

	return parseResponseJSON(res, result)
}

// PutJSONWithToken 使用令牌发送put请求
func PutJSONWithToken(ctx context.Context, urlStr string, tokenFunc func() (string, error), body, result interface{}) error {
	return RequestJSONWithToken(ctx, urlStr, http.MethodPut, tokenFunc, body, result)
}

// RequestFormDataWithToken 使用令牌发送post formdata请求
func RequestFormDataWithToken(ctx context.Context, urlStr, method string,
	tokenFunc func() (string, error), values map[string]interface{}, result interface{}) (err error) {

	payload := new(bytes.Buffer)
	w := multipart.NewWriter(payload)
	for key, r := range values {
		var fw io.Writer
		if x, ok := r.(io.Closer); ok {
			defer x.Close()
		}
		// Add file
		if x, ok := r.(*os.File); ok {
			if fw, err = w.CreateFormFile(key, x.Name()); err != nil {
				return
			}
		} else {
			// Add other fields
			if fw, err = w.CreateFormField(key); err != nil {
				return
			}
		}

		if rd, ok := r.(io.Reader); ok {
			if _, err = io.Copy(fw, rd); err != nil {
				return err
			}
		}
		if s, ok := r.(string); ok {
			if _, err = w.CreateFormField(s); err != nil {
				return err
			}
		}

	}
	resp, err := Request(ctx, urlStr, method, payload, func(req *http.Request) (*http.Request, error) {
		token, err := tokenFunc()
		if err != nil {
			return req, err
		}

		req.Header.Add(TokenKey, bearerToken(token))
		req.Header.Set("Content-Type", w.FormDataContentType())
		return req, nil
	})
	if err != nil {
		return err
	}

	return parseResponseJSON(resp, result)
}

// PostFormDataWithToken 使用令牌发送post formdata请求
func PostFormDataWithToken(ctx context.Context, urlStr string, tokenFun func() (string, error), values map[string]interface{}, result interface{}) error {
	return RequestFormDataWithToken(ctx, urlStr, http.MethodPost, tokenFun, values, result)
}
