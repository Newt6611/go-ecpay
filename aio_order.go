package ecpay

import (
	"crypto/sha256"
	"errors"
	"fmt"
	"net/url"
	"sort"
	"strings"
)

const ERROR_PREFIX = "ECPay "
const TRADE_DATE_FORMAT = "2006/02/01 15:04:05"

type AioOrder struct {
    merchantID string `json:"MerchantID"`
    paymentType string `json:"PaymentType"`
    choosePayment string `json:"ChoosePayment"`
    encryptType int `json:"EncryptType"`
    checkMacValue string `json:"CheckMacValue"`

    // required
    MerchantTradeNo string `json:"MerchantTradeNo"`
    MerchantTradeDate string `json:"MerchantTradeDate"`
    TotalAmount int  `json:"TotalAmount"`
    TradeDesc string `json:"TradeDesc"`
    ItemName string `json:"ItemName"`
    ReturnURL string `json:"ReturnURL"`

    StoreID string `json:"StoreID"`
    ClientBackURL string `json:"ClientBackURL"`
    ItemURL string `json:"ItemURL"`
    Remark string `json:"Remark"`
    ChooseSubPayment string `json:"ChooseSubPayment"`
    OrderResultURL string `json:"OrderResultURL"`
    NeedExtraPaidInfo string `json:"NeedExtraPaidInfo"`
    IgnorePayment string `json:"IgnorePayment"`
    PlatformID string `json:"PlatformID"`
    CustomField1 string `json:"CustomField1"`
    CustomField2 string `json:"CustomField2"`
    CustomField3 string `json:"CustomField3"`
    CustomField4 string `json:"CustomField4"`
    Language string `json:"Language"`
}

func generateCheckMacValue(m map[string]string, hashKey string, hashVI string) string {
    // sort keys
    keys := make([]string, len(m))
    i := 0
    for k := range m {
        keys[i] = k
        i++
    }
    
    sort.SliceStable(keys, func(i, j int) bool {
        return keys[i] < keys[j]
    })

    var result strings.Builder
    for _, val := range keys {
        s := val + "=" + m[val] + "&"
        result.WriteString(s)
    }
    finalString := result.String()
    // remove last &
    finalString = finalString[:len(finalString) - 1]
    finalString = fmt.Sprintf("HashKey=%s&%s&HashIV=%s", hashKey, finalString, hashVI)

    // url encode
    finalString = url.QueryEscape(finalString) 
    // to lowercase
    finalString = strings.ToLower(finalString)
    // sha265
    h := sha256.New()
    h.Write([]byte(finalString)) 
    bs := h.Sum(nil)
    // to string
    finalString = fmt.Sprintf("%x", bs)
    // to upper    
    finalString = strings.ToUpper(finalString)
    
    return finalString
}

func checkOrderField(o *AioOrder) error {
    if len(o.MerchantTradeNo) == 0 {
        return errors.New(ERROR_PREFIX + "MerchantTradeNo is required")
    }
    if len(o.MerchantTradeDate) == 0 {
        return errors.New(ERROR_PREFIX + "MerchantTradeDate is required")
    }
    if len(o.TradeDesc) == 0 {
        return errors.New(ERROR_PREFIX + "TradeDesc is required")
    }
    if len(o.ItemName) == 0 {
        return errors.New(ERROR_PREFIX + "ItemName is required")
    }
    if len(o.ReturnURL) == 0 {
        return errors.New(ERROR_PREFIX + "ReturnURL is required")
    }

    return nil
}
