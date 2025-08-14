package bereke_merchant

import (
	"context"
	"io"

	"github.com/bsagat/bereke-merchant-api/models/core"
	"github.com/bsagat/bereke-merchant-api/models/dto"
)

// Регистрация нового заказа в платежной системе
func (a *api) RegisterOrder(ctx context.Context, req core.RegisterOrderRequest) (core.RegisterOrderResponse, error) {
	reg, err := dto.FromCoreRegisterOrder(req)
	if err != nil {
		return core.RegisterOrderResponse{}, err
	}

	reqParams, err := reg.ToUrlValues()
	if err != nil {
		return core.RegisterOrderResponse{}, err
	}

	var response dto.RegisterOrderResponse
	if err = a.sendRequest(ctx, GET, "register.do", reqParams, &response); err != nil && err != io.EOF {
		return core.RegisterOrderResponse{}, err
	}

	return response.DtoToCore(), nil
}

// Возврат средств по заказу
func (a *api) RefundOrder(ctx context.Context, req core.RefundOrderRequest) (core.Response, error) {
	ref, err := dto.FromCoreRefundOrder(req)
	if err != nil {
		return core.Response{}, err
	}

	reqParams, err := ref.ToUrlValues()
	if err != nil {
		return core.Response{}, err
	}

	var response dto.Response
	if err := a.sendRequest(ctx, POST, "refund.do", reqParams, &response); err != nil && err != io.EOF {
		return core.Response{}, err
	}
	return response.DtoToCore(), nil
}

// Реверсирование (отмена после авторизации, но до завершения)
func (a *api) ReversalOrder(ctx context.Context, req core.ReversalOrderRequest) (core.Response, error) {
	rev, err := dto.FromCoreReversalOrder(req)
	if err != nil {
		return core.Response{}, err
	}

	reqParams, err := rev.ToUrlValues()
	if err != nil {
		return core.Response{}, err
	}

	var response dto.Response
	if err := a.sendRequest(ctx, POST, "reverse.do", reqParams, &response); err != nil && err != io.EOF {
		return core.Response{}, err
	}
	return response.DtoToCore(), nil
}

// Отклонение заказа, если он не был завершен
func (a *api) CancelOrder(ctx context.Context, req core.CancelOrderRequest) (core.Response, error) {
	reqParams, err := dto.FromCoreCancelOrder(req).ToUrlValues()
	if err != nil {
		return core.Response{}, err
	}

	var response dto.Response
	if err := a.sendRequest(ctx, POST, "decline.do", reqParams, &response); err != nil && err != io.EOF {
		return core.Response{}, err
	}
	return response.DtoToCore(), nil
}

// Получение расширенного статуса заказа
func (a *api) GetOrderStatus(ctx context.Context, req core.OrderStatusRequest) (core.OrderStatusResponse, error) {
	reqParams, err := dto.FromCoreOrderStatus(req).ToUrlValues()
	if err != nil {
		return core.OrderStatusResponse{}, err
	}

	var response dto.OrderStatusResponse
	if err = a.sendRequest(ctx, GET, "getOrderStatusExtended.do", reqParams, &response); err != nil {
		return core.OrderStatusResponse{}, err
	}
	return response.DtoToCore(), nil
}
