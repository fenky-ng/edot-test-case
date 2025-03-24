package model

import "github.com/google/uuid"

type ExtDeductStockItem struct {
	ProductId   uuid.UUID
	WarehouseId uuid.UUID
	Quantity    int64
}

type DeductStocksInput struct {
	UserId  uuid.UUID
	OrderNo string
	Items   []ExtDeductStockItem
	Release bool
}

type DeductStocksOutput struct {
	Successful bool
}

type HttpDeductStocksRequest struct {
	UserId  uuid.UUID             `json:"userId"`
	OrderNo string                `json:"orderNo"`
	Items   []HttpDeductStockItem `json:"items"`
	Release bool                  `json:"release"`
}

type HttpDeductStockItem struct {
	ProductId   uuid.UUID `json:"productId"`
	WarehouseId uuid.UUID `json:"warehouseId"`
	Quantity    int64     `json:"quantity"`
}

type HttpDeductStocksResponse struct {
	Error      string `json:"error"`
	Successful bool   `json:"successful"`
}
