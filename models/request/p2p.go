package request

type P2pPay struct {
	BankName     string `json:"bank_name"`      // 银行名称
	BankLinkNo   string `json:"bank_link_no"`   // 银联号
	OpBankName   string `json:"op_bank_name"`   //开户行名称
	AccountName  string `json:"account_name"`   //账户名称
	Account      string `json:"account"`        //卡号
	Amount       string `json:"amount"`         //金额
	Remark       string `json:"remark"`         //备注
	ModNo        string `json:"mod_no"`         // 交易编码
	BusNo        string `json:"bus_no"`         //业务单号
	SubAccountNo string `json:"sub_account_no"` // 记账单元编号
}
