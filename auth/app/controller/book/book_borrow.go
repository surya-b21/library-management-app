package book

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/surya-b21/library-management-app/auth/app/helper"
	"github.com/surya-b21/library-management-app/auth/app/pb"
)

// BookBorrow
// @Summary      BookBorrow
// @Description Borrow a book
// @Tags         Book
// @Accept       application/json
// @Produce		 application/json
// @Param id path string true "Book ID"
// @Success      200  {string} string "Successfully borrow book"
// @Router       /book/borrow/{id} [post]
// @Security BearerAuth
func BookBorrow(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	if id == "" {
		helper.NewErrorResponse(w, http.StatusBadRequest, "please fill book id")
		return
	}

	conn := helper.ServerDial("book")
	if conn == nil {
		return
	}
	defer conn.Close()

	client := pb.NewBookServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err := client.Borrow(ctx, &pb.BookId{Id: id})
	if err != nil {
		fmt.Println("error sending request:", err)
		helper.NewErrorResponse(w, http.StatusBadRequest, "Bad request")
		return
	}

	json, err := json.Marshal(map[string]string{
		"message": "Successfully borrow book",
	})

	if err != nil {
		fmt.Println(err)
	}

	helper.NewSuccessResponse(w, json)
}
