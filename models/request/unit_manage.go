package request

// AddSubAccount
//  @Description:   添加记账单元
//  @Author  ahKevinXy
//  @Date  2023-07-10 19:41:36
type AddSubAccount struct {
	//BaseRequest
	SubAccountName string `json:"sub_account_name"` //记账单元名称

	SubAccountNo string `json:"sub_account_no"` //记账单元编号 (可空默认生成)
}

type SubAccountBalance struct {
	//BaseRequest
	SubAccountNo string `json:"sub_account_no"` //子单元编号
}

type CloseSubAccount struct {
	SubAccountNo string `json:"sub_account_no"`
}
