package main

import "newoo/WinningDerbyAdmin/controller"
import "github.com/gin-gonic/gin"

var varRouter = new(WdRouter)

type WdRouter struct {
	e *gin.Engine
}

func (r *WdRouter) Init(engine *gin.Engine) {
	r.e = engine
}

func (r *WdRouter) LoginRouter() {
	r.e.GET("/", controller.LoginController)
	r.e.POST("/login_account", controller.PostLogin)
	r.e.GET("/logout", controller.LogOut)
	// r.e.GET("/error", controller.OnError)
	// r.e.POST("/data_toexcel", controller.ToExcelController)

}

func (r *WdRouter) UserRouter() {
	checker := controller.MenuCheckerNew()

	authGroup := r.e.Group("/auth")
	authGroup.Use(controller.RequireLogin)
	authGroup.GET("/user", controller.UserController)
	checker.AuthAdd("/user", "User")
	authGroup.POST("/user_search", controller.UserSearch)
	checker.AuthAdd("/user_search", "User")
	authGroup.POST("/user_save", controller.UserDataSave)
	checker.AuthAdd("/user_save", "User")

	authGroup.GET("/purchase", controller.PurchaseController)
	checker.AuthAdd("/purchase", "Purchase")
	authGroup.POST("/purchase_detail", controller.PurchaseDetail)
	checker.AuthAdd("/purchase_detail", "Purchase")

	authGroup.GET("/coin", controller.CoinController)
	checker.AuthAdd("/coin", "Coin")
	authGroup.POST("/coin_detail", controller.CoinDetail)
	checker.AuthAdd("/coin_detail", "Coin")

	authGroup.GET("/dia", controller.DiaController)
	checker.AuthAdd("/dia", "Dia")
	authGroup.POST("/dia_detail", controller.DiaDetail)
	checker.AuthAdd("/dia_detail", "Dia")

	authGroup.GET("/mileage", controller.MileageController)
	checker.AuthAdd("/mileage", "Mileage")
	authGroup.POST("/mileage_detail", controller.MileageDetail)
	checker.AuthAdd("/mileage_detail", "Mileage")

	authGroup.GET("/mailbox", controller.MailboxController)
	checker.AuthAdd("/mailbox", "MailBox")
	authGroup.POST("/mailbox_detail", controller.MailboxDetail)
	checker.AuthAdd("/mailbox_detail", "MailBox")

	authGroup.GET("/payout", controller.PayoutController)
	checker.AuthAdd("/payout", "PayOut")
	authGroup.POST("/payout_detail", controller.PayoutDetail)
	checker.AuthAdd("/payout_detail", "PayOut")

}

func (r *WdRouter) AdminRouter() {
	checker := controller.MenuCheckerNew()
	authGroup := r.e.Group("/auth")
	authGroup.Use(controller.RequireLogin)
	authGroup.GET("/sendmail", controller.SendMailController)
	checker.AuthAdd("/sendmail", "Send Mail")
	authGroup.POST("/send_mail", controller.SendMail)
	checker.AuthAdd("/send_mail", "Send Mail")

	authGroup.GET("/admin", controller.AdminController)
	checker.AuthAdd("/admin", "Admin")
	authGroup.POST("/account_list", controller.GetAccountList)
	checker.AuthAdd("/account_list", "Admin")
	authGroup.POST("/account_apply", controller.AccountApply)
	checker.AuthAdd("/account_apply", "Admin")
	authGroup.POST("/account_remove", controller.AccountRemove)
	checker.AuthAdd("/account_remove", "Admin")

	authGroup.GET("/admin_log", controller.AdminLogController)
	checker.AuthAdd("/admin_log", "Admin")
	authGroup.POST("/admin_log_search", controller.AdminLogSearch)
	checker.AuthAdd("/admin_log_search", "Admin")

}

func (r *WdRouter) StatistcsRouter() {
	checker := controller.MenuCheckerNew()
	authGroup := r.e.Group("/auth")
	authGroup.Use(controller.RequireLogin)
	authGroup.GET("/statistcs_dia", controller.StDiaController)
	checker.AuthAdd("/statistcs_dia", "Statistics_Dia")
	authGroup.POST("/stattistics_dai_search", controller.STDaiSearch)
	checker.AuthAdd("/stattistics_dai_search", "Statistics_Dia")

	authGroup.GET("/statistcs_coin", controller.StCoinController)
	checker.AuthAdd("/statistcs_coin", "Statistics_Coin")
	authGroup.POST("/stattistics_coin_search", controller.STCoinSearch)
	checker.AuthAdd("/stattistics_coin_search", "Statistics_Coin")

	authGroup.GET("/statistics_dau", controller.StDauController)
	checker.AuthAdd("/statistics_dau", "Statistics_Dau")
	authGroup.POST("/stattistics_dau_search", controller.STDAUSearch)
	checker.AuthAdd("/stattistics_dau_search", "Statistics_Dau")

	authGroup.GET("/statistics_ccu", controller.StCCUController)
	checker.AuthAdd("/statistics_ccu", "Statistics_CCU")
	authGroup.POST("/ccu_search", controller.STCCUSearch)
	checker.AuthAdd("/ccu_search", "Statistics_CCU")

}

func (r *WdRouter) SetRoute() {
	r.LoginRouter()
	r.UserRouter()
	r.AdminRouter()
	r.StatistcsRouter()
}
