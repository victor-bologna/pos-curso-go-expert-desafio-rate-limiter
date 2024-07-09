package middleware

import (
	"encoding/json"
	"net"
	"net/http"

	"github.com/pos-curso-go-expert-desafio-rate-limiter/config"
	"github.com/pos-curso-go-expert-desafio-rate-limiter/internal/dto"
	"github.com/pos-curso-go-expert-desafio-rate-limiter/internal/limiter"
)

const timeout_message = "you have reached the maximum number of requests or actions allowed within a certain time frame"

func RateLimiterMiddleware(rc limiter.RateLimiterStrategy, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		apiKey := r.Header.Get("API_KEY")
		maxReq := config.AppConfig.TokenMaximumReq
		ttl := config.AppConfig.Timeout
		if apiKey == "" {
			apiKey = getIpAddress(r)
			maxReq = config.AppConfig.IPMaximumReq
		}
		state, message := rc.NextRequest(apiKey, maxReq, ttl)
		if state == limiter.Allow {
			next.ServeHTTP(w, r)
			return
		}
		throwHttpError(w, message)
	})
}

func getIpAddress(r *http.Request) string {
	ipAddress := r.RemoteAddr
	if host, _, err := net.SplitHostPort(ipAddress); err == nil {
		ipAddress = host
	}
	return ipAddress
}

func throwHttpError(w http.ResponseWriter, message string) {
	middlewareError := dto.MiddlewareErrorDTO{Message: message}
	byteResp, _ := json.Marshal(middlewareError)
	if message == timeout_message {
		http.Error(w, string(byteResp), http.StatusTooManyRequests)
	} else {
		http.Error(w, string(byteResp), http.StatusInternalServerError)
	}
}
