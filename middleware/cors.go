package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Cors() gin.HandlerFunc {
	// CORS handler
	return func(c *gin.Context) {
		method := c.Request.Method
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Headers", "Access-Control-Allow-Headers,Authorization,User-Agent, Keep-Alive, Content-Type, X-Requested-With,X-CSRF-Token,AccessToken,Token")
		c.Header("Access-Control-Allow-Methods", "GET, POST, DELETE, PUT, PATCH, OPTIONS")
		c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
		c.Header("Access-Control-Allow-Credentials", "true")

		// allow all OPTIONS
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusAccepted)
		}
		c.Next()


	}

	//return func(c *gin.Context) {
	//	method := c.Request.Method
	//	origin := c.Request.Header.Get("Origin") //request Origin header
	//	if origin != "" {
	//		//echo client Origin
	//		c.Writer.Header().Set("Access-Control-Allow-Origin", origin)
	//		//allowed methods
	//		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE,UPDATE")
	//		//exposed headers
	//		c.Header("Access-Control-Allow-Headers", "Authorization, Content-Length, X-CSRF-Token, Token,session")
	//		// allowed request headers
	//		c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers")
	//		//preflight cache max-age
	//		c.Header("Access-Control-Max-Age", "172800")
	//		//allow credentials
	//		c.Header("Access-Control-Allow-Credentials", "true")
	//	}
	//
	//	//allow content-type checks
	//	if method == "OPTIONS" {
	//		c.JSON(http.StatusOK, "ok!")
	//	}
	//
	//	defer func() {
	//		if err := recover(); err != nil {
	//			log.Printf("Panic info is: %v", err)
	//		}
	//	}()
	//
	//	c.Next()
	//}
}
