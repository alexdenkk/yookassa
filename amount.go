package yookassa

// Amount - сумма платежа
type Amount struct {
	Value    string `json:"value"` // Сумма
	Currency string `json:"currency"` // Валюта
}
