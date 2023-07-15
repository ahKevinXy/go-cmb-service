package request

type BaseRequest struct {
	UserId      string `json:"user_id"`      //用户ID
	PriKey      string `json:"pri_key"`      //用户私钥
	Sm2Key      string `json:"sm2_key"`      //对称秘钥
	MainAccount string `json:"main_account"` //主卡号
}
