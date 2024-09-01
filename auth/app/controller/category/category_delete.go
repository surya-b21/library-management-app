package category

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/surya-b21/library-management-app/auth/app/helper"
	"github.com/surya-b21/library-management-app/auth/app/pb"
)

// CategoryDelete function
// @Summary      CategoryDelete
// @Description  Delete a category
// @Tags         Category
// @Accept       application/json
// @Produce		 application/json
// @Param id path string true "Category ID"
// @Success      200  {string} string "Successfully delete category"
// @Router       /category/{id} [delete]
// @Security BearerAuth
func CategoryDelete(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	if id == "" {
		helper.NewErrorResponse(w, http.StatusBadRequest, "please fill book id")
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

	_, err := client.Delete(ctx, &pb.CategoryId{Id: id})
	if err != nil {
		fmt.Println("error sending request:", err)
		helper.NewErrorResponse(w, http.StatusBadRequest, "Bad request")
		return
	}

	helper.NewSuccessResponse(w, []byte("Successfully delete category"))
}
