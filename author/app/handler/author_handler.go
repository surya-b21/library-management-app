package handler

import (
	"context"

	"github.com/surya-b21/library-management-app/author/app/model"
	"github.com/surya-b21/library-management-app/author/app/pb"
	"github.com/surya-b21/library-management-app/author/app/repo"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

type Handler struct {
	repo *repo.AuthorRepository
	pb.UnimplementedAuthorServiceServer
}

func NewHandler(repo *repo.AuthorRepository) *Handler {
	return &Handler{
		repo: repo,
	}
}

func (h *Handler) Create(ctx context.Context, req *pb.AuthorRequest) (*pb.Author, error) {
	if req.Name == "" {
		return nil, status.Errorf(codes.InvalidArgument, "title cannot empty")
	}

	author := model.AuthorAPI{
		Name: &req.Name,
	}

	result := h.repo.Create(ctx, author)

	return &pb.Author{
		Id:   result.ID.String(),
		Name: *result.Name,
	}, nil
}

func (h *Handler) Get(ctx context.Context, req *emptypb.Empty) (*pb.AuthorList, error) {
	result := h.repo.GetAll(ctx)

	authorList := make([]*pb.Author, len(result))

	for i, v := range result {
		author := pb.Author{
			Id:   v.ID.String(),
			Name: *v.Name,
		}

		authorList[i] = &author
	}

	return &pb.AuthorList{
		Author: authorList,
	}, nil
}

func (h *Handler) GetOne(ctx context.Context, req *pb.AuthorId) (*pb.Author, error) {
	if req.Id == "" {
		return nil, status.Errorf(codes.InvalidArgument, "id cannot empty")
	}

	result := h.repo.Get(ctx, req.Id)

	return &pb.Author{
		Id:   result.ID.String(),
		Name: *result.Name,
	}, nil
}

func (h *Handler) Update(ctx context.Context, req *pb.AuthorUpdate) (*pb.Author, error) {
	if req.Id == "" {
		return nil, status.Errorf(codes.InvalidArgument, "id cannot empty")
	}

	author := model.AuthorAPI{
		Name: &req.Name,
	}

	result := h.repo.Create(ctx, author)

	return &pb.Author{
		Id:   result.ID.String(),
		Name: *result.Name,
	}, nil
}

func (h *Handler) Delete(ctx context.Context, req *pb.AuthorId) (*emptypb.Empty, error) {
	if req.Id == "" {
		return nil, status.Errorf(codes.InvalidArgument, "id cannot empty")
	}

	h.repo.Delete(ctx, req.Id)

	return &emptypb.Empty{}, nil
}
