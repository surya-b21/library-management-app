package book

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/surya-b21/library-management-app/auth/app/helper"
	"github.com/surya-b21/library-management-app/auth/app/pb"
)

type body struct {
	Title      string `json:"title,omitempty"`
	Pages      int32  `json:"pages,omitempty"`
	AuthorID   string `json:"author_id,omitempty"`
	CategoryID string `json:"category_id,omitempty"`
	Stock      int32  `json:"stock,omitempty"`
}

func BookPost(w http.ResponseWriter, r *http.Request) {
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

	res, err := client.Create(ctx, &pb.BookRequest{
		Title:      body.Title,
		Pages:      body.Pages,
		AuthorId:   body.AuthorID,
		CategoryId: body.CategoryID,
		Stock:      body.Stock,
	})

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
		"stock":       res.Stock,
	})

	if err != nil {
		log.Fatal(err)
	}
	helper.NewSuccessResponse(w, json)
}
