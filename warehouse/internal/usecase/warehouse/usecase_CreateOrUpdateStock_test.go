package warehouse_test

import (
	"context"
	"errors"
	"testing"

	in_err "github.com/fenky-ng/edot-test-case/warehouse/internal/error"
	"github.com/fenky-ng/edot-test-case/warehouse/internal/model"
	db_warehouse "github.com/fenky-ng/edot-test-case/warehouse/internal/repository/db/warehouse"
	http_product "github.com/fenky-ng/edot-test-case/warehouse/internal/repository/http/product"
	http_shop "github.com/fenky-ng/edot-test-case/warehouse/internal/repository/http/shop"
	"github.com/fenky-ng/edot-test-case/warehouse/internal/usecase/warehouse"
	test_util "github.com/fenky-ng/edot-test-case/warehouse/internal/utility/test"
	"github.com/go-faker/faker/v4"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
)

func Test_WarehouseUsecase_CreateOrUpdateStock(t *testing.T) {
	type fields struct {
		RepoDbWarehouse  db_warehouse.RepoDbWarehouseInterface
		WarehouseUsecase warehouse.WarehouseUsecaseInterface
	}
	type args struct {
		ctx   context.Context
		input model.CreateOrUpdateStockInput
	}
	type test struct {
		name        string
		fields      fields
		args        args
		expectedRes model.CreateOrUpdateStockOutput
		assertErr   require.ErrorAssertionFunc
		mock        func(tt *test)
	}

	var (
		ctx      = context.Background()
		setInput = model.CreateOrUpdateStockInput{
			JWT:         faker.Jwt(),
			UserId:      uuid.MustParse(faker.UUIDHyphenated()),
			WarehouseId: uuid.MustParse(faker.UUIDHyphenated()),
			ProductId:   uuid.MustParse(faker.UUIDHyphenated()),
			Stock:       8,
		}
		transferInput = model.CreateOrUpdateStockInput{
			JWT:           faker.Jwt(),
			UserId:        uuid.MustParse(faker.UUIDHyphenated()),
			WarehouseId:   uuid.MustParse(faker.UUIDHyphenated()),
			ProductId:     uuid.MustParse(faker.UUIDHyphenated()),
			Stock:         8,
			ToWarehouseId: uuid.MustParse(faker.UUIDHyphenated()),
		}
	)
	tests := []test{
		{
			name: "should return error if request is invalid or error occurred when validating request",
			args: args{
				ctx:   ctx,
				input: setInput,
			},
			expectedRes: model.CreateOrUpdateStockOutput{},
			assertErr:   test_util.RequireErrorIs(in_err.ErrNotProductOwner),
			mock: func(tt *test) {
				mockCtrl := gomock.NewController(t)

				repoDbWarehouseMock := db_warehouse.NewMockRepoDbWarehouseInterface(mockCtrl)
				warehouseUsecaseMock := warehouse.NewMockWarehouseUsecaseInterface(mockCtrl)

				tt.fields.RepoDbWarehouse = repoDbWarehouseMock
				tt.fields.WarehouseUsecase = warehouseUsecaseMock

				warehouseUsecaseMock.EXPECT().ValidateCreateOrUpdateStock(
					test_util.ContextTypeMatcher,
					tt.args.input,
				).Return(
					in_err.ErrNotProductOwner,
				).Times(1)
			},
		},
		{
			name: "should return error if error occurred in repoDbWarehouse.Begin",
			args: args{
				ctx:   ctx,
				input: setInput,
			},
			expectedRes: model.CreateOrUpdateStockOutput{},
			assertErr:   test_util.RequireErrorIs(in_err.ErrDatabaseTransaction),
			mock: func(tt *test) {
				mockCtrl := gomock.NewController(t)

				repoDbWarehouseMock := db_warehouse.NewMockRepoDbWarehouseInterface(mockCtrl)
				warehouseUsecaseMock := warehouse.NewMockWarehouseUsecaseInterface(mockCtrl)

				tt.fields.RepoDbWarehouse = repoDbWarehouseMock
				tt.fields.WarehouseUsecase = warehouseUsecaseMock

				warehouseUsecaseMock.EXPECT().ValidateCreateOrUpdateStock(
					test_util.ContextTypeMatcher,
					tt.args.input,
				).Return(
					nil,
				).Times(1)

				repoDbWarehouseMock.EXPECT().Begin(
					test_util.ContextTypeMatcher,
					nil,
				).Return(
					tt.args.ctx,
					errors.New("expected Begin error"),
				).Times(1)
			},
		},
		{
			name: "should return error if error occurred in repoDbWarehouse.UpsertStock",
			args: args{
				ctx:   ctx,
				input: setInput,
			},
			expectedRes: model.CreateOrUpdateStockOutput{},
			assertErr:   test_util.RequireErrorIs(in_err.ErrUpsertStock),
			mock: func(tt *test) {
				mockCtrl := gomock.NewController(t)

				repoDbWarehouseMock := db_warehouse.NewMockRepoDbWarehouseInterface(mockCtrl)
				warehouseUsecaseMock := warehouse.NewMockWarehouseUsecaseInterface(mockCtrl)

				tt.fields.RepoDbWarehouse = repoDbWarehouseMock
				tt.fields.WarehouseUsecase = warehouseUsecaseMock

				warehouseUsecaseMock.EXPECT().ValidateCreateOrUpdateStock(
					test_util.ContextTypeMatcher,
					tt.args.input,
				).Return(
					nil,
				).Times(1)

				repoDbWarehouseMock.EXPECT().Begin(
					test_util.ContextTypeMatcher,
					nil,
				).Return(
					tt.args.ctx,
					nil,
				).Times(1)

				errUpsertStock := errors.Join(in_err.ErrUpsertStock, errors.New("expected UpsertStock error"))

				repoDbWarehouseMock.EXPECT().CommitOrRollback(
					test_util.ContextTypeMatcher,
					errUpsertStock,
				).Return(
					errUpsertStock,
				).Times(1)

				repoDbWarehouseMock.EXPECT().UpsertStock(
					test_util.ContextTypeMatcher,
					model.UpsertStockInput{
						WarehouseId: tt.args.input.WarehouseId,
						ProductId:   tt.args.input.ProductId,
						Stock:       tt.args.input.Stock,
						IsTransfer:  false,
						UpsertedBy:  tt.args.input.UserId.String(),
					}.Matcher(),
				).Return(
					model.UpsertStockOutput{},
					errUpsertStock,
				).Times(1)
			},
		},
		{
			name: "should return response successful if set stock was successful",
			args: args{
				ctx:   ctx,
				input: setInput,
			},
			expectedRes: model.CreateOrUpdateStockOutput{
				Successful: true,
			},
			assertErr: require.NoError,
			mock: func(tt *test) {
				mockCtrl := gomock.NewController(t)

				repoDbWarehouseMock := db_warehouse.NewMockRepoDbWarehouseInterface(mockCtrl)
				warehouseUsecaseMock := warehouse.NewMockWarehouseUsecaseInterface(mockCtrl)

				tt.fields.RepoDbWarehouse = repoDbWarehouseMock
				tt.fields.WarehouseUsecase = warehouseUsecaseMock

				warehouseUsecaseMock.EXPECT().ValidateCreateOrUpdateStock(
					test_util.ContextTypeMatcher,
					tt.args.input,
				).Return(
					nil,
				).Times(1)

				repoDbWarehouseMock.EXPECT().Begin(
					test_util.ContextTypeMatcher,
					nil,
				).Return(
					tt.args.ctx,
					nil,
				).Times(1)

				repoDbWarehouseMock.EXPECT().CommitOrRollback(
					test_util.ContextTypeMatcher,
					nil,
				).Return(
					nil,
				).Times(1)

				repoDbWarehouseMock.EXPECT().UpsertStock(
					test_util.ContextTypeMatcher,
					model.UpsertStockInput{
						WarehouseId: tt.args.input.WarehouseId,
						ProductId:   tt.args.input.ProductId,
						Stock:       tt.args.input.Stock,
						IsTransfer:  false,
						UpsertedBy:  tt.args.input.UserId.String(),
					}.Matcher(),
				).Return(
					model.UpsertStockOutput{
						Id: uuid.MustParse(faker.UUIDHyphenated()),
					},
					nil,
				).Times(1)
			},
		},
		{
			name: "should return error if error occurred in WarehouseUsecase.TransferStock",
			args: args{
				ctx:   ctx,
				input: transferInput,
			},
			expectedRes: model.CreateOrUpdateStockOutput{
				Successful: false,
			},
			assertErr: test_util.RequireErrorIs(in_err.ErrDeductStock),
			mock: func(tt *test) {
				mockCtrl := gomock.NewController(t)

				repoDbWarehouseMock := db_warehouse.NewMockRepoDbWarehouseInterface(mockCtrl)
				warehouseUsecaseMock := warehouse.NewMockWarehouseUsecaseInterface(mockCtrl)

				tt.fields.RepoDbWarehouse = repoDbWarehouseMock
				tt.fields.WarehouseUsecase = warehouseUsecaseMock

				warehouseUsecaseMock.EXPECT().ValidateCreateOrUpdateStock(
					test_util.ContextTypeMatcher,
					tt.args.input,
				).Return(
					nil,
				).Times(1)

				repoDbWarehouseMock.EXPECT().Begin(
					test_util.ContextTypeMatcher,
					nil,
				).Return(
					tt.args.ctx,
					nil,
				).Times(1)

				errTransferStock := errors.Join(in_err.ErrDeductStock, errors.New("expected TransferStock.DeductStock error"))

				repoDbWarehouseMock.EXPECT().CommitOrRollback(
					test_util.ContextTypeMatcher,
					errTransferStock,
				).Return(
					errTransferStock,
				).Times(1)

				warehouseUsecaseMock.EXPECT().TransferStock(
					test_util.ContextTypeMatcher,
					tt.args.input,
				).Return(
					errTransferStock,
				).Times(1)
			},
		},
		{
			name: "should return response successful if transfer stock was successful",
			args: args{
				ctx:   ctx,
				input: transferInput,
			},
			expectedRes: model.CreateOrUpdateStockOutput{
				Successful: true,
			},
			assertErr: require.NoError,
			mock: func(tt *test) {
				mockCtrl := gomock.NewController(t)

				repoDbWarehouseMock := db_warehouse.NewMockRepoDbWarehouseInterface(mockCtrl)
				warehouseUsecaseMock := warehouse.NewMockWarehouseUsecaseInterface(mockCtrl)

				tt.fields.RepoDbWarehouse = repoDbWarehouseMock
				tt.fields.WarehouseUsecase = warehouseUsecaseMock

				warehouseUsecaseMock.EXPECT().ValidateCreateOrUpdateStock(
					test_util.ContextTypeMatcher,
					tt.args.input,
				).Return(
					nil,
				).Times(1)

				repoDbWarehouseMock.EXPECT().Begin(
					test_util.ContextTypeMatcher,
					nil,
				).Return(
					tt.args.ctx,
					nil,
				).Times(1)

				repoDbWarehouseMock.EXPECT().CommitOrRollback(
					test_util.ContextTypeMatcher,
					nil,
				).Return(
					nil,
				).Times(1)

				warehouseUsecaseMock.EXPECT().TransferStock(
					test_util.ContextTypeMatcher,
					tt.args.input,
				).Return(
					nil,
				).Times(1)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.mock != nil {
				tt.mock(&tt)
			}

			u := warehouse.InitWarehouseUsecase(warehouse.InitWarehouseUsecaseOptions{
				RepoDbWarehouse: tt.fields.RepoDbWarehouse,
			})
			u.WarehouseUsecase = tt.fields.WarehouseUsecase
			gotRes, gotErr := u.CreateOrUpdateStock(tt.args.ctx, tt.args.input)
			tt.assertErr(t, gotErr)
			require.Equal(t, tt.expectedRes, gotRes)
		})
	}
}

func Test_WarehouseUsecase_ValidateCreateOrUpdateStock(t *testing.T) {
	type fields struct {
		RepoDbWarehouse db_warehouse.RepoDbWarehouseInterface
		RepoHttpShop    http_shop.RepoHttpShopInterface
		RepoHttpProduct http_product.RepoHttpProductInterface
	}
	type args struct {
		ctx   context.Context
		input model.CreateOrUpdateStockInput
	}
	type test struct {
		name      string
		fields    fields
		args      args
		assertErr require.ErrorAssertionFunc
		mock      func(tt *test)
	}

	var (
		ctx   = context.Background()
		input = model.CreateOrUpdateStockInput{
			JWT:           faker.Jwt(),
			UserId:        uuid.MustParse(faker.UUIDHyphenated()),
			WarehouseId:   uuid.MustParse(faker.UUIDHyphenated()),
			ProductId:     uuid.MustParse(faker.UUIDHyphenated()),
			Stock:         8,
			ToWarehouseId: uuid.MustParse(faker.UUIDHyphenated()),
		}
		selftTransferWarehouseId   = uuid.MustParse(faker.UUIDHyphenated())
		selfTransferWarehouseInput = model.CreateOrUpdateStockInput{
			JWT:           faker.Jwt(),
			UserId:        uuid.MustParse(faker.UUIDHyphenated()),
			WarehouseId:   selftTransferWarehouseId,
			ProductId:     uuid.MustParse(faker.UUIDHyphenated()),
			Stock:         8,
			ToWarehouseId: selftTransferWarehouseId,
		}

		userShopId        = uuid.MustParse(faker.UUIDHyphenated())
		anotherUserShopId = uuid.MustParse(faker.UUIDHyphenated())
	)

	tests := []test{
		{
			name: "should return error if error occurred in repoHttpShop.GetMyShop",
			args: args{
				ctx:   ctx,
				input: input,
			},
			assertErr: test_util.RequireErrorIs(in_err.ErrGetMyShop),
			mock: func(tt *test) {
				mockCtrl := gomock.NewController(t)

				repoDbWarehouse := db_warehouse.NewMockRepoDbWarehouseInterface(mockCtrl)
				repoHttpShopMock := http_shop.NewMockRepoHttpShopInterface(mockCtrl)
				repoHttpProductMock := http_product.NewMockRepoHttpProductInterface(mockCtrl)

				tt.fields.RepoDbWarehouse = repoDbWarehouse
				tt.fields.RepoHttpShop = repoHttpShopMock
				tt.fields.RepoHttpProduct = repoHttpProductMock

				repoHttpShopMock.EXPECT().GetMyShop(
					test_util.ContextTypeMatcher,
					model.GetMyShopInput{
						JWT: tt.args.input.JWT,
					},
				).Return(
					model.GetMyShopOutput{},
					errors.Join(in_err.ErrGetMyShop, errors.New("expected GetMyShop error")),
				).Times(1)
			},
		},
		{
			name: "should return error if error occurred in repoDbWarehouse.GetWarehouses",
			args: args{
				ctx:   ctx,
				input: input,
			},
			assertErr: test_util.RequireErrorIs(in_err.ErrGetWarehouses),
			mock: func(tt *test) {
				mockCtrl := gomock.NewController(t)

				repoDbWarehouse := db_warehouse.NewMockRepoDbWarehouseInterface(mockCtrl)
				repoHttpShopMock := http_shop.NewMockRepoHttpShopInterface(mockCtrl)
				repoHttpProductMock := http_product.NewMockRepoHttpProductInterface(mockCtrl)

				tt.fields.RepoDbWarehouse = repoDbWarehouse
				tt.fields.RepoHttpShop = repoHttpShopMock
				tt.fields.RepoHttpProduct = repoHttpProductMock

				repoHttpShopMock.EXPECT().GetMyShop(
					test_util.ContextTypeMatcher,
					model.GetMyShopInput{
						JWT: tt.args.input.JWT,
					},
				).Return(
					model.GetMyShopOutput{
						ExtShop: model.ExtShop{
							Id: userShopId,
						},
					},
					nil,
				).Times(1)

				repoDbWarehouse.EXPECT().GetWarehouses(
					test_util.ContextTypeMatcher,
					model.GetWarehousesInput{
						Ids: []uuid.UUID{
							tt.args.input.WarehouseId,
							tt.args.input.ToWarehouseId,
						},
					},
				).Return(
					model.GetWarehousesOutput{},
					errors.Join(in_err.ErrGetWarehouses, errors.New("expected GetWarehouses error")),
				).Times(1)

			},
		},
		{
			name: "should return error if warehouse not found",
			args: args{
				ctx:   ctx,
				input: input,
			},
			assertErr: test_util.RequireErrorIs(in_err.ErrWarehouseNotFound),
			mock: func(tt *test) {
				mockCtrl := gomock.NewController(t)

				repoDbWarehouse := db_warehouse.NewMockRepoDbWarehouseInterface(mockCtrl)
				repoHttpShopMock := http_shop.NewMockRepoHttpShopInterface(mockCtrl)
				repoHttpProductMock := http_product.NewMockRepoHttpProductInterface(mockCtrl)

				tt.fields.RepoDbWarehouse = repoDbWarehouse
				tt.fields.RepoHttpShop = repoHttpShopMock
				tt.fields.RepoHttpProduct = repoHttpProductMock

				repoHttpShopMock.EXPECT().GetMyShop(
					test_util.ContextTypeMatcher,
					model.GetMyShopInput{
						JWT: tt.args.input.JWT,
					},
				).Return(
					model.GetMyShopOutput{
						ExtShop: model.ExtShop{
							Id: userShopId,
						},
					},
					nil,
				).Times(1)

				repoDbWarehouse.EXPECT().GetWarehouses(
					test_util.ContextTypeMatcher,
					model.GetWarehousesInput{
						Ids: []uuid.UUID{
							tt.args.input.WarehouseId,
							tt.args.input.ToWarehouseId,
						},
					},
				).Return(
					model.GetWarehousesOutput{
						Warehouses: []model.Warehouse{
							{
								Id:     tt.args.input.ToWarehouseId,
								ShopId: userShopId,
							},
						},
						WarehouseById: map[uuid.UUID]model.Warehouse{
							tt.args.input.ToWarehouseId: {
								Id:     tt.args.input.ToWarehouseId,
								ShopId: userShopId,
							},
						},
					},
					nil,
				).Times(1)

			},
		},
		{
			name: "should return error if not warehouse owner",
			args: args{
				ctx:   ctx,
				input: input,
			},
			assertErr: test_util.RequireErrorIs(in_err.ErrNotWarehouseOwner),
			mock: func(tt *test) {
				mockCtrl := gomock.NewController(t)

				repoDbWarehouse := db_warehouse.NewMockRepoDbWarehouseInterface(mockCtrl)
				repoHttpShopMock := http_shop.NewMockRepoHttpShopInterface(mockCtrl)
				repoHttpProductMock := http_product.NewMockRepoHttpProductInterface(mockCtrl)

				tt.fields.RepoDbWarehouse = repoDbWarehouse
				tt.fields.RepoHttpShop = repoHttpShopMock
				tt.fields.RepoHttpProduct = repoHttpProductMock

				repoHttpShopMock.EXPECT().GetMyShop(
					test_util.ContextTypeMatcher,
					model.GetMyShopInput{
						JWT: tt.args.input.JWT,
					},
				).Return(
					model.GetMyShopOutput{
						ExtShop: model.ExtShop{
							Id: userShopId,
						},
					},
					nil,
				).Times(1)

				repoDbWarehouse.EXPECT().GetWarehouses(
					test_util.ContextTypeMatcher,
					model.GetWarehousesInput{
						Ids: []uuid.UUID{
							tt.args.input.WarehouseId,
							tt.args.input.ToWarehouseId,
						},
					},
				).Return(
					model.GetWarehousesOutput{
						Warehouses: []model.Warehouse{
							{
								Id:     tt.args.input.WarehouseId,
								ShopId: anotherUserShopId,
							},
							{
								Id:     tt.args.input.ToWarehouseId,
								ShopId: userShopId,
							},
						},
						WarehouseById: map[uuid.UUID]model.Warehouse{
							tt.args.input.WarehouseId: {
								Id:     tt.args.input.WarehouseId,
								ShopId: anotherUserShopId,
							},
							tt.args.input.ToWarehouseId: {
								Id:     tt.args.input.ToWarehouseId,
								ShopId: userShopId,
							},
						},
					},
					nil,
				).Times(1)
			},
		},
		{
			name: "should return error if transfer destination warehouse not found",
			args: args{
				ctx:   ctx,
				input: input,
			},
			assertErr: test_util.RequireErrorIs(in_err.ErrWarehouseNotFound),
			mock: func(tt *test) {
				mockCtrl := gomock.NewController(t)

				repoDbWarehouse := db_warehouse.NewMockRepoDbWarehouseInterface(mockCtrl)
				repoHttpShopMock := http_shop.NewMockRepoHttpShopInterface(mockCtrl)
				repoHttpProductMock := http_product.NewMockRepoHttpProductInterface(mockCtrl)

				tt.fields.RepoDbWarehouse = repoDbWarehouse
				tt.fields.RepoHttpShop = repoHttpShopMock
				tt.fields.RepoHttpProduct = repoHttpProductMock

				repoHttpShopMock.EXPECT().GetMyShop(
					test_util.ContextTypeMatcher,
					model.GetMyShopInput{
						JWT: tt.args.input.JWT,
					},
				).Return(
					model.GetMyShopOutput{
						ExtShop: model.ExtShop{
							Id: userShopId,
						},
					},
					nil,
				).Times(1)

				repoDbWarehouse.EXPECT().GetWarehouses(
					test_util.ContextTypeMatcher,
					model.GetWarehousesInput{
						Ids: []uuid.UUID{
							tt.args.input.WarehouseId,
							tt.args.input.ToWarehouseId,
						},
					},
				).Return(
					model.GetWarehousesOutput{
						Warehouses: []model.Warehouse{
							{
								Id:     tt.args.input.WarehouseId,
								ShopId: userShopId,
							},
						},
						WarehouseById: map[uuid.UUID]model.Warehouse{
							tt.args.input.WarehouseId: {
								Id:     tt.args.input.WarehouseId,
								ShopId: userShopId,
							},
						},
					},
					nil,
				).Times(1)

			},
		},
		{
			name: "should return error if not transfer destination warehouse owner",
			args: args{
				ctx:   ctx,
				input: input,
			},
			assertErr: test_util.RequireErrorIs(in_err.ErrNotWarehouseOwner),
			mock: func(tt *test) {
				mockCtrl := gomock.NewController(t)

				repoDbWarehouse := db_warehouse.NewMockRepoDbWarehouseInterface(mockCtrl)
				repoHttpShopMock := http_shop.NewMockRepoHttpShopInterface(mockCtrl)
				repoHttpProductMock := http_product.NewMockRepoHttpProductInterface(mockCtrl)

				tt.fields.RepoDbWarehouse = repoDbWarehouse
				tt.fields.RepoHttpShop = repoHttpShopMock
				tt.fields.RepoHttpProduct = repoHttpProductMock

				repoHttpShopMock.EXPECT().GetMyShop(
					test_util.ContextTypeMatcher,
					model.GetMyShopInput{
						JWT: tt.args.input.JWT,
					},
				).Return(
					model.GetMyShopOutput{
						ExtShop: model.ExtShop{
							Id: userShopId,
						},
					},
					nil,
				).Times(1)

				repoDbWarehouse.EXPECT().GetWarehouses(
					test_util.ContextTypeMatcher,
					model.GetWarehousesInput{
						Ids: []uuid.UUID{
							tt.args.input.WarehouseId,
							tt.args.input.ToWarehouseId,
						},
					},
				).Return(
					model.GetWarehousesOutput{
						Warehouses: []model.Warehouse{
							{
								Id:     tt.args.input.WarehouseId,
								ShopId: userShopId,
							},
							{
								Id:     tt.args.input.ToWarehouseId,
								ShopId: anotherUserShopId,
							},
						},
						WarehouseById: map[uuid.UUID]model.Warehouse{
							tt.args.input.WarehouseId: {
								Id:     tt.args.input.WarehouseId,
								ShopId: userShopId,
							},
							tt.args.input.ToWarehouseId: {
								Id:     tt.args.input.ToWarehouseId,
								ShopId: anotherUserShopId,
							},
						},
					},
					nil,
				).Times(1)
			},
		},
		{
			name: "should return error if transfer destination warehouse is the same as source warehouse",
			args: args{
				ctx:   ctx,
				input: selfTransferWarehouseInput,
			},
			assertErr: test_util.RequireErrorIs(in_err.ErrInvalidStockTransferDestination),
			mock: func(tt *test) {
				mockCtrl := gomock.NewController(t)

				repoDbWarehouse := db_warehouse.NewMockRepoDbWarehouseInterface(mockCtrl)
				repoHttpShopMock := http_shop.NewMockRepoHttpShopInterface(mockCtrl)
				repoHttpProductMock := http_product.NewMockRepoHttpProductInterface(mockCtrl)

				tt.fields.RepoDbWarehouse = repoDbWarehouse
				tt.fields.RepoHttpShop = repoHttpShopMock
				tt.fields.RepoHttpProduct = repoHttpProductMock

				repoHttpShopMock.EXPECT().GetMyShop(
					test_util.ContextTypeMatcher,
					model.GetMyShopInput{
						JWT: tt.args.input.JWT,
					},
				).Return(
					model.GetMyShopOutput{
						ExtShop: model.ExtShop{
							Id: userShopId,
						},
					},
					nil,
				).Times(1)

				repoDbWarehouse.EXPECT().GetWarehouses(
					test_util.ContextTypeMatcher,
					model.GetWarehousesInput{
						Ids: []uuid.UUID{
							tt.args.input.WarehouseId,
							tt.args.input.ToWarehouseId,
						},
					},
				).Return(
					model.GetWarehousesOutput{
						Warehouses: []model.Warehouse{
							{
								Id:     tt.args.input.WarehouseId,
								ShopId: userShopId,
							},
						},
						WarehouseById: map[uuid.UUID]model.Warehouse{
							tt.args.input.WarehouseId: {
								Id:     tt.args.input.WarehouseId,
								ShopId: userShopId,
							},
						},
					},
					nil,
				).Times(1)
			},
		},
		{
			name: "should return error if error occurred in repoHttpProduct.GetProductById",
			args: args{
				ctx:   ctx,
				input: input,
			},
			assertErr: test_util.RequireErrorIs(in_err.ErrGetProductById),
			mock: func(tt *test) {
				mockCtrl := gomock.NewController(t)

				repoDbWarehouse := db_warehouse.NewMockRepoDbWarehouseInterface(mockCtrl)
				repoHttpShopMock := http_shop.NewMockRepoHttpShopInterface(mockCtrl)
				repoHttpProductMock := http_product.NewMockRepoHttpProductInterface(mockCtrl)

				tt.fields.RepoDbWarehouse = repoDbWarehouse
				tt.fields.RepoHttpShop = repoHttpShopMock
				tt.fields.RepoHttpProduct = repoHttpProductMock

				repoHttpShopMock.EXPECT().GetMyShop(
					test_util.ContextTypeMatcher,
					model.GetMyShopInput{
						JWT: tt.args.input.JWT,
					},
				).Return(
					model.GetMyShopOutput{
						ExtShop: model.ExtShop{
							Id: userShopId,
						},
					},
					nil,
				).Times(1)

				repoDbWarehouse.EXPECT().GetWarehouses(
					test_util.ContextTypeMatcher,
					model.GetWarehousesInput{
						Ids: []uuid.UUID{
							tt.args.input.WarehouseId,
							tt.args.input.ToWarehouseId,
						},
					},
				).Return(
					model.GetWarehousesOutput{
						Warehouses: []model.Warehouse{
							{
								Id:     tt.args.input.WarehouseId,
								ShopId: userShopId,
							},
							{
								Id:     tt.args.input.ToWarehouseId,
								ShopId: userShopId,
							},
						},
						WarehouseById: map[uuid.UUID]model.Warehouse{
							tt.args.input.WarehouseId: {
								Id:     tt.args.input.WarehouseId,
								ShopId: userShopId,
							},
							tt.args.input.ToWarehouseId: {
								Id:     tt.args.input.ToWarehouseId,
								ShopId: userShopId,
							},
						},
					},
					nil,
				).Times(1)

				repoHttpProductMock.EXPECT().GetProductById(
					test_util.ContextTypeMatcher,
					model.GetProductByIdInput{
						Id: tt.args.input.ProductId,
					},
				).Return(
					model.GetProductByIdOutput{},
					errors.Join(in_err.ErrGetProductById, errors.New("expected GetProductById error")),
				).Times(1)
			},
		},
		{
			name: "should return error if not product owner",
			args: args{
				ctx:   ctx,
				input: input,
			},
			assertErr: test_util.RequireErrorIs(in_err.ErrNotProductOwner),
			mock: func(tt *test) {
				mockCtrl := gomock.NewController(t)

				repoDbWarehouse := db_warehouse.NewMockRepoDbWarehouseInterface(mockCtrl)
				repoHttpShopMock := http_shop.NewMockRepoHttpShopInterface(mockCtrl)
				repoHttpProductMock := http_product.NewMockRepoHttpProductInterface(mockCtrl)

				tt.fields.RepoDbWarehouse = repoDbWarehouse
				tt.fields.RepoHttpShop = repoHttpShopMock
				tt.fields.RepoHttpProduct = repoHttpProductMock

				repoHttpShopMock.EXPECT().GetMyShop(
					test_util.ContextTypeMatcher,
					model.GetMyShopInput{
						JWT: tt.args.input.JWT,
					},
				).Return(
					model.GetMyShopOutput{
						ExtShop: model.ExtShop{
							Id: userShopId,
						},
					},
					nil,
				).Times(1)

				repoDbWarehouse.EXPECT().GetWarehouses(
					test_util.ContextTypeMatcher,
					model.GetWarehousesInput{
						Ids: []uuid.UUID{
							tt.args.input.WarehouseId,
							tt.args.input.ToWarehouseId,
						},
					},
				).Return(
					model.GetWarehousesOutput{
						Warehouses: []model.Warehouse{
							{
								Id:     tt.args.input.WarehouseId,
								ShopId: userShopId,
							},
							{
								Id:     tt.args.input.ToWarehouseId,
								ShopId: userShopId,
							},
						},
						WarehouseById: map[uuid.UUID]model.Warehouse{
							tt.args.input.WarehouseId: {
								Id:     tt.args.input.WarehouseId,
								ShopId: userShopId,
							},
							tt.args.input.ToWarehouseId: {
								Id:     tt.args.input.ToWarehouseId,
								ShopId: userShopId,
							},
						},
					},
					nil,
				).Times(1)

				repoHttpProductMock.EXPECT().GetProductById(
					test_util.ContextTypeMatcher,
					model.GetProductByIdInput{
						Id: tt.args.input.ProductId,
					},
				).Return(
					model.GetProductByIdOutput{
						ExtProduct: model.ExtProduct{
							Id: tt.args.input.ProductId,
							Shop: model.ExtShop{
								Id: anotherUserShopId,
							},
						},
					},
					nil,
				).Times(1)
			},
		},
		{
			name: "should return no error if request is valid",
			args: args{
				ctx:   ctx,
				input: input,
			},
			assertErr: require.NoError,
			mock: func(tt *test) {
				mockCtrl := gomock.NewController(t)

				repoDbWarehouse := db_warehouse.NewMockRepoDbWarehouseInterface(mockCtrl)
				repoHttpShopMock := http_shop.NewMockRepoHttpShopInterface(mockCtrl)
				repoHttpProductMock := http_product.NewMockRepoHttpProductInterface(mockCtrl)

				tt.fields.RepoDbWarehouse = repoDbWarehouse
				tt.fields.RepoHttpShop = repoHttpShopMock
				tt.fields.RepoHttpProduct = repoHttpProductMock

				repoHttpShopMock.EXPECT().GetMyShop(
					test_util.ContextTypeMatcher,
					model.GetMyShopInput{
						JWT: tt.args.input.JWT,
					},
				).Return(
					model.GetMyShopOutput{
						ExtShop: model.ExtShop{
							Id: userShopId,
						},
					},
					nil,
				).Times(1)

				repoDbWarehouse.EXPECT().GetWarehouses(
					test_util.ContextTypeMatcher,
					model.GetWarehousesInput{
						Ids: []uuid.UUID{
							tt.args.input.WarehouseId,
							tt.args.input.ToWarehouseId,
						},
					},
				).Return(
					model.GetWarehousesOutput{
						Warehouses: []model.Warehouse{
							{
								Id:     tt.args.input.WarehouseId,
								ShopId: userShopId,
							},
							{
								Id:     tt.args.input.ToWarehouseId,
								ShopId: userShopId,
							},
						},
						WarehouseById: map[uuid.UUID]model.Warehouse{
							tt.args.input.WarehouseId: {
								Id:     tt.args.input.WarehouseId,
								ShopId: userShopId,
							},
							tt.args.input.ToWarehouseId: {
								Id:     tt.args.input.ToWarehouseId,
								ShopId: userShopId,
							},
						},
					},
					nil,
				).Times(1)

				repoHttpProductMock.EXPECT().GetProductById(
					test_util.ContextTypeMatcher,
					model.GetProductByIdInput{
						Id: tt.args.input.ProductId,
					},
				).Return(
					model.GetProductByIdOutput{
						ExtProduct: model.ExtProduct{
							Id: tt.args.input.ProductId,
							Shop: model.ExtShop{
								Id: userShopId,
							},
						},
					},
					nil,
				).Times(1)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.mock != nil {
				tt.mock(&tt)
			}

			u := warehouse.InitWarehouseUsecase(warehouse.InitWarehouseUsecaseOptions{
				RepoDbWarehouse: tt.fields.RepoDbWarehouse,
				RepoHttpShop:    tt.fields.RepoHttpShop,
				RepoHttpProduct: tt.fields.RepoHttpProduct,
			})
			gotErr := u.ValidateCreateOrUpdateStock(tt.args.ctx, tt.args.input)
			tt.assertErr(t, gotErr)
		})
	}
}

func Test_WarehouseUsecase_TransferStock(t *testing.T) {
	type fields struct {
		RepoDbWarehouse db_warehouse.RepoDbWarehouseInterface
	}
	type args struct {
		ctx   context.Context
		input model.CreateOrUpdateStockInput
	}
	type test struct {
		name      string
		fields    fields
		args      args
		assertErr require.ErrorAssertionFunc
		mock      func(tt *test)
	}

	var (
		ctx   = context.Background()
		input = model.CreateOrUpdateStockInput{
			JWT:           faker.Jwt(),
			UserId:        uuid.MustParse(faker.UUIDHyphenated()),
			WarehouseId:   uuid.MustParse(faker.UUIDHyphenated()),
			ProductId:     uuid.MustParse(faker.UUIDHyphenated()),
			Stock:         8,
			ToWarehouseId: uuid.MustParse(faker.UUIDHyphenated()),
		}
	)

	tests := []test{
		{
			name: "should return error if error occurred repoDbWarehouse.DeductStock",
			args: args{
				ctx:   ctx,
				input: input,
			},
			assertErr: test_util.RequireErrorIs(in_err.ErrDeductStock),
			mock: func(tt *test) {
				mockCtrl := gomock.NewController(t)

				repoDbWarehouseMock := db_warehouse.NewMockRepoDbWarehouseInterface(mockCtrl)

				tt.fields.RepoDbWarehouse = repoDbWarehouseMock

				repoDbWarehouseMock.EXPECT().DeductStock(
					test_util.ContextTypeMatcher,
					model.DeductStockInput{
						WarehouseId: tt.args.input.WarehouseId,
						ProductId:   tt.args.input.ProductId,
						Quantity:    tt.args.input.Stock,
						RequestedBy: tt.args.input.UserId.String(),
					}.Matcher(),
				).Return(
					model.DeductStockOutput{},
					errors.Join(in_err.ErrDeductStock, errors.New("expected DeductStock error")),
				).Times(1)
			},
		},
		{
			name: "should return error if error occurred repoDbWarehouse.UpsertStock",
			args: args{
				ctx:   ctx,
				input: input,
			},
			assertErr: test_util.RequireErrorIs(in_err.ErrUpsertStock),
			mock: func(tt *test) {
				mockCtrl := gomock.NewController(t)

				repoDbWarehouseMock := db_warehouse.NewMockRepoDbWarehouseInterface(mockCtrl)

				tt.fields.RepoDbWarehouse = repoDbWarehouseMock

				repoDbWarehouseMock.EXPECT().DeductStock(
					test_util.ContextTypeMatcher,
					model.DeductStockInput{
						WarehouseId: tt.args.input.WarehouseId,
						ProductId:   tt.args.input.ProductId,
						Quantity:    tt.args.input.Stock,
						RequestedBy: tt.args.input.UserId.String(),
					}.Matcher(),
				).Return(
					model.DeductStockOutput{
						Successful: true,
					},
					nil,
				).Times(1)

				repoDbWarehouseMock.EXPECT().UpsertStock(
					test_util.ContextTypeMatcher,
					model.UpsertStockInput{
						WarehouseId: tt.args.input.ToWarehouseId,
						ProductId:   tt.args.input.ProductId,
						Stock:       tt.args.input.Stock,
						IsTransfer:  true,
						UpsertedBy:  tt.args.input.UserId.String(),
					}.Matcher(),
				).Return(
					model.UpsertStockOutput{},
					errors.Join(in_err.ErrUpsertStock, errors.New("expected UpsertStock error")),
				).Times(1)
			},
		},
		{
			name: "should return no error if stock transfer was successful",
			args: args{
				ctx:   ctx,
				input: input,
			},
			assertErr: require.NoError,
			mock: func(tt *test) {
				mockCtrl := gomock.NewController(t)

				repoDbWarehouseMock := db_warehouse.NewMockRepoDbWarehouseInterface(mockCtrl)

				tt.fields.RepoDbWarehouse = repoDbWarehouseMock

				repoDbWarehouseMock.EXPECT().DeductStock(
					test_util.ContextTypeMatcher,
					model.DeductStockInput{
						WarehouseId: tt.args.input.WarehouseId,
						ProductId:   tt.args.input.ProductId,
						Quantity:    tt.args.input.Stock,
						RequestedBy: tt.args.input.UserId.String(),
					}.Matcher(),
				).Return(
					model.DeductStockOutput{
						Successful: true,
					},
					nil,
				).Times(1)

				repoDbWarehouseMock.EXPECT().UpsertStock(
					test_util.ContextTypeMatcher,
					model.UpsertStockInput{
						WarehouseId: tt.args.input.ToWarehouseId,
						ProductId:   tt.args.input.ProductId,
						Stock:       tt.args.input.Stock,
						IsTransfer:  true,
						UpsertedBy:  tt.args.input.UserId.String(),
					}.Matcher(),
				).Return(
					model.UpsertStockOutput{
						Id: uuid.MustParse(faker.UUIDHyphenated()),
					},
					nil,
				).Times(1)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.mock != nil {
				tt.mock(&tt)
			}

			u := warehouse.InitWarehouseUsecase(warehouse.InitWarehouseUsecaseOptions{
				RepoDbWarehouse: tt.fields.RepoDbWarehouse,
			})
			gotErr := u.TransferStock(tt.args.ctx, tt.args.input)
			tt.assertErr(t, gotErr)
		})
	}
}
