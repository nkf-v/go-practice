package tests

import (
	"context"
	"testing"
)

import (
	"github.com/brianvoe/gofakeit/v6"
	"github.com/gojuno/minimock/v3"
	"github.com/jackc/pgx/v4"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/require"
	"route256/libs/db"
	"route256/libs/db/transaction"
	dbMocks "route256/loms/internal/client/db/mocks"
	"route256/loms/internal/domain"
	domainMocks "route256/loms/internal/domain/mocks"
)

func TestFailOrder(t *testing.T) {
	var (
		mc = minimock.NewController(t)

		ctx    = context.Background()
		txOpts = pgx.TxOptions{IsoLevel: pgx.RepeatableRead}
		tx     = dbMocks.NewTxMock(mc)
		ctxTx  = context.WithValue(ctx, db.TxKey, tx)

		orderID               = gofakeit.Int64()
		amountOrderItemStocks = 5

		orderWithStatusAwaitingPayment *domain.Order
		orderWithStatusNew             *domain.Order
		orderWithStatusCancelled       *domain.Order
		orderWithStatusFailed          *domain.Order
		orderWithStatusPayed           *domain.Order
		orderItemStocks                []*domain.OrderItemStock
		stocks                         []*domain.Stock

		updateOrderError          = errors.New("error update orderID")
		getOrderItemStocksError   = errors.New("error get orderID item stocks")
		getStockError             = errors.New("error get stock")
		updateStockCountError     = errors.New("error update stock count")
		deleteOrderItemStockError = errors.New("error delete orderID item stock")
	)

	orderWithStatusAwaitingPayment = &domain.Order{
		ID:     orderID,
		Status: domain.AwaitingPayment,
	}
	orderWithStatusNew = &domain.Order{
		ID:     orderID,
		Status: domain.New,
	}
	orderWithStatusCancelled = &domain.Order{
		ID:     orderID,
		Status: domain.Cancelled,
	}
	orderWithStatusFailed = &domain.Order{
		ID:     orderID,
		Status: domain.Failed,
	}
	orderWithStatusPayed = &domain.Order{
		ID:     orderID,
		Status: domain.Payed,
	}

	orderItemStocks = make([]*domain.OrderItemStock, 0, amountOrderItemStocks)
	stocks = make([]*domain.Stock, 0, amountOrderItemStocks)
	for i := 0; i < amountOrderItemStocks; i++ {
		orderItemStock := &domain.OrderItemStock{
			OrderId:     orderID,
			Sku:         gofakeit.Uint32(),
			Count:       gofakeit.Uint64(),
			WarehouseID: gofakeit.Int64(),
		}
		orderItemStocks = append(orderItemStocks, orderItemStock)
		stocks = append(stocks, &domain.Stock{
			Sku:         orderItemStock.Sku,
			WarehouseID: orderItemStock.WarehouseID,
			Count:       gofakeit.Uint64(),
		})
	}

	testCases := []struct {
		name string
		args struct {
			order *domain.Order
		}
		want                         error
		dbMock                       dbMock
		stockRepositoryMock          stockRepositoryMock
		orderRepositoryMock          orderRepositoryMock
		orderItemRepositoryMock      orderItemRepositoryMock
		orderItemStockRepositoryMock orderItemStockRepositoryMock
	}{
		{
			name: "negative case orderID is new",
			args: struct {
				order *domain.Order
			}{
				order: orderWithStatusNew,
			},
			want: domain.ErrorOrderIsNew,
			dbMock: func(mc *minimock.Controller) db.DB {
				return nil
			},
			stockRepositoryMock: func(mc *minimock.Controller) domain.StockRepository {
				return nil
			},
			orderRepositoryMock: func(mc *minimock.Controller) domain.OrderRepository {
				return nil
			},
			orderItemRepositoryMock: func(mc *minimock.Controller) domain.OrderItemRepository {
				return nil
			},
			orderItemStockRepositoryMock: func(mc *minimock.Controller) domain.OrderItemStockRepository {
				return nil
			},
		},
		{
			name: "negative case orderID is cancelled",
			args: struct {
				order *domain.Order
			}{
				order: orderWithStatusCancelled,
			},
			want: domain.ErrorOrderIsCanceled,
			dbMock: func(mc *minimock.Controller) db.DB {
				return nil
			},
			stockRepositoryMock: func(mc *minimock.Controller) domain.StockRepository {
				return nil
			},
			orderRepositoryMock: func(mc *minimock.Controller) domain.OrderRepository {
				return nil
			},
			orderItemRepositoryMock: func(mc *minimock.Controller) domain.OrderItemRepository {
				return nil
			},
			orderItemStockRepositoryMock: func(mc *minimock.Controller) domain.OrderItemStockRepository {
				return nil
			},
		},
		{
			name: "negative case orderID is failed",
			args: struct {
				order *domain.Order
			}{
				order: orderWithStatusFailed,
			},
			want: domain.ErrorOrderIsFailed,
			dbMock: func(mc *minimock.Controller) db.DB {
				return nil
			},
			stockRepositoryMock: func(mc *minimock.Controller) domain.StockRepository {
				return nil
			},
			orderRepositoryMock: func(mc *minimock.Controller) domain.OrderRepository {
				return nil
			},
			orderItemRepositoryMock: func(mc *minimock.Controller) domain.OrderItemRepository {
				return nil
			},
			orderItemStockRepositoryMock: func(mc *minimock.Controller) domain.OrderItemStockRepository {
				return nil
			},
		},
		{
			name: "negative case error update orderID",
			args: struct {
				order *domain.Order
			}{
				order: orderWithStatusAwaitingPayment,
			},
			want: updateOrderError,
			dbMock: func(mc *minimock.Controller) db.DB {
				mock := dbMocks.NewDBMock(mc)
				mock.BeginTxMock.Expect(ctx, txOpts).Return(tx, nil)
				tx.RollbackMock.Expect(ctxTx).Return(nil)
				return mock
			},
			stockRepositoryMock: func(mc *minimock.Controller) domain.StockRepository {
				return nil
			},
			orderRepositoryMock: func(mc *minimock.Controller) domain.OrderRepository {
				mock := domainMocks.NewOrderRepositoryMock(mc)
				mock.UpdateMock.Expect(ctxTx, orderWithStatusAwaitingPayment).Return(updateOrderError)
				return mock
			},
			orderItemRepositoryMock: func(mc *minimock.Controller) domain.OrderItemRepository {
				return nil
			},
			orderItemStockRepositoryMock: func(mc *minimock.Controller) domain.OrderItemStockRepository {
				return nil
			},
		},
		{
			name: "negative case error get orderID item stocks",
			args: struct {
				order *domain.Order
			}{
				order: orderWithStatusPayed,
			},
			want: getOrderItemStocksError,
			dbMock: func(mc *minimock.Controller) db.DB {
				mock := dbMocks.NewDBMock(mc)
				mock.BeginTxMock.Expect(ctx, txOpts).Return(tx, nil)
				tx.RollbackMock.Expect(ctxTx).Return(nil)
				return mock
			},
			stockRepositoryMock: func(mc *minimock.Controller) domain.StockRepository {
				return nil
			},
			orderRepositoryMock: func(mc *minimock.Controller) domain.OrderRepository {
				mock := domainMocks.NewOrderRepositoryMock(mc)
				mock.UpdateMock.Expect(ctxTx, orderWithStatusPayed).Return(nil)
				return mock
			},
			orderItemRepositoryMock: func(mc *minimock.Controller) domain.OrderItemRepository {
				return nil
			},
			orderItemStockRepositoryMock: func(mc *minimock.Controller) domain.OrderItemStockRepository {
				mock := domainMocks.NewOrderItemStockRepositoryMock(mc)
				mock.GetListByOrderIDMock.Expect(ctxTx, orderWithStatusPayed.ID).Return(nil, getOrderItemStocksError)
				return mock
			},
		},
		{
			name: "negative case error get stock",
			args: struct {
				order *domain.Order
			}{
				order: orderWithStatusPayed,
			},
			want: getStockError,
			dbMock: func(mc *minimock.Controller) db.DB {
				mock := dbMocks.NewDBMock(mc)
				mock.BeginTxMock.Expect(ctx, txOpts).Return(tx, nil)
				tx.RollbackMock.Expect(ctxTx).Return(nil)
				return mock
			},
			stockRepositoryMock: func(mc *minimock.Controller) domain.StockRepository {
				mock := domainMocks.NewStockRepositoryMock(mc)
				mock.GetByWarehouseIDAndSkuMock.Return(nil, getStockError)
				return mock
			},
			orderRepositoryMock: func(mc *minimock.Controller) domain.OrderRepository {
				mock := domainMocks.NewOrderRepositoryMock(mc)
				mock.UpdateMock.Expect(ctxTx, orderWithStatusPayed).Return(nil)
				return mock
			},
			orderItemRepositoryMock: func(mc *minimock.Controller) domain.OrderItemRepository {
				return nil
			},
			orderItemStockRepositoryMock: func(mc *minimock.Controller) domain.OrderItemStockRepository {
				mock := domainMocks.NewOrderItemStockRepositoryMock(mc)
				mock.GetListByOrderIDMock.Expect(ctxTx, orderWithStatusPayed.ID).Return(orderItemStocks, nil)
				return mock
			},
		},
		{
			name: "negative case error update stock count",
			args: struct {
				order *domain.Order
			}{
				order: orderWithStatusPayed,
			},
			want: updateStockCountError,
			dbMock: func(mc *minimock.Controller) db.DB {
				mock := dbMocks.NewDBMock(mc)
				mock.BeginTxMock.Expect(ctx, txOpts).Return(tx, nil)
				tx.RollbackMock.Expect(ctxTx).Return(nil)
				return mock
			},
			stockRepositoryMock: func(mc *minimock.Controller) domain.StockRepository {
				mock := domainMocks.NewStockRepositoryMock(mc)
				for index, orderItemStock := range orderItemStocks {
					mock.GetByWarehouseIDAndSkuMock.
						When(ctxTx, orderItemStock.WarehouseID, orderItemStock.Sku).
						Then(stocks[index], nil)
				}
				mock.UpdateCountMock.Return(updateStockCountError)
				return mock
			},
			orderRepositoryMock: func(mc *minimock.Controller) domain.OrderRepository {
				mock := domainMocks.NewOrderRepositoryMock(mc)
				mock.UpdateMock.Expect(ctxTx, orderWithStatusPayed).Return(nil)
				return mock
			},
			orderItemRepositoryMock: func(mc *minimock.Controller) domain.OrderItemRepository {
				return nil
			},
			orderItemStockRepositoryMock: func(mc *minimock.Controller) domain.OrderItemStockRepository {
				mock := domainMocks.NewOrderItemStockRepositoryMock(mc)
				mock.GetListByOrderIDMock.Expect(ctxTx, orderWithStatusPayed.ID).Return(orderItemStocks, nil)
				return mock
			},
		},
		{
			name: "negative case error update stock count",
			args: struct {
				order *domain.Order
			}{
				order: orderWithStatusPayed,
			},
			want: deleteOrderItemStockError,
			dbMock: func(mc *minimock.Controller) db.DB {
				mock := dbMocks.NewDBMock(mc)
				mock.BeginTxMock.Expect(ctx, txOpts).Return(tx, nil)
				tx.RollbackMock.Expect(ctxTx).Return(nil)
				return mock
			},
			stockRepositoryMock: func(mc *minimock.Controller) domain.StockRepository {
				mock := domainMocks.NewStockRepositoryMock(mc)
				for index, orderItemStock := range orderItemStocks {
					mock.GetByWarehouseIDAndSkuMock.
						When(ctxTx, orderItemStock.WarehouseID, orderItemStock.Sku).
						Then(stocks[index], nil)
				}
				mock.UpdateCountMock.Return(nil)
				return mock
			},
			orderRepositoryMock: func(mc *minimock.Controller) domain.OrderRepository {
				mock := domainMocks.NewOrderRepositoryMock(mc)
				mock.UpdateMock.Expect(ctxTx, orderWithStatusPayed).Return(nil)
				return mock
			},
			orderItemRepositoryMock: func(mc *minimock.Controller) domain.OrderItemRepository {
				return nil
			},
			orderItemStockRepositoryMock: func(mc *minimock.Controller) domain.OrderItemStockRepository {
				mock := domainMocks.NewOrderItemStockRepositoryMock(mc)
				mock.GetListByOrderIDMock.Expect(ctxTx, orderWithStatusPayed.ID).Return(orderItemStocks, nil)
				mock.DeleteMock.Return(deleteOrderItemStockError)
				return mock
			},
		},
		{
			name: "positive case",
			args: struct {
				order *domain.Order
			}{
				order: orderWithStatusPayed,
			},
			want: nil,
			dbMock: func(mc *minimock.Controller) db.DB {
				mock := dbMocks.NewDBMock(mc)
				mock.BeginTxMock.Expect(ctx, txOpts).Return(tx, nil)
				tx.CommitMock.Expect(ctxTx).Return(nil)
				return mock
			},
			stockRepositoryMock: func(mc *minimock.Controller) domain.StockRepository {
				mock := domainMocks.NewStockRepositoryMock(mc)
				for index, orderItemStock := range orderItemStocks {
					mock.GetByWarehouseIDAndSkuMock.
						When(ctxTx, orderItemStock.WarehouseID, orderItemStock.Sku).
						Then(stocks[index], nil)
				}
				mock.UpdateCountMock.Return(nil)
				return mock
			},
			orderRepositoryMock: func(mc *minimock.Controller) domain.OrderRepository {
				mock := domainMocks.NewOrderRepositoryMock(mc)
				mock.UpdateMock.Expect(ctxTx, orderWithStatusPayed).Return(nil)
				return mock
			},
			orderItemRepositoryMock: func(mc *minimock.Controller) domain.OrderItemRepository {
				return nil
			},
			orderItemStockRepositoryMock: func(mc *minimock.Controller) domain.OrderItemStockRepository {
				mock := domainMocks.NewOrderItemStockRepositoryMock(mc)
				mock.GetListByOrderIDMock.Expect(ctxTx, orderWithStatusPayed.ID).Return(orderItemStocks, nil)
				for _, orderItemStock := range orderItemStocks {
					mock.DeleteMock.
						When(ctxTx, orderItemStock).
						Then(nil)
				}
				return mock
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			d := domain.NewDomain(
				transaction.NewTransactionManager(testCase.dbMock(mc)),
				testCase.stockRepositoryMock(mc),
				testCase.orderRepositoryMock(mc),
				testCase.orderItemRepositoryMock(mc),
				testCase.orderItemStockRepositoryMock(mc),
			)

			// Созраняем изначальный статус для переиспользования тестовых заказов
			initStatus := testCase.args.order.Status

			err := d.FailOrder(ctx, testCase.args.order)

			if testCase.want == nil {
				require.Equal(t, testCase.want, err)
			} else {
				require.ErrorContains(t, err, testCase.want.Error())
			}

			// После тестирования возвращаем изначальный статус тестовому заказу
			testCase.args.order.Status = initStatus
		})
	}
}