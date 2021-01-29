package middleware

import (
	"errors"

	"github.com/labstack/echo/v4"
	emid "github.com/labstack/echo/v4/middleware"
	"golang.org/x/time/rate"
)

type (
	// ReqLimitConfig 基于令牌桶的限流器配置项
	ReqLimitConfig struct {
		Skipper emid.Skipper

		Limiter *rate.Limiter
	}
)

var (
	// defaultReqLimitConfig 默认请求限制器配置
	defaultReqLimitConfig = ReqLimitConfig{
		Skipper: emid.DefaultSkipper,
		// limit = 1000 单位时间内（秒）放入到桶中的令牌数, 这个参数通常取决于你单机的qps
		// cap = 100 令牌桶的容量大小, 应对瞬时请求数
		Limiter: rate.NewLimiter(1000, 100),
	}

	// ErrAPIRequestLimit 请求过于频繁
	ErrAPIRequestLimit = errors.New("api request limit")
)

// ReqLimit 使用默认请求限制器
func ReqLimit() echo.MiddlewareFunc {
	return ReqLimitWithConfig(defaultReqLimitConfig)
}

// ReqLimitWithConfig 基于配置生成请求限制器
func ReqLimitWithConfig(config ReqLimitConfig) echo.MiddlewareFunc {

	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) (err error) {
			if config.Skipper(c) {
				return next(c)
			}

			if !config.Limiter.Allow() {
				return ErrAPIRequestLimit
			}

			return next(c)
		}
	}

}
