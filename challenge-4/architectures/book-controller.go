package architectures

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Get Books
func (InPg *PgDB) GetBooks(ctx *gin.Context) {
	books, err := GetBooks(InPg, 10)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusAccepted, books)
}

// Get Book
func (InPg *PgDB) GetBook(ctx *gin.Context) {

}

// Add Book
func (InPg *PgDB) AddBook(ctx *gin.Context) {

}

// Update Book
func (InPg *PgDB) UpdateBook(ctx *gin.Context) {

}

// Delete Book
func (InPg *PgDB) DeleteBook(ctx *gin.Context) {

}
