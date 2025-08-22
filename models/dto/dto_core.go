package dto

import (
	"strconv"

	money "github.com/bsagat/bereke-merchant-api/currency"
	"github.com/bsagat/bereke-merchant-api/models/core"
)

func (res *Response) DtoToCore() core.Response {
	convertedCode, _ := strconv.Atoi(res.ErrorCode)
	return core.Response{
		ErrorCode:    convertedCode,
		ErrorMessage: res.ErrorMessage,
	}
}

func (res *RegisterOrderResponse) DtoToCore() core.RegisterOrderResponse {
	return core.RegisterOrderResponse{
		Response: res.Response.DtoToCore(),
		FormURL:  res.FormURL,
		OrderID:  res.OrderID,
	}
}

func (res *OrderStatusResponse) DtoToCore() core.OrderStatusResponse {
	convertedCurrency := money.ToNumeric(res.Currency)

	return core.OrderStatusResponse{
		Response: res.Response.DtoToCore(),
		Amount:   money.ConvertFromMinorUnits(res.MinorAmount, convertedCurrency),
		Currency: convertedCurrency,

		OrderID:     res.OrderID,
		OrderNumber: res.OrderNumber,
		OrderStatus: res.OrderStatus,

		ActionCode:            res.ActionCode,
		ActionCodeDescription: res.ActionCodeDescription,

		AuthRefNum: res.AuthRefNum,
		TerminalID: res.TerminalID,

		Date:          res.Date,
		DepositedDate: res.DepositedDate,
		RefundedDate:  res.RefundedDate,
		ReversedDate:  res.ReversedDate,
		AuthDateTime:  res.AuthDateTime,

		BindingInfo:       res.BindingInfo.DtoToCore(),
		PaymentAmountInfo: res.PaymentAmountInfo.DtoToCore(),
		BankInfo:          res.BankInfo.DtoToCore(),
		CardInfo:          res.CardInfo.DtoToCore(),

		PaymentWay: res.PaymentWay,
		Refund:     res.Refund,
	}
}

func (res *BindingInfo) DtoToCore() core.BindingInfo {
	return core.BindingInfo{
		ClientID:     res.ClientID,
		BindingID:    res.BindingID,
		AuthDateTime: res.AuthDateTime,
		AuthRefNum:   res.AuthRefNum,
		TerminalID:   res.TerminalID,
	}
}

func (res *PaymentAmountInfo) DtoToCore() core.PaymentAmountInfo {
	return core.PaymentAmountInfo{
		ApprovedAmount:  res.ApprovedAmount,
		DepositedAmount: res.DepositedAmount,
		RefundedAmount:  res.RefundedAmount,
		PaymentState:    res.PaymentState,
	}
}

func (res *BankInfo) DtoToCore() core.BankInfo {
	countryCode, _ := strconv.Atoi(res.BankCountryCode)
	return core.BankInfo{
		BankName:        res.BankName,
		BankCountryCode: countryCode,
		BankCountryName: res.BankCountryName,
	}
}

func (res *CardInfo) DtoToCore() core.CardInfo {
	return core.CardInfo{
		MaskedPan:      res.MaskedPan,
		Expiration:     res.Expiration,
		CardholderName: res.CardholderName,
		Pan:            res.Pan,
		ApprovalCode:   res.ApprovalCode,
	}
}
