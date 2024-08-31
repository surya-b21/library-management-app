package author

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/surya-b21/library-management-app/auth/app/helper"
	pb "github.com/surya-b21/library-management-app/author/app/pb"
)

func AuthorDelete(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	if id == "" {
		helper.NewErrorResponse(w, http.StatusBadRequest, "please fill book id")
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

	_, err := client.Delete(ctx, &pb.AuthorId{Id: id})
	if err != nil {
		log.Println("error sending request:", err)
		helper.NewErrorResponse(w, http.StatusBadRequest, "Bad request")
		return
	}

	helper.NewSuccessResponse(w, []byte("Successfully delete author"))
}