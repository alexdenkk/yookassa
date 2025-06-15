package yookassa

type ConfirmationType string

const (
	TypeRedirect ConfirmationType = "redirect"
)

// Redirect - Оплата через перенаправление на сайт
type Redirect struct {
	// Код сценария оплаты
	Type            ConfirmationType `json:"type,omitempty"`
	// URL, на который пользователь будет перенаправлен для оплаты
	ConfirmationURL string           `json:"confirmation_url,omitempty"`
	// URL, на который пользователь будет перенаправлен после совершения оплаты
	ReturnURL       string           `json:"return_url,omitempty" binding:"max=2048"`
}
