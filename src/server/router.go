package server

import (
	"github.com/gin-gonic/gin"

	"fmt"
	"github.com/Electra-project/electrapay-api/src/authenticators"
	"github.com/Electra-project/electrapay-api/src/controllers"
	"github.com/Electra-project/electrapay-api/src/middlewares"
	"os"
	"strconv"
)

func Router() *gin.Engine {

	authenticator := authenticators.Authenticator()
	apiauthenticator := authenticators.BasicAuth()

	router := gin.Default()
	router.Use(middlewares.CORS())
	router.Use(middlewares.ResponseHeaders())

	var version = os.Getenv("VERSION")
	v, err := strconv.Atoi(version)
	if err != nil {
		v = 1
	}
	vdir := fmt.Sprint("/v", v, "/")

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"name":    "Electrapay API",
			"version": version,
		})
	})

	/**
	 * public routes
	 */
	accountController := new(controllers.AccountController)

	// register a new account - this will send an email with the login code
	router.POST(vdir+"/account/register/", accountController.Register)
	router.POST("/account/register/", accountController.Register)

	// login to create a JWT token
	router.POST(vdir+"/account/authenticate", authenticator.LoginHandler)
	router.POST("/account/authenticate", authenticator.LoginHandler)

	/**
	 * authenticated routes
	 */
	auth := router.Group("/")
	auth.Use(authenticator.MiddlewareFunc())
	{
		auth.GET(vdir+"/account/:accountid", accountController.Get)
		auth.GET("/account/:accountid", accountController.Get)
		auth.PUT(vdir+"/account/:accountid", accountController.Edit)
		auth.PUT("/account/:accountid", accountController.Edit)
		auth.POST(vdir+"/account/close/:accountid", accountController.Close)
		auth.POST("/account/close/:accountid", accountController.Close)
		auth.POST(vdir+"/account/apikey/:accountid", accountController.ApiKey)
		auth.POST("/account/apikey/:accountid", accountController.ApiKey)
		auth.POST(vdir+"/account/suspend/:accountid", accountController.Suspend)
		auth.POST("/account/suspend/:accountid", accountController.Suspend)

		auth.PUT("/account/:accountid/address/:addressid", accountController.AddressEdit)
		auth.PUT(vdir+"/account/:accountid/address/:addressid", accountController.AddressEdit)
		//auth.POST(vdir+"/account/:accountid/addressnew/", accountController.AddressAdd)
		//auth.POST("/account/:accountid/address", accountController.AddressAdd)
		auth.DELETE(vdir+"/account/:accountid/address/:addressid", accountController.AddressRemove)
		auth.DELETE("/account/:accountid/address/:addressid", accountController.AddressRemove)
		auth.PUT(vdir+"/account/:accountid/contact/:contactid", accountController.ContactEdit)
		auth.PUT("/account/:accountid/contact/:contactid", accountController.ContactEdit)
		//auth.POST(vdir+"/account/:accountid/contact/new", accountController.ContactAdd)
		//auth.POST("/account/:accountid/contact/new", accountController.ContactAdd)
		auth.DELETE(vdir+"/account/:accountid/contact/:contactid", accountController.ContactRemove)
		auth.DELETE("/account/:accountid/contact/:contactid", accountController.ContactRemove)
	}

	authapi := router.Group("/")
	authapi.Use(apiauthenticator)
	{

		orderController := new(controllers.OrderController)
		authapi.POST(vdir+"/order/", orderController.New)
		authapi.POST("/order/", orderController.New)
		authapi.GET(vdir+"/order/:uuid", orderController.Get)
		authapi.GET("/order/:uuid/", orderController.Get)
		authapi.POST(vdir+"/order/:uuid/cancel", orderController.Cancel)
		authapi.POST("/order/:uuid/cancel", orderController.Cancel)
		authapi.POST(vdir+"/order/:uuid/reverse", orderController.Reverse)
		authapi.POST("/order/:uuid/reverse", orderController.Reverse)
		authapi.GET("/paymentcategory/:accountid", orderController.PaymentCategory)
		authapi.GET(vdir+"/paymentcategory/:accountid", orderController.PaymentCategory)
		authapi.GET("/allowedcurrency/:accountid", orderController.AllowedCurrency)
		authapi.GET(vdir+"/allowedcurrency/:accountid", orderController.AllowedCurrency)
	}

	codeController := new(controllers.CodeController)
	router.GET(vdir+"/accounttype/", codeController.GetAccountType)
	router.GET("accounttype/", codeController.GetAccountType)
	router.GET(vdir+"/addresstype/", codeController.GetAddressType)
	router.GET("addresstype/", codeController.GetAddressType)
	router.GET(vdir+"/contacttype/", codeController.GetContactType)
	router.GET("contacttype/", codeController.GetContactType)
	router.GET(vdir+"/currencytype/", codeController.GetCurrencyType)
	router.GET("currencytype/", codeController.GetCurrencyType)
	router.GET("plugintype/", codeController.GetPluginType)
	router.GET(vdir+"/plugintype/", codeController.GetPluginType)
	router.GET(vdir+"/currency/", codeController.GetCurrency)
	router.GET("currency/", codeController.GetCurrency)
	router.GET(vdir+"/language/", codeController.GetLanguage)
	router.GET("language/", codeController.GetLanguage)
	router.GET(vdir+"/country/", codeController.GetCountry)
	router.GET("country/", codeController.GetCountry)
	router.GET(vdir+"/timezone/", codeController.GetTimeZone)
	router.GET("timezone/", codeController.GetTimeZone)

	router.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"message": "Not Found"})
	})

	return router
}
