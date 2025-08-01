package bereke_merchant

import (
	"context"
)

type OrderStatusRequest struct {
	// Обязательное поле — либо orderId, либо orderNumber. Приоритет — orderId
	OrderID     string `json:"orderId,omitempty"`     // [1..36]
	OrderNumber string `json:"orderNumber,omitempty"` // [1..36]

	// Необязательные поля
	Language      string `json:"language,omitempty"`      // [2], ISO 639-1: ru,en,by,kz,kk
	MerchantLogin string `json:"merchantLogin,omitempty"` // [1..255]
}

type OrderStatusResponse struct {
	// Error info
	Response

	// Order info
	OrderID     string      `json:"orderId,omitempty"`
	OrderNumber string      `json:"orderNumber,omitempty"`
	OrderStatus OrderStatus `json:"orderStatus,omitempty"`

	// Bank response info
	ActionCode            int    `json:"actionCode,omitempty"`
	ActionCodeDescription string `json:"actionCodeDescription,omitempty"`
	AuthRefNum            string `json:"authRefNum,omitempty"`
	TerminalID            string `json:"terminalId,omitempty"`

	// Transaction info
	Amount   int    `json:"amount,omitempty"`
	Currency string `json:"currency,omitempty"`

	// Time info (Unix ms)
	Date          int64 `json:"date,omitempty"`
	DepositedDate int64 `json:"depositedDate,omitempty"`
	RefundedDate  int64 `json:"refundedDate,omitempty"`
	ReversedDate  int64 `json:"reversedDate,omitempty"`
	AuthDateTime  int64 `json:"authDateTime,omitempty"`

	// Additional info
	BindingInfo       BindingInfo       `json:"bindingInfo,omitempty"`
	PaymentAmountInfo PaymentAmountInfo `json:"paymentAmountInfo,omitempty"`
	BankInfo          BankInfo          `json:"bankInfo,omitempty"`     // Информация о банке-эмитенте карты
	CardInfo          CardInfo          `json:"cardAuthInfo,omitempty"` // Информация о карте

	// Flags
	PaymentWay string `json:"paymentWay,omitempty"`
	Refund     bool   `json:"refund,omitempty"`
}

type BindingInfo struct {
	ClientID     string `json:"clientId,omitempty"`     // до 255 символов
	BindingID    string `json:"bindingId,omitempty"`    // до 255 символов
	AuthDateTime int64  `json:"authDateTime,omitempty"` // миллисекунды от Unix epoch
	AuthRefNum   string `json:"authRefNum,omitempty"`   // до 24 символов
	TerminalID   string `json:"terminalId,omitempty"`   // до 10 символов
}

type PaymentAmountInfo struct {
	ApprovedAmount  int64  `json:"approvedAmount,omitempty"`
	DepositedAmount int64  `json:"depositedAmount,omitempty"`
	RefundedAmount  int64  `json:"refundedAmount,omitempty"`
	PaymentState    string `json:"paymentState,omitempty"` // CREATED, APPROVED, etc.
}

type BankInfo struct {
	BankName        string `json:"bankName,omitempty"`        // до 50 символов
	BankCountryCode string `json:"bankCountryCode,omitempty"` // до 4 символов
	BankCountryName string `json:"bankCountryName,omitempty"` // до 160 символов
}

type CardInfo struct {
	MaskedPan      string `json:"maskedPan,omitempty"`      // до 19 символов
	Expiration     string `json:"expiration,omitempty"`     // до 6 символов, формат YYMMDD
	CardholderName string `json:"cardholderName,omitempty"` // до 50 символов
	Pan            string `json:"pan,omitempty"`            // до 19 символов
	ApprovalCode   string `json:"approvalCode,omitempty"`   // до 6 символов
}

type OrderStatus int

const (
	OrderStatusRegistered OrderStatus = iota // 0 - заказ зарегистрирован, но не оплачен
	OrderStatusAuthorized                    // 1 - заказ только авторизован и еще не завершен (для двухстадийных платежей)
	OrderStatusCompleted                     // 2 - заказ авторизован и завершен
	OrderStatusCancelled                     // 3 - авторизация отменена
	OrderStatusRefunded                      // 4 - по транзакции была проведена операция возврата
	OrderStatusPending                       // 5 - инициирована авторизация через ACS банка-эмитента
	OrderStatusDeclined                      // 6 - авторизация отклонена
	OrderStatusWaiting                       // 7 - ожидание оплаты заказа
	OrderStatusPartial                       // 8 - промежуточное завершение для многократного частичного завершения
)

func (a *api) OrderStatus(ctx context.Context, req OrderStatusRequest) (OrderStatusResponse, error) {
	reqParams, err := req.ToUrlValues()
	if err != nil {
		return OrderStatusResponse{}, err
	}

	var response OrderStatusResponse
	if err = a.do(ctx, GET, "getOrderStatusExtended.do", reqParams, &response); err != nil {
		return OrderStatusResponse{}, err
	}
	return response, nil
}
