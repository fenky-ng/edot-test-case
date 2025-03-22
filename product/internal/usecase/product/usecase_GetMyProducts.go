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

	stocksByProductId, err := u.getStocks(ctx, productsOut.Products)
	if err != nil {
		return output, err
	}

	for index := range productsOut.Products {
		// shop
		productsOut.Products[index].Shop.Name = shopOut.Name
		productsOut.Products[index].Shop.Status = shopOut.Status

		// stocks
		stocks := stocksByProductId[productsOut.Products[index].Id]
		for _, stock := range stocks {
			// show stock from all warehouse
			productsOut.Products[index].Stock.Total += stock.Stock
			productsOut.Products[index].Stock.Warehouses = append(productsOut.Products[index].Stock.Warehouses, stock)
		}
	}

	output.Products = productsOut.Products

	return output, nil
}
