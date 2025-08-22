package bereke_merchant

import (
	"context"

	"github.com/bsagat/bereke-merchant-api/models/core"
)

// RegisterOrderByNumber — упрощённая обёртка для создания заказа (одноэтапный платёж).
// Используется, когда средства сразу списываются с карты клиента.
// Аргументы:
//   - orderNumber — уникальный номер заказа в вашей системе
//   - amount — сумма к оплате
//   - currency — код валюты (например, 398 для KZT)
//   - returnURL — URL, куда клиент будет перенаправлен после успешной оплаты
//   - failURL — URL, куда клиент будет перенаправлен при ошибке оплаты
//
// Возвращает структуру RegisterOrderResponse, содержащую:
//   - идентификатор заказа в банке
//   - ссылку для редиректа клиента на страницу оплаты
func (a *api) RegisterOrderByNumber(
	ctx context.Context,
	orderNumber string,
	amount float64,
	currency int,
	returnURL, failURL string,
) (core.RegisterOrderResponse, error) {
	req := core.RegisterOrderRequest{
		Order: core.Order{
			OrderNumber: orderNumber,
			Amount:      amount,
			Currency:    currency,
			ReturnURL:   returnURL,
			FailURL:     failURL,
		},
	}
	return a.RegisterOrder(ctx, req)
}

// AuthOrderByNumber — упрощённая обёртка для двухэтапной авторизации (hold).
// В этом случае деньги на карте блокируются, но не списываются.
// Для фактического списания средств необходимо вызвать DepositOrder.
// Используется в сценариях, когда услуга/товар подтверждается позже.
func (a *api) AuthOrderByNumber(
	ctx context.Context,
	orderNumber string,
	amount float64,
	currency int,
	returnURL, failURL string,
) (core.RegisterOrderResponse, error) {
	req := core.RegisterOrderRequest{
		Order: core.Order{
			OrderNumber: orderNumber,
			Amount:      amount,
			Currency:    currency,
			ReturnURL:   returnURL,
			FailURL:     failURL,
		},
	}
	return a.AuthOrder(ctx, req)
}

// DepositOrderByNumber — подтверждение авторизованного заказа (capture).
// Используется только после AuthOrderByNumber для списания ранее заблокированных средств.
// Аргументы:
//   - orderNumber — идентификатор заказа
//   - amount — сумма списания (может быть меньше авторизованной); (если указать 0 — будет списана вся доступная сумма)
//   - currency — код валюты (например, 398 для KZT);
func (a *api) DepositOrderByNumber(
	ctx context.Context,
	orderNumber string,
	amount float64,
	currency int,
) (core.Response, error) {
	req := core.DepositOrderRequest{
		OrderID:  orderNumber,
		Amount:   amount,
		Currency: currency,
	}
	return a.DepositOrder(ctx, req)
}

// RefundOrderByID — возврат средств по успешному заказу.
// Средства возвращаются клиенту на карту.
// Аргументы:
//   - amount — сумма возврата (обязательно)
//   - currency — код валюты
//   - orderID — идентификатор заказа
func (a *api) RefundOrderByID(
	ctx context.Context,
	amount float64,
	currency int,
	orderID string,
) (core.Response, error) {
	req := core.RefundOrderRequest{
		OrderID:  orderID,
		Amount:   amount,
		Currency: currency,
	}
	return a.RefundOrder(ctx, req)
}

// ReversalOrderByID — аннулирование авторизованного платежа (reversal).
// Используется, если заказ был авторизован, но списание ещё не произошло.
// Фактически снимает блокировку с карты клиента.
// Аргументы:
//   - orderID — идентификатор заказа
//   - amount — сумма списания (может быть меньше авторизованной); (если указать 0 — будет аннулирована вся сумма)
//   - currency — код валюты (например, 398 для KZT);
func (a *api) ReversalOrderByID(
	ctx context.Context,
	amount float64,
	currency int,
	orderID string,
) (core.Response, error) {
	req := core.ReversalOrderRequest{
		OrderID:  orderID,
		Amount:   amount,
		Currency: currency,
	}
	return a.ReversalOrder(ctx, req)
}

// CancelOrderByID — отмена заказа.
// Используется до авторизации/списания средств.
// Если заказ ещё не был оплачен или подтверждён — переводит его в состояние CANCELED.
//
// ВАЖНО: Для вызова метода у пользователя/продавца должны быть соответствующие права
// на отклонение заказов (назначаются в настройках платёжной системы)
func (a *api) CancelOrderByID(
	ctx context.Context,
	orderID string,
) (core.Response, error) {
	req := core.CancelOrderRequest{
		OrderID: orderID,
	}
	return a.CancelOrder(ctx, req)
}

// GetOrderStatusByID — получение статуса заказа.
// Используется для проверки, в каком состоянии находится заказ (AUTHORIZED, DEPOSITED, REFUNDED и т.д.).
// Аргументы:
//   - orderID — идентификатор заказа
//
// Возвращает OrderStatusResponse с полным описанием состояния транзакции.
func (a *api) GetOrderStatusByID(
	ctx context.Context,
	orderID string,
) (core.OrderStatusResponse, error) {
	req := core.OrderStatusRequest{
		OrderID: orderID,
	}
	return a.GetOrderStatus(ctx, req)
}
