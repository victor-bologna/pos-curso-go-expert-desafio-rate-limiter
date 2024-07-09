package limiter_test

import (
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"strconv"
	"testing"

	"github.com/joho/godotenv"
	"github.com/pos-curso-go-expert-desafio-rate-limiter/config"
	"github.com/pos-curso-go-expert-desafio-rate-limiter/internal/limiter"
	"github.com/pos-curso-go-expert-desafio-rate-limiter/internal/middleware"
	"github.com/stretchr/testify/assert"
)

var AppConfig config.Config

func TestRateLimiter(t *testing.T) {
	testConfig()
	rc := limiter.NewRedisClient()

	myHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, World!"))
	})

	rateLimitedHandler := middleware.RateLimiterMiddleware(rc, myHandler)

	t.Run("Rate limiter - IP Test", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		tries := AppConfig.IPMaximumReq + 5

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
		tries := AppConfig.TokenMaximumReq + 5

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

func testConfig() {
	pwd, err := os.Getwd()
	if err != nil {
		log.Printf("Error getting current directory: %v", err)
	}

	rootPath := filepath.Join(pwd, "../..")
	godotenv.Load(filepath.Join(rootPath, ".env"))
	if err != nil {
		log.Printf("Error loading .env file: %v, using default values...", err)
	}

	redisAddr := os.Getenv("REDIS_ADDR")
	if redisAddr == "" {
		log.Printf("Redis Addr is empty, converting to default value: redis:6379")
		redisAddr = "localhost:6379"
	}
	redisDB, err := strconv.Atoi(os.Getenv("REDIS_DB"))
	if err != nil {
		log.Printf("Error converting REDIS_DB to integer: %v, converting to default value: 0", err)
		redisDB = 0
	}
	ipMaxReq, err := strconv.Atoi(os.Getenv("IP_MAX_REQ"))
	if err != nil {
		log.Printf("Error converting IP_MAX_REQ to integer: %v, converting to default value: 5", err)
		ipMaxReq = 5
	}
	tokenMaxReq, err := strconv.Atoi(os.Getenv("TOKEN_MAX_REQ"))
	if err != nil {
		log.Printf("Error converting TOKEN_MAX_REQ to integer: %v, converting to default value: 5", err)
		tokenMaxReq = 5
	}
	timeout, err := strconv.Atoi(os.Getenv("TIMEOUT"))
	if err != nil {
		log.Printf("Error converting TIMEOUT to integer: %v, converting to default value: 10", err)
		timeout = 10
	}

	AppConfig = config.Config{
		RedisAddr:       redisAddr,
		RedisPass:       os.Getenv("REDIS_PASS"),
		RedisDB:         redisDB,
		IPMaximumReq:    ipMaxReq,
		TokenMaximumReq: tokenMaxReq,
		Timeout:         timeout,
	}
}
