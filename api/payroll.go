package api

import (
	"fmt"
	"github.com/ahKevinXy/go-cmb/handler/payroll_old"
	"github.com/ahKevinXy/go-cmb/models"
	"github.com/gin-gonic/gin"
	"go-cmb-service/help"
	"go-cmb-service/models/request"
	"go-cmb-service/models/response"
	"go-cmb-service/pkg/config"
)

// SupPayroll
//  @Description:  超网代发
//  @param c
//  @Author  ahKevinXy
//  @Date  2023-07-11 16:48:16
func SupPayroll(c *gin.Context) {

	var req request.SupPayroll

	if err := c.BindJSON(&req); err != nil {

		return
	}

	var total []*models.Ntagcsaix1

	total = append(total, &models.Ntagcsaix1{
		Yurref: req.TotalInfo.BizNo,
		Bbknbr: config.CmbConfig.BankLinkNo,
		Payacc: config.CmbConfig.MainAccount,
		Bustya: req.TotalInfo.BusMode,
		Trxrmk: req.TotalInfo.Remark,
		Begtag: "Y",
		Endtag: "Y",
		Ttlamt: req.TotalInfo.TotalAmount,
		Ttlcnt: req.TotalInfo.TotalNum,
		Curamt: req.TotalInfo.TotalAmount,
		Curcnt: req.TotalInfo.TotalNum,
		Jzbnbr: req.TotalInfo.SubAccountNo,
	})
	var detail []*models.Ntagcsaix2

	for _, v := range req.DetailInfo {
		item := models.Ntagcsaix2{}

		item.Eacnbr = v.BankCardNo
		item.Eacnam = v.Name
		item.Rcvamt = v.Amount
		item.Trxtxt = v.Remark
		item.Trxseq = v.TrsNo

		detail = append(detail, &item)
	}

	sup, err := payroll_old.CreditHandleOtherBySup(config.CmbConfig.UserId, config.CmbConfig.Sm2Key, config.CmbConfig.PriKey, req.TotalInfo.BusMode, total, detail)
	if err != nil {
		return
	}

	var resp response.SupPayroll

	if len(sup.Response.Body.Ntagcagcz1) == 0 {
		return
	}

	resp.Status = sup.Response.Body.Ntagcagcz1[0].Reqsta
	resp.FlowId = sup.Response.Body.Ntagcagcz1[0].Reqnbr
	help.OK(c, resp, "ok")

}

// QuerySupPayrollByBizNo
//  @Description:  通过业务单号进行查询
//  @param c
//  @Author  ahKevinXy
//  @Date  2023-07-15 11:49:00
func QuerySupPayrollByBizNo(c *gin.Context) {

}

// QuerySupPayrollDetail
//  @Description:  查询交易明细
//  @param c
//  @Author  ahKevinXy
//  @Date  2023-07-15 11:49:39
func QuerySupPayrollDetail(c *gin.Context) {

	var req request.QuerySupPayrollDetail

	if err := c.BindJSON(&req); err != nil {

		return
	}

	detail, err := payroll_old.QueryOldPayRollOrderDetail(config.CmbConfig.UserId,
		config.CmbConfig.Sm2Key,
		config.CmbConfig.PriKey,
		req.FlowId,
		"",
		"",
		"",
	)
	if err != nil {
		return
	}

	// todo

	var resp []*response.QuerySupPayrollDetail

	for _, v := range detail.Response.Body.Ntagcdtly1 {
		var item response.QuerySupPayrollDetail

		item.BankCardNo = v.Accnbr
		item.Name = v.Accnam
		item.TrsNo = v.Trxseq
		item.Amount = v.Trsamt
		item.Remark = v.Trsdsp
		item.Status = v.Stscod

		if item.Status != "S" {
			item.ErrorMsg = "支付失败"
		}

		resp = append(resp, &item)
	}

	help.OK(c, resp, "ok")
}

// QuerySupPayRefund
//  @Description:   获取跨行退票
//  @param c
//  @Author  ahKevinXy
//  @Date  2023-07-15 11:50:15
func QuerySupPayRefund(c *gin.Context) {

	var req request.QuerySupPayRefund

	if err := c.BindJSON(&req); err != nil {

		return
	}

	//payroll_old.Refund(config.CmbConfig.UserId, config.CmbConfig.Sm2Key, config.CmbConfig.PriKey, )
}

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
