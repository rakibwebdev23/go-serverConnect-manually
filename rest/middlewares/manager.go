package middleware

import "net/http"

// Middleware function signature
type Middleware func(http.Handler) http.Handler

// Middleware Manager
type Manager struct {
	globalMiddlewares []Middleware
}

// Create a new manager
func NewManager() *Manager {
	return &Manager{
		globalMiddlewares: make([]Middleware, 0),
	}
}

// Register global middlewares
func (m *Manager) Use(middlewares ...Middleware) {
	m.globalMiddlewares = append(m.globalMiddlewares, middlewares...)
}

// Apply global + route middlewares to a handler
func (m *Manager) With(handler http.Handler, middlewares ...Middleware) http.Handler {
	h := handler

	// Combine global and route-specific middlewares
	allMiddlewares := append(m.globalMiddlewares, middlewares...)

	// Apply in reverse order
	for i := len(allMiddlewares) - 1; i >= 0; i-- {
		h = allMiddlewares[i](h)
	}

	return h
}




// package middleware

// import "net/http"

// type Middleware func(http.Handler) http.Handler //signature of middleware function

// // Middleware Manager
// type Manager struct {
// 	globalMiddlewares []Middleware
// }

// // Create a new manager
// func NewManager() *Manager {
// 	return &Manager{
// 		globalMiddlewares: make([]Middleware, 0),
// 	}
// }

// //Recever function 
// func (mnger *Manager) Use(middlewares ...Middleware) {
// 	mnger.globalMiddlewares = append(mnger.globalMiddlewares, middlewares...)
// }

// func (mnger *Manager) With(handler http.Handler, middlewares ...Middleware) http.Handler {
// 	h := handler
// 	for _, middleware := range middlewares {
// 		h = middleware(h)
// 	}
// 	return h;
// }