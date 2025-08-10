package money

import (
	"errors"
	"math"
	"strconv"
)

var ErrCurrencyNotSupported = errors.New("currency is not supported")

// Доступные валюты
const (
	KZT string = "KZT" // 398
	USD string = "USD" // 840
	RUB string = "RUB" // 643
	EUR string = "EUR" // 978
)

// ISO currency map
var (
	alphaToNumeric = map[string]int{
		KZT: 398,
		USD: 840,
		RUB: 643,
		EUR: 978,
	}

	numericToAlpha = map[int]string{
		398: KZT,
		840: USD,
		643: RUB,
		978: EUR,
	}

	minorUnits = map[int]int{
		398: 100,
		840: 100,
		643: 100,
		978: 100,
	}
)

// Конвертирует ISO currency code в string
func ToAlpha(code int) string {
	return numericToAlpha[code]
}

// Конвертирует ISO currency code в int
func ToNumeric(code string) int {
	return alphaToNumeric[code]
}

// Выполняет обработку string ISO code ("398", "KZT") и возвращает string
// FromString("KZT") |  returns "KZT"
// FromString("398") | 	returns "KZT"
func FromString(value string) string {
	// Try as numeric
	if num, err := strconv.Atoi(value); err == nil {
		return ToAlpha(num)
	}
	// Try as alpha
	return value
}

// ToMinorUnit переводит сумму в основных единицах валюты (например, 10.50 USD)
// в количество минорных единиц (например, 1050 центов).
func ToMinorUnit(amount float64, currency int) (int, error) {
	minor, ok := minorUnits[currency]
	if !ok {
		// по умолчанию считаем 2 знака
		return 0, ErrCurrencyNotSupported
	}

	return int(math.Round(amount * float64(minor))), nil
}

// ConvertFromMinorUnits переводит сумму в минорных единцах валюты (например, 1050 центов)
// в основых единицах валюты (например, 10.50 USD)
func ConvertFromMinorUnits(minorAmount int, currency int) (float64, error) {
	minor, ok := minorUnits[currency]
	if !ok {
		// по умолчанию считаем 2 знака
		return 0, ErrCurrencyNotSupported
	}

	return float64(minorAmount) / float64(minor), nil
}
