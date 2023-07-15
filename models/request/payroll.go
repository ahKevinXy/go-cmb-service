package request

type PayrollMods struct {
	//BaseRequest
	BusCode string `json:"bus_code"`
}

type P2PMods struct {
	//BaseRequest
	BusCode string `json:"bus_code"`
}

// SupPayroll
//  @Description:  超网代发
//  @Author  ahKevinXy
//  @Date  2023-07-15 18:00:16
type SupPayroll struct {
	TotalInfo struct {
		TotalNum     string `json:"total_num"`      //总笔数
		SubAccountNo string `json:"sub_account_no"` //记账单元编号
		BusCode      string `json:"bus_code"`       // 业务编号
		Remark       string `json:"remark"`         // 备注
		BizNo        string `json:"biz_no"`         //业务单号
	} `json:"total_info"` // 汇总信息
	DetailInfo []*DetailInfo `json:"detail_info"` // 明细信息
}

type DetailInfo struct {
	BankCardNo string `json:"bank_card_no"` // 卡号
	Name       string `json:"name"`         //姓名
	Amount     string `json:"amount"`       //金额
	Remark     string `json:"remark"`       //备注
	TrsNo      string `json:"trs_no"`       //交易编号
}
