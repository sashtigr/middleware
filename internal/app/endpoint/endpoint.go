package endpoint

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
)

type Service interface {
	DaysLeft() int64
}

type Endpoint struct {
	service Service
}

func New(svc Service) *Endpoint {
	return &Endpoint{
		service: svc,
	}
}

func (endpoint *Endpoint) Status(ctx echo.Context) error {
	days := endpoint.service.DaysLeft()
	s := fmt.Sprintf("Days left: %d", days)

	err := ctx.String(http.StatusOK, s)
	if err != nil {
		return err
	}

	return nil
}
