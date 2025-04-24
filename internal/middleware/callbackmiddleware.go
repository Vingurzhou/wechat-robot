package middleware

import (
	"net/http"
)

type CallbackMiddleware struct {
}

func NewCallbackMiddleware() *CallbackMiddleware {
	return &CallbackMiddleware{}
}

func (m *CallbackMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// TODO generate middleware implement function, delete after code implementation
		// body, err := io.ReadAll(r.Body)
		// if err != nil {
		// 	logx.Errorf("Failed to read request body: %v", err)
		// 	http.Error(w, "Failed to read request body", http.StatusBadRequest)
		// 	return
		// }
		// logx.Infof("Request Payload: %s", string(body))
		// r.Body = io.NopCloser(bytes.NewBuffer(body))
		// Passthrough to next handler if need
		next(w, r)
	}
}
