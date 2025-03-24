package warehouse_test

import (
	"context"
	"errors"
	"testing"

	"github.com/fenky-ng/edot-test-case/warehouse/internal/constant"
	in_err "github.com/fenky-ng/edot-test-case/warehouse/internal/error"
	"github.com/fenky-ng/edot-test-case/warehouse/internal/model"
	db_warehouse "github.com/fenky-ng/edot-test-case/warehouse/internal/repository/db/warehouse"
	"github.com/fenky-ng/edot-test-case/warehouse/internal/usecase/warehouse"
	test_util "github.com/fenky-ng/edot-test-case/warehouse/internal/utility/test"
	"github.com/go-faker/faker/v4"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
)

func Test_WarehouseUsecase_DeductStocks(t *testing.T) {
	type fields struct {
		RepoDbWarehouse  db_warehouse.RepoDbWarehouseInterface
		WarehouseUsecase warehouse.WarehouseUsecaseInterface
	}
	type args struct {
		ctx   context.Context
		input model.DeductStocksInput
	}
	type test struct {
		name        string
		fields      fields
		args        args
		expectedRes model.DeductStocksOutput
		assertErr   require.ErrorAssertionFunc
		mock        func(tt *test)
	}

	var (
		ctx = context.Background()

		item1 = model.DeductStockItem{
			ProductId:   uuid.MustParse(faker.UUIDHyphenated()),
			WarehouseId: uuid.MustParse(faker.UUIDHyphenated()),
			Quantity:    2,
		}
		item2 = model.DeductStockItem{
			ProductId:   uuid.MustParse(faker.UUIDHyphenated()),
			WarehouseId: uuid.MustParse(faker.UUIDHyphenated()),
			Quantity:    8,
		}
		deductInput = model.DeductStocksInput{
			UserId:  uuid.MustParse(faker.UUIDHyphenated()),
			OrderNo: "ORD/280808182838/1234567890",
			Items: []model.DeductStockItem{
				item1,
				item2,
			},
			Release: true,
		}
	)

	tests := []test{
		{
			name: "should return error if request is invalid",
			args: args{
				ctx:   ctx,
				input: deductInput,
			},
			expectedRes: model.DeductStocksOutput{},
			assertErr:   test_util.RequireErrorIs(in_err.ErrGetWarehouses),
			mock: func(tt *test) {
				mockCtrl := gomock.NewController(t)

				repoDbWarehouseMock := db_warehouse.NewMockRepoDbWarehouseInterface(mockCtrl)
				warehouseUsecaseMock := warehouse.NewMockWarehouseUsecaseInterface(mockCtrl)

				tt.fields.RepoDbWarehouse = repoDbWarehouseMock
				tt.fields.WarehouseUsecase = warehouseUsecaseMock

				warehouseUsecaseMock.EXPECT().ValidateDeductStocks(
					test_util.ContextTypeMatcher,
					tt.args.input,
				).Return(
					errors.Join(in_err.ErrGetWarehouses, errors.New("expected ValidateDeductStocks.GetWarehouses error")),
				).Times(1)

			},
		},
		{
			name: "should return error if error occurred in repoDbWarehouse.Begin",
			args: args{
				ctx:   ctx,
				input: deductInput,
			},
			expectedRes: model.DeductStocksOutput{},
			assertErr:   test_util.RequireErrorIs(in_err.ErrDatabaseTransaction),
			mock: func(tt *test) {
				mockCtrl := gomock.NewController(t)

				repoDbWarehouseMock := db_warehouse.NewMockRepoDbWarehouseInterface(mockCtrl)
				warehouseUsecaseMock := warehouse.NewMockWarehouseUsecaseInterface(mockCtrl)

				tt.fields.RepoDbWarehouse = repoDbWarehouseMock
				tt.fields.WarehouseUsecase = warehouseUsecaseMock

				warehouseUsecaseMock.EXPECT().ValidateDeductStocks(
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
			name: "should return error if error occurred in repoDbWarehouse.DeductStock",
			args: args{
				ctx:   ctx,
				input: deductInput,
			},
			expectedRes: model.DeductStocksOutput{},
			assertErr:   test_util.RequireErrorIs(in_err.ErrDeductStock),
			mock: func(tt *test) {
				mockCtrl := gomock.NewController(t)

				repoDbWarehouseMock := db_warehouse.NewMockRepoDbWarehouseInterface(mockCtrl)
				warehouseUsecaseMock := warehouse.NewMockWarehouseUsecaseInterface(mockCtrl)

				tt.fields.RepoDbWarehouse = repoDbWarehouseMock
				tt.fields.WarehouseUsecase = warehouseUsecaseMock

				warehouseUsecaseMock.EXPECT().ValidateDeductStocks(
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

				errDeductStock := errors.Join(in_err.ErrDeductStock, errors.New("expected DeductStock error"))

				repoDbWarehouseMock.EXPECT().CommitOrRollback(
					test_util.ContextTypeMatcher,
					errDeductStock,
				).Return(
					errDeductStock,
				).Times(1)

				repoDbWarehouseMock.EXPECT().DeductStock(
					test_util.ContextTypeMatcher,
					model.DeductStockInput{
						WarehouseId: tt.args.input.Items[0].WarehouseId,
						ProductId:   tt.args.input.Items[0].ProductId,
						Quantity:    tt.args.input.Items[0].Quantity,
						Release:     tt.args.input.Release,
						RequestedBy: tt.args.input.UserId.String(),
						NoWait:      true,
					}.Matcher(),
				).Return(
					model.DeductStockOutput{},
					errDeductStock,
				).Times(1)
			},
		},
		{
			name: "should return response successful if stock deduction/release was successful",
			args: args{
				ctx:   ctx,
				input: deductInput,
			},
			expectedRes: model.DeductStocksOutput{
				Successful: true,
			},
			assertErr: require.NoError,
			mock: func(tt *test) {
				mockCtrl := gomock.NewController(t)

				repoDbWarehouseMock := db_warehouse.NewMockRepoDbWarehouseInterface(mockCtrl)
				warehouseUsecaseMock := warehouse.NewMockWarehouseUsecaseInterface(mockCtrl)

				tt.fields.RepoDbWarehouse = repoDbWarehouseMock
				tt.fields.WarehouseUsecase = warehouseUsecaseMock

				warehouseUsecaseMock.EXPECT().ValidateDeductStocks(
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

				repoDbWarehouseMock.EXPECT().DeductStock(
					test_util.ContextTypeMatcher,
					model.DeductStockInput{
						WarehouseId: tt.args.input.Items[0].WarehouseId,
						ProductId:   tt.args.input.Items[0].ProductId,
						Quantity:    tt.args.input.Items[0].Quantity,
						Release:     tt.args.input.Release,
						RequestedBy: tt.args.input.UserId.String(),
						NoWait:      true,
					}.Matcher(),
				).Return(
					model.DeductStockOutput{
						Successful: true,
					},
					nil,
				).Times(1)

				repoDbWarehouseMock.EXPECT().DeductStock(
					test_util.ContextTypeMatcher,
					model.DeductStockInput{
						WarehouseId: tt.args.input.Items[1].WarehouseId,
						ProductId:   tt.args.input.Items[1].ProductId,
						Quantity:    tt.args.input.Items[1].Quantity,
						Release:     tt.args.input.Release,
						RequestedBy: tt.args.input.UserId.String(),
						NoWait:      true,
					}.Matcher(),
				).Return(
					model.DeductStockOutput{
						Successful: true,
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
			u.WarehouseUsecase = tt.fields.WarehouseUsecase
			gotRes, gotErr := u.DeductStocks(tt.args.ctx, tt.args.input)
			tt.assertErr(t, gotErr)
			require.Equal(t, tt.expectedRes, gotRes)
		})
	}
}

func Test_WarehouseUsecase_ValidateDeductStocks(t *testing.T) {
	type fields struct {
		RepoDbWarehouse db_warehouse.RepoDbWarehouseInterface
	}
	type args struct {
		ctx   context.Context
		input model.DeductStocksInput
	}
	type test struct {
		name      string
		fields    fields
		args      args
		assertErr require.ErrorAssertionFunc
		mock      func(tt *test)
	}

	var (
		ctx = context.Background()

		warehouseId = uuid.MustParse(faker.UUIDHyphenated())
		item1       = model.DeductStockItem{
			ProductId:   uuid.MustParse(faker.UUIDHyphenated()),
			WarehouseId: warehouseId,
			Quantity:    2,
		}
		item2 = model.DeductStockItem{
			ProductId:   uuid.MustParse(faker.UUIDHyphenated()),
			WarehouseId: warehouseId,
			Quantity:    8,
		}
		deductInput = model.DeductStocksInput{
			UserId:  uuid.MustParse(faker.UUIDHyphenated()),
			OrderNo: "ORD/280808182838/1234567890",
			Items: []model.DeductStockItem{
				item1,
				item2,
			},
			Release: false,
		}
	)

	releaseInput := deductInput
	releaseInput.Release = true

	tests := []test{
		{
			name: "if request release then skip validation",
			args: args{
				ctx:   ctx,
				input: releaseInput,
			},
			assertErr: require.NoError,
			mock: func(tt *test) {
				mockCtrl := gomock.NewController(t)

				repoDbWarehouseMock := db_warehouse.NewMockRepoDbWarehouseInterface(mockCtrl)

				tt.fields.RepoDbWarehouse = repoDbWarehouseMock
			},
		},
		{
			name: "if error occurred in repoDbWarehouse.GetWarehouses then return error",
			args: args{
				ctx:   ctx,
				input: deductInput,
			},
			assertErr: test_util.RequireErrorIs(in_err.ErrGetWarehouses),
			mock: func(tt *test) {
				mockCtrl := gomock.NewController(t)

				repoDbWarehouseMock := db_warehouse.NewMockRepoDbWarehouseInterface(mockCtrl)

				tt.fields.RepoDbWarehouse = repoDbWarehouseMock

				repoDbWarehouseMock.EXPECT().GetWarehouses(
					test_util.ContextTypeMatcher,
					model.GetWarehousesInput{
						Ids: []uuid.UUID{
							warehouseId,
						},
					},
				).Return(
					model.GetWarehousesOutput{},
					errors.Join(in_err.ErrGetWarehouses, errors.New("expected GetWarehouses error")),
				).Times(1)
			},
		},
		{
			name: "if warehouse not found then return error",
			args: args{
				ctx:   ctx,
				input: deductInput,
			},
			assertErr: test_util.RequireErrorIs(in_err.ErrWarehouseNotFound),
			mock: func(tt *test) {
				mockCtrl := gomock.NewController(t)

				repoDbWarehouseMock := db_warehouse.NewMockRepoDbWarehouseInterface(mockCtrl)

				tt.fields.RepoDbWarehouse = repoDbWarehouseMock

				repoDbWarehouseMock.EXPECT().GetWarehouses(
					test_util.ContextTypeMatcher,
					model.GetWarehousesInput{
						Ids: []uuid.UUID{
							warehouseId,
						},
					},
				).Return(
					model.GetWarehousesOutput{},
					nil,
				).Times(1)
			},
		},
		{
			name: "if warehouse is inactive then return error",
			args: args{
				ctx:   ctx,
				input: deductInput,
			},
			assertErr: test_util.RequireErrorIs(in_err.ErrWarehouseInactive),
			mock: func(tt *test) {
				mockCtrl := gomock.NewController(t)

				repoDbWarehouseMock := db_warehouse.NewMockRepoDbWarehouseInterface(mockCtrl)

				tt.fields.RepoDbWarehouse = repoDbWarehouseMock

				repoDbWarehouseMock.EXPECT().GetWarehouses(
					test_util.ContextTypeMatcher,
					model.GetWarehousesInput{
						Ids: []uuid.UUID{
							warehouseId,
						},
					},
				).Return(
					model.GetWarehousesOutput{
						Warehouses: []model.Warehouse{
							{
								Id:     warehouseId,
								Status: constant.WarehouseStatus_Inactive,
							},
						},
						WarehouseById: map[uuid.UUID]model.Warehouse{
							warehouseId: {
								Id:     warehouseId,
								Status: constant.WarehouseStatus_Inactive,
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
			})
			gotErr := u.ValidateDeductStocks(tt.args.ctx, tt.args.input)
			tt.assertErr(t, gotErr)
		})
	}
}
