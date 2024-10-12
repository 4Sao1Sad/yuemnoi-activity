package handler

import (
	"context"
	"time"

	"github.com/4Sao1Sad/yuemnoi-activity/internal/model"
	"github.com/4Sao1Sad/yuemnoi-activity/internal/repository"
	pb "github.com/4Sao1Sad/yuemnoi-activity/proto/activity"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
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

func (g *ActivityLogGRPC) CreateActivityLogGRPC(ctx context.Context, input *pb.CreateActivityLogRequest) (*pb.CreateActivityLogResponse, error) {
	// Construct the activity log entry from the request input
	data := model.ActivityLog{
			LogDetail: input.LogDetail,
			UserId:    input.UserId,
			Timestamp: time.Now(),  // Current timestamp
	}

	// Insert the new activity log into the repository
	activityLog, err := g.repository.CreateActivityLog(data)
	if err != nil {
			return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	// Create the response message with the inserted activity log
	resp := pb.CreateActivityLogResponse{
			ActivityLog: &pb.ActivityLog{
					Id:        activityLog.ID, 
					LogDetail: activityLog.LogDetail,
					UserId:    activityLog.UserId,
					Timestamp: timestamppb.New(activityLog.Timestamp),
			},
	}

	return &resp, nil
}

func (g *ActivityLogGRPC) ViewActivityHistoryGRPC(ctx context.Context, input *pb.ViewActivityHistoryRequest) (*pb.ViewActivityHistoryResponse, error) {
	// Fetch logs by user_id from the repository
	history, err := g.repository.ViewActivityHistoryByUserId(input.UserId)
    if err != nil {
        return nil, err
    }

    // Convert logs from the repository format to the protobuf format
    var activityHistory []*pb.ActivityLog
    for _, logEntry := range history {
			activityHistory = append(activityHistory, &pb.ActivityLog{
            Id:        logEntry.ID,
            LogDetail: logEntry.LogDetail,
            UserId:    logEntry.UserId,
            Timestamp: timestamppb.New(logEntry.Timestamp),
        })
    }

    // Create the response message
    resp := pb.ViewActivityHistoryResponse{
        ActivityLogs: activityHistory,
    }

    return &resp, nil
}