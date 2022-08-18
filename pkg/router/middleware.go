package router

import (
	"boosters/api/post"
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

type Validator struct {
	v *validator.Validate
}

func NewValidator() *Validator {
	return &Validator{v: validator.New()}
}
func (cv *Validator) Validate(i interface{}) error {
	if err := cv.v.Struct(i); err != nil {

		return err
	}
	return nil
}

type MiddlewareInstance func(h echo.HandlerFunc) echo.HandlerFunc

type Middleware struct {
	v    Validator
	list []MiddlewareInstance
}

func NewMiddleware() *Middleware {
	var list []MiddlewareInstance
	return &Middleware{
		v:    *NewValidator(),
		list: list,
	}
}

func (middleware *Middleware) Use(instance MiddlewareInstance) {
	middleware.list = append(middleware.list, instance)
}
func (m *Middleware) ValidateStruct(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		v := NewValidator()
		var bodyBytes []byte
		bodyBytes, _ = ioutil.ReadAll(c.Request().Body)
		// write back to request body
		c.Request().Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))
		// parse json data
		var post post.Post
		if err := json.Unmarshal(bodyBytes, &post); err != nil {
			return c.JSON(400, "error json.")
		}
		if err := v.Validate(post); err != nil {
			logrus.Error(err)
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		return next(c)
	}
}
