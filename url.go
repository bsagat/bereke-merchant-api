package bereke_merchant

import (
	"net/url"
	"strconv"
)

func (r *RegisterOrderRequest) ToUrlValues() (url.Values, error) {
	values := url.Values{}

	values.Set("orderNumber", r.OrderNumber)
	values.Set("amount", strconv.Itoa(r.Amount))

	if r.Currency != 0 {
		values.Set("currency", strconv.Itoa(r.Currency))
	}
	if r.ReturnURL != "" {
		values.Set("returnUrl", r.ReturnURL)
	}
	if r.FailURL != "" {
		values.Set("failUrl", r.FailURL)
	}
	if r.DynamicCallbackURL != "" {
		values.Set("dynamicCallbackUrl", r.DynamicCallbackURL)
	}
	if r.Description != "" {
		values.Set("description", r.Description)
	}
	if r.Language != "" {
		values.Set("language", r.Language)
	}
	if r.IP != "" {
		values.Set("ip", r.IP)
	}
	if r.ClientId != "" {
		values.Set("clientId", r.ClientId)
	}
	if r.CardholderName != "" {
		values.Set("cardholderName", r.CardholderName)
	}
	if r.SessionTimeoutSecs != 0 {
		values.Set("sessionTimeoutSecs", strconv.Itoa(r.SessionTimeoutSecs))
	}
	if r.ExpirationDate != "" {
		values.Set("expirationDate", r.ExpirationDate)
	}
	if r.BindingId != "" {
		values.Set("bindingId", r.BindingId)
	}
	if r.PostAddress != "" {
		values.Set("postAddress", r.PostAddress)
	}
	if r.FeeInput != 0 {
		values.Set("feeInput", strconv.Itoa(r.FeeInput))
	}
	if r.Email != "" {
		values.Set("email", r.Email)
	}

	return values, nil
}

func (r OrderStatusRequest) ToUrlValues() (url.Values, error) {
	values := url.Values{}
	if r.OrderID != "" {
		values.Set("orderId", r.OrderID)
	}
	if r.OrderNumber != "" {
		values.Set("orderNumber", r.OrderNumber)
	}
	if r.Language != "" {
		values.Set("language", r.Language)
	}
	if r.MerchantLogin != "" {
		values.Set("merchantLogin", r.MerchantLogin)
	}
	return values, nil
}

func (r RefundOrderRequest) ToUrlValues() (url.Values, error) {
	values := url.Values{}
	values.Set("orderId", r.OrderID)
	values.Set("amount", strconv.Itoa(r.Amount))

	if r.Language != "" {
		values.Set("language", r.Language)
	}
	if r.JSONParams != "" {
		values.Set("jsonParams", r.JSONParams)
	}
	if r.ExpectedDepositedAmount != 0 {
		values.Set("expectedDepositedAmount", strconv.Itoa(r.ExpectedDepositedAmount))
	}
	if r.ExternalRefundID != "" {
		values.Set("externalRefundId", r.ExternalRefundID)
	}
	if r.Currency != 0 {
		values.Set("currency", strconv.Itoa(r.Currency))
	}

	return values, nil
}

func (r ReversalOrderRequest) ToUrlValues() (url.Values, error) {
	values := url.Values{}
	values.Set("orderId", r.OrderID)

	if r.Language != "" {
		values.Set("language", r.Language)
	}
	if r.OrderNumber != "" {
		values.Set("orderNumber", r.OrderNumber)
	}
	if r.JSONParams != "" {
		values.Set("jsonParams", r.JSONParams)
	}
	if r.MerchantLogin != "" {
		values.Set("merchantLogin", r.MerchantLogin)
	}
	if r.Amount != 0 {
		values.Set("amount", strconv.Itoa(r.Amount))
	}
	if r.Currency != 0 {
		values.Set("currency", strconv.Itoa(r.Currency))
	}

	return values, nil
}

func (r CancelOrderRequest) ToUrlValues() (url.Values, error) {
	values := url.Values{}
	values.Set("orderId", r.OrderID)
	values.Set("language", r.Language)
	values.Set("orderNumber", r.OrderNumber)

	return values, nil
}
