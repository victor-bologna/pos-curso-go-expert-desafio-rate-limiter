package limiter

type RateLimiterStrategy interface {
	NextRequest(key string, maximumReq int, ttl int) State
}
