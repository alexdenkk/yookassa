# yookassa
 Yookassa SDK with basic functions

<br>

# Использование

<br>

## Подключение
```go
package main

import (
       "github.com/alexdenkk/yookassa"
)

func main() {
       client := yookassa.NewClient("id-магазина", "секретный-ключ")
}
```

<br>

## Создание платежа
```go
paymentReq := yookassa.Payment{
       Amount: yookassa.Amount{
              Value:    "100.00",
              Currency: "RUB",
       },
       Description: "Покупка товара",
       Capture:     false, // если false - требуется подтверждение через метод CapturePayment
       Confirmation: yookassa.Redirect{
              Type: yookassa.TypeRedirect,
              ReturnURL: "https://github.com/alexdenkk",
       },
       PaymentMethod: yookassa.PaymentTypeBankCard,
}

payment, err := client.CreatePayment(paymentReq)

if err != nil {
       log.Fatalf("Error creating payment: %v", err)
}
```

<br>

## Получение информации о платеже
```go
paymentInfo, err := client.GetPayment(payment.ID)
if err != nil {
       log.Fatalf("Error getting payment: %v", err)
}
```

<br>

## Подтверждение платежа
 Подтвердить можно только платеж в статусе `"waiting_for_capture"`
```go
capturedPayment, err := client.CapturePayment(payment.ID)

if err != nil {
       log.Fatalf("Error capturing payment: %v", err)
}
```

<br>

## Отмена платежа
```go
canceledPayment, err := client.CancelPayment(payment.ID)

if err != nil {
       log.Fatalf("Error canceling payment: %v", err)
}
```
