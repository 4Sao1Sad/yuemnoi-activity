package handler

import (
	"context"

	"github.com/4Sao1Sad/yuemnoi-activity/internal/model"
	"github.com/4Sao1Sad/yuemnoi-activity/internal/repository"
	pb "github.com/4Sao1Sad/yuemnoi-activity/proto/activity"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type ActivityLogGRPC struct {
	pb.UnimplementedActivityLogServiceServer
	repository repository.ActivityLogRepository
}

func NewActivityLogGRPC(repo repository.ActivityLogRepository) *ActivityLogGRPC {
	return &ActivityLogGRPC{
		repository: repo,
	}
}

func (g *ActivityLogGRPC) CreateActivityLog(ctx context.Context, input *pb.CreateActivityLogRequest) (*pb.CreateActivityLogResponse, error) {
	data := model.ActivityLog{
		LogDetail: input.LogDetail,
		UserId:    uint(input.UserId),
	}

	activityLog, err := g.repository.CreateActivityLog(ctx, data.LogDetail, data.UserId)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	resp := pb.CreateActivityLogResponse{
		ActivityLog: &pb.ActivityLog{
			LogDetail: activityLog.LogDetail,
			UserId:    uint64(activityLog.UserId),
			Timestamp: activityLog.Timestamp,
		},
	}

	return &resp, nil
}

func (g *ActivityLogGRPC) ViewActivityHistory(ctx context.Context, input *pb.ViewActivityHistoryRequest) (*pb.ViewActivityHistoryResponse, error) {
	history, err := g.repository.ViewActivityHistoryByUserId(ctx, uint(input.UserId))
	if err != nil {
		return nil, err
	}

	var activityHistory []*pb.ActivityLog
	for _, logEntry := range *history {
		activityHistory = append(activityHistory, &pb.ActivityLog{
			LogDetail: logEntry.LogDetail,
			UserId:    uint64(logEntry.UserId),
			Timestamp: logEntry.Timestamp,
		})
	}

	resp := pb.ViewActivityHistoryResponse{
		ActivityLog: activityHistory,
	}

	return &resp, nil
}
