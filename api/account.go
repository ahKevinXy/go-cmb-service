package api

import (
	"fmt"
	"github.com/ahKevinXy/go-cmb/handler/account"
	"github.com/gin-gonic/gin"
	"go-cmb-service/constant"
	"go-cmb-service/help"
	"go-cmb-service/models/request"
	"go-cmb-service/models/response"
	"go-cmb-service/pkg/config"
)

// GetMainAccountInfo
//  @Description:   获取主账户信息 余额等信息
//  @param c
//  @Author  ahKevinXy
//  @Date  2023-07-15 11:45:04
func GetMainAccountInfo(c *gin.Context) {
	info, err := account.MainAccountInfo(config.CmbConfig.UserId, config.CmbConfig.Sm2Key, config.CmbConfig.PriKey, config.CmbConfig.MainAccount, "")
	if err != nil {
		return
	}
	var resp []*response.MainAccountInfo

	for _, v := range info.Response.Body.Ntqacinfz {
		var item response.MainAccountInfo
		item.Account = v.Accnbr
		item.AvailableBalance = v.Accblv //可用余额
		resp = append(resp, &item)
	}
	help.OK(c, resp, "ok")

}

// GetMainAccountInfoHistory
//  @Description:   交易流水模式
//  @param c
//  @Author  ahKevinXy
//  @Date  2023-07-15 11:48:13
func GetMainAccountInfoHistory(c *gin.Context) {

}

// GetMainAccountTrans
//  @Description:   获取交易流水
//  @param c
//  @Author  ahKevinXy
//  @Date  2023-07-15 11:52:50
func GetMainAccountTrans(c *gin.Context) {

}

// PayByP2P
//  @Description:  对公支付
//  @param c
//  @Author  ahKevinXy
//  @Date  2023-07-15 11:45:53
func PayByP2P(c *gin.Context) {
	var req request.P2pPay
	//todo
	if err := c.BindJSON(&req); err != nil {
		return
	}
	single, err := account.MainAccountPaySingle(config.CmbConfig.UserId,
		config.CmbConfig.Sm2Key,
		config.CmbConfig.PriKey,
		config.CmbConfig.MainAccount,
		"",
		constant.PayrollBusCode,
		req.ModNo,
		req.AccountName,
		req.Account,
		req.BankName,
		"",
		"",
		req.Amount,
		req.BankLinkNo,
		req.Remark,
		req.BusNo,
		"",
		"",
	)
	if err != nil {
		return
	}

	fmt.Println(single)

}

// P2PReceipt
//  @Description:  公对公单个回单查询
//  @param c
//  @Author  ahKevinXy
//  @Date  2023-07-15 11:46:27
func P2PReceipt(c *gin.Context) {

}
