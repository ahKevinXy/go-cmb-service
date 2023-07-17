package response

// PayrollMods
//  @Description:   获取交易模式
//  @Author  ahKevinXy
//  @Date  2023-07-12 22:45:51
type PayrollMods struct {
	ModNo   string `json:"mod_no"`   // 交易编号
	ModName string `json:"mod_name"` //模式名称
	Status  string `json:"status"`   //状态 A 有效 B 关闭
	EftAt   string `json:"eft_at"`   //生效时间
	ExpAt   string `json:"exp_at"`   //过期时间
}

type SupPayroll struct {
	Status string `json:"status"`
	FlowId string `json:"flow_id"`
}

type QuerySupPayrollDetail struct {
	TrsNo      string `json:"trs_no"`
	BankCardNo string `json:"account"`
	Name       string `json:"name"`
	Amount     string `json:"amount"` //金额
	Remark     string `json:"remark"` //备注
	ErrorMsg   string `json:"error_msg"`
	Status     string `json:"status"` //状态
}
