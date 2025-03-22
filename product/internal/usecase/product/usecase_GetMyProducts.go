package product

import (
	"context"

	"github.com/fenky-ng/edot-test-case/product/internal/model"
)

func (u *ProductUsecase) GetMyProducts(ctx context.Context, input model.GetMyProductsInput) (output model.GetMyProductsOutput, err error) {
	shopOut, err := u.repoHttpShop.GetMyShop(ctx, model.GetMyShopInput{
		JWT: input.JWT,
	})
	if err != nil {
		return output, err
	}

	productsOut, err := u.repoDbProduct.GetProducts(ctx, model.GetProductsInput{
		ShopId: shopOut.Id,
	})
	if err != nil {
		return output, err
	}

	warehousesByProductId, err := u.getProductWarehouses(ctx, output.Products)
	if err != nil {
		return output, err
	}

	for index := range productsOut.Products {
		// shop
		productsOut.Products[index].Shop.Name = shopOut.Name
		productsOut.Products[index].Shop.Status = shopOut.Status

		// warehouse
		warehouses := warehousesByProductId[output.Products[index].Id]
		for _, warehouse := range warehouses {
			// show stock from all warehouse
			output.Products[index].Stock.Total += warehouse.Stock
			output.Products[index].Stock.Warehouses = append(output.Products[index].Stock.Warehouses, warehouse)
		}
	}

	output.Products = productsOut.Products

	return output, nil
}
