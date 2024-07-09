package main

import (
	"net/http"

	"github.com/pos-curso-go-expert-desafio-rate-limiter/config"
	"github.com/pos-curso-go-expert-desafio-rate-limiter/internal/limiter"
	"github.com/pos-curso-go-expert-desafio-rate-limiter/internal/middleware"
)

func main() {
	config.LoadConfig("")
	rc := limiter.NewRedisClient()

	myHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, World!"))
	})

	http.Handle("/", middleware.RateLimiterMiddleware(rc, myHandler))
	http.ListenAndServe(":8080", nil)
}
