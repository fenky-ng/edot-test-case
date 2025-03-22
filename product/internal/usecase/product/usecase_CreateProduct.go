package product

import (
	"context"
	"time"

	"github.com/fenky-ng/edot-test-case/product/internal/model"
	"github.com/google/uuid"
)

func (u *ProductUsecase) CreateProduct(ctx context.Context, input model.CreateProductInput) (output model.CreateProductOutput, err error) {
	shopOut, err := u.repoHttpShop.GetMyShop(ctx, model.GetMyShopInput{
		JWT: input.JWT,
	})
	if err != nil {
		return output, err
	}

	insertOut, err := u.repoDbProduct.InsertProduct(ctx, model.InsertProductInput{
		Id:          uuid.New(),
		ShopId:      shopOut.Id,
		Name:        input.Name,
		Description: input.Description,
		Price:       input.Price,
		Status:      input.Status,
		CreatedAt:   time.Now().UnixMilli(),
		CreatedBy:   input.UserId.String(),
	})
	if err != nil {
		return output, err
	}

	output.Id = insertOut.Id

	return output, nil
}
