package middleware

import (
	"task/app/core"
	"task/app/presentation/responce"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

func Request(logger *zap.Logger) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			req := c.Request()
			res := c.Response()

			uuidObj, _ := uuid.NewUUID()
			uuid := uuidObj.String()
			c.Set("id", uuid)

			// アクセスログ
			logger.Info("Access Start",
				zap.String("RequestId", uuid),
				zap.String("Host", req.Host),
				zap.String("Protocl", req.Proto),
				zap.String("Path", req.URL.Path),
				zap.String("Method", req.Method),
			)

			// 次の処理を実行
			err := next(c)
			if err != nil {
				dstErr := core.AsAppError(err)
				c.JSON(responce.ConvertErrorCode(dstErr.Code()), responce.NewDefaultRespoce(dstErr.Error()))

				// エラーログ
				msg := dstErr.Error()
				if dstErr.Code() == core.SystemError {
					logger.Error(string(dstErr.Code()),
						zap.String("RequestId", uuid),
						zap.String("detail", msg))
				}
			}

			// レスポンスログ
			logger.Info("Access End",
				zap.String("RequestId", uuid),
				zap.Int("Status", res.Status))

			return nil
		}
	}
}
