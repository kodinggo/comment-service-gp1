package grpcsvc

import (
	pb "github.com/kodinggo/comment-service-gp1/pb/comment"
	"context"
)

type CommentService struct {
	pb.UnimplementedCommentServiceServer
}

func NewCommentService() *CommentService {
	return &CommentService{}
}

func (s *CommentService) FindAllCommentsByStoryID(context.Context, *pb.FindAllCommentsByStoryIDRequest) (*pb.Comments, error) {
	comments := []*pb.Comment{
		{
			Id:      1,
			Comment: "Ini contoh komentar",
		},
		{
			Id:      2,
			Comment: "hehehe",
		},
		{
			Id:      2,
			Comment: "Hanya coba saja",
		},
	}

	return &pb.Comments{
		Comments: comments,
	}, nil
}
