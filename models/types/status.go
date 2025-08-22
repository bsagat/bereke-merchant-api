package types

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

type PaymentState string

const (
	OrderCreated   PaymentState = "CREATED"   // Заказ создан (но не оплачен)
	OrderApproved  PaymentState = "APPROVED"  // Заказ одобрен (средства на счету покупателя заблокированы)
	OrderDeposited PaymentState = "DEPOSITED" // Заказ завершен (деньги списаны со счета покупателя)
	OrderDeclined  PaymentState = "DECLINED"  // Заказ отклонен
	OrderReversed  PaymentState = "REVERSED"  // Авторизованный заказ отклонен
	OrderRefunded  PaymentState = "REFUNDED"  // Возврат средств
)
