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

// BookPut
// @Summary      BookPut
// @Description  Edit book
// @Tags         Book
// @Accept       application/json
// @Produce		 application/json
// @Param id path string true "Book ID"
// @Param data body book.body true "Book Payload"
// @Success      200 {object} pb.Book
// @Router       /book/{id} [put]
// @Security BearerAuth
func BookPut(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	if id == "" {
		helper.NewErrorResponse(w, http.StatusBadRequest, "please fill book id")
		return
	}

	var body body
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		helper.NewErrorResponse(w, http.StatusBadRequest, err.Error())
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

	res, err := client.Update(ctx, &pb.Book{
		Id:         id,
		Title:      body.Title,
		Pages:      body.Pages,
		AuthorId:   body.AuthorID,
		CategoryId: body.CategoryID,
		Stock:      body.Stock,
	})
	if err != nil {
		fmt.Println("error sending request:", err)
		helper.NewErrorResponse(w, http.StatusBadRequest, "Bad request")
		return
	}

	json, err := json.Marshal(map[string]interface{}{
		"id":          res.Id,
		"title":       res.Title,
		"pages":       res.Pages,
		"author_id":   res.AuthorId,
		"category_id": res.CategoryId,
		"stock":       res.Stock,
	})

	if err != nil {
		fmt.Println(err)
	}
	helper.NewSuccessResponse(w, json)
}
