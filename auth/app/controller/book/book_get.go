package book

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
	"time"

	"github.com/surya-b21/library-management-app/auth/app/helper"
	"github.com/surya-b21/library-management-app/auth/app/pb"
)

// BookGet function
// @Summary      BookGet
// @Description  Get list of book
// @Tags         Book
// @Accept       application/json
// @Produce		 application/json
// @Param title query string false "Book title"
// @Success      200 {object} []pb.Book
// @Router       /book [get]
func BookGet(w http.ResponseWriter, r *http.Request) {
	search := r.URL.Query().Get("title")

	conn := helper.ServerDial("book")
	if conn == nil {
		return
	}
	defer conn.Close()

	client := pb.NewBookServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	res, err := client.Get(ctx, &pb.BookParam{
		Title: search,
	})
	if err != nil {
		fmt.Println("error sending request:", err)
		helper.NewErrorResponse(w, http.StatusBadRequest, "Bad request")
		return
	}

	listBook := []map[string]interface{}{}
	for _, book := range res.Book {
		// look for author and category
		wg := sync.WaitGroup{}
		author := map[string]interface{}{}
		category := map[string]interface{}{}

		wg.Add(2)
		go func() {
			defer wg.Done()

			conn := helper.ServerDial("author")
			if conn == nil {
				return
			}
			defer conn.Close()

			client := pb.NewAuthorServiceClient(conn)

			ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
			defer cancel()

			res, err := client.GetOne(ctx, &pb.AuthorId{Id: book.AuthorId})
			if err != nil {
				fmt.Println("error sending request:", err)
				return
			}

			author = map[string]interface{}{
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

			client := pb.NewCategoryServiceClient(conn)

			ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
			defer cancel()

			res, err := client.GetOne(ctx, &pb.CategoryId{Id: book.CategoryId})
			if err != nil {
				fmt.Println("error sending request:", err)
				return
			}

			category = map[string]interface{}{
				"id":   res.Id,
				"name": res.Name,
			}
		}()

		wg.Wait()

		listBook = append(listBook, map[string]interface{}{
			"id":          book.Id,
			"title":       book.Title,
			"pages":       book.Pages,
			"author_id":   book.AuthorId,
			"author":      author,
			"category_id": book.CategoryId,
			"category":    category,
			"stock":       book.Stock,
		})
	}

	json, err := json.Marshal(listBook)
	if err != nil {
		fmt.Println(err)
	}

	helper.NewSuccessResponse(w, json)
}
