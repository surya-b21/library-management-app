package book

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/surya-b21/library-management-app/auth/app/helper"
	pbauthor "github.com/surya-b21/library-management-app/author/app/pb"
	pb "github.com/surya-b21/library-management-app/book/app/pb"
	pbcategory "github.com/surya-b21/library-management-app/category/app/pb"
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

	mapResult := map[string]interface{}{
		"id":          res.Id,
		"title":       res.Title,
		"pages":       res.Pages,
		"author_id":   res.AuthorId,
		"category_id": res.CategoryId,
		"stock":       res.Stock,
	}

	// look for author and category
	wg := sync.WaitGroup{}

	wg.Add(2)
	go func() {
		defer wg.Done()

		conn := helper.ServerDial("author")
		if conn == nil {
			return
		}
		defer conn.Close()

		client := pbauthor.NewAuthorServiceClient(conn)

		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		res, err := client.GetOne(ctx, &pbauthor.AuthorId{Id: res.AuthorId})
		if err != nil {
			log.Println("error sending request:", err)
			return
		}

		mapResult["author"] = map[string]interface{}{
			"id":   res.Id,
			"name": res.Name,
		}
	}()

	go func() {
		defer wg.Done()

		conn := helper.ServerDial("category")
		if conn == nil {
			return
		}
		defer conn.Close()

		client := pbcategory.NewCategoryServiceClient(conn)

		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		res, err := client.GetOne(ctx, &pbcategory.CategoryId{Id: res.CategoryId})
		if err != nil {
			log.Println("error sending request:", err)
			return
		}

		mapResult["category"] = map[string]interface{}{
			"id":   res.Id,
			"name": res.Name,
		}
	}()

	wg.Wait()

	json, err := json.Marshal(mapResult)

	if err != nil {
		log.Fatal(err)
	}

	helper.NewSuccessResponse(w, json)
}
