package yookassa

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"github.com/google/uuid"
)

// Базовый url API
const (
	baseURL = "https://api.yookassa.ru/v3/"
)

// Client - клиент для работы с API YooKassa
type Client struct {
	shopID         string // ID магазина
	shopSecret     string // Секретный ключ
	httpClient     *http.Client // Клиент http
}

// NewClient - новый клиент YooKassa
//
// Принимает:
// shopID: string - ID магазина
// shopSecret: string - Секретный ключ
//
// Возвращает:
// Клиент API: *Client
func NewClient(shopID, shopSecret string) *Client {
	return &Client{
		shopID:     shopID,
		shopSecret: shopSecret,
		httpClient: &http.Client{},
	}
}

// CreatePayment - создание нового платежа
//
// Принимает:
// request: Payment - Запрос на создание платежа
//
// Возвращает:
// Объект платежа: *Payment
// Ошибка: error
func (c *Client) CreatePayment(request Payment) (*Payment, error) {
	// URL для запроса
	url := baseURL + "payments"

	// Создание тела запроса json
	requestBody, err := json.Marshal(request)

	if err != nil {
		return nil, fmt.Errorf("error marshaling request: %v", err)
	}

	// Запрос к yookassa
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(requestBody))

	if err != nil {
		return nil, fmt.Errorf("error creating request: %v", err)
	}

	// Подключение авторизации по id магазина и секретному ключу
	req.SetBasicAuth(c.shopID, c.shopSecret)

	// Тип тела запроса
	req.Header.Set("Content-Type", "application/json")

	// Добваление в запрос ключа идемпотентности
	req.Header.Set("Idempotence-Key", uuid.New().String())

	// Выполнение запроса
	resp, err := c.httpClient.Do(req)

	if err != nil {
		return nil, fmt.Errorf("error sending request: %v", err)
	}

	defer resp.Body.Close()

	// Проверка статуса ответа
	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("API error: %s, body: %s", resp.Status, string(body))
	}

	// Декодирование тела ответа
	var payment Payment

	if err := json.NewDecoder(resp.Body).Decode(&payment); err != nil {
		return nil, fmt.Errorf("error decoding response: %v", err)
	}

	return &payment, nil
}

// GetPayment - получение информации о платеже по ID
//
// Принимает:
// paymentID: string - id платежа
//
// Возвращает:
// Объект платежа: *Payment
// Ошибка: error
func (c *Client) GetPayment(paymentID string) (*Payment, error) {
	// Проверка наличия id платежа
	if paymentID == "" {
		return nil, errors.New("payment ID cannot be empty")
	}

	// URL для запроса
	url := baseURL + "payments/" + paymentID

	// Запрос к yookassa
	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		return nil, fmt.Errorf("error creating request: %v", err)
	}

	// Подключение авторизации по id магазина и секретному ключу
	req.SetBasicAuth(c.shopID, c.shopSecret)

	// Тип тела запроса
	req.Header.Set("Content-Type", "application/json")

	// Добваление в запрос ключа идемпотентности
	req.Header.Set("Idempotence-Key", uuid.New().String())

	// Выполнение запроса
	resp, err := c.httpClient.Do(req)

	if err != nil {
		return nil, fmt.Errorf("error sending request: %v", err)
	}

	defer resp.Body.Close()

	// Проверка статуса ответа
	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("API error: %s, body: %s", resp.Status, string(body))
	}

	// Декодирование тела ответа
	var payment Payment

	if err := json.NewDecoder(resp.Body).Decode(&payment); err != nil {
		return nil, fmt.Errorf("error decoding response: %v", err)
	}

	return &payment, nil
}

// CapturePayment - подтверждение платежа
//
// Принимает:
// paymentID: string - id платежа
//
// Возвращает:
// Объект платежа: *Payment
// Ошибка: error
func (c *Client) CapturePayment(paymentID string) (*Payment, error) {
	// Проверка наличия id платежа
	if paymentID == "" {
		return nil, errors.New("payment ID cannot be empty")
	}

	// URL для запроса
	url := baseURL + "payments/" + paymentID + "/capture"

	// Запрос к yookassa
	req, err := http.NewRequest("POST", url, bytes.NewBuffer([]byte("{}")))

	if err != nil {
		return nil, fmt.Errorf("error creating request: %v", err)
	}

	// Подключение авторизации по id магазина и секретному ключу
	req.SetBasicAuth(c.shopID, c.shopSecret)

	// Тип тела запроса
	req.Header.Set("Content-Type", "application/json")

	// Добваление в запрос ключа идемпотентности
	req.Header.Set("Idempotence-Key", uuid.New().String())

	// Выполнение запроса
	resp, err := c.httpClient.Do(req)

	if err != nil {
		return nil, fmt.Errorf("error sending request: %v", err)
	}

	defer resp.Body.Close()

	// Проверка статуса ответа
	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("API error: %s, body: %s", resp.Status, string(body))
	}

	// Декодирование тела ответа
	var payment Payment

	if err := json.NewDecoder(resp.Body).Decode(&payment); err != nil {
		return nil, fmt.Errorf("error decoding response: %v", err)
	}

	return &payment, nil
}

// CancelPayment отменяет платеж
func (c *Client) CancelPayment(paymentID string) (*Payment, error) {
	// Проверка наличия id платежа
	if paymentID == "" {
		return nil, errors.New("payment ID cannot be empty")
	}

	// URL для запроса
	url := baseURL + "payments/" + paymentID + "/cancel"

	// Запрос к yookassa
	req, err := http.NewRequest("POST", url, nil)

	if err != nil {
		return nil, fmt.Errorf("error creating request: %v", err)
	}

	// Подключение авторизации по id магазина и секретному ключу
	req.SetBasicAuth(c.shopID, c.shopSecret)

	// Тип тела запроса
	req.Header.Set("Content-Type", "application/json")

	// Добваление в запрос ключа идемпотентности
	req.Header.Set("Idempotence-Key", uuid.New().String())

	// Выполнение запроса
	resp, err := c.httpClient.Do(req)

	if err != nil {
		return nil, fmt.Errorf("error sending request: %v", err)
	}

	defer resp.Body.Close()

	// Проверка статуса ответа
	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("API error: %s, body: %s", resp.Status, string(body))
	}

	// Декодирование тела ответа
	var payment Payment

	if err := json.NewDecoder(resp.Body).Decode(&payment); err != nil {
		return nil, fmt.Errorf("error decoding response: %v", err)
	}

	return &payment, nil
}
