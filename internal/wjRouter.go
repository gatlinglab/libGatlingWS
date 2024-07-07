package iWSServer

import "net/http"

type cWJWSRouter struct {
	handler         map[string]http.HandlerFunc
	homeHandler     http.HandlerFunc
	notFoundHandler http.HandlerFunc
}

func newWJWSRouter() *cWJWSRouter {
	return &cWJWSRouter{handler: make(map[string]http.HandlerFunc), homeHandler: pageEmpty, notFoundHandler: pageEmpty}
}

func (pInst *cWJWSRouter) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		pInst.homeHandler(w, r)
		return
	}
	http.NotFound(w, r)
	return
}

func (pInst *cWJWSRouter) HandlerFunc(pattern string, fn http.HandlerFunc) {
	newMap := make(map[string]http.HandlerFunc)
	for key, value := range pInst.handler {
		newMap[key] = value
	}
	newMap[pattern] = fn
}

/////////////////////////////// default function for empty call /////////////////////////////

func pageEmpty(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("not found"))
}
