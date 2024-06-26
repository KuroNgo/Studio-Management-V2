package model

import (
	"net/http"

	"github.com/gin-gonic/gin"

	// limiter là một thư viện hỗ trợ giới hạn tốc độ truy cập
	// nó hỗ trợ nhiều loại giới hạn khác nhau
	// - Local rate limiting
	// - Distributed rate limiting
	// - Dynamic rate limiting
	// - Custom rate limiting
	// - Rate limiting with custom storage
	// - Rate limiting with custom matcher
	// - Rate limiting with custom identifier
	// - Rate limiting with custom extractor
	// - Rate limiting with custom limiter
	// - Rate limiting with custom response
	// - Rate limiting with custom error handler
	// - Rate limiting with custom store
	// - Rate limiting with custom algorithm
	"github.com/ulule/limiter/v3"
)

type Server struct {
	Limiter    *limiter.Limiter
	Router     *gin.Engine
	HttpServer *http.Server
}
