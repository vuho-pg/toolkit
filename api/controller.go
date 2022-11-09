package api

import (
	"github.com/labstack/echo/v4"
	"github.com/vuho-pg/toolkit/errors"
	"net/http"
)

type EchoController interface {
	Register(g *echo.Group)
}

type EchoControllerImp struct {
}

func (e EchoControllerImp) Serve(c echo.Context, resp Response, err error) error {
	if err != nil {
		unwrap, ok := err.(*errors.Error)
		if !ok {
			return c.JSON(http.StatusInternalServerError, InternalError("internal error"))
		}
		return c.JSON(unwrap.StatusCode(), CodeAndMessage(unwrap.StatusCode(), unwrap.Error()))
	}
	return c.JSON(resp.StatusCode(), resp.GetResponse())
}

func (e EchoControllerImp) Error(c echo.Context, err error) error {
	return e.Serve(c, nil, err)
}
