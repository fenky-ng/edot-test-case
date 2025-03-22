package product

import (
	"context"

	"github.com/fenky-ng/edot-test-case/product/internal/constant"
	"github.com/fenky-ng/edot-test-case/product/internal/model"
)

func (u *ProductUsecase) GetProductById(ctx context.Context, input model.GetProductByIdInput) (output model.GetProductByIdOutput, err error) {
	productOut, err := u.repoDbProduct.GetProduct(ctx, model.GetProductInput{
		Id: input.Id,
	})
	if err != nil {
		return output, err
	}

	// shop
	shopById, err := u.getShops(ctx, []model.Product{
		productOut.Product,
	})
	if err != nil {
		return output, err
	}

	shop := shopById[productOut.Product.Shop.Id]
	productOut.Product.Shop.Name = shop.Name
	productOut.Product.Shop.Status = shop.Status

	// stock
	stocksByProductId, err := u.getStocks(ctx, []model.Product{
		productOut.Product,
	})
	if err != nil {
		return output, err
	}

	stocks := stocksByProductId[productOut.Product.Id]
	for _, stock := range stocks {
		if stock.WarehouseStatus == constant.WarehouseStatus_Inactive { // only show stock from active warehouse
			continue
		}
		productOut.Product.Stock.Total += stock.Stock
		productOut.Product.Stock.Warehouses = append(productOut.Product.Stock.Warehouses, stock)
	}

	output.Product = productOut.Product

	return output, nil
}
