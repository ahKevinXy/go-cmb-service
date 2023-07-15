package help

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct {
	Code int `json:"code" example:"200"` // 验证码

	Data interface{} `json:"data"` // 数据集

	Msg string `json:"msg"` // 消息
}

type Page struct {
	List      interface{} `json:"list"`
	Count     int         `json:"count"`
	PageIndex int         `json:"pageIndex"`
	PageSize  int         `json:"pageSize"`
}
type ListCountItem struct {
	List  interface{} `json:"list"`
	Count int64       `json:"count"`
}
type ListSumItem struct {
	List interface{} `json:"list"`
	Sum  float64     `json:"sum"`
}
type PageResponse struct {
	// 代码
	Code int `json:"code" example:"200"`
	// 数据集
	Data Page `json:"data"`
	// 消息
	Msg      string `json:"msg"`
	PageSize string `json:"page_size"`
}
type ListCountResponse struct {
	Code int           `json:"code"`
	Data ListCountItem `json:"data"`

	Msg string `json:"msg"`
}
type ListSumsResponse struct {
	Code int         `json:"code"`
	Data ListSumItem `json:"data"`

	Msg string `json:"msg"`
}

func (res *Response) ReturnOK() *Response {
	res.Code = 200
	return res
}

func (res *Response) ReturnError(code int) *Response {
	res.Code = code
	return res
}

func (res *PageResponse) ReturnOK() *PageResponse {
	res.Code = 200
	return res
}

// Error
//  @Description:  错误信息
//  @param c
//  @param code
//  @param err
//  @param msg
//  @Author  ahKevinXy
//  @Date  2023-07-11 23:30:37
func Error(c *gin.Context, code int, err error, msg string) {
	var res Response
	res.Msg = err.Error()
	if msg != "" {
		res.Msg = msg
	}
	c.JSON(http.StatusOK, res.ReturnError(code))
}

// ErrorCallBack
//  @Description:  错误直接返回
//  @param c
//  @param code
//  @param msg
//  @Author  ahKevinXy
//  @Date  2023-07-11 23:30:44
func ErrorCallBack(c *gin.Context, code int, msg string) {
	var res Response
	if msg != "" {
		res.Msg = msg
	}
	c.JSON(http.StatusOK, res.ReturnError(code))
}

func SystemError(c *gin.Context) {
	var res Response

	res.Msg = "系统错误"

	c.JSON(http.StatusOK, res.ReturnError(500))
}

// OK
//  @Description:  成功返回
//  @param c
//  @param data
//  @param msg
//  @Author  ahKevinXy
//  @Date  2023-07-11 23:30:57
func OK(c *gin.Context, data interface{}, msg string) {
	var res Response
	res.Data = data
	if msg != "" {
		res.Msg = msg
	} else {
		res.Msg = "ok"
	}
	c.JSON(http.StatusOK, res.ReturnOK())
}

func NotLogin(c *gin.Context) {
	var res Response
	res.Msg = "用户没有登录"
	c.JSON(http.StatusOK, res.ReturnError(401))
}
func Send(c *gin.Context, data interface{}, msg string) {
	var res Response
	res.Data = data
	if msg != "" {
		res.Msg = msg
	}
	c.JSON(http.StatusOK, data)
}

// PageOK 分页数据处理
func PageOK(c *gin.Context, result interface{}, count int, pageIndex int, pageSize int, msg string) {
	var res PageResponse
	res.Data.List = result
	res.Data.Count = count
	res.Data.PageIndex = pageIndex
	res.Data.PageSize = pageSize
	if msg != "" {
		res.Msg = msg
	}
	c.JSON(http.StatusOK, res.ReturnOK())
}

func ListCount(c *gin.Context, result interface{}, counts int64) {
	var res ListCountResponse
	res.Data.List = result
	res.Data.Count = counts
	res.Code = 200
	res.Msg = "ok"
	c.JSON(http.StatusOK, res)
}
func ListSums(c *gin.Context, result interface{}, sum float64) {
	var res ListSumsResponse
	res.Data.List = result
	res.Data.Sum = sum
	res.Code = 200
	res.Msg = "ok"
	c.JSON(http.StatusOK, res)
}

type PagePro struct {
	Current  int         `json:"current"`
	Data     interface{} `json:"data"`
	PageSize string      `json:"pageSize"`
	Success  bool        `json:"success"`
	Total    int         `json:"total"`
}

// TablePro
//  @Description:  ant-design table
//  @param c
//  @param result
//  @param counts
//  @param pageSize
//  @param current
//  @Author  ahKevinXy
//  @Date  2023-07-11 23:29:57
func TablePro(c *gin.Context, result interface{}, counts int, pageSize string, current int) {
	var res PagePro
	res.Current = current
	res.Data = result
	res.Total = counts
	res.PageSize = pageSize
	res.Success = true
	c.JSON(http.StatusOK, res)
}

// PageOKOpt
//  @Description:   分页
//  @param c
//  @param result
//  @param count
//  @param pageIndex
//  @param pageSize
//  @param msg
//  @param opt
//  @Author  ahKevinXy
//  @Date  2023-07-11 23:29:49
func PageOKOpt(c *gin.Context, result interface{}, count int, pageIndex int, pageSize int, msg string, opt ...interface{}) {
	var res PageResponse
	res.Data.List = result
	res.Data.Count = count
	res.Data.PageIndex = pageIndex
	res.Data.PageSize = pageSize
	if msg != "" {
		res.Msg = msg
	}
	c.JSON(http.StatusOK, res.ReturnOK())
}

// Custum
//  @Description:  通用
//  @param c
//  @param data
//  @Author  ahKevinXy
//  @Date  2023-07-11 23:29:16
func Custum(c *gin.Context, data gin.H) {
	c.JSON(http.StatusOK, data)
}

// DirectBack
//  @Description:  直接返回
//  @param c
//  @param data
//  @Author  ahKevinXy
//  @Date  2023-07-11 23:29:23
func DirectBack(c *gin.Context, data interface{}) {

	c.JSON(http.StatusOK, data)
}

// PageOKSpe
//  @Description:   分页
//  @param c
//  @param result
//  @param count
//  @param pageIndex
//  @param pageSize
//  @param msg
//  @param item
//  @Author  ahKevinXy
//  @Date  2023-07-11 23:29:00
func PageOKSpe(c *gin.Context, result interface{}, count int, pageIndex int, pageSize int, msg string, item interface{}) {
	var res PageResponse
	res.Data.List = result
	res.Data.Count = count
	res.Data.PageIndex = pageIndex
	res.Data.PageSize = pageSize
	//res.Data.Item = item
	if msg != "" {
		res.Msg = msg
	}
	c.JSON(http.StatusOK, res.ReturnOK())
}
