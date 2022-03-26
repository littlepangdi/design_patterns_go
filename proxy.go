package main

import (
	"fmt"
	"net/http"
)

/*
	The proxy pattern is a structural design pattern
	that enables you to provide a substitute for an object or its placeholder.
	The proxy controls access to the original object and allows some processing before and after submitting the request to the object.
	1. Lazy initialization (virtual proxy).
	2. Access control(protect proxy)
	3. Execute the remote service locally (remote proxy).

	A web server like Nginx acts as a proxy for the application server
*/
var (
	urlStatus        = "/system/status"
	urlCreate        = "/system/create"
	resultOK         = "ok"
	resultNotOK      = "not ok"
	resultNotAllowed = "not allowed"
	resultCreated    = "system created"
)

type pServer interface {
	handlerRequest(string, string) (int, string)
}

type nginx struct {
	application       *application
	maxAllowedRequest int
	rateLimiter       map[string]int
}

func newNginxServer() *nginx {
	return &nginx{
		application:       &application{},
		maxAllowedRequest: 2,
		rateLimiter:       make(map[string]int),
	}
}

func (n *nginx) handlerRequest(url, method string) (int, string) {
	if !n.allowRateLimiting(url) {
		return http.StatusForbidden, resultNotAllowed
	}
	return n.application.handlerRequest(url, method)
}
func (n *nginx) allowRateLimiting(url string) bool {
	if n.rateLimiter[url] == 0 {
		n.rateLimiter[url] = 1
	}
	if n.rateLimiter[url] > n.maxAllowedRequest {
		return false
	}
	n.rateLimiter[url]++
	return true
}

type application struct{}

func (a *application) handlerRequest(url, method string) (int, string) {
	if url == urlStatus && method == http.MethodGet {
		return http.StatusOK, resultOK
	}
	if url == urlCreate && method == http.MethodPut {
		return http.StatusCreated, resultCreated
	}
	return http.StatusNotFound, resultNotOK
}

func RunProxy() {
	nginxServer := newNginxServer()
	code, body := nginxServer.handlerRequest(urlStatus, http.MethodGet)
	fmt.Printf("Url: %s\tHttpCode: %d\tBody: %s\n", urlStatus, code, body)
	code, body = nginxServer.handlerRequest(urlStatus, http.MethodGet)
	fmt.Printf("Url: %s\tHttpCode: %d\tBody: %s\n", urlStatus, code, body)
	code, body = nginxServer.handlerRequest(urlStatus, http.MethodGet)
	fmt.Printf("Url: %s\tHttpCode: %d\tBody: %s\n", urlStatus, code, body)

	code, body = nginxServer.handlerRequest(urlStatus, http.MethodPost)
	fmt.Printf("Url: %s\tHttpCode: %d\tBody: %s\n", urlStatus, code, body)
	code, body = nginxServer.handlerRequest(urlStatus, http.MethodGet)
	fmt.Printf("Url: %s\tHttpCode: %d\tBody: %s\n", urlStatus, code, body)
}
