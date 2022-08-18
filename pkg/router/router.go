package router

import (
	"boosters/api/pkg/handler"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
	_ "github.com/swaggo/echo-swagger/example/docs"
)

type Router struct {
	hand handler.Handler
}

func NewRouter(h handler.Handler) *Router {
	return &Router{hand: h}
}

func (r *Router) StartRouter() {
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	m := NewMiddleware()

	// Routes
	e.GET("/", r.hand.Hello)
	e.GET("/swagger/*", echoSwagger.WrapHandler)
	posts := e.Group("/posts")
	posts.POST("", r.hand.CreatePost, m.ValidateStruct)
	posts.GET("", r.hand.GetAllPosts)
	posts.GET("/:id", r.hand.GetById)
	posts.PUT("/:id", r.hand.UpdatePost, m.ValidateStruct)
	posts.DELETE("/:id", r.hand.DeletePost)
	// Start server
	e.Logger.Fatal(e.Start(":8000"))

}
