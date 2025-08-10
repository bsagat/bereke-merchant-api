package bereke_merchant

import (
	"context"
	"io"

	"github.com/bsagat/bereke-merchant-api/models/dto"
)

// Регистрация нового заказа в платежной системе
func (a *api) RegisterOrder(ctx context.Context, req dto.RegisterOrderRequest) (dto.RegisterOrderResponse, error) {
	reqParams, err := req.ToUrlValues()
	if err != nil {
		return dto.RegisterOrderResponse{}, err
	}

	var response dto.RegisterOrderResponse
	if err = a.sendRequest(ctx, GET, "register.do", reqParams, &response); err != nil && err != io.EOF {
		return dto.RegisterOrderResponse{}, err
	}

	return response, nil
}

// Возврат средств по заказу
func (a *api) RefundOrder(ctx context.Context, req dto.RefundOrderRequest) (dto.Response, error) {
	reqParams, err := req.ToUrlValues()
	if err != nil {
		return dto.Response{}, err
	}

	var response dto.Response
	if err := a.sendRequest(ctx, POST, "refund.do", reqParams, &response); err != nil && err != io.EOF {
		return dto.Response{}, err
	}
	return response, nil
}

// Реверсирование (отмена после авторизации, но до завершения)
func (a *api) ReversalOrder(ctx context.Context, req dto.ReversalOrderRequest) (dto.Response, error) {
	reqParams, err := req.ToUrlValues()
	if err != nil {
		return dto.Response{}, err
	}

	var response dto.Response
	if err := a.sendRequest(ctx, POST, "reverse.do", reqParams, &response); err != nil && err != io.EOF {
		return dto.Response{}, err
	}
	return response, nil
}

// Отклонение заказа, если он не был завершен
func (a *api) CancelOrder(ctx context.Context, req dto.CancelOrderRequest) (dto.Response, error) {
	reqParams, err := req.ToUrlValues()
	if err != nil {
		return dto.Response{}, err
	}

	var response dto.Response
	if err := a.sendRequest(ctx, POST, "decline.do", reqParams, &response); err != nil && err != io.EOF {
		return dto.Response{}, err
	}
	return response, nil
}

// Получение расширенного статуса заказа
func (a *api) GetOrderStatus(ctx context.Context, req dto.OrderStatusRequest) (dto.OrderStatusResponse, error) {
	reqParams, err := req.ToUrlValues()
	if err != nil {
		return dto.OrderStatusResponse{}, err
	}

	var response dto.OrderStatusResponse
	if err = a.sendRequest(ctx, GET, "getOrderStatusExtended.do", reqParams, &response); err != nil {
		return dto.OrderStatusResponse{}, err
	}
	return response, nil
}
