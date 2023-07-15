package api

import (
	"fmt"
	"github.com/ahKevinXy/go-cmb/handler/payroll_old"
	"github.com/gin-gonic/gin"
	"go-cmb-service/models/request"
	"go-cmb-service/pkg/config"
)

// GetCallback
//  @Description:  获取回单
//  @param c
//  @Author  ahKevinXy
//  @Date  2023-07-11 16:49:30
func GetCallback(c *gin.Context) {

}

// GetCallbackDownload
//  @Description:  获取回单下载地址
//  @param c
//  @Author  ahKevinXy
//  @Date  2023-07-11 16:49:54
func GetCallbackDownload(c *gin.Context) {
	var req request.Callback

	if err := c.BindJSON(&req); err != nil {
		return
	}

	downloadUrl, err := payroll_old.QueryPayrollStatementDownloadUrl(
		config.CmbConfig.UserId, config.CmbConfig.Sm2Key, config.CmbConfig.PriKey, req.TaskId,
	)
	if err != nil {
		return
	}

	fmt.Println("下载地址 ", downloadUrl)
}
