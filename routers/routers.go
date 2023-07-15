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
	group.POST("payroll/mods", api.GetPayrollMods)
	// 超网代发
	group.POST("payroll/pay/sup", api.SupPayroll)

	//超网代发

	// 添加记账单元

	// 代发

	// 申请回单

	// 获取回单

	// 交易管家

	//添加 记账单元
	group.POST("unit_manage/sub_account/add", api.AddSubAccount) // ok
	// 关闭 记账单元
	group.POST("unit_manage/sub_account/close", api.CloseSubAccount) //ok
	// 获取 记账单元账户编号余额及其状态
	group.POST("unit_manage/sub_accounts", api.GetSubAccounts)
	// 获取余额
	group.POST("unit_manage/sub_accounts", api.GetSubAccounts)

	// 账户管理

	//获取余额
	group.POST("/main_account/info", api.GetMainAccountInfo)

	// 获取对公支付的模式
	group.POST("p2p/mods", api.P2PMods)
	// 交易流水
	group.POST("/account/trans")
	// 对公支付
	group.POST("/p2p/pay")

	// 支付结果查询
	group.POST("/p2p/pay/query")

	// 单个回单查询
	group.POST("/p2p/pay/receipt")

	return router
}
