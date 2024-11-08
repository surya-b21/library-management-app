package category

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/surya-b21/library-management-app/auth/app/helper"
	"github.com/surya-b21/library-management-app/auth/app/pb"
	"google.golang.org/protobuf/types/known/emptypb"
)

// CategoryGet
// @Summary      CategoryGet
// @Description  Get list category
// @Tags         Category
// @Accept       application/json
// @Produce		 application/json
// @Success      200 {object} []pb.Category
// @Router       /category [get]
// @Security BearerAuth
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
		fmt.Println("error sending request:", err)
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
		fmt.Print(err)
	}

	helper.NewSuccessResponse(w, json)
}
