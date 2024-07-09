package limiter

type RateLimiterStrategy interface {
	NextRequest(key string, maximumReq int, timeout int) (State, string)
}
