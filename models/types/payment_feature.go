package types

type PaymentFeature string

const (
	AUTO_PAYMENT         PaymentFeature = "AUTO_PAYMENT"         // Платеж без проверки подлинности владельца карты
	VERIFY               PaymentFeature = "VERIFY"               // Верификация владельца карты без списания средств
	FORCE_TDS            PaymentFeature = "FORCE_TDS"            // Принудительное проведение платежа с использованием 3-D Secure
	FORCE_SSL            PaymentFeature = "FORCE_SSL"            // Принудительное проведение платежа через SSL
	FORCE_FULL_TDS       PaymentFeature = "FORCE_FULL_TDS"       // Полное проведение аутентификации с 3-D Secure
	FORCE_CREATE_BINDING PaymentFeature = "FORCE_CREATE_BINDING" // Принудительное создание связки
)
