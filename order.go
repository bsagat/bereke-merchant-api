package bereke_merchant

import (
	"context"
	"io"

	"github.com/bsagat/bereke-merchant-api/models/core"
	"github.com/bsagat/bereke-merchant-api/models/dto"
)

// Регистрация нового заказа в платежной системе
func (a *api) RegisterOrder(ctx context.Context, req core.RegisterOrderRequest) (core.RegisterOrderResponse, error) {
	reqParams := dto.FromCoreRegisterOrder(req).ToUrlValues()

	var response dto.RegisterOrderResponse
	if err := a.sendRequest(ctx, POST, "register.do", reqParams, &response); err != nil && err != io.EOF {
		return core.RegisterOrderResponse{}, err
	}

	return response.DtoToCore(), nil
}

// Авторизация нового заказа в платежной системе
func (a *api) AuthOrder(ctx context.Context, req core.RegisterOrderRequest) (core.RegisterOrderResponse, error) {
	reqParams := dto.FromCoreRegisterOrder(req).ToUrlValues()

	var response dto.RegisterOrderResponse
	if err := a.sendRequest(ctx, POST, "registerPreAuth.do", reqParams, &response); err != nil && err != io.EOF {
		return core.RegisterOrderResponse{}, err
	}

	return response.DtoToCore(), nil
}

// Списание средств по заказу со счета
// ПРИМЕЧАНИЕ:
// Заказ должен быть в статусе (APPROVED)

func (a *api) DepositOrder(ctx context.Context, req core.DepositOrderRequest) (core.Response, error) {
	reqParams := dto.FromCoreDepositOrder(req).ToUrlValues()

	var response dto.Response
	if err := a.sendRequest(ctx, POST, "deposit.do", reqParams, &response); err != nil && err != io.EOF {
		return core.Response{}, err
	}

	return response.DtoToCore(), nil
}

// Возврат средств по заказу
func (a *api) RefundOrder(ctx context.Context, req core.RefundOrderRequest) (core.Response, error) {
	reqParams := dto.FromCoreRefundOrder(req).ToUrlValues()

	var response dto.Response
	if err := a.sendRequest(ctx, POST, "refund.do", reqParams, &response); err != nil && err != io.EOF {
		return core.Response{}, err
	}
	return response.DtoToCore(), nil
}

// Реверсирование (отмена после авторизации, но до завершения)
func (a *api) ReversalOrder(ctx context.Context, req core.ReversalOrderRequest) (core.Response, error) {
	reqParams := dto.FromCoreReversalOrder(req).ToUrlValues()

	var response dto.Response
	if err := a.sendRequest(ctx, POST, "reverse.do", reqParams, &response); err != nil && err != io.EOF {
		return core.Response{}, err
	}
	return response.DtoToCore(), nil
}

// Отклонение заказа, если он не был завершен
func (a *api) CancelOrder(ctx context.Context, req core.CancelOrderRequest) (core.Response, error) {
	reqParams := dto.FromCoreCancelOrder(req).ToUrlValues()

	var response dto.Response
	if err := a.sendRequest(ctx, POST, "decline.do", reqParams, &response); err != nil && err != io.EOF {
		return core.Response{}, err
	}
	return response.DtoToCore(), nil
}

// Получение расширенного статуса заказа
func (a *api) GetOrderStatus(ctx context.Context, req core.OrderStatusRequest) (core.OrderStatusResponse, error) {
	reqParams := dto.FromCoreOrderStatus(req).ToUrlValues()

	var response dto.OrderStatusResponse
	if err := a.sendRequest(ctx, GET, "getOrderStatusExtended.do", reqParams, &response); err != nil {
		return core.OrderStatusResponse{}, err
	}
	return response.DtoToCore(), nil
}
