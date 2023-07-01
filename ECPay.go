package ecpay

import (
	"strconv"
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

func (ec *ECPay) CreateOrder(order *Order) (string, error) {
    order.merchantID = ec.merchantId
    order.paymentType = "aio"
    order.choosePayment = "ALL"
    order.encryptType = 1

    err := checkOrderField(order)
    if err != nil {
        return "", err
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
        return "", err 
    }

    html := "<form id=\"data_set\" action=" + string(ec.endpoint) + " method=\"post\">"
    for k, v := range m {
        html += "<input type=\"hidden\" name=" + k + " value='" + v + "' />"
    } 
    html += "<script type=\"text/javascript\">document.getElementById(\"data_set\").submit();</script>"
    html += "</form>"

    return html, nil
}
