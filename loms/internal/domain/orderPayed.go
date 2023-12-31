package domain

import (
	"context"
)

import (
	"github.com/pkg/errors"
)

var (
	ErrorOrderIsPayed    = errors.New("order already mark as payed")
	ErrorOrderIsCanceled = errors.New("order is canceled")
	ErrorOrderIsFailed   = errors.New("order is failed")
	ErrorNotEnoughItems  = errors.New("Not enough item")
)

func (d *domain) OrderPayedMark(ctx context.Context, order *Order) error {
	// Не обрабатываем оплаченные заказы
	if order.Status == Payed {
		return ErrorOrderIsPayed
	}

	// Не обрабатываем отмененные заказы
	if order.Status == Cancelled {
		return ErrorOrderIsCanceled
	}

	// Не обрабатываем провальные заказы
	if order.Status == Failed {
		return ErrorOrderIsFailed
	}

	err := d.manager.RepeatableRead(ctx, func(ctxTx context.Context) error {
		// Получаем резервации заказа
		orderItemStocks, err := d.orderItemStockRepository.GetListByOrderID(ctxTx, order.ID)
		if err != nil {
			return err
		}

		// Делаем списания товара со склада
		for _, orderItemStock := range orderItemStocks {
			// Получаем остатки товара на складе с которого хотим списать разерервированный товар
			stock, err := d.stockRepository.GetByWarehouseIDAndSku(ctxTx, orderItemStock.WarehouseID, orderItemStock.Sku)
			if err != nil {
				return err
			}
			stock.Count -= orderItemStock.Count
			err = d.stockRepository.UpdateCount(ctxTx, stock)
			if err != nil {
				return err
			}
		}

		// Если все прошло укспешно то помечаем заказ как оплаченный
		order.Status = Payed
		err = d.orderRepository.Update(ctxTx, order)
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return err
	}

	err = d.orderStatusNotifier.Notify(order)
	if err != nil {
		return err
	}

	return nil
}
