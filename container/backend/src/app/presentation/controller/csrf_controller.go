package controller

import (
	"net/http"
	"task/app/application/schema"

	"github.com/labstack/echo/v4"
)

type csrfController struct {
}

// 認証コントローラーの作成
func NewCsrfController() *csrfController {
	return &csrfController{}
}

// CSRFトークンの取得
func (ac *csrfController) GetCsrfToken(c echo.Context) error {
	token := c.Get("csrf").(string)

	resonseBody := schema.CSRFModel{
		Csrf: token,
	}
	return c.JSON(http.StatusOK, resonseBody)
}
