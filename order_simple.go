package bereke_merchant

import (
	"context"

	money "github.com/bsagat/bereke-merchant-api/currency"
	"github.com/bsagat/bereke-merchant-api/models/dto"
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
) (dto.RegisterOrderResponse, error) {

	// Преобразуем сумму в минорные единицы (например, 10.50 USD -> 1050 центов)
	amountMinor, err := money.ToMinorUnit(amount, currency)
	if err != nil {
		return dto.RegisterOrderResponse{}, err
	}

	// Формируем упрощённый запрос
	req := dto.RegisterOrderRequest{
		Order: dto.Order{
			OrderNumber: orderNumber,
			Amount:      amountMinor,
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
) (dto.Response, error) {

	// Преобразуем сумму в минорные единицы (например, 10.50 USD -> 1050 центов)
	amountMinor, err := money.ToMinorUnit(amount, currency)
	if err != nil {
		return dto.Response{}, err
	}

	// Формируем упрощённый запрос
	req := dto.RefundOrderRequest{
		OrderID:  orderID,
		Amount:   amountMinor,
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
) (dto.Response, error) {

	// Преобразуем сумму в минорные единицы (например, 10.50 USD -> 1050 центов)
	amountMinor, err := money.ToMinorUnit(amount, currency)
	if err != nil {
		return dto.Response{}, err
	}

	// Формируем упрощённый запрос
	req := dto.ReversalOrderRequest{
		OrderID:  orderID,
		Amount:   amountMinor,
		Currency: currency,
	}

	// Вызываем основной метод
	return a.ReversalOrder(ctx, req)
}

func (a *api) CancelOrderByID(
	ctx context.Context,
	orderID string,
) (dto.Response, error) {

	// Формируем основной запрос
	req := dto.CancelOrderRequest{
		OrderID: orderID,
	}

	// Вызываем основной метод
	return a.CancelOrder(ctx, req)
}

func (a *api) GetOrderStatusByID(
	ctx context.Context,
	orderID string,
) (dto.OrderStatusResponse, error) {

	// Формируем основной запрос
	req := dto.OrderStatusRequest{
		OrderID: orderID,
	}

	// Вызываем основной метод
	return a.GetOrderStatus(ctx, req)
}
