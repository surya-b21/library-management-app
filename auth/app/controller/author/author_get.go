package author

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
		log.Fatalln("error sending request:", err)
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
		log.Print(err)
	}

	helper.NewSuccessResponse(w, json)
}
