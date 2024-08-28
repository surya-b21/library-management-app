package handler

import (
	"book/app/pb"
	"book/app/repo"
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Handler struct {
	repo *repo.BookRepository
	pb.UnimplementedBookServiceServer
}

func NewHandler(repo *repo.BookRepository) *Handler {
	return &Handler{
		repo: repo,
	}
}

func (h *Handler) Create(ctx context.Context, req *pb.BookRequest) (*pb.Book, error) {
	if req.Title == "" {
		return nil, status.Errorf(codes.InvalidArgument, "title cannot empty")
	}

	if req.Pages == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "pages cannot empty")
	}

	return &pb.Book{
		Title: "Success test it",
	}, nil
}
