package middleware

import "github.com/gin-gonic/gin"

func CORSPublic() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type,Content-Length,Accept-Encoding,X-CSRF-Token,Authorization,accept,origin,Cache-Control,X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET")
		c.Next()
	}
}

func CORSForDev() gin.HandlerFunc {
	return func(c *gin.Context) {
		allowedOrigins := "https://localhost"
		c.Writer.Header().Set("Access-Control-Allow-Origin", allowedOrigins)
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, PUT, DELETE, OPTIONS")
		c.Writer.Header().Set("Access-Control-Max-Age", "86400") //24h

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
