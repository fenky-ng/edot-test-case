package product

import (
	"context"
	"fmt"

	"github.com/fenky-ng/edot-test-case/product/internal/constant"
	"github.com/fenky-ng/edot-test-case/product/internal/model"
	"github.com/google/uuid"
)

func (u *ProductUsecase) GetProducts(ctx context.Context, input model.GetProductsInput) (output model.GetProductsOutput, err error) {
	output, err = u.repoDbProduct.GetProducts(ctx, input)
	if err != nil {
		return output, err
	}

	shopById, err := u.getShops(ctx, output.Products)
	if err != nil {
		return output, err
	}

	warehousesByProductId, err := u.getProductWarehouses(ctx, output.Products)
	if err != nil {
		return output, err
	}

	for index := range output.Products {
		// shop
		shop := shopById[output.Products[index].Shop.Id]
		output.Products[index].Shop.Name = shop.Name
		output.Products[index].Shop.Status = shop.Status

		// warehouse
		warehouses := warehousesByProductId[output.Products[index].Id]
		for _, warehouse := range warehouses {
			if warehouse.Status == constant.ShopWarehouseStatus_Inactive { // only show stock from active warehouse
				continue
			}
			output.Products[index].Stock.Total += warehouse.Stock
			output.Products[index].Stock.Warehouses = append(output.Products[index].Stock.Warehouses, warehouse)
		}
	}

	return output, err
}

func (u *ProductUsecase) getShops(
	ctx context.Context,
	products []model.Product,
) (shopById map[uuid.UUID]model.Shop, err error) {
	shopById = make(map[uuid.UUID]model.Shop)

	shopIds := []uuid.UUID{}
	uniqueShopId := make(map[uuid.UUID]struct{})
	for _, product := range products {
		if _, exists := uniqueShopId[product.Shop.Id]; exists {
			continue
		}
		shopIds = append(shopIds, product.Shop.Id)
		uniqueShopId[product.Shop.Id] = struct{}{}
	}

	if len(shopIds) == 0 {
		return
	}

	shopsOut, err := u.repoHttpShop.GetShops(ctx, model.GetShopsInput{
		Ids: shopIds,
	})
	if err != nil {
		return
	}

	for _, shop := range shopsOut.Shops {
		shopById[shop.Id] = model.Shop{
			Id:     shop.Id,
			Name:   shop.Name,
			Status: shop.Status,
		}
	}

	return
}

func (u *ProductUsecase) getProductWarehouses(
	ctx context.Context,
	products []model.Product,
) (warehousesByProductId map[uuid.UUID][]model.ProductWarehouse, err error) {
	warehousesByProductId = make(map[uuid.UUID][]model.ProductWarehouse)

	productIds := []uuid.UUID{}
	uniqueProductId := make(map[uuid.UUID]struct{})
	for _, product := range products {
		if _, exists := uniqueProductId[product.Shop.Id]; exists {
			continue
		}
		productIds = append(productIds, product.Shop.Id)
		uniqueProductId[product.Shop.Id] = struct{}{}
	}

	if len(productIds) == 0 {
		return
	}

	productWarehousesOut, err := u.repoHttpWarehouse.GetProductWarehouses(ctx, model.GetProductWarehousesInput{
		ProductIds: productIds,
	})
	if err != nil {
		return
	}

	// TODO
	fmt.Println(productWarehousesOut)

	// for _, shop := range shopsOut.Shops {
	// 	shopById[shop.Id] = model.Shop{
	// 		Id:     shop.Id,
	// 		Name:   shop.Name,
	// 		Status: shop.Status,
	// 	}
	// }

	return
}
