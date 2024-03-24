package router

import (
	"net/http"
	"os"
	"task/app/application/interface/database"
	"task/app/application/schema"
	"task/app/presentation/controller"
	appmiddleware "task/app/presentation/middleware"

	echojwt "github.com/labstack/echo-jwt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.uber.org/zap"
)

func SetUp(e *echo.Echo, logger *zap.Logger, databaseHandller database.IDatabaseHandller) error {
	csrfController := controller.NewCsrfController()

	// apiグループ
	api := e.Group("/api")

	// CORSミドルウェアの設定
	api.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:3000", os.Getenv("FE_URL")},
		AllowHeaders: []string{
			echo.HeaderOrigin,
			echo.HeaderContentType,
			echo.HeaderAccept,
			echo.HeaderAccessControlAllowHeaders,
			echo.HeaderXCSRFToken,
		},
		AllowMethods:     []string{"POST", "GET", "PUT", "DELETE"},
		AllowCredentials: true,
	}))

	// CSRFミドルウェアの設定
	api.Use(middleware.CSRFWithConfig(middleware.CSRFConfig{
		CookiePath:     "/",
		CookieDomain:   os.Getenv("MY_DOMAIN"),
		CookieHTTPOnly: true,
		CookieSameSite: http.SameSiteNoneMode,
		CookieMaxAge:   60,
	}))

	// リクエストミドルウェアの設定
	api.Use(appmiddleware.Request(logger))

	// default
	api.GET("", func(c echo.Context) error {
		return c.JSON(http.StatusOK, schema.DefaultResponseModel{
			Message: "Hello, world!",
		})
	})

	// CSRF関連
	api.GET("/csrf", csrfController.GetCsrfToken)

	// v1グループ
	v1 := api.Group("/v1")

	// JWTミドルウェアの設定
	v1.Use(echojwt.WithConfig(echojwt.Config{
		SigningKey:  []byte(os.Getenv("JWT_KEY")),
		TokenLookup: "cookie:token",
		ErrorHandler: func(c echo.Context, err error) error {
			return c.JSON(http.StatusUnauthorized, schema.DefaultResponseModel{
				Message: err.Error(),
			})
		},
	}))

	return nil
}
