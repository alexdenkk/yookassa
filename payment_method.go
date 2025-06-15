package yookassa

// Методы оплаты

type PaymentMethodType string

const (
	PaymentTypeBankCard PaymentMethodType = "bank_card"

)

type PaymentMethod interface {}

type paymentMethod struct {
	// Код метода оплаты
	Type PaymentMethodType `json:"type,omitempty"`

	// ID метода оплаты
	ID string `json:"id,omitempty"`
	Saved bool `json:"saved,omitempty"`
	// Название метода оплаты
	Title string `json:"title,omitempty"`
}

type BankCard struct {
	paymentMethod

	// Данные банковской карты
	Card Card `json:"card,omitempty"`
}
