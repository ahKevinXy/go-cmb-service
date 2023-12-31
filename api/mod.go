package api

import (
	"github.com/ahKevinXy/go-cmb/handler/account"
	"github.com/ahKevinXy/go-cmb/handler/payroll_old"
	"github.com/gin-gonic/gin"
	"go-cmb-service/help"
	"go-cmb-service/models/request"
	"go-cmb-service/models/response"
	"go-cmb-service/pkg/config"
)

// GetPayrollMods
//  @Description:   获取代发业务模式
//  @param c
//  @Author  ahKevinXy
//  @Date  2023-07-10 19:12:29
func GetPayrollMods(c *gin.Context) {
	var req request.PayrollMods

	if err := c.BindJSON(&req); err != nil {

		return
	}
	// 获取业务模式
	mods, err := payroll_old.PayMods(config.CmbConfig.UserId, config.CmbConfig.Sm2Key, config.CmbConfig.PriKey, req.BusCode, config.CmbConfig.MainAccount)
	if err != nil {
		help.ErrorCallBack(c, 500, "服务器错误")
		return
	}
	if len(mods.Response.Body.Ntagtls2Z) == 0 {

		return
	}
	var payrollMods []*response.PayrollMods

	for _, v := range mods.Response.Body.Ntagtls2Z {
		payrollMod := &response.PayrollMods{}

		payrollMod.ModNo = v.Trstyp
		payrollMod.ModName = v.CTrstyp
		payrollMod.EftAt = v.Eftdat
		payrollMod.ExpAt = v.Expdat
		payrollMods = append(payrollMods, payrollMod)

	}
	help.OK(c, payrollMods, "ok")
}

// P2PMods
//  @Description:  公对公业务模式
//  @param c
//  @Author  ahKevinXy
//  @Date  2023-07-15 11:47:32
func P2PMods(c *gin.Context) {
	var req request.P2PMods

	if err := c.BindJSON(&req); err != nil {

		return
	}

	// 获取业务模式
	mods, err := account.PayMods(config.CmbConfig.UserId, config.CmbConfig.Sm2Key, config.CmbConfig.PriKey, req.BusCode)
	if err != nil {
		help.ErrorCallBack(c, 500, "服务器错误")
		return
	}
	if len(mods.Response.Body.Ntqmdlstz) == 0 {

		return
	}
	var payrollMods []*response.PayrollMods

	for _, v := range mods.Response.Body.Ntqmdlstz {
		payrollMod := &response.PayrollMods{}

		payrollMod.ModNo = v.Busmod
		payrollMod.ModName = v.Modals

		payrollMods = append(payrollMods, payrollMod)

	}
	help.OK(c, payrollMods, "ok")

}
