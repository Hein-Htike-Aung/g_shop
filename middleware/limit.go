package middleware

import (
	"net/http"
	"yixiang.co/go-mall/pkg/global"
	"yixiang.co/go-mall/pkg/limiter"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
)

// LimitIP: global rate limit by IP
// limit format string, e.g. "5-S". Examples:
//
// * 5 reqs/second: "5-S"
// * 10 reqs/minute: "10-M"
// * 1000 reqs/hour: "1000-H"
// * 2000 reqs/day: "2000-D"
//
func LimitIP(limit string) gin.HandlerFunc {

	return func(c *gin.Context) {
		// rate limit by IP
		key := limiter.GetKeyIP(c)
		if ok := limitHandler(c, key, limit); !ok {
			return
		}
		c.Next()
	}
}

// LimitPerRoute: per-route rate limit
func LimitPerRoute(limit string) gin.HandlerFunc {
	return func(c *gin.Context) {
		// rate limit by IP and route
		key := limiter.GetKeyRouteWithIP(c)
		if ok := limitHandler(c, key, limit); !ok {
			return
		}
		c.Next()
	}
}

func limitHandler(c *gin.Context, key string, limit string) bool {

	// check rate limit
	rate, err := limiter.CheckRate(c, key, limit)
	if err != nil {
		global.YSHOP_LOG.Error(err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": "internal server error, please try again later",
		})

		return false
	}

	// ---- rate limit headers -----
	// X-RateLimit-Limit: max requests
	// X-RateLimit-Remaining: remaining requests
	// X-RateLimit-Reset: reset time (counter returns to limit)
	c.Header("X-RateLimit-Limit", cast.ToString(rate.Limit))
	c.Header("X-RateLimit-Remaining", cast.ToString(rate.Remaining))
	c.Header("X-RateLimit-Reset", cast.ToString(rate.Reset))

	// limit reached
	if rate.Reached {
		// tell client limit was exceeded
		c.AbortWithStatusJSON(http.StatusTooManyRequests, gin.H{
			"message": "too many requests",
		})
		return false
	}

	return true
}