package middleware

import "net/http"

type CallbackMiddleware struct {
}

func NewCallbackMiddleware() *CallbackMiddleware {
	return &CallbackMiddleware{}
}

func (m *CallbackMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// TODO generate middleware implement function, delete after code implementation

		// Passthrough to next handler if need
		next(w, r)
	}
}
