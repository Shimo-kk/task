package controller

import (
	"net/http"
	"task/app/application/interface/usecase"
	"task/app/application/schema"
	"task/app/core"
	"task/app/presentation/responce"

	"github.com/labstack/echo/v4"
)

type workspaceController struct {
	usecase usecase.IWorkspaceUsecase
}

// ワークスペースコントローラーの作成
func NewWorkspaceController(usecase usecase.IWorkspaceUsecase) *workspaceController {
	return &workspaceController{usecase: usecase}
}

// ワークスペースの作成
func (c *workspaceController) CreateWorkspace(context echo.Context) error {
	requestBody := schema.WorkspaceCreateModel{}
	if err := context.Bind(&requestBody); err != nil {
		return context.JSON(http.StatusBadRequest, responce.NewDefaultRespoce(err.Error()))
	}

	if err := c.usecase.CreateWorkspace(requestBody); err != nil {
		dstErr := core.AsAppError(err)
		if dstErr.Code() == core.SystemError {
			return err
		} else {
			return context.JSON(responce.ConvertErrorCode(dstErr.Code()), responce.NewDefaultRespoce(dstErr.Error()))
		}
	}

	return context.JSON(http.StatusOK, responce.NewDefaultRespoce("ワークスペースを作成しました。"))
}
