package ecpay

import (
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

type ENDPOINT string
const (
    DEV_ENDPOINT  ENDPOINT = "https://payment-stage.ecpay.com.tw/Cashier/AioCheckOut/V5"
	PROD_ENDPOINT ENDPOINT = "https://payment.ecpay.com.tw/Cashier/AioCheckOut/V5"
)

type ECPay struct {
	endpoint ENDPOINT 
    merchantId string
    hashKey string
    hashIV string
}

func NewECPay(merchantId string, endpoint ENDPOINT, hashKey string, hashIV string) *ECPay {
    e := &ECPay {
        merchantId: merchantId,
        endpoint: endpoint,
        hashKey: hashKey,
        hashIV: hashIV,
    }
	return e
}

func (ec *ECPay) CreateOrder(order *Order) error {
    order.merchantID = ec.merchantId
    order.paymentType = "aio"
    order.choosePayment = "ALL"
    order.encryptType = 1

    err := checkOrderField(order)
    if err != nil {
        return err
    }

    // required
    m := map[string]string {
        "MerchantID": order.merchantID,
        "PaymentType": order.paymentType,
        "ChoosePayment": order.choosePayment,
        "EncryptType": strconv.Itoa(order.encryptType),
        "MerchantTradeDate": order.MerchantTradeDate,
        "MerchantTradeNo": order.MerchantTradeNo,
        "TotalAmount": strconv.Itoa(order.TotalAmount),
        "TradeDesc": order.TradeDesc,
        "ItemName": order.ItemName,
        "ReturnURL": order.ReturnURL,
    }

    if len(order.StoreID) != 0 {
        m["StoreID"] = order.StoreID
    }
    if len(order.ClientBackURL) != 0 {
        m["ClientBackURL"] = order.ClientBackURL
    }
    if len(order.ItemURL) != 0 {
        m["ItemURL"] = order.ItemURL
    }
    if len(order.Remark) != 0 {
        m["Remark"] = order.Remark
    }
    if len(order.ChooseSubPayment) != 0 {
        m["ChooseSubPayment"] = order.ChooseSubPayment
    }
    if len(order.OrderResultURL) != 0 {
        m["OrderResultURL"] = order.OrderResultURL
    }
    if len(order.NeedExtraPaidInfo) != 0 {
        m["NeedExtraPaidInfo"] = order.NeedExtraPaidInfo
    }
    if len(order.IgnorePayment) != 0 {
        m["IgnorePayment"] = order.IgnorePayment
    }
    if len(order.PlatformID) != 0 {
        m["PlatformID"] = order.PlatformID
    }
    if len(order.CustomField1) != 0 {
        m["CustomField1"] = order.CustomField1
    }
    if len(order.CustomField2) != 0 {
        m["CustomField2"] = order.CustomField2
    }
    if len(order.CustomField3) != 0 {
        m["CustomField3"] = order.CustomField3
    }
    if len(order.CustomField4) != 0 {
        m["CustomField4"] = order.CustomField4
    }
    if len(order.Language) != 0 {
        m["Language"] = order.Language
    }

    m["CheckMacValue"] = generateCheckMacValue(m, ec.hashKey, ec.hashIV)

    if err != nil {
        return err 
    }

    data := url.Values{}
    for k, v := range m {
        data.Add(k, v)
    }
    // data.Add("MerchantID", order.merchantID)
    // data.Add("PaymentType", order.paymentType)
    // data.Add("ChoosePayment", order.choosePayment)
    // data.Add("EncryptType", strconv.Itoa(order.encryptType))
    // data.Add("CheckMacValue", order.checkMacValue)
    // data.Add("MerchantTradeNo", order.MerchantTradeNo)
    // data.Add("MerchantTradeDate", order.MerchantTradeDate)
    // data.Add("TotalAmount", strconv.Itoa(order.TotalAmount))
    // data.Add("TradeDesc", order.TradeDesc)
    // data.Add("ItemName", order.ItemName)
    // data.Add("ReturnURL", order.ReturnURL)

    // data.Add("StoreID", order.StoreID)
    // data.Add("ClientBackURL", order.ClientBackURL)
    // data.Add("ItemURL", order.ItemURL)
    // data.Add("Remark", order.Remark)
    // data.Add("ChooseSubPayment", order.ChooseSubPayment)
    // data.Add("OrderResultURL", order.OrderResultURL)
    // data.Add("NeedExtraPaidInfo", order.NeedExtraPaidInfo)
    // data.Add("IgnorePayment", order.IgnorePayment)
    // data.Add("PlatformID", order.PlatformID)
    // data.Add("CustomField1", order.CustomField1)
    // data.Add("CustomField2", order.CustomField2)
    // data.Add("CustomField3", order.CustomField3)
    // data.Add("CustomField4", order.CustomField4)
    // data.Add("Language", order.Language)

    ec.send(data)

    return nil
}

func (ec *ECPay) send(form url.Values) {
    req, err := http.NewRequest("POST", string(ec.endpoint), strings.NewReader(form.Encode()))
    if err != nil {
        log.Fatalln(err)
    }

    req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
    res, err := http.DefaultClient.Do(req)
    if err != nil {
        log.Fatal(err)
    }

    defer res.Body.Close()
}
