package rest

import (
	"context"

	"github.com/labstack/echo/v4"
)

type InquiryUseCase interface {
	Inquiry(ctx context.Context)
}

type inquiryHandler struct {
	usecase InquiryUseCase
}

func NewInquiry(inquiryUseCase InquiryUseCase) *inquiryHandler {
	return &inquiryHandler{
		usecase: inquiryUseCase,
	}
}

func (i *inquiryHandler) InitRoutes(e *echo.Group) {
}

func (i *inquiryHandler) Inquiry(c echo.Context) error {

}
