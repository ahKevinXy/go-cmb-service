package response

type MainAccountInfo struct {
	AvailableBalance string `json:"available_balance"` // 可用余额
	Account          string `json:"account"`           //账号
}
