package dto

import "github.com/bsagat/bereke-merchant-api/models/types"

// ------------------------------------------------------------
// Базовая структура заказа
// ------------------------------------------------------------

type Order struct {
	// Идентификатор заказа в системе мерчанта (вашей системе)
	OrderNumber string `json:"orderNumber"`

	// Сумма заказа в минимальных единицах валюты
	// Например: 1000 = 10.00 RUB
	Amount int `json:"amount"`

	// Код валюты по стандарту ISO 4217
	// Например: 643 — RUB, 840 — USD
	Currency int `json:"currency,omitempty"`

	// URL, на который будет перенаправлен пользователь после успешной оплаты
	ReturnURL string `json:"returnUrl"`

	// URL, на который будет перенаправлен пользователь в случае ошибки или отмены
	FailURL string `json:"failUrl,omitempty"`

	// Описание заказа (может отображаться пользователю или в отчетах)
	Description string `json:"description,omitempty"`

	// Язык интерфейса оплаты (ISO 639-1: ru, en, by, kz, kk)
	Language string `json:"language,omitempty"`

	// Время жизни сессии оплаты в секундах
	SessionTimeoutSecs int `json:"sessionTimeoutSecs,omitempty"`

	// Дата и время, когда заказ перестанет быть доступен для оплаты
	// Формат: YYYY-MM-DDThh:mm:ss
	ExpirationDate string `json:"expirationDate,omitempty"`

	// Дополнительные возможности платежа (определяется в types.PaymentFeature)
	Features types.PaymentFeature `json:"features,omitempty"`

	// Комиссия, которую вводит мерчант (если применяется)
	FeeInput int `json:"feeInput,omitempty"`
}

// ------------------------------------------------------------
// Запрос на регистрацию заказа
// ------------------------------------------------------------

type RegisterOrderRequest struct {
	Order

	// Информация о клиенте
	IP             string `json:"ip,omitempty"`             // IP-адрес клиента
	ClientId       string `json:"clientId,omitempty"`       // Идентификатор клиента в вашей системе
	CardholderName string `json:"cardholderName,omitempty"` // Имя держателя карты
	Email          string `json:"email,omitempty"`          // Email клиента

	// Детали платежа
	BindingId string `json:"bindingId,omitempty"` // Идентификатор привязки карты

	// Дополнительная информация
	PostAddress        string `json:"postAddress,omitempty"`        // Почтовый адрес клиента
	DynamicCallbackURL string `json:"dynamicCallbackUrl,omitempty"` // URL для динамического колбэка (заменяет статичный)
}

// ------------------------------------------------------------
// Запрос на возврат средств
// ------------------------------------------------------------

type RefundOrderRequest struct {
	OrderID string `json:"orderId"` // Номер заказа в платежном шлюзе

	// Сумма возврата в минимальных единицах валюты
	// Пример: 500 = 5.00 RUB
	Amount int `json:"amount"`

	// Код валюты платежа (ISO 4217)
	Currency int `json:"currency,omitempty"`

	// Язык ответа (ISO 639-1)
	Language string `json:"language,omitempty"`

	// Дополнительные данные в формате JSON
	JSONParams string `json:"jsonParams,omitempty"`

	// Для определения повторного запроса возврата (чтобы избежать дублирования)
	ExpectedDepositedAmount int `json:"expectedDepositedAmount,omitempty"`

	// Внешний идентификатор возврата (уникальный в системе мерчанта)
	ExternalRefundID string `json:"externalRefundId,omitempty"`
}

// ------------------------------------------------------------
// Запрос на получение статуса заказа
// ------------------------------------------------------------

type OrderStatusRequest struct {
	// Обязательное поле — либо orderId, либо orderNumber.
	// Приоритет — orderId.
	OrderID     string `json:"orderId,omitempty"`     // Идентификатор заказа в шлюзе [1..36]
	OrderNumber string `json:"orderNumber,omitempty"` // Идентификатор заказа у мерчанта [1..36]

	// Необязательные поля
	Language      string `json:"language,omitempty"`      // Язык ответа (ISO 639-1: ru,en,by,kz,kk)
	MerchantLogin string `json:"merchantLogin,omitempty"` // Логин мерчанта [1..255]
}

// ------------------------------------------------------------
// Запрос на сторнирование (отмену проведенного платежа)
// ------------------------------------------------------------

type ReversalOrderRequest struct {
	OrderID string `json:"orderId"` // Номер заказа в платежном шлюзе

	// Дополнительно можно указать:
	OrderNumber   string `json:"orderNumber,omitempty"`   // Номер заказа в системе мерчанта
	Amount        int    `json:"amount,omitempty"`        // Сумма отмены
	Currency      int    `json:"currency,omitempty"`      // Код валюты платежа (ISO 4217)
	Language      string `json:"language,omitempty"`      // Язык ответа
	JSONParams    string `json:"jsonParams,omitempty"`    // Дополнительные параметры в JSON
	MerchantLogin string `json:"merchantLogin,omitempty"` // Логин мерчанта (для отмены от его имени)
}

// ------------------------------------------------------------
// Запрос на отмену заказа
// ------------------------------------------------------------

type CancelOrderRequest struct {
	OrderID     string `json:"orderId"`     // Номер заказа в платежном шлюзе
	OrderNumber string `json:"orderNumber"` // Номер заказа в системе мерчанта
	Language    string `json:"language"`    // Язык ответа (ISO 639-1)
}
