package dto

import "github.com/bsagat/bereke-merchant-api/models/types"

// ------------------------------------------------------------
// Базовый ответ API
// ------------------------------------------------------------

type Response struct {
	// Код ошибки:
	// 0  — успех
	// 1-99 — ошибка
	ErrorCode string `json:"errorCode,omitempty"`

	// Человеко-читаемое описание ошибки
	// Язык текста зависит от параметров запроса
	ErrorMessage string `json:"errorMessage,omitempty"`
}

// ------------------------------------------------------------
// Ответ при регистрации заказа
// ------------------------------------------------------------

type RegisterOrderResponse struct {
	Response

	// URL, на который нужно перенаправить покупателя для оплаты
	FormURL string `json:"formUrl,omitempty"`

	// Идентификатор заказа в платёжном шлюзе (до 36 символов)
	OrderID string `json:"orderId,omitempty"`
}

// ------------------------------------------------------------
// Ответ со статусом заказа
// ------------------------------------------------------------

type OrderStatusResponse struct {
	// --- Информация об ошибке ---
	Response

	// --- Данные заказа ---
	OrderID     string            `json:"orderId,omitempty"`     // ID заказа в шлюзе
	OrderNumber string            `json:"orderNumber,omitempty"` // ID заказа в системе мерчанта
	OrderStatus types.OrderStatus `json:"orderStatus,omitempty"` // Статус заказа (enum)

	// --- Ответ банка ---
	ActionCode            int    `json:"actionCode,omitempty"`            // Код ответа банка
	ActionCodeDescription string `json:"actionCodeDescription,omitempty"` // Текстовое описание кода
	AuthRefNum            string `json:"authRefNum,omitempty"`            // Номер авторизации
	TerminalID            string `json:"terminalId,omitempty"`            // ID терминала банка

	// --- Финансовая информация ---
	MinorAmount int    `json:"amount,omitempty"`   // Сумма заказа в минорных единицах валюты (например, копейки)
	Currency    string `json:"currency,omitempty"` // Код валюты (ISO 4217)

	// --- Временные метки (Unix ms) ---
	Date          int64 `json:"date,omitempty"`          // Дата создания заказа
	DepositedDate int64 `json:"depositedDate,omitempty"` // Дата депозита
	RefundedDate  int64 `json:"refundedDate,omitempty"`  // Дата возврата
	ReversedDate  int64 `json:"reversedDate,omitempty"`  // Дата сторнирования
	AuthDateTime  int64 `json:"authDateTime,omitempty"`  // Дата авторизации

	// --- Дополнительная информация ---
	BindingInfo       BindingInfo       `json:"bindingInfo,omitempty"`       // Информация о привязке карты
	PaymentAmountInfo PaymentAmountInfo `json:"paymentAmountInfo,omitempty"` // Детализация по суммам
	BankInfo          BankInfo          `json:"bankInfo,omitempty"`          // Данные о банке
	CardInfo          CardInfo          `json:"cardAuthInfo,omitempty"`      // Данные о карте

	// --- Флаги и способы оплаты ---
	PaymentWay string `json:"paymentWay,omitempty"` // Способ оплаты (card, sbp и т.д.)
	Refund     bool   `json:"refund,omitempty"`     // true — если был возврат
}

// ------------------------------------------------------------
// Вложенные структуры
// ------------------------------------------------------------

// Информация о привязке карты
type BindingInfo struct {
	ClientID     string `json:"clientId,omitempty"`     // ID клиента в системе мерчанта (до 255 символов)
	BindingID    string `json:"bindingId,omitempty"`    // ID привязки карты (до 255 символов)
	AuthDateTime int64  `json:"authDateTime,omitempty"` // Дата авторизации (мс от Unix epoch)
	AuthRefNum   string `json:"authRefNum,omitempty"`   // Номер авторизации (до 24 символов)
	TerminalID   string `json:"terminalId,omitempty"`   // ID терминала (до 10 символов)
}

// Информация о суммах транзакции
type PaymentAmountInfo struct {
	ApprovedAmount  int64  `json:"approvedAmount,omitempty"`  // Одобренная сумма
	DepositedAmount int64  `json:"depositedAmount,omitempty"` // Депонированная сумма
	RefundedAmount  int64  `json:"refundedAmount,omitempty"`  // Возвращенная сумма
	PaymentState    string `json:"paymentState,omitempty"`    // Состояние платежа (CREATED, APPROVED и т.д.)
}

// Информация о банке
type BankInfo struct {
	BankName        string `json:"bankName,omitempty"`        // Название банка (до 50 символов)
	BankCountryCode string `json:"bankCountryCode,omitempty"` // Код страны банка (до 4 символов)
	BankCountryName string `json:"bankCountryName,omitempty"` // Название страны банка (до 160 символов)
}

// Информация о карте
type CardInfo struct {
	MaskedPan      string `json:"maskedPan,omitempty"`      // Маскированный номер карты (до 19 символов)
	Expiration     string `json:"expiration,omitempty"`     // Дата окончания (YYMMDD, до 6 символов)
	CardholderName string `json:"cardholderName,omitempty"` // Имя держателя карты (до 50 символов)
	Pan            string `json:"pan,omitempty"`            // Полный номер карты (до 19 символов)
	ApprovalCode   string `json:"approvalCode,omitempty"`   // Код авторизации (до 6 символов)
}
