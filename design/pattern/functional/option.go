package functional

import (
	"net/http"
	"strings"
	"time"
)

/**
函数式编程模式
option模式，类似于面向对象里的Builder模式
 */

//************************配置对象
type requestOptions struct {
	time time.Duration
	data string
	headers map[string]string
}
func defaultRequestOptions() *requestOptions {
	return &requestOptions{ //默认请求
		time:    5 * time.Second,
		data:    "",
		headers: make(map[string]string),
	}
}

//************************函数赋值的包装
type Option interface {
	apply(opt *requestOptions)
}
type funcOption struct {
	f func(*requestOptions)
}
func (p *funcOption) apply(opt *requestOptions) {
	p.f(opt)
}
func NewFuncOption(f func(*requestOptions)) *funcOption {
	return &funcOption{
		f:f,
	}
}


//************************定义一堆With函数，可赋值修改配置
func WithTimeout(time time.Duration) Option{
	return NewFuncOption(func(options *requestOptions) {
		options.time = time
	})
}
func WithData(data string) Option{
	return NewFuncOption(func(options *requestOptions) {
		options.data = data
	})
}
func WithHeaders(headers map[string]string) Option{
	return NewFuncOption(func(options *requestOptions) {
		for key,val := range headers{
			options.headers[key] = val
		}
	})
}

//使用举例
func httpRequest(method string,url string,options ...Option) error{
	reqOpts :=defaultRequestOptions() //默认配置
	for _,opt := range options{
		opt.apply(reqOpts)
	}
	// 创建请求对象
	req, err := http.NewRequest(method, url, strings.NewReader(reqOpts.data))
	if err != nil{
		return err
	}
	// 设置请求头
	for key, value := range reqOpts.headers {
		req.Header.Add(key, value)
	}
	return nil
}






