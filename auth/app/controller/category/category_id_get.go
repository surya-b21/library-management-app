package category

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/surya-b21/library-management-app/auth/app/helper"
	"github.com/surya-b21/library-management-app/auth/app/pb"
)

func CategoryIdGet(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	if id == "" {
		helper.NewErrorResponse(w, http.StatusBadRequest, "please fill category id")
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

	res, err := client.GetOne(ctx, &pb.CategoryId{Id: id})
	if err != nil {
		fmt.Println("error sending request:", err)
		helper.NewErrorResponse(w, http.StatusBadRequest, "Bad request")
		return
	}

	json, err := json.Marshal(map[string]interface{}{
		"id":   res.Id,
		"name": res.Name,
	})

	if err != nil {
		fmt.Println(err)
	}
	helper.NewSuccessResponse(w, json)
}
