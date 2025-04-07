package handler

import (
	pb "cat_alog/internal/api/grpc"
	"cat_alog/internal/domain/model"
	"cat_alog/internal/domain/service"
	"github.com/google/uuid"
	"golang.org/x/net/context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"time"
)

type GrpcCatHandler struct {
	catService service.CatService
	pb.UnimplementedCatServiceServer
}

func NewGrpcCatHandler(catService service.CatService) *GrpcCatHandler {
	return &GrpcCatHandler{catService: catService}
}

func (g *GrpcCatHandler) GetCatById(ctx context.Context, request *pb.GetCatByIdRequest) (*pb.GetCatByIdResponse, error) {
	cat, err := g.catService.GetById(request.GetId())
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &pb.GetCatByIdResponse{
		Cat: &pb.Cat{
			Id:          cat.Id,
			Name:        cat.Name,
			DateOfBirth: cat.DateOfBirth.String(),
			ImageUrl:    cat.ImageUrl,
		},
	}, nil
}

func (g *GrpcCatHandler) CreateCat(ctx context.Context, request *pb.CreateCatRequest) (*pb.CreateCatResponse, error) {
	id, err := uuid.NewUUID()
	if err != nil {
		return nil, status.Error(codes.Internal, "failed to generate id")
	}
	dateOfBirth, err := time.Parse("2006-01-02", request.DateOfBirth)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	cat := model.Cat{
		Id:          id.String(),
		Name:        request.Name,
		DateOfBirth: dateOfBirth,
		ImageUrl:    request.ImageUrl,
	}
	err = g.catService.Create(&cat)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	return &pb.CreateCatResponse{Id: id.String()}, nil
}

func (g *GrpcCatHandler) SearchCat(ctx context.Context, search *pb.SearchCatByText) (*pb.Cats, error) {
	searchResult, err := g.catService.Search(search.GetSearchText())
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	cats := make([]*pb.Cat, len(searchResult))
	for i, item := range searchResult {
		cats[i] = &pb.Cat{
			Id:          item.Id,
			Name:        item.Name,
			DateOfBirth: item.DateOfBirth.String(),
			ImageUrl:    item.ImageUrl,
		}
	}
	return &pb.Cats{Cats: cats}, nil
}
