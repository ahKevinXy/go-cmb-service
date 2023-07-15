package routers

import (
	"github.com/gin-gonic/gin"
	"go-cmb-service/api"
)

func Routers() *gin.Engine {

	// 获取银联cmd 模式
	router := gin.New()

	group := router.Group("v1/cmb")

	// 获取招商银行代发模式
	group.GET("payroll/mods", api.GetMods)
	// 超网代发
	group.GET("payroll/sup/pay", api.SupPayroll)

	//超网代发

	// 添加记账单元

	// 代发

	// 申请回单

	// 获取回单

	// 交易管家

	//添加 记账单元
	group.POST("unit_manage/sub_account/add", api.AddSubAccount)
	// 关闭 记账单元
	group.POST("unit_manage/sub_account/close", api.CloseSubAccount)
	// 获取 记账单元账户编号余额及其状态
	group.GET("unit_manage/sub_accounts", api.GetSubAccounts)

	return router
}
