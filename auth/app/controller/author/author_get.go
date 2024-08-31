package author

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

// AuthorGet
// @Summary      AuthorGet
// @Description  Get list author
// @Tags         Author
// @Accept       application/json
// @Produce		 application/json
// @Success      200 {object} []pb.Author
// @Router       /author [get]
func AuthorGet(w http.ResponseWriter, r *http.Request) {
	conn := helper.ServerDial("author")
	if conn == nil {
		return
	}
	defer conn.Close()

	client := pb.NewAuthorServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	res, err := client.Get(ctx, &emptypb.Empty{})
	if err != nil {
		fmt.Println("error sending request:", err)
		helper.NewErrorResponse(w, http.StatusBadRequest, "Bad request")
		return
	}

	listAuthor := []map[string]interface{}{}
	for _, author := range res.Author {
		listAuthor = append(listAuthor, map[string]interface{}{
			"id":   author.Id,
			"name": author.Name,
		})
	}

	json, err := json.Marshal(listAuthor)
	if err != nil {
		fmt.Print(err)
	}

	helper.NewSuccessResponse(w, json)
}
