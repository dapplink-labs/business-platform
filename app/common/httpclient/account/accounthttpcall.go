package account

import (
	"errors"
	"fmt"
	"log"
	"time"

	gresty "github.com/go-resty/resty/v2"
)

type AccountHttpCallClient interface {
}

type callclient struct {
	grestyClient *gresty.Client
}

var (
	errHTTPError       = errors.New("http error")
	errInvalidAddress  = errors.New("invalid address")
	errInvalidResponse = errors.New("invalid response")
)

const (
	defaultRequestTimeout   = 10 * time.Second
	defaultRetryWaitTime    = 10 * time.Second
	defaultRetryMaxWaitTime = 30 * time.Second
	defaultRetryCount       = 3
	defaultWithDebug        = false
)

func NewHttpClient(baseUrl, apiKey string) (AccountHttpCallClient, error) {
	return NewHttpClientAll(baseUrl, apiKey, defaultWithDebug)
}

func NewHttpClientAll(baseUrl, apiKey string, withDebug bool) (AccountHttpCallClient, error) {
	grestyClient := gresty.New()
	grestyClient.SetBaseURL(baseUrl)
	grestyClient.SetTimeout(defaultRequestTimeout)
	grestyClient.SetRetryCount(defaultRetryCount)
	grestyClient.SetRetryWaitTime(defaultRetryWaitTime)
	grestyClient.SetRetryMaxWaitTime(defaultRetryMaxWaitTime)
	grestyClient.SetDebug(withDebug)
	grestyClient.OnBeforeRequest(func(c *gresty.Client, r *gresty.Request) error {
		log.Printf("Making request to %s (Attempt %d)", r.URL, r.Attempt)
		return nil
	})

	grestyClient.OnAfterResponse(func(c *gresty.Client, r *gresty.Response) error {
		statusCode := r.StatusCode()
		attempt := r.Request.Attempt
		method := r.Request.Method
		url := r.Request.URL
		log.Printf("Response received: Method=%s, URL=%s, Status=%d, Attempt=%d",
			method, url, statusCode, attempt)

		if statusCode >= 400 {
			if statusCode == 404 {
				return fmt.Errorf("%d resource not found %s %s: %w",
					statusCode, method, url, errHTTPError)
			}
			if statusCode >= 500 {
				return fmt.Errorf("%d server error %s %s: %w",
					statusCode, method, url, errHTTPError)
			}
			return fmt.Errorf("%d cannot %s %s: %w",
				statusCode, method, url, errHTTPError)
		}
		return nil
	})
	return &callclient{grestyClient: grestyClient}, nil
}
