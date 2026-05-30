package core_middleware

import (
	"net/http"
	"time"

	core_logger "github.com/cephalopagus/bkv-intalant-task/internal/core/logger"
	"go.uber.org/zap"
)

func Logging(next http.Handler, logger *core_logger.Logger) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		logger.Info("request",
			zap.String("method", r.Method),
			zap.String("path", r.URL.Path),
			zap.String("duration", time.Since(start).String()),
		)
	})
}
