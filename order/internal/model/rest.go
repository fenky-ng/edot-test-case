package model

import (
	"github.com/fenky-ng/edot-test-case/order/internal/constant"
	"github.com/google/uuid"
)

type RestAPIErrorResponse struct {
	Error string `json:"error"`
}

type RestAPIOrderItem struct {
	ProductId   uuid.UUID `json:"productId"`
	WarehouseId uuid.UUID `json:"warehouseId"`
	Quantity    int64     `json:"quantity"`
}

type RestAPICreateOrderRequest struct {
	Items []RestAPIOrderItem `json:"items"`
}

type RestAPICreateOrderResponse struct {
	OrderNo string               `json:"orderNo"`
	Status  constant.OrderStatus `json:"status"`
}

type RestAPIOrder struct {
	Id           uuid.UUID            `json:"id"`
	UserId       uuid.UUID            `json:"userId"`
	OrderNo      string               `json:"orderNo"`
	Status       constant.OrderStatus `json:"status"`
	PaymentRefNo string               `json:"paymentRefNo,omitempty"`
}

type RestAPIGetOrdersResponse struct {
	Orders []RestAPIOrder `json:"data"`
}

type RestAPIConfirmPaymentRequest struct {
	OrderNo      string `json:"orderNo"`
	PaymentRefNo string `json:"paymentRefNo"`
}

type RestAPIConfirmPaymentResponse struct {
	Successful bool `json:"successful"`
}
