package book

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/surya-b21/library-management-app/auth/app/helper"
	pb "github.com/surya-b21/library-management-app/book/app/pb"
)

func BookIdGet(w http.ResponseWriter, r *http.Request) {
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

	res, err := client.GetOne(ctx, &pb.BookId{Id: id})
	if err != nil {
		log.Fatalln("error sending request:", err)
		helper.NewErrorResponse(w, http.StatusBadRequest, "Bad request")
		return
	}

	json, err := json.Marshal(map[string]interface{}{
		"id":          res.Id,
		"title":       res.Title,
		"pages":       res.Pages,
		"author_id":   res.AuthorId,
		"category_id": res.CategoryId,
	})

	if err != nil {
		log.Fatal(err)
	}
	helper.NewSuccessResponse(w, json)
}
