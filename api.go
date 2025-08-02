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
)

type (
	method string
	Mode   string
)

const (
	GET  method = "GET"
	POST method = "POST"
	TEST Mode   = "TEST"
	PROD Mode   = "PROD"
)

var (
	testURL = "https://3dsec.berekebank.kz/payment/rest/"
	prodURL = "https://securepayments.berekebank.kz/payment/rest/"
)

type API interface {
	RegisterOrder(ctx context.Context, req RegisterOrderRequest) (RegisterOrderResponse, error)
	OrderStatus(ctx context.Context, req OrderStatusRequest) (OrderStatusResponse, error)
	RefundOrder(ctx context.Context, req RefundOrderRequest) (Response, error)
	ReversalOrder(ctx context.Context, req ReversalOrderRequest) (Response, error)
	CancelOrder(ctx context.Context, req CancelOrderRequest) (Response, error)
	Ping() error
}

type authType int

const (
	authLogin authType = iota
	authToken
	authCertificate
)

type api struct {
	authType       authType
	credentials    url.Values
	baseURL        string
	mode           Mode
	certPath       string
	certPassphrase string
}

func NewWithLogin(login, password string, mode Mode) (API, error) {
	creds := url.Values{}
	creds.Set("userName", login)
	creds.Set("password", password)
	return newAPI(mode, creds, authLogin, "", "")
}

func NewWithToken(token string, mode Mode) (API, error) {
	creds := url.Values{}
	creds.Set("token", token)
	return newAPI(mode, creds, authToken, "", "")
}

func NewWithCertificate(certPath, passphrase string, mode Mode) (API, error) {
	return newAPI(mode, url.Values{}, authCertificate, certPath, passphrase)
}

func newAPI(mode Mode, creds url.Values, at authType, certPath, passphrase string) (API, error) {
	var baseURL string
	switch mode {
	case TEST:
		baseURL = testURL
	case PROD:
		baseURL = prodURL
	default:
		return nil, fmt.Errorf("invalid mode: %s", mode)
	}

	return &api{
		authType:       at,
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

func (a *api) do(ctx context.Context, method method, path string, params url.Values, result interface{}) error {
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
	if a.mode == PROD && a.authType == authCertificate {
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
