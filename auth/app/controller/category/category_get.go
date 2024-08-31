package category

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/surya-b21/library-management-app/auth/app/helper"
	"github.com/surya-b21/library-management-app/auth/app/pb"
	"google.golang.org/protobuf/types/known/emptypb"
)

func CategoryGet(w http.ResponseWriter, r *http.Request) {
	conn := helper.ServerDial("category")
	if conn == nil {
		return
	}
	defer conn.Close()

	client := pb.NewCategoryServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	res, err := client.Get(ctx, &emptypb.Empty{})
	if err != nil {
		log.Println("error sending request:", err)
		helper.NewErrorResponse(w, http.StatusBadRequest, "Bad request")
		return
	}

	listCategory := []map[string]interface{}{}
	for _, category := range res.Category {
		listCategory = append(listCategory, map[string]interface{}{
			"id":   category.Id,
			"name": category.Name,
		})
	}

	json, err := json.Marshal(listCategory)
	if err != nil {
		log.Print(err)
	}

	helper.NewSuccessResponse(w, json)
}
