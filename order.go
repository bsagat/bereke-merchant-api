package bereke_merchant

import (
	"context"
	"io"
)

type ClientInfo struct {
	IP             string `json:"ip,omitempty"`
	ClientId       string `json:"clientId,omitempty"`
	CardholderName string `json:"cardholderName,omitempty"`
	Email          string `json:"email,omitempty"`
}

type RegisterOrderRequest struct {
	Order
	ClientInfo
	PaymentDetails
	AdditionalInfo
}

type RegisterOrderResponse struct {
	Response
	FormURL string `json:"formUrl,omitempty"` // URL для редиректа покупателя на оплату
	OrderID string `json:"orderId,omitempty"` // ID заказа в шлюзе (UUID или строка до 36 символов)
}

func (a *api) RegisterOrder(ctx context.Context, req RegisterOrderRequest) (RegisterOrderResponse, error) {
	reqParams, err := req.ToUrlValues()
	if err != nil {
		return RegisterOrderResponse{}, err
	}

	var response RegisterOrderResponse
	if err = a.do(ctx, GET, "register.do", reqParams, &response); err != nil && err != io.EOF {
		return RegisterOrderResponse{}, err
	}
	return response, nil
}

type RefundOrderRequest struct {
	OrderID                 string `json:"orderId"`                           // Номер заказа в платежном шлюзе
	Amount                  int    `json:"amount"`                            // Сумма возврата в минимальных единицах валюты (например, в копейках)
	Language                string `json:"language,omitempty"`                // Ключ языка по ISO 639-1
	JSONParams              string `json:"jsonParams,omitempty"`              // Дополнительные данные в формате JSON
	ExpectedDepositedAmount int    `json:"expectedDepositedAmount,omitempty"` // Для определения повторного запроса
	ExternalRefundID        string `json:"externalRefundId,omitempty"`        // Идентификатор возврата
	Currency                int    `json:"currency,omitempty"`                // Код валюты платежа ISO 4217
}

type Response struct {
	ErrorCode    string `json:"errorCode,omitempty"`    // 0 = успех, 1-99 = ошибка
	ErrorMessage string `json:"errorMessage,omitempty"` // Текст ошибки (язык зависит от запроса)
}

func (a *api) RefundOrder(ctx context.Context, req RefundOrderRequest) (Response, error) {
	reqParams, err := req.ToUrlValues()
	if err != nil {
		return Response{}, err
	}

	var response Response
	if err := a.do(ctx, POST, "refund.do", reqParams, &response); err != nil && err != io.EOF {
		return Response{}, err
	}
	return response, nil
}

type ReversalOrderRequest struct {
	OrderID       string `json:"orderId"`                 // Номер заказа в платежном шлюзе
	Language      string `json:"language,omitempty"`      // Ключ языка по ISO 639-1
	OrderNumber   string `json:"orderNumber,omitempty"`   // Номер заказа в системе мерчанта
	JSONParams    string `json:"jsonParams,omitempty"`    // Дополнительные данные в формате JSON
	MerchantLogin string `json:"merchantLogin,omitempty"` // Логин мерчанта для отмены от его имени
	Amount        int    `json:"amount,omitempty"`        // Сумма отмены в минимальных единицах валюты
	Currency      int    `json:"currency,omitempty"`      // Код валюты платежа ISO 4217
}

func (a *api) ReversalOrder(ctx context.Context, req ReversalOrderRequest) (Response, error) {
	reqParams, err := req.ToUrlValues()
	if err != nil {
		return Response{}, err
	}

	var response Response
	if err := a.do(ctx, POST, "reverse.do", reqParams, &response); err != nil && err != io.EOF {
		return Response{}, err
	}
	return response, nil
}

type CancelOrderRequest struct {
	OrderID     string `json:"orderId"`     // Номер заказа в платежном шлюзе
	Language    string `json:"language"`    // Ключ языка по ISO 639-1
	OrderNumber string `json:"orderNumber"` // Номер заказа в системе мерчанта
}

// Отклонить можно только заказ, который не был завершен.
func (a *api) CancelOrder(ctx context.Context, req CancelOrderRequest) (Response, error) {
	reqParams, err := req.ToUrlValues()
	if err != nil {
		return Response{}, err
	}

	var response Response
	if err := a.do(ctx, POST, "decline.do", reqParams, &response); err != nil && err != io.EOF {
		return Response{}, err
	}
	return response, nil
}

type Order struct {
	OrderNumber        string         `json:"orderNumber"`
	Amount             int            `json:"amount"`
	Currency           int            `json:"currency,omitempty"`
	ReturnURL          string         `json:"returnUrl"`
	FailURL            string         `json:"failUrl,omitempty"`
	Description        string         `json:"description,omitempty"`
	Language           string         `json:"language,omitempty"`
	SessionTimeoutSecs int            `json:"sessionTimeoutSecs,omitempty"`
	ExpirationDate     string         `json:"expirationDate,omitempty"`
	Features           PaymentFeature `json:"features,omitempty"`
	FeeInput           int            `json:"feeInput,omitempty"`
}

type PaymentDetails struct {
	BindingId string `json:"bindingId,omitempty"`
}

type AdditionalInfo struct {
	PostAddress        string `json:"postAddress,omitempty"`
	DynamicCallbackURL string `json:"dynamicCallbackUrl,omitempty"`
}

type PaymentFeature string

const (
	AUTO_PAYMENT         PaymentFeature = "AUTO_PAYMENT"         // Платеж без проверки подлинности владельца карты
	VERIFY               PaymentFeature = "VERIFY"               // Верификация владельца карты без списания средств
	FORCE_TDS            PaymentFeature = "FORCE_TDS"            // Принудительное проведение платежа с использованием 3-D Secure
	FORCE_SSL            PaymentFeature = "FORCE_SSL"            // Принудительное проведение платежа через SSL
	FORCE_FULL_TDS       PaymentFeature = "FORCE_FULL_TDS"       // Полное проведение аутентификации с 3-D Secure
	FORCE_CREATE_BINDING PaymentFeature = "FORCE_CREATE_BINDING" // Принудительное создание связки
)
