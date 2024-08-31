package category

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/surya-b21/library-management-app/auth/app/helper"
	"github.com/surya-b21/library-management-app/auth/app/pb"
)

func CategoryPut(w http.ResponseWriter, r *http.Request) {
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

	conn := helper.ServerDial("category")
	if conn == nil {
		return
	}
	defer conn.Close()

	client := pb.NewCategoryServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	res, err := client.Update(ctx, &pb.CategoryUpdate{
		Id:   id,
		Name: body.Name,
	})
	if err != nil {
		log.Println("error sending request:", err)
		helper.NewErrorResponse(w, http.StatusBadRequest, "Bad request")
		return
	}

	json, err := json.Marshal(map[string]interface{}{
		"id":   res.Id,
		"name": res.Name,
	})

	if err != nil {
		log.Println(err)
	}
	helper.NewSuccessResponse(w, json)
}
