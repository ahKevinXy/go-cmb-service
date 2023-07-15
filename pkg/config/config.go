package config

import (
	"fmt"
	"github.com/spf13/viper"
	"io/ioutil"
	"os"
	"strings"
)

var cfgApplication *viper.Viper

var cfgCmb *viper.Viper

func InitConfig(configPath string) {
	var err error
	viper.SetConfigFile(configPath)

	content, err := ioutil.ReadFile(configPath)
	fmt.Println(configPath, "999999")
	if err != nil {
		fmt.Println("读取配置文件失败:%+v", err)
	}
	err = viper.ReadConfig(strings.NewReader(os.ExpandEnv(string(content))))

	if err != nil {
		fmt.Println("读取配置文件失败:---", err)
		return
	}

	err = viper.ReadConfig(strings.NewReader(os.ExpandEnv(string(content))))

	if err != nil {
		fmt.Println("读取配置文件失败:---", err)
		return
	}

	// 初始化应用配置文件
	cfgApplication = viper.Sub("application")

	if cfgApplication == nil {
		panic("初始化应用配置文件失败")
	}
	ApplicationConfig = InitApplication(cfgApplication)

	// 初始化 cmb 配置
	cfgCmb = viper.Sub("cmb")

	if cfgCmb == nil {
		panic("config not found database")
	}
	CmbConfig = InitCmb(cfgCmb)

}
