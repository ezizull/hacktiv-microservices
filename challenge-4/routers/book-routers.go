package routers

import (
	"microservices/challenge-4/architectures"

	"github.com/gin-gonic/gin"
)

func BookRouters(router *gin.Engine, InPg *architectures.PgDB) {
	RBook := router.Group("/books")
	{
		RBook.GET("", InPg.GetBooks)
		RBook.GET("/:id", InPg.GetBook)
		RBook.POST("", InPg.AddBook)
		RBook.PUT("/:id", InPg.UpdateBook)
		RBook.DELETE("/:id", InPg.DeleteBook)
	}
}
