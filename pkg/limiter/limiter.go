package limiter

import (
	"strings"
	"yixiang.co/go-mall/pkg/global"
	v9 "yixiang.co/go-mall/pkg/redis/v9"

	"github.com/gin-gonic/gin"
	limiterlib "github.com/ulule/limiter/v3"
	sredis "github.com/ulule/limiter/v3/drivers/store/redis"
)

// GetKeyIP rate limiter key: IP
func GetKeyIP(c *gin.Context) string {
	return c.ClientIP()
}

// GetKeyRouteWithIP rate limiter key: route+IP
func GetKeyRouteWithIP(c *gin.Context) string {
	return routeToKeyString(c.FullPath()) + c.ClientIP()
}

// CheckRate check if rate limit exceeded
func CheckRate(c *gin.Context, key string, formatted string) (limiterlib.Context, error) {

	// create limiter.Rate
	var context limiterlib.Context
	rate, err := limiterlib.NewRateFromFormatted(formatted)
	if err != nil {
		global.YSHOP_LOG.Error(err)
		return context, err
	}

	//v8.Redis.Client

	// init store with app Redis
	store, err := sredis.NewStoreWithOptions(v9.Redis.Client, limiterlib.StoreOptions{
		// key prefix for tidy Redis keys
		Prefix: "yshop:limiter",
	})
	if err != nil {
		global.YSHOP_LOG.Error(err)
		return context, err
	}

	// build limiter with rate and store
	limiterObj := limiterlib.New(store, rate)

	// get rate limit result
	return limiterObj.Get(c, key)
}

// routeToKeyString format URL slashes as dashes
func routeToKeyString(routeName string) string {
	routeName = strings.ReplaceAll(routeName, "/", "-")
	routeName = strings.ReplaceAll(routeName, ":", "_")
	return routeName
}
