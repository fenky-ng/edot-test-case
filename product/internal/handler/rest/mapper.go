package rest

import "github.com/fenky-ng/edot-test-case/product/internal/model"

func mapProductStockWarehouses(
	in []model.ProductWarehouse,
) (out []model.RestAPIProductWarehouse) {
	out = make([]model.RestAPIProductWarehouse, 0)
	for _, item := range in {
		out = append(out, model.RestAPIProductWarehouse{
			Id:     item.Id,
			Stock:  item.Stock,
			Status: item.Status,
		})
	}
	return out
}
