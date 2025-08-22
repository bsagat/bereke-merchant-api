package dto

import (
	money "github.com/bsagat/bereke-merchant-api/currency"
	"github.com/bsagat/bereke-merchant-api/models/core"
)

func FromCoreOrder(req core.Order) Order {
	return Order{
		OrderNumber:        req.OrderNumber,
		Amount:             money.ToMinorUnit(req.Amount, req.Currency),
		Currency:           req.Currency,
		ReturnURL:          req.ReturnURL,
		FailURL:            req.FailURL,
		Description:        req.Description,
		Language:           req.Language,
		SessionTimeoutSecs: req.SessionTimeoutSecs,
		ExpirationDate:     req.ExpirationDate,
		Features:           req.Features,
		FeeInput:           req.FeeInput,
	}
}

func FromCoreRegisterOrder(req core.RegisterOrderRequest) RegisterOrderRequest {
	return RegisterOrderRequest{
		Order:              FromCoreOrder(req.Order),
		IP:                 req.IP,
		ClientId:           req.ClientId,
		CardholderName:     req.CardholderName,
		Email:              req.Email,
		BindingId:          req.BindingId,
		PostAddress:        req.PostAddress,
		DynamicCallbackURL: req.DynamicCallbackURL,
	}
}

func FromCoreDepositOrder(req core.DepositOrderRequest) DepositOrderRequest {
	return DepositOrderRequest{
		OrderID:  req.OrderID,
		Amount:   money.ToMinorUnit(req.Amount, req.Currency),
		Language: req.Language,
		Currency: req.Currency,
	}
}

func FromCoreRefundOrder(req core.RefundOrderRequest) RefundOrderRequest {
	return RefundOrderRequest{
		OrderID:                 req.OrderID,
		Amount:                  money.ToMinorUnit(req.Amount, req.Currency),
		Currency:                req.Currency,
		Language:                req.Language,
		JSONParams:              req.JSONParams,
		ExpectedDepositedAmount: req.ExpectedDepositedAmount,
		ExternalRefundID:        req.ExternalRefundID,
	}
}

func FromCoreOrderStatus(req core.OrderStatusRequest) OrderStatusRequest {
	return OrderStatusRequest{
		OrderID:       req.OrderID,
		OrderNumber:   req.OrderNumber,
		Language:      req.Language,
		MerchantLogin: req.MerchantLogin,
	}
}

func FromCoreReversalOrder(req core.ReversalOrderRequest) ReversalOrderRequest {
	return ReversalOrderRequest{
		OrderID:       req.OrderID,
		OrderNumber:   req.OrderNumber,
		Amount:        money.ToMinorUnit(req.Amount, req.Currency),
		Currency:      req.Currency,
		Language:      req.Language,
		JSONParams:    req.JSONParams,
		MerchantLogin: req.MerchantLogin,
	}
}

func FromCoreCancelOrder(req core.CancelOrderRequest) CancelOrderRequest {
	return CancelOrderRequest{
		OrderID:     req.OrderID,
		OrderNumber: req.OrderNumber,
		Language:    req.Language,
	}
}
