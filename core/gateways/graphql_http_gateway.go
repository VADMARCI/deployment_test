package gateways

import (
	"net/http"
)

type AddHeaderTransport struct {
	T                    http.RoundTripper
	SessionToken         string
	InternalServiceToken string
	Cookies              []*http.Cookie
}

func (adt *AddHeaderTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	req.Header.Add("Authorization", adt.SessionToken)
	req.Header.Add("service-jwt-token", adt.InternalServiceToken)
	for _, cookie := range adt.Cookies {
		req.AddCookie(cookie)
	}
	return adt.T.RoundTrip(req)
}

func NewHttpClientWithAuthTokens(cookies []*http.Cookie, sessionToken, internalServiceToken string) *http.Client {
	return &http.Client{
		Transport: &AddHeaderTransport{T: http.DefaultTransport, Cookies: cookies, SessionToken: sessionToken, InternalServiceToken: internalServiceToken},
	}
}
