package core

import (
	"net/url"
)

type Request struct {
	Method string
	FullURL string `eg:"https://github.com/og?type=public&language=go"`
	Path string `eg:"/og"`
	URL string `eg:"/og?type=public&language=go"`
	Query url.Values
}
type Body struct {
	Type string
	String string
	HTML string
	JSON interface{}
}
type Response struct {
	Status int
	Body Body
}
type App struct {

}
type Content struct {
	Req Request
	Res Response
}
func (c Content) Cookie() Cookie {

}
// 这是一个具有争议的api end 类似php的 die
// 不考虑开发效率的情况下，应该每个函数明确的返回错误，而不是调用方无法预知函数运行中断
func (c Content) EndJSON(v interface{}) {
	panic(v)
}
type Cookie struct {}
func (cookie Cookie) Set(key string, value string) {}

func (c Content) Session() Session {

}
type Session struct {}
func (session Session) Set(key string, value string) {}

func (c Content) Next() ResponseError {
	return ResponseError{}
}
func (c Content) BindQuery(v interface{}) {}
func (c Content) BindForm(v interface{}) {}

func (c Content) SetString(text string) {
	c.SetBody(Response{
		Status: 200,
		Body: Body{
			Type: "string",
			String: text
		},
	})
}
func (c Content) SetJSON(v interface{}) {
	c.SetBody(Response{
		Status: 200,
		Body: Body{
			Type: "json",
			JSON: v,
		},
	})
}
func (c Content) SetBody (response Response) {

}
type Action func(c Content) ResponseError
func (app App) Use(Action) {

}
type RenderConfig struct {
	Dir string
}
func (app App) Render(config RenderConfig) {

}
type RouterConfig struct {
	Name string
}
func (app App) Get(router string,config RouterConfig, action Action) {

}
func (app App) Post(router string,config RouterConfig, action Action) {

}
type ResponseError struct {

}