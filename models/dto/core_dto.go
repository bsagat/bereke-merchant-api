package dto

import (
	money "github.com/bsagat/bereke-merchant-api/currency"
	"github.com/bsagat/bereke-merchant-api/models/core"
)

func FromCoreOrder(req core.Order) (Order, error) {
	minorAmount, err := money.ToMinorUnit(req.Amount, req.Currency)
	if err != nil {
		return Order{}, err
	}

	return Order{
		OrderNumber:        req.OrderNumber,
		Amount:             minorAmount,
		Currency:           req.Currency,
		ReturnURL:          req.ReturnURL,
		FailURL:            req.FailURL,
		Description:        req.Description,
		Language:           req.Language,
		SessionTimeoutSecs: req.SessionTimeoutSecs,
		ExpirationDate:     req.ExpirationDate,
		Features:           req.Features,
		FeeInput:           req.FeeInput,
	}, nil
}

func FromCoreRegisterOrder(req core.RegisterOrderRequest) (RegisterOrderRequest, error) {
	dtoOrder, err := FromCoreOrder(req.Order)
	if err != nil {
		return RegisterOrderRequest{}, err
	}

	return RegisterOrderRequest{
		Order:              dtoOrder,
		IP:                 req.IP,
		ClientId:           req.ClientId,
		CardholderName:     req.CardholderName,
		Email:              req.Email,
		BindingId:          req.BindingId,
		PostAddress:        req.PostAddress,
		DynamicCallbackURL: req.DynamicCallbackURL,
	}, nil
}

func FromCoreRefundOrder(req core.RefundOrderRequest) (RefundOrderRequest, error) {
	minorAmount, err := money.ToMinorUnit(req.Amount, req.Currency)
	if err != nil {
		return RefundOrderRequest{}, err
	}

	return RefundOrderRequest{
		OrderID:                 req.OrderID,
		Amount:                  minorAmount,
		Currency:                req.Currency,
		Language:                req.Language,
		JSONParams:              req.JSONParams,
		ExpectedDepositedAmount: req.ExpectedDepositedAmount,
		ExternalRefundID:        req.ExternalRefundID,
	}, nil
}

func FromCoreOrderStatus(req core.OrderStatusRequest) OrderStatusRequest {
	return OrderStatusRequest{
		OrderID:       req.OrderID,
		OrderNumber:   req.OrderNumber,
		Language:      req.Language,
		MerchantLogin: req.MerchantLogin,
	}
}

func FromCoreReversalOrder(req core.ReversalOrderRequest) (ReversalOrderRequest, error) {
	minorAmount, err := money.ToMinorUnit(req.Amount, req.Currency)
	if err != nil {
		return ReversalOrderRequest{}, err
	}

	return ReversalOrderRequest{
		OrderID:       req.OrderID,
		OrderNumber:   req.OrderNumber,
		Amount:        minorAmount,
		Currency:      req.Currency,
		Language:      req.Language,
		JSONParams:    req.JSONParams,
		MerchantLogin: req.MerchantLogin,
	}, nil
}

func FromCoreCancelOrder(req core.CancelOrderRequest) CancelOrderRequest {
	return CancelOrderRequest{
		OrderID:     req.OrderID,
		OrderNumber: req.OrderNumber,
		Language:    req.Language,
	}
}
