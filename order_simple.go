package bereke_merchant

import (
	"context"

	"github.com/bsagat/bereke-merchant-api/models/core"
)

// RegisterOrder — упрощённая обёртка для создания заказа.
// Принимает номер заказа, сумму, валюту, а также URL для возврата при успешной и неуспешной оплате.
// Возвращает RegisterOrderResponse, содержащий ссылку для редиректа покупателя и ID заказа.
func (a *api) RegisterOrderByNumber(
	ctx context.Context,
	orderNumber string,
	amount float64,
	currency int,
	returnURL, failURL string,
) (core.RegisterOrderResponse, error) {
	// Формируем упрощённый запрос
	req := core.RegisterOrderRequest{
		Order: core.Order{
			OrderNumber: orderNumber,
			Amount:      amount,
			Currency:    currency,
			ReturnURL:   returnURL,
			FailURL:     failURL,
		},
	}

	// Вызываем основной метод регистрации заказа
	return a.RegisterOrder(ctx, req)
}

func (a *api) RefundOrderByID(
	ctx context.Context,
	amount float64,
	currency int,
	orderID string,
) (core.Response, error) {
	// Формируем упрощённый запрос
	req := core.RefundOrderRequest{
		OrderID:  orderID,
		Amount:   amount,
		Currency: currency,
	}

	// Вызываем основной метод возврата средств
	return a.RefundOrder(ctx, req)
}

func (a *api) ReversalOrderByID(
	ctx context.Context,
	amount float64,
	currency int,
	orderID string,
) (core.Response, error) {
	// Формируем упрощённый запрос
	req := core.ReversalOrderRequest{
		OrderID:  orderID,
		Amount:   amount,
		Currency: currency,
	}

	// Вызываем основной метод
	return a.ReversalOrder(ctx, req)
}

func (a *api) CancelOrderByID(
	ctx context.Context,
	orderID string,
) (core.Response, error) {

	// Формируем основной запрос
	req := core.CancelOrderRequest{
		OrderID: orderID,
	}

	// Вызываем основной метод
	return a.CancelOrder(ctx, req)
}

func (a *api) GetOrderStatusByID(
	ctx context.Context,
	orderID string,
) (core.OrderStatusResponse, error) {

	// Формируем основной запрос
	req := core.OrderStatusRequest{
		OrderID: orderID,
	}

	// Вызываем основной метод
	return a.GetOrderStatus(ctx, req)
}
