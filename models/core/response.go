package core

import "github.com/bsagat/bereke-merchant-api/models/types"

// ------------------------------------------------------------
// Базовый ответ API
// ------------------------------------------------------------

type Response struct {
	// Код ошибки:
	// 0  — успех
	// 1-99 — ошибка
	ErrorCode int

	// Человеко-читаемое описание ошибки
	// Язык текста зависит от параметров запроса
	ErrorMessage string
}

// ------------------------------------------------------------
// Ответ при регистрации заказа
// ------------------------------------------------------------

type RegisterOrderResponse struct {
	Response

	// URL, на который нужно перенаправить покупателя для оплаты
	FormURL string

	// Идентификатор заказа в платёжном шлюзе (до 36 символов)
	OrderID string
}

// ------------------------------------------------------------
// Ответ со статусом заказа
// ------------------------------------------------------------

type OrderStatusResponse struct {
	// --- Информация об ошибке ---
	Response

	// --- Данные заказа ---
	OrderID     string            // ID заказа в шлюзе
	OrderNumber string            // ID заказа в системе мерчанта
	OrderStatus types.OrderStatus // Статус заказа (enum)

	// --- Ответ банка ---
	ActionCode            int    // Код ответа банка
	ActionCodeDescription string // Текстовое описание кода
	AuthRefNum            string // Номер авторизации
	TerminalID            string // ID терминала банка

	// --- Финансовая информация ---
	Amount   float64 // Сумма заказа в основных единицах валюты
	Currency int     // Код валюты (ISO 4217)

	// --- Временные метки (Unix ms) ---
	Date          int64 // Дата создания заказа
	DepositedDate int64 // Дата депозита
	RefundedDate  int64 // Дата возврата
	ReversedDate  int64 // Дата сторнирования
	AuthDateTime  int64 // Дата авторизации

	// --- Дополнительная информация ---
	BindingInfo       BindingInfo       // Информация о привязке карты
	PaymentAmountInfo PaymentAmountInfo // Детализация по суммам
	BankInfo          BankInfo          // Данные о банке
	CardInfo          CardInfo          // Данные о карте

	// --- Флаги и способы оплаты ---
	PaymentWay string // Способ оплаты (card, sbp и т.д.)
	Refund     bool   // true — если был возврат
}

// ------------------------------------------------------------
// Вложенные структуры
// ------------------------------------------------------------

// Информация о привязке карты
type BindingInfo struct {
	ClientID     string // ID клиента в системе мерчанта (до 255 символов)
	BindingID    string // ID привязки карты (до 255 символов)
	AuthDateTime int64  // Дата авторизации (мс от Unix epoch)
	AuthRefNum   string // Номер авторизации (до 24 символов)
	TerminalID   string // ID терминала (до 10 символов)
}

// Информация о суммах транзакции
type PaymentAmountInfo struct {
	ApprovedAmount  int64  // Одобренная сумма
	DepositedAmount int64  // Депонированная сумма
	RefundedAmount  int64  // Возвращенная сумма
	PaymentState    string // Состояние платежа (CREATED, APPROVED и т.д.)
}

// Информация о банке
type BankInfo struct {
	BankName        string // Название банка (до 50 символов)
	BankCountryCode int    // Код страны банка (до 4 символов)
	BankCountryName string // Название страны банка (до 160 символов)
}

// Информация о карте
type CardInfo struct {
	MaskedPan      string // Маскированный номер карты (до 19 символов)
	Expiration     string // Дата окончания (YYMMDD, до 6 символов)
	CardholderName string // Имя держателя карты (до 50 символов)
	Pan            string // Полный номер карты (до 19 символов)
	ApprovalCode   string // Код авторизации (до 6 символов)
}
