package yookassa

// Payment - Структура платежа
type Payment struct {
	ID            string                 `json:"id,omitempty"`
	Status        string                 `json:"status,omitempty"`
	Amount        Amount                 `json:"amount"`
	Description   string                 `json:"description,omitempty"`
	Metadata      map[string]interface{} `json:"metadata,omitempty"`
	Confirmation  Redirect               `json:"confirmation"`
	Capture       bool                   `json:"capture,omitempty"`
	PaymentMethod PaymentMethod          `json"payment_method,omitempty"`
}
