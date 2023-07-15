package api

import (
	"context"
	"fmt"
	cmbConfig "github.com/ahKevinXy/go-cmb/config"
	"github.com/ahKevinXy/go-web-tools/help/utils"
	"log"

	"github.com/spf13/cobra"
	"go-cmb-service/pkg/config"
	"go-cmb-service/routers"
	"net/http"
	"os"
	"os/signal"
	"time"
)

var rootCmd = &cobra.Command{}

var (
	configPath string //配置文件
	port       string
	mode       string
	err        error
	StartCmd   = &cobra.Command{
		Use:   "cmb",
		Short: "Start Cmb Server",
		// 预先处理命令
		PreRun: func(cmd *cobra.Command, args []string) {

			setup()
			usage()
		},
		//
		RunE: func(cmd *cobra.Command, args []string) error {
			return run()
		},
	}
)

func setup() {

	// 初始化 配置文件
	config.InitConfig(configPath)

	// 初始化 配置文件
	cmbConfig.InitConfig(config.CmbConfig.SassName, config.CmbConfig.Url, config.CmbConfig.SassKey, config.CmbConfig.DefaultText)

	if err != nil {

		return
	}

}

func run() error {

	r := routers.Routers()

	srv := &http.Server{
		Addr:    config.ApplicationConfig.Host + ":" + config.ApplicationConfig.Port,
		Handler: r,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %+v", err)
		}
	}()
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		fmt.Println("关闭 ---->", time.Now().Format(utils.GoLangTimeFormat))
		log.Println("Server exiting")
		log.Fatal("Server Shutdown:", err)
	}
	return nil

}
func usage() {
	usageStr := `starting daily report server`
	fmt.Println(usageStr)
}

// 初始化
func init() {
	StartCmd.PersistentFlags().StringVarP(&configPath, "config", "c", "configs/cmb.yaml", "Start server with provided configuration file")
	StartCmd.PersistentFlags().StringVarP(&port, "port", "p", "8000", "Tcp port server listening on")
	StartCmd.PersistentFlags().StringVarP(&mode, "mode", "m", "dev", "server mode ; eg:dev,test,prod")
}
