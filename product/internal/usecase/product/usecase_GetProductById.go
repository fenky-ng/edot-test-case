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

	// product warehouses
	warehousesByProductId, err := u.getProductWarehouses(ctx, []model.Product{
		productOut.Product,
	})
	if err != nil {
		return output, err
	}

	warehouses := warehousesByProductId[productOut.Product.Shop.Id]
	for _, warehouse := range warehouses {
		if warehouse.Status == constant.ShopWarehouseStatus_Inactive { // only show stock from active warehouse
			continue
		}
		productOut.Product.Stock.Total += warehouse.Stock
		productOut.Product.Stock.Warehouses = append(productOut.Product.Stock.Warehouses, warehouse)
	}

	output.Product = productOut.Product

	return output, nil
}
