package constant

import "time"

type OrderStatus string

const (
	OrderStatus_WaitingForPayment   OrderStatus = "WAITING_FOR_PAYMENT"
	OrderStatus_CancelledStockIssue OrderStatus = "CANCELLED_STOCK_ISSUE"
	OrderStatus_Expired             OrderStatus = "EXPIRED"
	OrderStatus_Paid                OrderStatus = "PAID"
)

func (t OrderStatus) String() string {
	return string(t)
}

const (
	OrderPaymentExpiryDuration = 5 * time.Minute
)
