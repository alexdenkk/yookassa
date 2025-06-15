package yookassa

// Данные банковской карты
type Card struct {
	// Первые 6 цифр номера карты
	First6        string `json:"first6,omitempty"`
	// Последние 4 цифры номера карты
	Last4         string `json:"last4,omitempty"`
	// Срок действия до (год)
	ExpiryYear    string `json:"expiry_year,omitempty"`
	// Срок действия до (месяц)
	ExpiryMonth   string `json:"expiry_month,omitempty"`
	// Тип карты (МИР, Visa, MasterCard, PayPal, и т.д.)
	CardType      string `json:"card_type,omitempty"`
	// Код страны (Например: RU)
	IssuerCountry string `json:"issuer_country,omitempty"`
	// Название банка
	IssuerName    string `json:"issuer_name,omitempty"`
	// Откуда была использована карта (apple_pay, google_pay, и т.д.)
	Source        string `json:"source,omitempty"`
}
