package limiter_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/pos-curso-go-expert-desafio-rate-limiter/config"
	"github.com/pos-curso-go-expert-desafio-rate-limiter/internal/limiter"
	"github.com/pos-curso-go-expert-desafio-rate-limiter/internal/middleware"
	"github.com/stretchr/testify/assert"
)

func TestRateLimiter(t *testing.T) {
	config.LoadConfig("../..")
	rc := limiter.NewRedisClient()

	myHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, World!"))
	})

	rateLimitedHandler := middleware.RateLimiterMiddleware(rc, myHandler)

	t.Run("Rate limiter - IP Test", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		tries := config.AppConfig.IPMaximumReq + 5

		for i := 0; i < tries; i++ {
			w := httptest.NewRecorder()
			rateLimitedHandler.ServeHTTP(w, req)
			resp := w.Result()
			if i < config.AppConfig.IPMaximumReq {
				assert.Equal(t, 200, resp.StatusCode)
			} else {
				assert.Equal(t, 429, resp.StatusCode)
			}
		}
	})

	t.Run("Rate Limiter - Token Test", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		req.Header.Set("API_KEY", "test-token")
		tries := config.AppConfig.TokenMaximumReq + 5

		for i := 0; i < tries; i++ {
			w := httptest.NewRecorder()
			rateLimitedHandler.ServeHTTP(w, req)
			resp := w.Result()
			if i < config.AppConfig.TokenMaximumReq {
				assert.Equal(t, 200, resp.StatusCode)
			} else {
				assert.Equal(t, 429, resp.StatusCode)
			}
		}
	})
}
