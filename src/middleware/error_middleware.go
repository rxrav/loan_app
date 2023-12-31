package middleware

import (
	"encoding/json"
	"net/http"

	"github.com/rxrav/loan_app/src/dto"
)

func ErrorMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		// Do error handling here
		defer func() {
			if err := recover(); err != nil {
				appErr := err.(dto.LoanApplicationError)
				writer.Header().Set("Content-Type", "application/json")
				// should be last to set
				writer.WriteHeader(whatHttpCode(appErr.ErrCode))
				_ = json.NewEncoder(writer).Encode(appErr)
			}
		}()

		// Call the next handler, which can be another middleware in the chain, or the final handler.
		next.ServeHTTP(writer, request)
	})
}

func whatHttpCode(businessErrCode int) int {
	httpCode := http.StatusBadRequest
	switch {
	case businessErrCode == 20004:
		httpCode = http.StatusUnauthorized
	case businessErrCode >= 20000:
		httpCode = http.StatusInternalServerError
	case businessErrCode >= 10000 && businessErrCode < 20000:
		httpCode = http.StatusBadRequest
	case businessErrCode >= 9000 && businessErrCode < 10000:
		httpCode = http.StatusBadRequest
	default:
		httpCode = http.StatusBadRequest
	}
	return httpCode
}
