package middleware

import "net/http"

type Middleware func(http.Handler) http.Handler //signature of middleware function

type Manager struct {
	globalMiddlewares []Middleware
}

func NewManager() *Manager {
	return &Manager{
		globalMiddlewares: make([]Middleware, 0),
	}
}

//Recever function 
func (mnger *Manager) Use(middlewares ...Middleware) {
	mnger.globalMiddlewares = append(mnger.globalMiddlewares, middlewares...)
}

func (mnger *Manager) With(handler http.Handler, middlewares ...Middleware) http.Handler {
	h := handler
	for _, middleware := range middlewares {
		h = middleware(h)
	}
	return h;
}