package controllers

import "github.com/gin-gonic/gin"
import (
	"encoding/json"
	"github.com/ruannelloyd/electrapay-api/src/helpers"
	"github.com/ruannelloyd/electrapay-api/src/queue"
	"strings"
)

type AccountType struct {
	Id          int64  `json:"id"`
	Code        string `json:"code"`
	Description string `json:"description"`
}

type AddressType struct {
	Id   int64  `json:"id"`
	Code string `json:"code"`
}

type ContactType struct {
	Id   int64  `json:"id"`
	Code string `json:"code"`
}

type CurrencyType struct {
	Id   int64  `json:"id"`
	Code string `json:"code"`
}

type PluginType struct {
	Id   int64  `json:"id"`
	Code string `json:"code"`
}

type Currency struct {
	Id           int64  `json:"id"`
	Code         string `json:"code"`
	Description  string `json:"description"`
	Iban         string `json:"iban"`
	Symbol       string `json:"symbol"`
	Codeapi      string `json:"codeapi"`
	CurrencyType string `json:"currencytype"`
}

type Language struct {
	Id          int64  `json:"id"`
	Code        string `json:"code"`
	Description string `json:"description"`
}

type Country struct {
	Id          int64  `json:"id"`
	Code        string `json:"code"`
	Description string `json:"description"`
}

type Timezone struct {
	Id          int64  `json:"id"`
	Code        string `json:"code"`
	CountryCode string `json:"countrycode"`
	UTCOffset   string `json:"utcoffset"`
}

type CodeController struct{}

func (s CodeController) GetAccountType(c *gin.Context) {
	//API to retrieve static data information

	var queueinfo queue.Queue
	queueinfo.Category = "ACCOUNTTYPE_FIND"
	queueinfo.APIType = "GET"
	URLArray := strings.Split(c.Request.RequestURI, "/")
	if len(URLArray[2]) >= 1 {
		queueinfo.APIURL = c.Request.RequestURI
		queueinfo.Parameters = ""
		queueinfo.Version = URLArray[1]
	}
	version := helpers.GetVersion()
	if len(URLArray[2]) < 1 {
		queueinfo.APIURL = c.Request.RequestURI
		queueinfo.Parameters = ""
		queueinfo.Version = version
	}
	queueinfo.RequestInfo = "{}"
	queueinfo, err := queue.QueueProcess(queueinfo)
	if err != nil {
		c.AbortWithError(404, err)
		return
	}

	var accounttypes []AccountType
	queueResult := queueinfo.ResponseInfo
	json.Unmarshal([]byte(queueResult), &accounttypes)

	c.JSON(200, accounttypes)

}

func (s CodeController) GetAddressType(c *gin.Context) {
	//API to retrieve static data information

	var queueinfo queue.Queue
	version := helpers.GetVersion()
	queueinfo.Category = "ADDRESSTYPE_FIND"
	queueinfo.QueueCategory = "queue"
	queueinfo.APIType = "GET"
	URLArray := strings.Split(c.Request.RequestURI, "/")
	if len(URLArray[2]) >= 1 {
		queueinfo.APIURL = c.Request.RequestURI
		queueinfo.Parameters = ""
		queueinfo.Version = URLArray[1]
	}
	if len(URLArray[2]) < 1 {
		queueinfo.APIURL = c.Request.RequestURI
		queueinfo.Parameters = ""
		queueinfo.Version = version
	}
	queueinfo.RequestInfo = "{}"
	queueinfo, err := queue.QueueProcess(queueinfo)
	if err != nil {
		c.AbortWithError(404, err)
		return
	}

	var addresstypes []AddressType
	queueResult := queueinfo.ResponseInfo
	json.Unmarshal([]byte(queueResult), &addresstypes)

	c.JSON(200, addresstypes)

}

func (s CodeController) GetContactType(c *gin.Context) {
	//API to retrieve static data information

	var queueinfo queue.Queue
	queueinfo.Category = "CONTACTTYPE_FIND"
	queueinfo.QueueCategory = "queue"
	queueinfo.APIType = "GET"
	version := helpers.GetVersion()
	URLArray := strings.Split(c.Request.RequestURI, "/")
	if len(URLArray[2]) >= 1 {
		queueinfo.APIURL = c.Request.RequestURI
		queueinfo.Parameters = ""
		queueinfo.Version = URLArray[1]
	}
	if len(URLArray[2]) < 1 {
		queueinfo.APIURL = c.Request.RequestURI
		queueinfo.Parameters = ""
		queueinfo.Version = version
	}
	queueinfo.RequestInfo = "{}"
	queueinfo, err := queue.QueueProcess(queueinfo)
	if err != nil {
		c.AbortWithError(404, err)
		return
	}

	var contacttypes []ContactType
	queueResult := queueinfo.ResponseInfo
	json.Unmarshal([]byte(queueResult), &contacttypes)

	c.JSON(200, contacttypes)

}

func (s CodeController) GetCurrencyType(c *gin.Context) {
	//API to retrieve static data information

	var queueinfo queue.Queue
	version := helpers.GetVersion()
	queueinfo.Category = "CURRENCYTYPE_FIND"
	queueinfo.QueueCategory = "queue"
	queueinfo.APIType = "GET"
	URLArray := strings.Split(c.Request.RequestURI, "/")
	if len(URLArray[2]) >= 1 {
		queueinfo.APIURL = c.Request.RequestURI
		queueinfo.Parameters = ""
		queueinfo.Version = URLArray[1]
	}
	if len(URLArray[2]) < 1 {
		queueinfo.APIURL = c.Request.RequestURI
		queueinfo.Parameters = ""
		queueinfo.Version = version
	}
	queueinfo.RequestInfo = "{}"
	queueinfo, err := queue.QueueProcess(queueinfo)
	if err != nil {
		c.AbortWithError(404, err)
		return
	}
	var currencytypes []CurrencyType
	queueResult := queueinfo.ResponseInfo
	json.Unmarshal([]byte(queueResult), &currencytypes)

	c.JSON(200, currencytypes)

}

func (s CodeController) GetPluginType(c *gin.Context) {
	//API to retrieve static data information

	var queueinfo queue.Queue
	queueinfo.Category = "PLUGINTYPE_FIND"
	queueinfo.QueueCategory = "queue"
	queueinfo.APIType = "GET"
	version := helpers.GetVersion()
	URLArray := strings.Split(c.Request.RequestURI, "/")
	if len(URLArray[2]) >= 1 {
		queueinfo.APIURL = c.Request.RequestURI
		queueinfo.Parameters = ""
		queueinfo.Version = URLArray[1]
	}
	if len(URLArray[2]) < 1 {
		queueinfo.APIURL = c.Request.RequestURI
		queueinfo.Parameters = ""
		queueinfo.Version = version
	}
	queueinfo.RequestInfo = "{}"
	queueinfo, err := queue.QueueProcess(queueinfo)
	if err != nil {
		c.AbortWithError(404, err)
		return
	}
	var plugintypes []PluginType
	queueResult := queueinfo.ResponseInfo
	json.Unmarshal([]byte(queueResult), &plugintypes)

	c.JSON(200, plugintypes)

}

func (s CodeController) GetCurrency(c *gin.Context) {
	//API to retrieve static data information

	var queueinfo queue.Queue
	queueinfo.Category = "CURRENCY_FIND"
	queueinfo.QueueCategory = "queue"
	queueinfo.APIType = "GET"
	version := helpers.GetVersion()
	URLArray := strings.Split(c.Request.RequestURI, "/")
	if len(URLArray[2]) >= 1 {
		queueinfo.APIURL = c.Request.RequestURI
		queueinfo.Parameters = ""
		queueinfo.Version = URLArray[1]
	}
	if len(URLArray[2]) < 1 {
		queueinfo.APIURL = c.Request.RequestURI
		queueinfo.Parameters = ""
		queueinfo.Version = version
	}
	queueinfo.RequestInfo = "{}"
	queueinfo.ResponseInfo = ""
	queueinfo, err := queue.QueueProcess(queueinfo)
	if err != nil {
		c.AbortWithError(404, err)
		return
	}
	var currencies []Currency
	queueResult := queueinfo.ResponseInfo

	json.Unmarshal([]byte(queueResult), &currencies)

	c.JSON(200, currencies)

}

func (s CodeController) GetCountry(c *gin.Context) {
	//API to retrieve static data information

	var queueinfo queue.Queue
	queueinfo.Category = "COUNTRY_FIND"
	queueinfo.QueueCategory = "queue"
	queueinfo.APIType = "GET"
	version := helpers.GetVersion()
	URLArray := strings.Split(c.Request.RequestURI, "/")
	if len(URLArray[2]) >= 1 {
		queueinfo.APIURL = c.Request.RequestURI
		queueinfo.Parameters = ""
		queueinfo.Version = URLArray[1]
	}
	if len(URLArray[2]) < 1 {
		queueinfo.APIURL = c.Request.RequestURI
		queueinfo.Parameters = ""
		queueinfo.Version = version
	}
	queueinfo.RequestInfo = "{}"
	queueinfo, err := queue.QueueProcess(queueinfo)
	if err != nil {
		c.AbortWithError(404, err)
		return
	}
	var countries []Country
	queueResult := queueinfo.ResponseInfo
	json.Unmarshal([]byte(queueResult), &countries)

	c.JSON(200, countries)

}

func (s CodeController) GetLanguage(c *gin.Context) {
	//API to retrieve static data information

	var queueinfo queue.Queue
	queueinfo.Category = "LANGUAGE_FIND"
	queueinfo.QueueCategory = "queue"
	queueinfo.APIType = "GET"
	version := helpers.GetVersion()
	URLArray := strings.Split(c.Request.RequestURI, "/")
	if len(URLArray[2]) >= 1 {
		queueinfo.APIURL = c.Request.RequestURI
		queueinfo.Parameters = ""
		queueinfo.Version = URLArray[1]
	}
	if len(URLArray[2]) < 1 {
		queueinfo.APIURL = c.Request.RequestURI
		queueinfo.Parameters = ""
		queueinfo.Version = version
	}
	queueinfo.RequestInfo = "{}"
	queueinfo.ResponseInfo = ""
	queueinfo, err := queue.QueueProcess(queueinfo)
	if err != nil {
		c.AbortWithError(404, err)
		return
	}
	var languages []Language
	queueResult := queueinfo.ResponseInfo
	json.Unmarshal([]byte(queueResult), &languages)

	c.JSON(200, languages)

}

func (s CodeController) GetTimeZone(c *gin.Context) {
	//API to retrieve static data information

	var queueinfo queue.Queue
	queueinfo.Category = "TIMEZONE_FIND"
	queueinfo.QueueCategory = "queue"
	queueinfo.APIType = "GET"
	URLArray := strings.Split(c.Request.RequestURI, "/")
	version := helpers.GetVersion()
	if len(URLArray[2]) >= 1 {
		queueinfo.APIURL = c.Request.RequestURI
		queueinfo.Parameters = ""
		queueinfo.Version = URLArray[1]
	}
	if len(URLArray[2]) < 1 {
		queueinfo.APIURL = c.Request.RequestURI
		queueinfo.Parameters = ""
		queueinfo.Version = version
	}
	queueinfo.RequestInfo = "{}"
	queueinfo, err := queue.QueueProcess(queueinfo)
	if err != nil {
		c.AbortWithError(404, err)
		return
	}
	var timezones []Timezone
	queueResult := queueinfo.ResponseInfo
	json.Unmarshal([]byte(queueResult), &timezones)

	c.JSON(200, timezones)

}
