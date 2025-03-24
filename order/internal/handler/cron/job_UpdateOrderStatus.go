package cron

import (
	"context"
	"log"
	"time"

	"github.com/fenky-ng/edot-test-case/order/internal/constant"
	"github.com/fenky-ng/edot-test-case/order/internal/model"
	pointer_util "github.com/fenky-ng/edot-test-case/order/internal/utility/pointer"
)

func (h *CronHandler) UpdateOrderStatus() {
	log.Println("[JOB] CRON.UpdateOrderStatus started")

	// room for improvement: add distributed lock on cron level

	h.updateExpiredOrderStatus()

	// put another order status updater here

	log.Println("[JOB] CRON.UpdateOrderStatus ended")
}

func (h *CronHandler) updateExpiredOrderStatus() {
	ctx := context.Background()

	ordersOut, err := h.usecase.OrderUsecase.GetOrders(ctx, model.GetOrdersInput{
		GetExpiredOrders: true,
	})
	if err != nil {
		return
	}

	for _, order := range ordersOut.Orders {
		// room for improvement: add distributed lock on order level

		h.processExpiredOrder(ctx, order)
	}
}

func (h *CronHandler) processExpiredOrder(
	ctx context.Context,
	order model.Order,
) {
	_, err := h.usecase.OrderUsecase.ReleaseStocks(ctx, model.ReleaseStocksInput{
		UserId:  order.UserId,
		OrderNo: order.OrderNo,
	})
	if err != nil {
		log.Printf("[JOB] error occurred at UpdateOrderStatus.updateExpiredOrderStatus.processExpiredOrder.ReleaseStocks orderNo: %s, error: %v", order.OrderNo, err)
		return
	}

	_, err = h.usecase.OrderUsecase.UpdateOrder(ctx, model.UpdateOrderInput{
		OrderNo:   order.OrderNo,
		Status:    pointer_util.PointerOf(constant.OrderStatus_Expired),
		UpdatedAt: time.Now().UnixMilli(),
		UpdatedBy: "CRON.UpdateOrderStatus",
	})
	if err != nil {
		log.Printf("[JOB] error occurred at UpdateOrderStatus.updateExpiredOrderStatus.processExpiredOrder.UpdateOrder orderNo: %s, error: %v", order.OrderNo, err)
	}
}
