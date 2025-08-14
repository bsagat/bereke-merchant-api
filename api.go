package bereke_merchant

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"time"

	"github.com/bsagat/bereke-merchant-api/models/core"
	"github.com/bsagat/bereke-merchant-api/models/types"
)

type method string

const (
	GET  method = "GET"
	POST method = "POST"
)

var (
	testURL = "https://3dsec.berekebank.kz/payment/rest/"
	prodURL = "https://securepayments.berekebank.kz/payment/rest/"
)

// API — основной интерфейс работы с Bereke Merchant API.
type API interface {
	// --- Заказы ---
	RegisterOrder(ctx context.Context, req core.RegisterOrderRequest) (core.RegisterOrderResponse, error)
	RegisterOrderByNumber(ctx context.Context, orderNumber string, amount float64, currency int, returnURL, failURL string) (core.RegisterOrderResponse, error)

	GetOrderStatus(ctx context.Context, req core.OrderStatusRequest) (core.OrderStatusResponse, error)
	GetOrderStatusByID(ctx context.Context, orderID string) (core.OrderStatusResponse, error)

	// --- Операции с заказами ---
	RefundOrder(ctx context.Context, req core.RefundOrderRequest) (core.Response, error)
	RefundOrderByID(ctx context.Context, amount float64, currency int, orderID string) (core.Response, error)

	ReversalOrder(ctx context.Context, req core.ReversalOrderRequest) (core.Response, error)
	ReversalOrderByID(ctx context.Context, amount float64, currency int, orderID string) (core.Response, error)

	CancelOrder(ctx context.Context, req core.CancelOrderRequest) (core.Response, error)
	CancelOrderByID(ctx context.Context, orderID string) (core.Response, error)

	// --- Системное ---
	Ping() error
}

type api struct {
	authType       types.Auth
	credentials    url.Values
	baseURL        string
	mode           types.Mode
	certPath       string
	certPassphrase string
}

func NewWithLogin(login, password string, mode types.Mode) (API, error) {
	creds := url.Values{}
	creds.Set("userName", login)
	creds.Set("password", password)
	return newAPI(mode, creds, types.AuthLogin, "", "")
}

func NewWithToken(token string, mode types.Mode) (API, error) {
	creds := url.Values{}
	creds.Set("token", token)
	return newAPI(mode, creds, types.AuthToken, "", "")
}

func NewWithCertificate(certPath, passphrase string, mode types.Mode) (API, error) {
	return newAPI(mode, url.Values{}, types.AuthCertificate, certPath, passphrase)
}

func newAPI(mode types.Mode, creds url.Values, authType types.Auth, certPath, passphrase string) (API, error) {
	var baseURL string
	switch mode {
	case types.TEST:
		baseURL = testURL
	case types.PROD:
		baseURL = prodURL
	default:
		return nil, fmt.Errorf("invalid mode: %s", mode)
	}

	return &api{
		authType:       authType,
		credentials:    creds,
		baseURL:        baseURL,
		mode:           mode,
		certPath:       certPath,
		certPassphrase: passphrase,
	}, nil
}

func (a *api) Ping() error {
	client := http.Client{
		Timeout: 3 * time.Second,
	}

	resp, err := client.Get(a.baseURL)
	if err != nil {
		return fmt.Errorf("server is unreachable: %v", err)
	}
	defer resp.Body.Close()

	return nil
}

func (a *api) sendRequest(ctx context.Context, method method, path string, params url.Values, result interface{}) error {
	endpoint := fmt.Sprintf("%s/%s", a.baseURL, path)
	req, err := http.NewRequestWithContext(ctx, string(method), endpoint, nil)
	if err != nil {
		log.Printf("Error creating request: %v", err)
		return err
	}

	req.Header.Set("Accept", "*/*")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	// Добавление параметров авторизации
	query := url.Values{}
	for key, vals := range a.credentials {
		for _, val := range vals {
			query.Add(key, val)
		}
	}
	for key, vals := range params {
		for _, val := range vals {
			query.Add(key, val)
		}
	}

	// PROD-режим: подписываем или шифруем
	if a.mode == types.PROD && a.authType == types.AuthCertificate {
		body := query.Encode()
		if err := a.signAndSetHeaders(req, body); err != nil {
			log.Printf("Error signing request: %v", err)
			return err
		}
	} else {
		req.URL.RawQuery = query.Encode()
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Printf("Error making request: %v", err)
		return err
	}
	defer resp.Body.Close()

	if result != nil {
		return json.NewDecoder(resp.Body).Decode(result)
	}
	_, err = io.Copy(io.Discard, resp.Body)
	return err
}
