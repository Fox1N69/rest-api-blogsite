package routers

import (
	"blog/internal/handlers"

	"github.com/labstack/echo/v4"
)

type PostRouter struct {
	handler *handlers.PostHandler
}

func NewPostRouter(postHander *handlers.PostHandler) *PostRouter {
	return &PostRouter{
		handler: postHander,
	}
}

func (r *PostRouter) RouterSetup(e *echo.Echo) {
	api := e.Group("/api/post")
	{
		api.GET("/", r.handler.GetAllPosts)
		api.POST("/", r.handler.CreatePost)
		api.PUT(":id", r.handler.UpdatePost)
		api.DELETE(":id", r.handler.DeletePost)
	}
}