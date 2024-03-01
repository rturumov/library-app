package handler

import (
	"github.com/gin-gonic/gin"
	"library-app/pkg/service"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
	}

	api := router.Group("/api")
	{
		lists := api.Group("/lists")
		{
			lists.POST("/", h.createList)
			lists.GET("/", h.getAllLists)
			lists.GET("/:id", h.getListById)
			lists.PUT("/:id", h.getListById)
			lists.DELETE("/:id", h.deleteList)

			books := lists.Group(":id/books")
			{
				books.GET("/", h.getAllBooks)
				books.GET("/:book_id", h.getBooksById)
				books.PUT("/:book_id", h.putBook)
				books.DELETE("/:book_id", h.deleteBook)
			}
		}
	}

	return router
}
