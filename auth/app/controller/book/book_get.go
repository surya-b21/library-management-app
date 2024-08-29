package book

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/surya-b21/library-management-app/auth/app/helper"
	pb "github.com/surya-b21/library-management-app/book/app/pb"
	"google.golang.org/protobuf/types/known/emptypb"
)

func BookGet(w http.ResponseWriter, r *http.Request) {
	conn := helper.ServerDial("book")
	if conn == nil {
		return
	}
	defer conn.Close()

	client := pb.NewBookServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	res, err := client.Get(ctx, &emptypb.Empty{})
	if err != nil {
		log.Fatalln("error sending request:", err)
		helper.NewErrorResponse(w, http.StatusBadRequest, "Bad request")
		return
	}

	listBook := []map[string]interface{}{}
	for _, book := range res.Book {
		listBook = append(listBook, map[string]interface{}{
			"id":          book.Id,
			"title":       book.Title,
			"pages":       book.Pages,
			"author_id":   book.AuthorId,
			"category_id": book.CategoryId,
		})
	}

	json, err := json.Marshal(listBook)
	if err != nil {
		log.Fatal(err)
	}

	helper.NewSuccessResponse(w, json)
}
