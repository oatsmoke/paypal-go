# Go PayPal

[PayPal REST API Docs](https://developer.paypal.com/api/rest/)

[Payment Method Tokens API](https://developer.paypal.com/docs/multiparty/checkout/save-payment-methods/purchase-later/payment-tokens-api/cards/)

## Install
```
go get -u github.com/oatsmoke/paypal-go
```

## Init
```
import (
    "github.com/oatsmoke/paypal-go"
    "github.com/oatsmoke/paypal-go/model"
)

//...

ctx := context.Background()
client := paypal.NewClient(ClientID, Secret, URL)
```

## CreateSetupToken
```
data := &model.PayPal{
	Customer: &model.Customer{
		Id:                 UserId,
		MerchantCustomerId: Email,
	},
	PaymentSource: &model.PaymentSource{
		Card: &model.Card{
			Number: Number,
			Expiry: Expiry,
			Name:   Name,
		},
	},
}

res, err := client.CreateSetupToken(ctx, data)
```

## CreatePaymentToken
```
data := &model.PayPal{
	PaymentSource: &model.PaymentSource{
		Token: &model.Token{
			Id:   SetupToken,
			Type: "SETUP_TOKEN",
		},
	},
}

res, err := client.CreatePaymentToken(ctx, data)
```

## DeletePaymentToken
```
err := client.DeletePaymentToken(ctx, PaymentToken)
```

## RetrievePaymentToken
```
res, err := client.RetrievePaymentToken(ctx, PaymentToken)
```

## CreateOrder
```
data := &model.PayPal{
	Intent: "CAPTURE",
	PaymentSource: &model.PaymentSource{
		Card: &model.Card{
			VaultId: PaymentToken,
		},
	},
	PurchaseUnits: []*model.PurchaseUnit{},
}
data.PurchaseUnits = append(data.PurchaseUnits, &model.PurchaseUnit{
	Amount: &model.Amount{
		CurrencyCode: CurrencyCode,
		Value:        Value,
	},
})

res, err := client.CreateOrder(ctx, data)
```

## RefundCapturedPayment
```
res, err := client.RefundCapturedPayment(ctx, PaymentId)
```

## ShowRefundDetails
```
res, err := client.ShowRefundDetails(ctx, RefandId)
```
