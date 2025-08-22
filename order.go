package bereke_merchant

import (
	"context"
	"io"

	"github.com/bsagat/bereke-merchant-api/models/core"
	"github.com/bsagat/bereke-merchant-api/models/dto"
)

// RegisterOrder — регистрация нового заказа (одноэтапный платёж).
// Отправляет запрос в endpoint `register.do`.
// Используется, если необходимо сразу списать деньги с карты клиента.
// Аргументы:
//   - req — структура RegisterOrderRequest с номером заказа, суммой, валютой и URL для редиректов.
//
// Возвращает RegisterOrderResponse с ID заказа и ссылкой для оплаты.
func (a *api) RegisterOrder(ctx context.Context, req core.RegisterOrderRequest) (core.RegisterOrderResponse, error) {
	reqParams := dto.FromCoreRegisterOrder(req).ToUrlValues()

	var response dto.RegisterOrderResponse
	if err := a.sendRequest(ctx, POST, "register.do", reqParams, &response); err != nil && err != io.EOF {
		return core.RegisterOrderResponse{}, err
	}

	return response.DtoToCore(), nil
}

// AuthOrder — авторизация нового заказа (двухэтапный платёж).
// Отправляет запрос в endpoint `registerPreAuth.do`.
// В этом случае средства блокируются, но не списываются.
// Чтобы завершить платёж, необходимо вызвать DepositOrder.
func (a *api) AuthOrder(ctx context.Context, req core.RegisterOrderRequest) (core.RegisterOrderResponse, error) {
	reqParams := dto.FromCoreRegisterOrder(req).ToUrlValues()

	var response dto.RegisterOrderResponse
	if err := a.sendRequest(ctx, POST, "registerPreAuth.do", reqParams, &response); err != nil && err != io.EOF {
		return core.RegisterOrderResponse{}, err
	}

	return response.DtoToCore(), nil
}

// DepositOrder — списание средств по заказу (capture).
// Endpoint: `deposit.do`.
// Применяется только к заказам, находящимся в статусе APPROVED (после AuthOrder).
// Аргументы:
//   - req — структура с ID заказа, суммой и валютой.
//
// Возвращает Response с кодом результата.
func (a *api) DepositOrder(ctx context.Context, req core.DepositOrderRequest) (core.Response, error) {
	reqParams := dto.FromCoreDepositOrder(req).ToUrlValues()

	var response dto.Response
	if err := a.sendRequest(ctx, POST, "deposit.do", reqParams, &response); err != nil && err != io.EOF {
		return core.Response{}, err
	}

	return response.DtoToCore(), nil
}

// RefundOrder — возврат средств по завершённому заказу.
// Endpoint: `refund.do`.
// Используется, если деньги уже списаны и нужно вернуть клиенту всю или часть суммы.
func (a *api) RefundOrder(ctx context.Context, req core.RefundOrderRequest) (core.Response, error) {
	reqParams := dto.FromCoreRefundOrder(req).ToUrlValues()

	var response dto.Response
	if err := a.sendRequest(ctx, POST, "refund.do", reqParams, &response); err != nil && err != io.EOF {
		return core.Response{}, err
	}
	return response.DtoToCore(), nil
}

// ReversalOrder — реверсирование авторизованного заказа (reverse).
// Endpoint: `reverse.do`.
// Используется, если заказ был авторизован (средства заблокированы),
// но списание ещё не произошло. Фактически снимает блокировку.
func (a *api) ReversalOrder(ctx context.Context, req core.ReversalOrderRequest) (core.Response, error) {
	reqParams := dto.FromCoreReversalOrder(req).ToUrlValues()

	var response dto.Response
	if err := a.sendRequest(ctx, POST, "reverse.do", reqParams, &response); err != nil && err != io.EOF {
		return core.Response{}, err
	}
	return response.DtoToCore(), nil
}

// CancelOrder — отмена/отклонение заказа.
// Endpoint: `decline.do`.
// Используется до завершения оплаты, если заказ необходимо аннулировать.
//
// ВАЖНО: Для вызова метода у пользователя/продавца должны быть соответствующие права
// на отклонение заказов (назначаются в настройках платёжной системы)
func (a *api) CancelOrder(ctx context.Context, req core.CancelOrderRequest) (core.Response, error) {
	reqParams := dto.FromCoreCancelOrder(req).ToUrlValues()

	var response dto.Response
	if err := a.sendRequest(ctx, POST, "decline.do", reqParams, &response); err != nil && err != io.EOF {
		return core.Response{}, err
	}
	return response.DtoToCore(), nil
}

// GetOrderStatus — получение расширенного статуса заказа.
// Endpoint: `getOrderStatusExtended.do`.
// Возвращает OrderStatusResponse, содержащий текущее состояние заказа
// (например: REGISTERED, AUTHORIZED, DEPOSITED, REFUNDED, REVERSED и т.д.).
func (a *api) GetOrderStatus(ctx context.Context, req core.OrderStatusRequest) (core.OrderStatusResponse, error) {
	reqParams := dto.FromCoreOrderStatus(req).ToUrlValues()

	var response dto.OrderStatusResponse
	if err := a.sendRequest(ctx, GET, "getOrderStatusExtended.do", reqParams, &response); err != nil {
		return core.OrderStatusResponse{}, err
	}
	return response.DtoToCore(), nil
}
