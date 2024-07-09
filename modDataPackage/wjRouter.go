package modDataPackage

import (
	"fmt"
	"net/http"
)

type cWJWSRouter struct {
	handler          map[string]http.HandlerFunc
	homeHandler      http.HandlerFunc
	notFoundHandler  http.HandlerFunc
	upgradeRouterKey string
	server           *CGatlingWSServer
}

func newWJWSRouter(serverInst *CGatlingWSServer) *cWJWSRouter {
	return &cWJWSRouter{handler: make(map[string]http.HandlerFunc),
		homeHandler:      pageEmpty,
		notFoundHandler:  pageEmpty,
		upgradeRouterKey: "/ws",
		server:           serverInst,
	}
}

func (pInst *cWJWSRouter) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println("serveHttp route enter:", r.URL.Path)
	if r.URL.Path == "/" {
		pInst.homeHandler(w, r)
		return
	} else if r.URL.Path == pInst.upgradeRouterKey {
		pInst.server.Upgrade(w, r)
		return
	}

	fn1, exists := pInst.handler[r.URL.Path]
	if !exists {
		http.NotFound(w, r)
	} else {
		fn1(w, r)
	}
}

func (pInst *cWJWSRouter) HandlerFunc(pattern string, fn http.HandlerFunc) {
	if pattern == "/" {
		pInst.homeHandler = fn
		return
	}
	newMap := make(map[string]http.HandlerFunc)
	for key, value := range pInst.handler {
		newMap[key] = value
	}
	newMap[pattern] = fn
	pInst.handler = newMap
}

/////////////////////////////// default function for empty call /////////////////////////////

func pageEmpty(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("not found"))
}
