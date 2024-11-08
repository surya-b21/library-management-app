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

type body struct {
	Name string `json:"name,omitempty"`
}

// CategoryPost function
// @Summary      CategoryPost
// @Description  Create a new category
// @Tags         Category
// @Accept       application/json
// @Produce		 application/json
// @Param data body category.body true "Category Payload"
// @Success      200 {object} pb.Category
// @Router       /category [post]
// @Security BearerAuth
func CategoryPost(w http.ResponseWriter, r *http.Request) {
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

	res, err := client.Create(ctx, &pb.CategoryRequest{
		Name: body.Name,
	})

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
