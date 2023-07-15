package config

import "github.com/spf13/viper"

type Cmb struct {
	SassName    string `json:"sass_name"`    // Sass 名称
	SassKey     string `json:"sass_key"`     // Sass key
	Url         string `json:"url"`          // 招商银行地址
	DefaultText string `json:"default_text"` // 默认请求参数
	PriKey      string `json:"pri_key"`      //用户私钥
	Sm2Key      string `json:"sm2_key"`      //国密对称秘钥
	MainAccount string `json:"main_account"` //主账户
	UserId      string `json:"user_id"`      //用户ID
	PayrollMod  string `json:"payroll_mod"`  //代发模式
	BankLinkNo  string `json:"bank_link_no"` //银联号
}

// InitCmb
//  @Description:  初始化 cmb 配置
//  @param cfg
//  @return *Cmb
//  @Author  ahKevinXy
//  @Date  2023-05-16 10:47:45
func InitCmb(cfg *viper.Viper) *Cmb {
	return &Cmb{

		SassKey:     cfg.GetString("sass_key"),
		SassName:    cfg.GetString("sass_name"),
		Url:         cfg.GetString("url"),
		DefaultText: cfg.GetString("default_text"),
		PriKey:      cfg.GetString("pri_key"),
		Sm2Key:      cfg.GetString("sm2_key"),
		MainAccount: cfg.GetString("main_account"),
		UserId:      cfg.GetString("user_id"),
		PayrollMod:  cfg.GetString("payroll_mod"),
	}
}

var CmbConfig = new(Cmb)
