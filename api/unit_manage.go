package api

import (
	"fmt"
	"github.com/ahKevinXy/go-cmb/handler/unit_manager"
	"github.com/gin-gonic/gin"
	"go-cmb-service/models/request"
	"go-cmb-service/pkg/config"
)

// AddSubAccount
//  @Description:   添加记账单元
//  @param c
//  @Author  ahKevinXy
//  @Date  2023-07-11 16:39:28
func AddSubAccount(c *gin.Context) {
	var req request.AddSubAccount

	if err := c.BindJSON(&req); err != nil {

		return
	}

	uni, err := unit_manager.AddUnitAccountV1(config.CmbConfig.UserId, config.CmbConfig.Sm2Key, config.CmbConfig.PriKey, config.CmbConfig.MainAccount, req.SubAccountName, req.SubAccountNo)
	if err != nil {
		return
	}
	// 返回记账单元
	fmt.Println(uni.Response.Body.Ntdumaddz1[0].Dyanbr)
}

// CloseSubAccount
//  @Description:   关闭记账单元
//  @param c
//  @Author  ahKevinXy
//  @Date  2023-07-11 16:44:34
func CloseSubAccount(c *gin.Context) {
	var req request.CloseSubAccount
	v1, err := unit_manager.CloseUnitAccountV1(config.CmbConfig.UserId, config.CmbConfig.Sm2Key, config.CmbConfig.PriKey, config.CmbConfig.MainAccount, req.SubAccountNo)
	if err != nil {
		return
	}
	fmt.Println(v1)
}

// GetSubAccounts
//  @Description:   获取记账单元列表
//  @param c
//  @Author  ahKevinXy
//  @Date  2023-07-11 17:20:57
func GetSubAccounts(c *gin.Context) {
	var req request.SubAccountBalance

	if err := c.BindJSON(&req); err != nil {

		return
	}

	info, err := unit_manager.QueryUnitAccountInfo(config.CmbConfig.UserId, config.CmbConfig.Sm2Key, config.CmbConfig.PriKey, config.CmbConfig.MainAccount, req.SubAccountNo)
	if err != nil {
		return
	}

	fmt.Println(info)
}

// GetSubAccountHistoryBalance
//  @Description:   获取记账单元历史余额
//  @param c
//  @Author  ahKevinXy
//  @Date  2023-07-11 17:24:50
func GetSubAccountHistoryBalance(c *gin.Context) {

}
