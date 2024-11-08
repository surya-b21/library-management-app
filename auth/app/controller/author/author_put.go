package author

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/surya-b21/library-management-app/auth/app/helper"
	"github.com/surya-b21/library-management-app/auth/app/pb"
)

// AuthorPut function
// @Summary      AuthorPut
// @Description  Update a author
// @Tags         Author
// @Accept       application/json
// @Produce		 application/json
// @Param id path string true "Author ID"
// @Param data body author.body true "Author Payload"
// @Success      200 {object} pb.Author
// @Router       /author/{id} [put]
// @Security BearerAuth
func AuthorPut(w http.ResponseWriter, r *http.Request) {
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

	conn := helper.ServerDial("author")
	if conn == nil {
		return
	}
	defer conn.Close()

	client := pb.NewAuthorServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	res, err := client.Update(ctx, &pb.AuthorUpdate{
		Id:   id,
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
