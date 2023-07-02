this is only work for "all in one"

a little example
```
ec := ecpay.NewECPay("3002607", ecpay.DEV_ENDPOINT, "pwFHCqoQZGmho4w6", "EkRm7iFT261dpevs")

order := ecpay.AioOrder{
    TradeDesc: "促銷方案",
    MerchantTradeDate: ec.GetFormatedTime(),
    MerchantTradeNo: ecpay.GenerateMerchantNo(),
    ReturnURL: "http://localhost:8080/returnUrl",
    ItemName: "Apple iphone 7 手機殼",
    TotalAmount: 1000,
    IgnorePayment: ecpay.IgnorePayment_ATM + "#" + ecpay.IgnorePayment_CVS + "#" + ecpay.IgnorePayment_TWQR,
    Language: ecpay.Language_JPN,
}
str, _ := ec.CreateOrder(&order)
fmt.Println(str)
```
