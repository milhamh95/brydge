package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type restHandler struct {
	usecase useCase
}

func NewRest(inquiryUseCase useCase) *restHandler {
	return &restHandler{
		usecase: inquiryUseCase,
	}
}

func (h *restHandler) InitRoutes(e *echo.Group) {
	e.POST("/inquiry", h.Confirm)
}

func (h *restHandler) Confirm(c echo.Context) error {
	var req request

	err := c.Bind(&req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	h.usecase.Confirm(c.Request().Context())
	return c.JSON(http.StatusOK, nil)
}
